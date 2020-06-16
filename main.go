package main

func main() {
	purchase := 33_33
	percent := 1
	limit := 100
	bonus := (purchase * percent) / limit

	if bonus < limit {
		println(bonus)
	} else if bonus > limit {
		bonus := limit
		println(bonus)
	}
}
