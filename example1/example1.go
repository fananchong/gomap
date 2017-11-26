package main

import (
	"fmt"

	"github.com/fananchong/gomap"
)

func less(v1, v2 interface{}) bool {
	return v1.(int) < v2.(int)
}

func main() {
	// Init new OrderedMap
	m := gomap.NewOrderedMap(less)

	// Set key
	m.Set("a", 1)
	m.Set("c", 3)
	m.Set("b", 2)
	m.Set("d", 4)
	m.Set("e", 2)
	m.Set("f", 2)

	// Get key
	if val, ok := m.Get("a"); ok == true {
		// Found key "a"
		fmt.Println(val)
	}

	// Delete a key
	m.Delete("e")

	// Failed Get lookup becase we deleted "e"
	if _, ok := m.Get("e"); ok == false {
		// Did not find key "e"
		fmt.Println("e not found")
	}

	// Iterator
	for it := m.Iterator(); it.HasNext(); {
		fmt.Println(it.Next())
	}

	// First
	fmt.Println(m.First())

	// Back
	fmt.Println(m.Back())
}
