package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	ClassName string `json:"class_name"`
}

func PrintStudent(st Student) string {
	return fmt.Sprintf("FirstName: %s, LastName: %s, Age: %d, ClasName: %s", st.FirstName, st.LastName, st.Age, st.ClassName)
}

func main() {

	jsonStr := `[
		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
	]`

	var students []Student
	json.Unmarshal([]byte(jsonStr), &students)
	for _, st := range students {
		fmt.Println(PrintStudent(st))
	}
}
