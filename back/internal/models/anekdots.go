package models

type Anekdot struct {
	Id         int
	CategoryId int
	Text       string
}

type AnekdotDto struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Text         string `json:"text"`
}
