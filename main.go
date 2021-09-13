package main

import (
	"fmt"
)

func main() {
	var n int
	var r []int
	binario := 0
	counter := 1
	remainder := 0

	fmt.Println("Escriba el valor")
	fmt.Scanf("%d", &n)

	for n != 0 {
		remainder = n % 2
		n = n / 2
		binario += remainder * counter
		counter *= 10
	}
	
	r = append(r, binario)

	fmt.Println(r)
}