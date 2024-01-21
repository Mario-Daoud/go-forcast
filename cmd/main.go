package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type WeatherForecast struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float32 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int32   `json:"time_epoch"`
				TempC     float32 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float32 `json:"chance_of_rain"`
				FeelsLikeC   float32 `json:"feelslike_c"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	godotenv.Load(".env")

	var city string
	var forecast24h bool

	flag.StringVar(&city, "city", "Kortrijk", "City name")
	flag.BoolVar(&forecast24h, "forecast24h", false, "Set to true for a 24-hour forecast")

	flag.Parse()

	apiKey := os.Getenv("WEATHER_API_KEY")

	apiUrl := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&aqi=no&alerts=no", apiKey, city)

	res, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available.")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var forecast WeatherForecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		panic(err)
	}

	location, current, hours := forecast.Location, forecast.Current, forecast.Forecast.ForecastDay[0].Hour
	fmt.Printf("Current weather: %s, %s: %.1fC, %s \n \n", location.Name, location.Country, current.TempC, current.Condition.Text)

	// Get the current time
	now := time.Now()

	for _, hour := range hours {
		date := time.Unix(int64(hour.TimeEpoch), 0)

		// if not full day forecast skip past hours that have already occurred
		if !forecast24h && date.Before(now) {
			continue
		}

		msg := fmt.Sprintf(
			"* %s ⇨ %.1f°C | %.1f%% | %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.TempC <= 0 {
			fmt.Print(msg)
		} else if hour.TempC >= 20 {
			color.Yellow(msg)
		} else {
			color.Blue(msg)
		}
	}
}
