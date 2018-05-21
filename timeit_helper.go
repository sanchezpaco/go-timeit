package timeit

import "math"

const millisHaveASecond = 1000.0
const secondsHaveAMinute = 60.0
const minutesHaveAnHour = 60.0
const hoursHaveADay = 24.0
const daysHaveAWeek = 7.0
const weeksHaveAMonth = 4.34524
const monthsHaveAYear = 12.0

// TimePassed struct with all the time units available.
type TimePassed struct {
	Years   float64
	Months  float64
	Weeks   float64
	Days    float64
	Hours   float64
	Minutes float64
	Seconds float64
}

func parseTime(timeInMillis float64) TimePassed {
	var timePassed = TimePassed{}

	totalTime := timeInMillis / millisHaveASecond
	totalTime = totalTime / secondsHaveAMinute
	totalTime, timePassed = getTimeInUnits(totalTime)

	timePassed = getRawTimePassed(totalTime, timePassed)
	timePassed = balanceTime(timePassed)

	return timePassed
}

func getTimeInUnits(totalTime float64) (float64, TimePassed) {

	var timeInUnits = TimePassed{}

	timeInUnits.Seconds = calculateRawTime(totalTime, secondsHaveAMinute)
	totalTime, timeInUnits.Minutes = calculateInUnits(totalTime, minutesHaveAnHour)
	totalTime, timeInUnits.Hours = calculateInUnits(totalTime, hoursHaveADay)
	totalTime, timeInUnits.Days = calculateInUnits(totalTime, daysHaveAWeek)
	totalTime, timeInUnits.Weeks = calculateInUnits(totalTime, weeksHaveAMonth)
	totalTime, timeInUnits.Months = calculateInUnits(totalTime, monthsHaveAYear)

	return totalTime, timeInUnits
}

func getRawTimePassed(totalTime float64, timeInUnits TimePassed) TimePassed {
	var rawTimePassed = TimePassed{}

	rawTimePassed.Years = math.Floor(totalTime)
	rawTimePassed.Months = timeInUnits.Months + calculateRawTime(totalTime, monthsHaveAYear)
	rawTimePassed.Weeks = timeInUnits.Weeks + calculateRawTime(rawTimePassed.Months, weeksHaveAMonth)
	rawTimePassed.Days = timeInUnits.Days + calculateRawTime(rawTimePassed.Weeks, daysHaveAWeek)
	rawTimePassed.Hours = timeInUnits.Hours + calculateRawTime(rawTimePassed.Days, hoursHaveADay)
	rawTimePassed.Minutes = timeInUnits.Minutes + calculateRawTime(rawTimePassed.Hours, minutesHaveAnHour)
	rawTimePassed.Seconds = timeInUnits.Seconds + calculateRawTime(rawTimePassed.Minutes, secondsHaveAMinute)

	return rawTimePassed
}

func balanceTime(rawTimePassed TimePassed) TimePassed {

	var balancedTime = TimePassed{}

	realSeconds := getRealDate(rawTimePassed.Seconds, secondsHaveAMinute)
	balancedTime.Seconds = realSeconds["time"]

	realMinutes := getRealDate(rawTimePassed.Minutes+realSeconds["timeToAddToNext"], minutesHaveAnHour)
	balancedTime.Minutes = realMinutes["time"]

	realHours := getRealDate(rawTimePassed.Hours+realMinutes["timeToAddToNext"], hoursHaveADay)
	balancedTime.Hours = realHours["time"]

	realDays := getRealDate(rawTimePassed.Days+realHours["timeToAddToNext"], daysHaveAWeek)
	balancedTime.Days = realDays["time"]

	realWeeks := getRealDate(rawTimePassed.Weeks+realDays["timeToAddToNext"], weeksHaveAMonth)
	balancedTime.Weeks = realWeeks["time"]

	realMonths := getRealDate(rawTimePassed.Months+realWeeks["timeToAddToNext"], monthsHaveAYear)
	balancedTime.Months = realMonths["time"]

	balancedTime.Years = rawTimePassed.Years + realMonths["timeToAddToNext"]

	return balancedTime
}

func calculateInUnits(time float64, units float64) (float64, float64) {
	time = math.Floor(time)
	time = time / units
	totalInUnits := calculateRawTime(time, units)

	return time, totalInUnits
}

func getRealDate(time float64, divisor float64) map[string]float64 {
	var realDate = make(map[string]float64)

	realTime := time / divisor
	timeToAddToNext := 0.

	if realTime >= 1 {
		timeToAddToNext = math.Floor(realTime)
		realTime = calculateRawTime(realTime, divisor)
	} else {
		realTime = time
	}

	realDate["time"] = math.Floor(realTime)
	realDate["timeToAddToNext"] = timeToAddToNext

	return realDate
}

func calculateRawTime(x float64, y float64) float64 {
	totalTimeRound := math.Floor(x)

	return (x - totalTimeRound) * y
}
