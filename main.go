package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func square(x int) int {
	return x * x
}

func anything() {
	fmt.Println("...")
}

func main() {
	fmt.Println("Git Stash!")
}
