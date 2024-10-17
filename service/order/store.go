package order

import (
	"database/sql"

	"github.com/damarteplok/golang-jwt-mysql-test/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	_, err := s.db.Exec("INSERT INTO orders (userId, total, status, address) VALUES (?,?,?,?)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
	}
}
