package utils

func isLeftBracket(r rune) bool {
	if r == '(' || r == '{' || r == '[' {
		return true
	}

	return false
}

func isRightBracket(r rune) bool {
	if r == ')' || r == '}' || r == ']' {
		return true
	}

	return false
}

func isBarcketPairMatch(leftBracket, rightBracket rune) bool {
	if leftBracket == '(' && rightBracket == ')' {
		return true
	}

	if leftBracket == '{' && rightBracket == '}' {
		return true
	}
	if leftBracket == '[' && rightBracket == ']' {
		return true
	}

	return false
}

//BracketPairValidation is used to check if bracket pair in string match or not
//Assume string should have at least on pair of brackets, for example, "" and "123" return false
func BracketPairValidation(s string) bool {
	length := len(s)

	if length == 0 {
		return false
	}

	if length == 1 {
		return false
	}

	var leftBracketStack []rune

	//Avoid too much slice auto expanding when dealing with long string, the 64, 8, 256 and 32 just are estimating number
	//Assume the average lenth of content enclosed by bracket is around 6, plus two brackets is 8
	var initStackLength int
	if length > 64 {
		initStackLength = length/8 + 1
		if initStackLength > 256 {
			initStackLength = 256
		} else if initStackLength < 32 {
			initStackLength = 32
		}
	} else {
		initStackLength = length/2 + 1
	}

	leftBracketStack = make([]rune, initStackLength)

	leftStackPoint := -1

	var leftBracketCount, rightBracketCount int

	for _, item := range s {
		if isLeftBracket(item) {
			leftBracketCount++
			leftStackPoint++
			if leftStackPoint == len(leftBracketStack) {
				leftBracketStack = append(leftBracketStack, item)
			} else {
				leftBracketStack[leftStackPoint] = item
			}
		} else {
			if isRightBracket(item) {
				rightBracketCount++
				if leftStackPoint == -1 {
					return false
				}
				if isBarcketPairMatch(leftBracketStack[leftStackPoint], item) {
					leftStackPoint--
				} else {
					return false
				}
			}
		}
	}

	if leftBracketCount == 0 || rightBracketCount == 0 {
		return false
	}

	if leftStackPoint == -1 {
		return true
	}

	return false
}
