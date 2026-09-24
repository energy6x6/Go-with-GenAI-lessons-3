// Package orders створює замовлення через передане сховище.
package orders

// OrderStore — мінімальний інтерфейс залежності, потрібної OrderService.
type OrderStore interface {
	Exec(query string, args ...any) error
}

// OrderService залежить від інтерфейсу сховища, а не від конкретної БД.
type OrderService struct {
	store OrderStore
}

// NewOrderService — конструктор із впровадженням залежності.
func NewOrderService(store OrderStore) *OrderService {
	return &OrderService{store: store}
}

// PlaceOrder передає параметри замовлення сховищу й повертає його помилку.
func (s *OrderService) PlaceOrder(orderID string, amount float64) error {
	return s.store.Exec("INSERT INTO orders (id, amount) VALUES (?, ?)", orderID, amount)
}
