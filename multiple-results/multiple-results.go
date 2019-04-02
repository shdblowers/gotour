package main

import "fmt"

// this function is returning multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "there")
	fmt.Println(a, b)
}
