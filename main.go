package main

import (
	"fmt"
	"time"
)

func goroute(a int, b int) {
	fmt.Printf("%d+%d=%d\n", a, b-a, b)
}
func calc(n int) {
	for i := 0; i <= n; i++ {
		go goroute(i, n)
	}
}
func main() {
	calc(5)
	time.Sleep(3 * time.Second)
}
