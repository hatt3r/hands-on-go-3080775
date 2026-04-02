// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	// var authors map[string]author

	// initialize the map with make
	//authors = make(map[string]author)

	authors := map[string]author{
		"jj": {name: "james joyce"},
		"vm": {name: "virginia woolf"},
	}

	// add authors to the map
	// authors["jj"] = author{name: "James Joyce"}
	// authors["vm"] = author{name: "Virginia Woolf"}

	// print the map with fmt.Printf
	fmt.Printf("%#v\n",authors)

	// read a value from the map with a known key
	fmt.Println(authors["jj"])
}
