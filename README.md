<!-- ![go](/img.gif) -->
# Phonenumber
This library validates phone numbers 
# Installation
To install the sort package, you need to first install [Go](https://golang.org/) and set your Go workspace.
1. You can use the below Go command to install sort
```sh
$ go get -u github.com/iqquee/phonenumber
```
2. Import it in your code:
```sh
import "github.com/iqquee/phonenumber"
```
# Note: 
This only works with nigerian number at the moment. But I intend to cover all other country codes as well.
# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```
## Validate Phone Number
To validate a phone number, use the Parse() method 
```go
package main

import (
	"fmt"

	"github.com/iqquee/phonenumber"
)

func main() {

	number := "09132600841"
	num, err := phonenumber.Parse(number)
	if err != nil {
		fmt.Printf("This is the error: %v\n", err)
	}

	fmt.Printf("This is the new formatted number: %v\n", num)
    //Response: 2349132600841

}
```