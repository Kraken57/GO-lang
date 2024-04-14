package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learn Go Language ")

	var name string = "Ahmad"
	fmt.Println(name)

	var a = 65
	fmt.Println(a)


	person := 123
	fmt.Println(person)


	// Public variable (exported)
	var PublicVariable int = 42
	// Private variable (unexported)
	var privateVariable int = 10
	fmt.Println(PublicVariable, privateVariable)
	// ● In this example, PublicVariable is visible and can be accessed from other
	// packages, while privateVariable is only visible within the same package
}