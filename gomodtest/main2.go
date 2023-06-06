package main

import (
	//"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	list := []ColorGroup{group, group}

	b, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	ioutil.WriteFile("json", b, 0644)
}