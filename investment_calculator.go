package main

import (
	"fmt"
	"math"
)

func main() {
	// var years = 10 // float64(10) to convert int to float64
	var investmentAmount, years, expectedReturnRate float64
	const inflationRate = 2.5

	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount) // & pointer to the var

	fmt.Println("Enter years: ")
	fmt.Scan(&years)

	fmt.Println("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value: ")
	fmt.Println(futureValue)
	
	fmt.Println("Future real value: ")
	fmt.Println(futureRealValue)
}

// go mod init github.com/yourname/yourrepo

// go build => to create the executable

// ./first-app or double click the .exe file
