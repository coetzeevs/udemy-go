package main

import "fmt"

func main() {
	i := makeRange(1, 10)
	for _, v := range i {
		if v%2 == 0 {
			fmt.Printf("%v, Even\n", v)
		} else {
			fmt.Printf("%v, Odd\n", v)
		}
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
