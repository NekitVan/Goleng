package main

import (
	"fmt"
)

func main() {
	Task2()
	Task3()
	Task4()
}

// Task2
func Task2() {
	var (
		name string = "NIK"
		age  int    = 21
	)
	fmt.Println(name)
	fmt.Println(age)

}

// Task3
func Task3() {
	var number [5]int = [5]int{1, 2, 3, 4, 5}
	//срез
	var nomber2 = [...]int{1, 2, 3}
	fmt.Println(number)
	fmt.Println(nomber2[1])

}
func Task4() {
	array := []int{1, 3, 5, 7, 1, 5, 7, 21, 1, 1, 0, 3, 5}
	sortedArr := sort(array)
	fmt.Println("массив:", array)
	fmt.Println("Отсортированный массив:", sortedArr)
}

func sort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	pivot := array[0]

	var min []int
	var max []int

	for _, v := range array[1:] {
		if v < pivot {
			min = append(min, v)
		} else {
			max = append(max, v)
		}
	}
	return append(append(sort(min), pivot), sort(max)...)
}
