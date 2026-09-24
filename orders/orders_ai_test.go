package orders_test

import (
	"errors"
	"reflect"
	"testing"

	"example.com/lesson03/orders"
)

// mockOrderStore замінює сховище: записує виклик і повертає задану помилку.
type mockOrderStore struct {
	calls int
	query string
	args  []any
	err   error
}

func (s *mockOrderStore) Exec(query string, args ...any) error {
	s.calls++
	s.query = query
	s.args = append([]any(nil), args...)
	return s.err
}

var _ orders.OrderStore = (*mockOrderStore)(nil)

func TestPlaceOrderAI_Contract(t *testing.T) {
	dbErr := errors.New("database unavailable")
	for _, tc := range []struct {
		name     string
		orderID  string
		amount   float64
		storeErr error
	}{
		{name: "success", orderID: "order-ai-1", amount: 42.5},
		{name: "original error is returned", orderID: "order-ai-3", amount: 10, storeErr: dbErr},
	} {
		t.Run(tc.name, func(t *testing.T) {
			store := &mockOrderStore{err: tc.storeErr}
			service := orders.NewOrderService(store)
			gotErr := service.PlaceOrder(tc.orderID, tc.amount)
			if gotErr != tc.storeErr {
				t.Fatalf("PlaceOrder() error = %v, want original error %v", gotErr, tc.storeErr)
			}
			if store.calls != 1 {
				t.Fatalf("Exec calls = %d, want 1", store.calls)
			}
			const query = "INSERT INTO orders (id, amount) VALUES (?, ?)"
			if store.query != query {
				t.Errorf("Exec query = %q, want %q", store.query, query)
			}
			if want := []any{tc.orderID, tc.amount}; !reflect.DeepEqual(store.args, want) {
				t.Errorf("Exec args = %#v, want %#v", store.args, want)
			}
		})
	}
}
