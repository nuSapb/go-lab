package main

import "fmt"

func main() {
	fmt.Printf("input a fruit: ")
	var fruit string
	fmt.Scanln(&fruit)
	if fruit == "" {
		fmt.Printf("meh! 😣")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("🍎")
	case "banana":
		fmt.Println("🍌")
	case "lemon":
		fmt.Println("🍋")
	default:
		fmt.Println("😱")
	}

}
