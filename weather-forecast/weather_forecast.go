// Package weather: Returns a string with current weather conditon for the given loation.
package weather

// CurrentCondition: current weather condition.
var CurrentCondition string

// CurrentLocation: current user location.
var CurrentLocation string

// Forecast: Returns a string with current weather conditon for the given loation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
