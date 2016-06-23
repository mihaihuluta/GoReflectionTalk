// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows how to reflect over a slice of struct type values.
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

	// Create a slice of struct type user values.
	us := []user{
		{
			name:     "Cindy",
			age:      27,
			building: 321.45,
			secure:   true,
			roles:    []string{"admin", "developer"},
		},
		{
			name:     "Bill",
			age:      40,
			building: 456.21,
			secure:   false,
			roles:    []string{"developer"},
		},
	}

	// Display the information about the slice of users values.
	v := reflect.ValueOf(us)
	fmt.Printf("Kind: %v\tType: %v\n", v.Kind(), v.Type())

	// Iterate over the slice via reflection.
	for i := 0; i < v.Len(); i++ {
		fmt.Println(v.Index(i))
	}
}
