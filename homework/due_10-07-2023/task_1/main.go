package main

import "fmt"

func fibonacci() func() int {
	i1 := 0
	i2 := 1
	return func() int {
		sum := i1 + i2
		i1 = i2
		i2 = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
