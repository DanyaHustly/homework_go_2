package order

import "fmt"

// Блок order — объединяющая задача.
// Здесь нужно закрепить:
// constants/iota, switch, int conversion, bool,
// вычисление цены в копейках и сборку итоговой строки.

// Статусы заказа.
//
// TODO: задайте StatusNew=0, StatusPaid=1 и StatusCanceled=2 как последовательные значения в указанном порядке.
const (
	StatusNew = iota
	StatusPaid
	StatusCanceled
)

// OrderSummary собирает краткое описание заказа.
//
// TODO: верните строку "status=<status_text> payment=<payment_text> price_kop=<price>". Статусы: new, paid, canceled, неизвестный — unknown; paid=true даёт paid, иначе not_paid; отрицательную цену считайте 0, остальные рубли переведите в копейки.
func OrderSummary(status int, priceRub int, paid bool) string {
	var statusText string

	switch status {
	case StatusNew:
		statusText = "new"

	case StatusPaid:
		statusText = "paid"

	case StatusCanceled:
		statusText = "canceled"

	default:
		statusText = "unknown"
	}

	paymentText := "not_paid"

	if paid {
		paymentText = "paid"
	}

	if priceRub < 0 {
		priceRub = 0
	}

	priceKop := priceRub * 100

	return fmt.Sprintf(
		"status=%s payment=%s price_kop=%d",
		statusText,
		paymentText,
		priceKop,
	)
}
