package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) { // if return multiple things
	// specific every return type
	return a, b
}

func main() {
	// var num1, num2 float64 = 5.6, 9.5 // num1 will have a fixed type if complied
	num1, num2 := 5.6, 9.5 //short hand,
	//will get a default type automatically when complied
	// var num2 float64 = 9.5
	fmt.Println(add(num1, num2))
	// if variables and imports are not used, then error are thrown

	w1, w2 := "Word 1", "Word 2"
	fmt.Println(multiple(w1, w2))
	var a int = 62
	var b float64 = float64(a) //type conversion
	fmt.Print(b)
}
