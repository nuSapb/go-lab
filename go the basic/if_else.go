package main

import "fmt"

func main() {
	fmt.Printf("input a score: ")
	var score int
	fmt.Scanln(&score)

	if score < 50 {
		fmt.Println("F")
	} else if score < 60 {
		fmt.Println("D")
	} else if score < 70 {
		fmt.Println("C")
	} else {
		fmt.Println("S")
	}

	// TODO: if else to switch case below***

	switch {
	case score < 50:
		fmt.Println("F")
	case score < 60:
		fmt.Println("D")
	case score < 70:
		fmt.Println("C")
	default:
		fmt.Println("S")
	}

}
