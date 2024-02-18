package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 float64
	var operator string

	// Prompt the user to enter the first number
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	// Prompt the user to enter the operator
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Prompt the user to enter the second number
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	// Perform the calculation based on the operator provided
	result := calculate(num1, num2, operator)

	// Print the result
	fmt.Printf("Result: %f\n", result)
}

// Function to perform calculation based on operator
func calculate(num1, num2 float64, operator string) float64 {
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero")
			// Returning NaN (Not a Number) for division by zero
			return 0
		}
	default:
		fmt.Println("Error: Invalid operator")
		// Returning NaN (Not a Number) for invalid operator
		return 0
	}
	return result
}
