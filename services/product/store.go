package product

import (
	"database/sql"

	"github.com/norrico31/rest-api-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	return nil, nil
}

func (s *Store) GetProduct(productId int) (types.Product, error) {
	return nil, nil
}
