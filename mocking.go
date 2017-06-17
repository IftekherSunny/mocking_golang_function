package example

////////////////////////////////////////////
// Define payment process type
////////////////////////////////////////////
type PaymentProcess func(amount float32) bool

////////////////////////////////////////////
//  Shopping cart function
////////////////////////////////////////////
func ShoppingCart(amount float32, paymentProcess PaymentProcess) string {
	if paymentProcess(amount) {
		return "Payment has been accepted"
	}

	return "Something wrong!!!"
}
