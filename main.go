package main

import (
	"fmt"
)

func main() {
	g := New(5)

	for i := 0; i < 10; i++ {
		g.Set(i, i*i)
	}

	fmt.Println("[+]Contents :")
	for i := 0; i < 10; i++ {
		if val, err := g.Get(i); err == nil {
			fmt.Println(i, val)
		}
	}
}
