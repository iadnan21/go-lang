// Package weather provides functionality to store and report weather conditions.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location for the weather forecast.
var CurrentLocation string

// Forecast updates the weather information and returns a formatted forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}