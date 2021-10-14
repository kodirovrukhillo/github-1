package payment

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {

	// loop range
	var payment []types.PaymentSource
	for _, card := range cards {

		if card.Balance < 0 || !card.Active {
			continue
		} else {
			for i := 0; i < len(payment); i++ {

				pay := types.PaymentSource{
					PAN:     card.PAN,
					Balance: card.Balance,
				}
				payment = append(payment, pay)
				return payment

			}

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
