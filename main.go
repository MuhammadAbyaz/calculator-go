package main

import "fmt"

func main(){
	var operation string
	var num1, num2 int

	fmt.Println("CALCULATOR")
	fmt.Println("==========")
	fmt.Println("Which Operation to perform (add, sub, mul, div): ")
	fmt.Scan(&operation)
	fmt.Print("Enter First Number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number: ")
	fmt.Scan(&num2)
	switch operation{
		case "add":
			fmt.Println(num1 + num2)
		case "sub":
			fmt.Println(num1 - num2)
		case "mul":
			fmt.Println(num1 * num2)
		case "div":
			fmt.Println(num1 / num2)
		default:
			fmt.Println("Please enter valid operation")
	}
}