package card
import (
	"fmt"
	"bank/pkg/bank/types"
	)
	func ExamplePaymentSources() {
		payment:=[]types.Card{
			{
				Balance: 10_000_00,
				PAN: "5058 xxxx xxxx 9999",
				Active: true,
			},
			{
				Balance: 10_000_00,
				PAN: "5058 xxxx xxxx 8888",
				Active: true,
			},
					
		}
		max:=PaymentSources(payment)
		fmt.Println(max[0].Number)
		//Output:
		//5058 xxxx xxxx 9999
	}



