package simple

func IntersetCalculation(principal float32, rate int, years float32) float32 {
	var interest float32
	interest = (principal * float32(rate) * years) / 100

	return interest
}
