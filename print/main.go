package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.8234567

	fmt.Println("Name:", name, "Age:", age, "Height:", height)

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("Type of age: %T\n", age)

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}