# Weather 
Command line tool for getting daily weather forecasts.

## How to use?

##### Prequisites
- Have GO installed

##### Run program
- Possible flags
    - city -> (string) name of the city
    - forecast24h -> (true/false) see forecast for the coming 24 hours or till the end of the day

- example usage
```bash
$ go build cmd/main.go
$ ./weather -city=los-angeles -forecast24h=false
```

- example output
```
Current weather: Los Angeles, United States of America: 17.8C, Overcast

* 00:00 ⇨ 15.6°C | 72.0% | Patchy rain possible
* 01:00 ⇨ 15.6°C | 77.0% | Patchy rain possible
* 02:00 ⇨ 15.4°C | 80.0% | Patchy rain possible
* 03:00 ⇨ 15.1°C | 0.0% | Overcast
* 04:00 ⇨ 14.8°C | 62.0% | Patchy rain possible
* 05:00 ⇨ 14.6°C | 71.0% | Patchy rain possible
* 06:00 ⇨ 14.5°C | 100.0% | Patchy rain possible
* 07:00 ⇨ 14.4°C | 87.0% | Patchy rain possible
* 08:00 ⇨ 14.3°C | 79.0% | Patchy rain possible
```
