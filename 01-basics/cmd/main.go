package main

import (
	"fmt"
	"gitlab.com/salivare/go-learn/01-basics/internal/api"
)

const testConst = "hello"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {

	t := api.Test()
	fmt.Println(t)
	m := map[string]int{
		"James": 32,
		"Jack":  42,
		"At":    100,
	}

	for k, v := range m {

		println(k, v)
	}

}

func number(n int) (int, string) {
	return n, testConst
}
