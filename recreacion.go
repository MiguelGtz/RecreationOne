package main

import (
	"fmt"
)

func sumacuadrados(num int) int {
	var suma int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			suma += i * i
		}
	}
	return suma
}

func ListSquared(m, n int) [][]int {

}

func main() {
	fmt.Println(ListSquared(1, 250))
	fmt.Println(ListSquared(42, 250))
}
