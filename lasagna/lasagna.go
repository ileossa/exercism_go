package lasagna

const OvenTime = 40

func RemainingOvenTime(timer int) int {
	if timer > OvenTime {
		return 0
	}
	return OvenTime - timer
}

func PreparationTime(layer int) int {
	return layer * 2
}

func ElapsedTime(layer int, cookingTime int) int {
	return PreparationTime(layer) + cookingTime
}
