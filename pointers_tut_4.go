package main

import (
	"fmt"
)

func main() {
	x := 15
	a := &x // memory address
	fmt.Println(a)
	fmt.Println(*a) // will print the value of variable that a is pointing to
	*a = 5          // reading through of memory address stored at a
	fmt.Println(x)
	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)
}
