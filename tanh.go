package main

func tanh(x float64) (result float64) {
	const (
		LeftCutoffPoint  = -10.0
		LeftLimit        = -1.0
		RightCutoffPoint = 10.0
		RightLimit       = 1.0
		Seed             = 1.0
		DepthOfNesting   = 30
	)

	if x < LeftCutoffPoint {
		return LeftLimit
	}

	if x > RightCutoffPoint {
		return RightLimit
	}

	result = Seed

	var xSquare float64 = x * x
	var depth uint8 = DepthOfNesting

	for {
		var isDenominatorZero bool = (result == 0)
	    
		if isDenominatorZero {
			result = Seed
		}
		
		offset := float64(2*depth + 1)
		result = offset + xSquare/result
		
		var reachedTopOfFraction bool = (depth == 0)
		
		if reachedTopOfFraction {
			result = x / result
			return
		}
		depth--
	}
	return
}
