package main

import (
	"fmt"
)

func main() {
	fmt.Println("===================Dibawah Merupakan Tugas Slice Array========================")
	tugas3()
	fmt.Println("=========================================================================================")
}

func tugas3() {
	array1 := make([]int, 0)
	maksimal := 30
	for i := 1; i <= maksimal; i++ {
		count := 0
		for y := 1; y <= i; y++ {
			if i%y == 0 {
				count++
			}
		}
		if count == 2 && i > 1 {
			array1 = append(array1, i)
		}
	}
	fmt.Println(array1)
}
