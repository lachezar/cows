package main

func permutations(state []uint8, rank int, candidates []candidate) []candidate {
	if rank < 0 {
		candidates = append(candidates, candidate{state[3], state[2], state[1], state[0]})
	} else {
		var digit uint8 = 0
		if rank == len(state)-1 {
			digit = 1
		}
		for digit < 10 {
			var shouldSkip = false
			for i := len(state) - 1; i > rank; i-- {
				if state[i] == digit {
					shouldSkip = true
					break
				}
			}
			if !shouldSkip {
				state[rank] = digit
				candidates = permutations(state, rank-1, candidates)
			}
			digit += 1
		}
	}
	return candidates
}

func filterNonMatchingCandidates(candidates []candidate, query candidate, answer answer) []candidate {
	newCandidates := make([]candidate, 0, cap(candidates))
	for _, c := range candidates {
		if c.Compare(query) == answer {
			newCandidates = append(newCandidates, c)
		}
	}
	return newCandidates
}

func (self candidate) IsValid() bool {
	if self[0] == 0 {
		return false
	}
	for i := 0; i < len(self)-1; i++ {
		for j := i + 1; j < len(self); j++ {
			if self[i] == self[j] {
				return false
			}
		}
	}
	return true
}

func (self candidate) Compare(other candidate) answer {
	answer := answer{}
	for i, selfDigit := range self {
		for j, otherDigit := range other {
			if selfDigit == otherDigit {
				if i == j {
					answer.bulls++
				} else {
					answer.cows++
				}
			}
		}
	}
	return answer
}

func (self candidate) Show() string {
	for i := 0; i < len(self); i++ {
		self[i] += 0x30
	}
	return string(self[:])
}
