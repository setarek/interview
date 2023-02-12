package main

import (
	"fmt"
	"math"
)

func Calculate(nums []float64, res chan float64) {
	var result float64
	for _, num := range nums {
		result = result + math.Pow(num, 2)
	}
	res <- result
}

func main() {

	nums := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12}

	firstResult := make(chan float64)
	secondResult := make(chan float64)

	go Calculate(nums[0:5], firstResult)
	go Calculate(nums[5:], secondResult)

	a := <-firstResult
	b := <-secondResult
	fmt.Println(a + b)

}
