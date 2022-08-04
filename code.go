package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func main() {

	answer := Add(4, 5)
	fmt.Println(answer)
}
