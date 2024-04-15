package main

import "fmt"

func multiply(a,b int) int {
	return a*b
}


func add (a,b int) (int){
	return a + b
}

func simpleFunction(){
	fmt.Println("Simple function")
}

func main() {
	fmt.Println("This is the main function")

	simpleFunction()

	ans := add(5,6)
	fmt.Println("Sum is", ans)

	data := multiply(7,8)
	fmt.Println("Product is", data)

}