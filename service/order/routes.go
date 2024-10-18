package order

import (
	"net/http"

	"github.com/damarteplok/golang-jwt-mysql-test/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
}
