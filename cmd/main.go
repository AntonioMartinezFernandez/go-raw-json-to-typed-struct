package main

import (
	"encoding/json"
	"fmt"
	"go-raw-json-to-typed-struct/internal"
	"go-raw-json-to-typed-struct/pkg"
)

func main() {
	/* RAW STUDENT */
	rawStudent :=
		`
		{
			"name":"Albert",
			"age": 28,
			"subjects": [
				{
					"name": "Electrical engineering",
					"level": 5
				},
				{
					"name": "Thermodynamics",
					"level": 4
				}
			]
		}
		`

	/* CLASSIC UNMARSHALL */
	var unmarshalled_student internal.Student
	err := json.Unmarshal([]byte(rawStudent), &unmarshalled_student)
	if err != nil {
		fmt.Println("unmarshall error: ", err)
	}
	fmt.Println(unmarshalled_student)

	/* UNMARSHALL USING CUSTOM LIBRARY */
	st, err := pkg.RawJsonToTypedStruct[internal.Student](rawStudent)
	if err != nil {
		fmt.Println("unmarshall with custom library error: ", err)
	}
	fmt.Println(*st)
}
