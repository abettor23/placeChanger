package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 153
	fmt.Println("До обмена значений местами")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// обмен местами значений переменных

	transitb := 0
	transitb += b
	transita := 0
	transita += a
	b = (b - b) + transita
	a = (a - a) + transitb

	fmt.Println("После обмена значений местами")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
