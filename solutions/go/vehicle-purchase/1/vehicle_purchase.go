package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return strings.ToLower(kind) == "car"  || strings.ToLower(kind) == "truck" 
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var betterChoice string 

	if option1 < option2 {
		betterChoice = option1 
	} else {
		betterChoice = option2
	}

	return betterChoice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var percentage float64 = 100

	if age >= 1 && age < 3 {
		percentage = 80
	} else if age >= 3 && age < 10 {
		percentage = 70
	} else {
		percentage = 50
	}

	return originalPrice * (percentage/100.0)
}
