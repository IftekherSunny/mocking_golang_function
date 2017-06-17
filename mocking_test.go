package example

import "testing"

////////////////////////////////////////////
// Mocking payment process function
////////////////////////////////////////////
func mockPaymentProcess(amount float32) bool {
	if amount == 100 {
		return true
	}

	return false
}

func TestShoppingCart(t *testing.T) {
	t.Run("Testing shopping cart feature", func(t *testing.T) {

		if ShoppingCart(100, mockPaymentProcess) != "Payment has been accepted" {
			t.Error("Message must be 'Payment has been accepted'")
		}

		if ShoppingCart(0, mockPaymentProcess) != "Something wrong!!!" {
			t.Error("Message must be 'Something wrong!!!'")
		}
	})
}
