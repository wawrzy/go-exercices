package main

import (
	"fmt"
	"gostoredb/src/skv"
)

func main() {
	store, err := skv.Open("./data.db")

	if err != nil {
		panic(err)
	}

	store.Put("hel", "eee")
	store.Put("hello", 42)

	fmt.Println(store.Get("hel"))

	store.Delete("hel")

	fmt.Println(store.Get("hel"))
}
