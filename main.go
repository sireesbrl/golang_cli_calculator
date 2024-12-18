package main

import "fmt"

func main() {
	var input1 int
	var input2 int
	var operator string
	fmt.Println("Welcome to the GoLang Calculator!")
	fmt.Println("Valid operators: +, -, *, /")
	fmt.Println("Press Ctrl+C to quit.")

	for {
		fmt.Println("Enter an expression:")
		fmt.Scanf("%d %s %d", &input1, &operator, &input2)

		switch operator {
		case "+":
			fmt.Println("Sum:", add(input1, input2))
		case "-":
			fmt.Println("Difference:", subtract(input1, input2))
		case "*":
			fmt.Println("Product:", multiply(input1, input2))
		case "/":
			fmt.Println(divide(input1, input2))
		default:
			fmt.Println("Invalid operator")
		}
	}

}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) string {
	if b == 0 {
		return "Cannot divide by zero"
	}
	quotient := a / b
	remainder := a % b
	return fmt.Sprintf("Quotient: %d, Remainder: %d", quotient, remainder)
}
