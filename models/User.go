package models

type User struct {
	ID        int
	UserName  string
	FirstName string
	LastName  string
	Interests []Interest
}
