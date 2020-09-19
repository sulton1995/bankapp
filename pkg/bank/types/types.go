package types

	type Currency string
	const (
		TJS Currency = "TJS"
		RUB Currency = "RUB"
		USD Currency = "USD"
	)
	type Money int

	type Payment struct {
		ID int
		Amount Money
	}
	type Card struct{
		ID int
		PAN string
		Balance Money
		Currency Currency
		Color string
		Name string
		Active bool
		MinBalance Money
	}
	type PaymentSource struct{
		Type string //card
		Number string //number of type "5058 xxxx xxxx 8888"
		Balance Money
		
	}
	