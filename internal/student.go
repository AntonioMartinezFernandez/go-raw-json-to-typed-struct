package internal

type Student struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Subjects []Subject `json:"subjects"`
}

type Subject struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}
