package main

import (
	"fmt"

	"github.com/hisyntax/phonenumber"
)

func main() {
	number := "09132600841"
	num, err := phonenumber.Parse(number)
	if err != nil {
		fmt.Printf("This is the error: %v\n", err)
	}

	fmt.Printf("This is the new formatted number: %v\n", num)
}
