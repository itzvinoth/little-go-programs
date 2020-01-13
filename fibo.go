package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	f := []int{0, 1, 1}
	return func(i int) int {
		t := f[i-1] + f[i-2]
		f = append(f, t)
		return t
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		if i > 2 {
			fmt.Println(f(i))
		}
	}
}
