package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100
	toDollars := 100
	oneHundredPercent := 100

	bonus := (purchase * percent / oneHundredPercent) / toDollars // вместо нуля ваша формула
	if bonus > limit {
		bonus = limit
	}
	println(bonus) // должно быть 33* (см. объяснение ниже)
}
