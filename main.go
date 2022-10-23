package main

import (
	"fmt"
	"math"
)

//====================================================== Math for each of the operations ============================================================

func addition() {
	var num1, num2 float32
	var answer float32

	fmt.Print("Please enter 2 numbers: ")
	fmt.Scan(&num1, &num2)

	answer = num1 + num2
	fmt.Println("The sum of", num1, "&", num2, "is", answer)
}

func subtraction() {
	var num1, num2 float32
	var answer float32

	fmt.Print("Please enter 2 numbers: ")
	fmt.Scan(&num1, &num2)

	answer = num1 - num2
	fmt.Println("The difference between", num1, "&", num2, "is", answer)
}

func multiply() {
	var num1, num2 float32
	var answer float32

	fmt.Print("Please enter 2 numbers: ")
	fmt.Scan(&num1, &num2)

	answer = num1 * num2
	fmt.Println("The product of", num1, "&", num2, "is", answer)
}

func division() {
	var num1, num2 float32
	var answer float32

	fmt.Print("Please enter 2 numbers: ")
	fmt.Scan(&num1, &num2)

	answer = num1 / num2
	fmt.Println("The quotent of", num1, "&", num2, "is", answer)
}

func squareNumber() {
	var num1, num2 float64
	var answer float64

	fmt.Print("Please enter 2 numbers: ")
	fmt.Scan(&num1, &num2)

	answer = math.Pow(num1, num2)
	fmt.Println(num1, "to the power of", num2, "is", answer)
}

//============================================================== Greeting the User ==============================================================

func greetUser() {
	var answer string

	fmt.Println("Welcome to the Interactive Calculator!")
	fmt.Print("What operation do you want to preform? (+, -, *, /, sqr) ")
	fmt.Scan(&answer)

	if answer == "+" {
		addition()
	} else if answer == "-" {
		subtraction()
	} else if answer == "*" {
		multiply()
	} else if answer == "/" {
		division()
	} else if answer == "sqr" {
		squareNumber()
	} else {
		fmt.Println("Invalid Operation")
	}
}

//============================================================== Main function ======================================================================

func main() {
	greetUser()
}
