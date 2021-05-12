package weight

func KToP(k Kilogram) Pound {
	return Pound(k * 0.454)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p / 0.454)
}
