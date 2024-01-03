// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewProduct struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Value       int     `json:"value"`
	Quanity     int     `json:"quanity"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Order struct {
	ID      string   `json:"id"`
	User    *User    `json:"user"`
	Product *Product `json:"product"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Value       int     `json:"value"`
	Quanity     int     `json:"quanity"`
}

type Query struct {
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}