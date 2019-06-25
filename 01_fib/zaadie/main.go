package main

import "fmt"

func Fib(n int) {
	// Start index
	start := 1

	fmt.Println(start)

	for i := 1; i < n; i++ {

		if i == 1 {
			fmt.Println(i)
		}

		prev := n - (n - i)
		nxt := prev + 1
		
		fmt.Println(prev + nxt)
	}

}

func main() {

	Fib(7)
}