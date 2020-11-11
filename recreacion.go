package main

import (
	"fmt"
	"math"
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
	resultado := [][]int{}
	for i := m; i <= n; i++ {
		raiz := math.Sqrt(float64(sumacuadrados(i)))
		raizentera := float64(int(math.Sqrt(float64(sumacuadrados(i)))))
		if raiz == raizentera {
			aux := []int{i,sumacuadrados(i)}
			resultado = append(resultado, aux)
		}
	}
	return resultado
}

func main() {
	fmt.Println(ListSquared(1, 250))
	fmt.Println(ListSquared(42, 250))
}
