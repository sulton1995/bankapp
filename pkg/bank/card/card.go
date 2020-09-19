package payment
import(
	"bank/pkg/bank/types"
	
)


func PaymentSources(cards[] types.Card) []types.PaymentSource{
	payment:= []types.PaymentSource{}
	for _, card := range cards {
		if card.Active ==true {
			if card.Balance > 0{
			 payment = append (payment, types.PaymentSource{ Number:card.PAN, Balance:card.Balance})
			 
				
			}
				
		}
	}
	return payment
}

