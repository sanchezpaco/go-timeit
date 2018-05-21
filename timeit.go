package timeit

// TimeIt function that returns a TimePassed struct given a time in milliseconds.
func TimeIt(timeInMillis float64) TimePassed {
	return parseTime(timeInMillis)
}
