/*
Package weather is the official meteorological package for the kingdom of Goblinocus.
It provides accurate (well, mostly) weather forecasting services for all goblin cities.
The package maintains global state of the last queried location and its weather condition,
allowing goblins to check both current and historical (last queried) weather data.
*/
package weather

// CurrentCondition stores the weather condition (sunny, rainy, dragon-fog, etc.)
var CurrentCondition string

// CurrentLocation remembers which city was last checked for weather.
var CurrentLocation string

/*
Forecast creates a human-readable weather report for any goblin city.
It takes a city name and weather condition, updates the package's global variables,
and returns a formatted string that even the goblin council can understand.

Parameters:
  - city: the name of the goblin settlement to forecast
  - condition: the current weather phenomenon

Returns:
  A formatted string containing the city name and its weather condition
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}