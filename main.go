package main

import "fmt"

func main() {
	fmt.Println(somar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func somar(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}
