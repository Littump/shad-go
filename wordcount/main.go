//go:build !solution

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	m := make(map[string]int)
	for _, arg := range args {
		dat, _ := os.ReadFile(arg)
		s := string(dat)

		for _, word := range strings.Fields(s) {
			m[word]++
		}
	}

	for k, v := range m {
		if v >= 2 {
			fmt.Printf("%v\t%s\n", v, k)
		}
	}
}
