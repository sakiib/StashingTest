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
	if y == 0 {
		fmt.Println("denominator can't be zero!")
		return -1
	}
	return x / y
}

func square(x int) int {
	return x * x
}

func anything() {
	fmt.Println("...")
}

func sakibWorks() {
	fmt.Println("sakib worked here!")
}

func doShit() {
	fmt.Println("some shitty code added!")
}

func alaminWorks() {
	fmt.Println("alamin worked here!")
}

func main() {
	fmt.Println("Git Stash!")
}
