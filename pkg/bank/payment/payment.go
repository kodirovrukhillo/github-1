package payment

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {

	var payment []types.PaymentSource
	// loop range

	for _, card := range cards {
		if card.Balance < 0 || !card.Active {
			continue
		} else {
		}

	}
	return payment

}

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{

		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: types.TJS,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}
