package timeit

import "testing"

func TestTimeIT(t *testing.T) {
	expectedTime := TimePassed{
		Years:   0,
		Months:  0,
		Weeks:   0,
		Days:    0,
		Hours:   0,
		Minutes: 10,
		Seconds: 0,
	}

	var timeInMillis = 600000.0

	result := TimeIt(timeInMillis)

	if result != expectedTime {
		t.Error("Expected ", expectedTime, " got ", result)
	}
}
