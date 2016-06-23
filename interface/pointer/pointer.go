// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows how to reflect over a struct type pointer that is stored
// inside an interface value.
package main

import (
	"fmt"
	"reflect"
)

// user represents a basic user in the system.
type user struct {
	name     string
	age      int
	building float32
	secure   bool
	roles    []string
}

func main() {

	// Create a struct type user value.
	u := user{
		name:     "Cindy",
		age:      27,
		building: 321.45,
		secure:   true,
		roles:    []string{"admin", "developer"},
	}

	// Store a pointer to the user value inside an empty interface value.
	var i interface{} = &u

	// Display information about the pointer that is stored inside the
	// empty interface value.
	v := reflect.ValueOf(i)
	fmt.Printf("Kind: %v\tType: %v\n", v.Kind(), v.Type())

	// Inspect the value that the pointer points to.
	v = v.Elem()
	fmt.Printf("Kind: %v\tType: %v\t\tNumFields: %v\n", v.Kind(), v.Type(), v.NumField())
}
