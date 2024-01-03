package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Product struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	Value       int16
	Quantity    int16
}

func NewProduct(db *sql.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Create(name string, description string, value int16, quantity int16) (Product, error) {
	id := uuid.New().String()
	_, err := p.db.Exec("INSERT INTO product (id, name, description) VALUES ($1, $2, $3)", id, name, description, value, quantity)
	if err != nil {
		return Product{}, err
	}
	return Product{
		ID:          id,
		Name:        name,
		Description: description,
		Value:       value,
		Quantity:    quantity,
	}
}
