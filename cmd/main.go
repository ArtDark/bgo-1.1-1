package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100
	const toDollars = 100
	const oneHundredPercent = 100

	bonus := (purchase * percent / oneHundredPercent) / toDollars
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
