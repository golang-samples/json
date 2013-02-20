package main

import (
	"encoding/json"
	"fmt"
)

const text = `
{
	"foo": "hello",
	"bar": "golang!"
}
`

func main() {
	var m = make(map[string]string)
	json.Unmarshal([]byte(text), &m)
	fmt.Println(m["foo"], m["bar"])
}
