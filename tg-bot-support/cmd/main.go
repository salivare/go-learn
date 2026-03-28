package main

import "fmt"

func main() {
	var exampleArr = [4]string{"2", "5", "7", "9"}

	for _, str := range exampleArr {
		fmt.Println(str)
	}
}
