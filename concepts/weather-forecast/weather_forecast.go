// Package weather provide tools for weather forecasting.
package weather
// CurrentCondition var stores your current condition.
var CurrentCondition string
// CurrentLocation var stores your current location.
var CurrentLocation string
// Forecast func provides a forecast based on CurrentCondition and CurrentLocation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
