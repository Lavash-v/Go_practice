package main

import "fmt"

func reverse(arr *[]int) {
	if arr == nil || len(*arr) <= 1 {
		return
	}

	s := *arr

	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("До реверса:", s)

	reverse(&s)

	fmt.Println("После реверса:", s)
}
