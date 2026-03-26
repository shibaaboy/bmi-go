package main

import (
	"fmt"
	"math"
)

const IMT_POWER = 2

func main() {
	userWeight, userHeight := getUserValues()

	if userHeight == 0 || userWeight == 0 {
		fmt.Println("Invalid values. Please enter correct your parameters.")
		return
	}

	bmi := calculateBmi(userWeight, userHeight)
	outputResultBmi(bmi)
}

func outputResultBmi(bmi float64) {
	fmt.Printf("Your BMI: %.0f\n", bmi)
	fmt.Println("Category:", getBmiCategory(bmi))
}

func calculateBmi(userWeight, userHeight float64) float64 {
	heightMeters := userHeight / 100
	result := userWeight / math.Pow(heightMeters, IMT_POWER)
	return result
}

func getUserValues() (float64, float64) {
	var userHeight float64
	var userWeight float64

	fmt.Print("Your height (cm): ")
	fmt.Scan(&userHeight)

	fmt.Print("Your weight (kg): ")
	fmt.Scan(&userWeight)

	return userWeight, userHeight
}

func getBmiCategory(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi < 25:
		return "Normal"
	case bmi < 30:
		return "Overweight"
	default:
		return "Obese"
	}
}
