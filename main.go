package main

func main() {
	var (
		purchase float64   = 33_33.0
		percent  float64   = 1.0
		limit    float64   = 100.0
		bonus    float64 = (purchase * percent) / limit
	)

	if bonus < limit {
		println(bonus)
	} else if bonus > limit {
		bonus := limit
		println(bonus)
	}
}
