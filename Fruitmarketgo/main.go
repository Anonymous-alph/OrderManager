package main

import (
	"fmt"
)

func main() {

	fruitsavailable := map[string]int{
		"Apple":  10,
		"Mango":  20,
		"Banana": 100,
	}
	for key, value := range fruitsavailable {
		fmt.Printf("%s $%v\n", key, value)

	}
	var input string
	var inputfruits string
	fmt.Print("Enter the fruits you want from the list shown: ")
	fmt.Scanln(&inputfruits)
	fmt.Print("Do you want to continue shopping (y/n): ")
	fmt.Scanln(&input)

	//Switch and case for "Fruits Selected from menu"

	switch inputfruits {
	case "Apple":
		fmt.Printf("you have chosen %s (y/n) ", inputfruits)
		fmt.Scanln(&inputfruits)

	case "Banana":
		fmt.Printf("you have chosen %s (y/n) ", inputfruits)
		fmt.Scanln(&inputfruits)

	case "Mango":
		fmt.Printf("you have chosen %s (y/n) ", inputfruits)
		fmt.Scanln(&inputfruits)
	default:
		fmt.Printf("Invalid input: %v\n", inputfruits)

	}

	switch input {
	case "y":
		fmt.Print("Enter the fruits you want from the list shown: ")
		fmt.Scanln(&inputfruits)
		fmt.Print("Do you want to continue shopping (y/n): ")
		fmt.Scanln(&input)
		switch inputfruits {
		case "Apple":
			fmt.Printf("you have chosen %s (y/n) ", inputfruits)
			fmt.Scanln(&inputfruits)

		case "Banana":
			fmt.Printf("you have chosen %s (y/n) ", inputfruits)
			fmt.Scanln(&inputfruits)

		case "Mango":
			fmt.Printf("you have chosen %s (y/n) ", inputfruits)
			fmt.Scanln(&inputfruits)
		default:
			fmt.Printf("Invalid input: %v\n", inputfruits)

		}
	case "n":
		fmt.Print("Thank you for shopping with us")
	default:
		fmt.Print("Invalid Input")
		fmt.Print("Enter the fruits you want from the list shown: ")
		fmt.Scanln(&inputfruits)
		fmt.Print("Do you want to continue shopping (y/n): ")
		fmt.Scanln(&input)
		switch inputfruits {
		case "Apple":
			fmt.Printf("you have chosen %s (y/n) ", inputfruits)
			fmt.Scanln(&inputfruits)

		case "Banana":
			fmt.Printf("you have chosen %s (y/n) ", inputfruits)
			fmt.Scanln(&inputfruits)

		case "Mango":
			fmt.Printf("you have chosen %s (y/n) ", inputfruits)
			fmt.Scanln(&inputfruits)
		default:
			fmt.Printf("Invalid input: %v\n", inputfruits)

		}

	}
}
