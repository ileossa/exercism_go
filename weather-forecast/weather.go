// Package weather provides ability to check weather for a city/location
package weather

var (
	// CurrentCondition represents current weather condition
	CurrentCondition string
	// CurrentLocation represents current geographical location
	CurrentLocation  string
)

// Forecast return a string that shows current weather conditions for a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
