package main

import (
	"fmt"

	"math/rand"
	"time"

	"github.com/fananchong/gomap"
)

func main() {
	// Init new RandomMap
	m := gomap.NewRandomMap()

	// Set key
	m.Set("a", 1)
	m.Set("c", 2)
	m.Set("b", 3)
	m.Set("d", 4)
	m.Set("e", 5)
	m.Set("f", 6)

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

	rand.Seed(time.Now().UnixNano())

	// Random
	for i := 0; i < 10; i++ {
		fmt.Println(m.Random())
	}
}
