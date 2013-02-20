package main

import (
	"encoding/json"
	"fmt"
)

const text = `
{
	"id": 1,
	"name": "プログラミング言語Goフレーズブック",
	"categories": [
		{ "id": 3, "name": "プログラミング" },
		{ "id": 4, "name": "技術" }
	]
}
`
type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Categories []Category `json:"categories"`
}

func main() {
	var book Book
	json.Unmarshal([]byte(text), &book)
	fmt.Printf("%v\n", book)
}
