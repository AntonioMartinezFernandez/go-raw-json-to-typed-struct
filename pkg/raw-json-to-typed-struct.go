package pkg

import "encoding/json"

func RawJsonToTypedStruct[T any](rawJSON string) (*T, error) {
	var typedStruct T
	if err := json.Unmarshal([]byte(rawJSON), &typedStruct); err != nil {
		return nil, err
	}
	return &typedStruct, nil
}
