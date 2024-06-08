package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames"`
	T         T        `json:"T,omitempty"`
	Y         *T       `json:"Y,omitempty"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func main() {
	b := []byte(`{"name":"mike","age":"20", "nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
