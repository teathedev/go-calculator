package main

import (
	"fmt"

	"github.com/teathedev/go-calculator/src/cli"
)

func addition(number1 int, number2 int) int {
	return number1 + number2
}

func subtraction(number1 int, number2 int) int {
	return number1 - number2
}

func multiplication(number1 int, number2 int) int {
	return number1 * number2
}

func division(number1 int, number2 int) int {
	return number1 / number2
}

func main() {
	fmt.Println(`
        Welcome to calculator!
        You can do those actions:
        [1] Addition
        [2] Subtraction
        [3] Multiplication
        [4] Division
    `)

	action, err := cli.GetNumberInput("Select action: ")

	if err != nil {
		return
	}

	number1, err := cli.GetNumberInput("enter number: ")

	if err != nil {
		return
	}

	number2, err2 := cli.GetNumberInput("enter number: ")

	if err2 != nil {
		return
	}

	var result int

	switch action {
	case 1:
		result = addition(number1, number2)
	case 2:
		result = subtraction(number1, number2)
	case 3:
		result = multiplication(number1, number2)
	case 4:
		result = division(number1, number2)
	default:
		fmt.Println("You choose undefined action")
		return
	}

	fmt.Println("Your result is: ", result)
}
