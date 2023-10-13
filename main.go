package main

import "fmt"

func main() {

	var x float64
	var y float64
	var operation string

	fmt.Print("Enter First number: ")
	fmt.Scanln(&x)

	fmt.Print("Enter Second number:")
	fmt.Scanln(&y)

	fmt.Print("Enter Operation:")
	fmt.Scanln(&operation)

	if operation == "+" {
		fmt.Println(x + y)

	} else if operation == "*" {
		fmt.Println(x * y)
	} else if operation == "/" {
		fmt.Println(x / y)
	} else if operation == "-" {
		fmt.Println(x - y)
	} else {
		fmt.Println("invalid symbol")
	}

}
