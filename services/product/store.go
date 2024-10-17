package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/norrico31/rest-api-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*types.Product, 0)
	for rows.Next() {
		product := &types.Product{}

		err := scanRowIntoProduct(rows, product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) GetProductsById(productIds []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productIds)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	// convert productIds to []interface{}
	args := make([]interface{}, len(productIds))
	for i, v := range productIds {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		product := &types.Product{}
		err = scanRowIntoProduct(rows, product)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}

func (s *Store) GetProduct(productId int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", productId)
	if err != nil {
		return nil, err
	}

	product := &types.Product{}
	for rows.Next() {
		err = scanRowIntoProduct(rows, product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (s *Store) CreateProduct(payload types.ProductCreatePayload) (*types.Product, error) {
	res, err := s.db.Exec("INSERT INTO products (name, price, image, description, qty) VALUES (?, ?, ?, ?, ?)", payload.Name, payload.Price, payload.Image, payload.Description, payload.Qty)
	if err != nil {
		return nil, err
	}
	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	product := &types.Product{
		ID:          int(lastInsertedId),
		Name:        payload.Name,
		Price:       payload.Price,
		Image:       payload.Image,
		Description: payload.Description,
		Qty:         payload.Qty,
	}
	return product, nil
}

func scanRowIntoProduct(rows *sql.Rows, product *types.Product) error {
	return rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Qty,
		&product.CreatedAt,
	)
}
