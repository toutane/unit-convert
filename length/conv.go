package length

func MToF(m Meter) Foot {
	return Foot(m * 3.281)
}

func FToM(ft Foot) Meter {
	return Meter(ft / 3.281)
}
