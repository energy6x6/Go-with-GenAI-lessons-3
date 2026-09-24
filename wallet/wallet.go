// Package wallet реалізує Розділ 3, Крок 1 домашньої роботи:
// зміну оригінальних елементів зрізу через приймачі-вказівники.
package wallet

// SecureWallet зберігає баланс; його методи узгоджено використовують вказівник.
type SecureWallet struct {
	balance float64
}

// Balance повертає поточний баланс без копіювання гаманця.
func (w *SecureWallet) Balance() float64 {
	return w.balance
}

func (w *SecureWallet) Deposit(amt float64) {
	w.balance += amt
}

// ApplyDeposits додає amt до кожного гаманця у зрізі wallets.
//
// Це — саме той сценарій із заняття: wallets це []SecureWallet
// (зріз ЗНАЧЕНЬ, а не вказівників). Наївний `for _, w := range
// wallets { w.Deposit(amt) }` НЕ подіє, бо w — копія елемента
// циклу.
//
// Індексний цикл звертається до справжнього елемента зрізу.
func ApplyDeposits(wallets []SecureWallet, amt float64) {
	for i := range wallets {
		wallets[i].Deposit(amt)
	}
}
