package main

func main() {
	purchase := 3333
	percent := 1
	limit := 100
	bonus := (purchase * percent) / limit

	if bonus < limit || bonus >= 0 {
		println(bonus)
	} else if bonus > limit {
		bonus := limit
		println(bonus)
	} else {
		println ("No Bonus this time")
	}
}
