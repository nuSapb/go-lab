package main

import "fmt"

func main() {
	/* Style 1
	   var variable_name type
	*/
	var gopher string
	gopher = "Gopher"
	fmt.Printf("Hello, %s.\n", gopher)

	/* Style 2
	var variable_name value
	*/
	var name = "Anu Sakpibal"
	fmt.Printf("My name is %s.\n", name)

	/* Style 3
	variable_name := value
	*/
	nickName := "nuSapb"
	fmt.Printf("My nick name is %s.\n", nickName)

}
