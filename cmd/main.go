package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
				TempC     float32 `json:"temp_s"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float32 `json:"chance_of_rain"`
				FeelsLikeC   float32 `json:"feelslike_c"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func readUserInput(m string) (*string, error) {
	fmt.Print(m)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	input = strings.TrimSuffix(strings.ToLower(input), "\n")

	if input == "" {
		input = "Kortrijk"
	} else {
		if err != nil {
			return nil, err
		}
	}

	return &input, nil
}

func main() {
	godotenv.Load(".env")

	city, err := readUserInput("Enter city (Kortrijk): ")
	fmt.Println(*city)
	if err != nil {
		panic(err)
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	fmt.Println(apiKey)
	fmt.Println(*city)

	apiUrl := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, *city)
    fmt.Println(apiUrl)

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
	fmt.Println(forecast)
}
