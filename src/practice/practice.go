package main

import (
	"fmt"

	"example.com/backend"
	"rsc.io/quote"
)

type dim struct {
	l int
	b int
	h int
}

func dimensions(l, b, h int) (int, int) {

	return l * b, l * b * h
}

func main() {
	fmt.Println(dimensions(1, 2, 3))
	backend.Run(quote.Go())
}
