package temporal_test

import (
	"testing"
	"time"

	"github.com/maniartech/temporal"
	"github.com/stretchr/testify/assert"
)

func TestYear(t *testing.T) {
	// YearStart
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate := temporal.YearStart(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	now := time.Now()
	expectedDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	functionDate = temporal.YearStart()
	assert.Equal(t, expectedDate, functionDate)

	// YearEnd
	expectedDate = time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC)
	functionDate = temporal.YearEnd(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	expectedDate = time.Date(now.Year(), 12, 31, 23, 59, 59, 999999999, time.Local)
	functionDate = temporal.YearEnd()
	assert.Equal(t, expectedDate, functionDate)

	// LastYear
	expectedDate = time.Now().AddDate(-1, 0, 0)
	functionDate = temporal.LastYear()

	assert.Equal(t, expectedDate, functionDate)

	// NextYear
	expectedDate = time.Now().AddDate(1, 0, 0)
	functionDate = temporal.NextYear()

	assert.Equal(t, expectedDate, functionDate)

	// Years
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, year := range testRange {
		expectedDate := fixedDate.AddDate(year, 0, 0)
		functionDate := temporal.Years(year, fixedDate)

		assert.Equal(t, expectedDate, functionDate)
	}

	expectedDate = truccateSecond(now.AddDate(1, 0, 0))
	functionDate = truccateSecond(temporal.Years(1))
	assert.Equal(t, expectedDate, functionDate)

	assert.Panics(t, func() { temporal.Years(0, fixedDate) })
}

func TestMonth(t *testing.T) {
	// MonthStart - write a test that loops through all the months and checks that the first day of the month is correct
	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	totalDaysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for _, month := range months {
		expectedDate := time.Date(2023, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		functionDate := temporal.MonthStart(time.Date(2023, time.Month(month), 1, 11, 2, 10, 0, time.UTC))

		assert.Equal(t, expectedDate, functionDate)
	}
	now := time.Now()
	expectedDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	functionDate := temporal.MonthStart()
	assert.Equal(t, expectedDate, functionDate)
	// MonthEnd
	for _, month := range months {
		expectedDate := time.Date(2023, time.Month(month), totalDaysInMonth[month-1], 23, 59, 59, 999999999, time.UTC)
		functionDate := temporal.MonthEnd(time.Date(2023, time.Month(month), 1, 11, 22, 10, 0, time.UTC))

		assert.Equal(t, expectedDate, functionDate)
	}
	expectedDate = time.Date(now.Year(), now.Month(), totalDaysInMonth[now.Month()-1], 23, 59, 59, 999999999, time.Local)
	functionDate = temporal.MonthEnd()
	assert.Equal(t, expectedDate, functionDate)

	// LastMonth
	expectedDate = time.Now().AddDate(0, -1, 0)
	functionDate = temporal.LastMonth()

	assert.Equal(t, expectedDate, functionDate)

	// NextMonth
	expectedDate = time.Now().AddDate(0, 1, 0)
	functionDate = temporal.NextMonth()

	assert.Equal(t, expectedDate, functionDate)

	// Months
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, month := range testRange {
		expectedDate := fixedDate.AddDate(0, month, 0)
		functionDate := temporal.Months(month, fixedDate)

		assert.Equal(t, expectedDate, functionDate)
	}

	expectedDate = truccateSecond(now.AddDate(0, 1, 0))
	functionDate = truccateSecond(temporal.Months(1))
	assert.Equal(t, expectedDate, functionDate)

	assert.Panics(t, func() { temporal.Months(0, fixedDate) })

}

func truccateSecond(t time.Time) time.Time {
	return t.Truncate(time.Second)
}

func TestWeek(t *testing.T) {
	// WeekStart
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate := temporal.WeekStart(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	now := time.Now()
	expectedDate = truccateSecond(time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday()), 0, 0, 0, 0, time.Local))
	functionDate = truccateSecond(temporal.WeekStart())
	assert.Equal(t, expectedDate, functionDate)

	// WeekEnd
	expectedDate = time.Date(2023, 1, 7, 23, 59, 59, 999999999, time.UTC)
	functionDate = temporal.WeekEnd(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	expectedDate = time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday())+6, 23, 59, 59, 999999999, time.Local)
	functionDate = temporal.WeekEnd()
	assert.Equal(t, expectedDate, functionDate)

	// LastWeek
	expectedDate = time.Now().AddDate(0, 0, -7)
	functionDate = temporal.LastWeek()

	assert.Equal(t, expectedDate, functionDate)

	// NextWeek
	expectedDate = time.Now().AddDate(0, 0, 7)
	functionDate = temporal.NextWeek()

	assert.Equal(t, expectedDate, functionDate)

	// Weeks
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, week := range testRange {
		expectedDate := fixedDate.AddDate(0, 0, week*7)
		functionDate := temporal.Weeks(week, fixedDate)

		assert.Equal(t, expectedDate, functionDate)
	}

	expectedDate = truccateSecond(now.AddDate(0, 0, 7))
	functionDate = truccateSecond(temporal.Weeks(1))

	assert.Equal(t, expectedDate, functionDate)

	assert.Panics(t, func() { temporal.Weeks(0, fixedDate) })

	// WeekStartOn
	expectedDate = time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	functionDate = temporal.WeekStartOn(time.Sunday, time.Date(2024, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	expectedDate = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate = temporal.WeekStartOn(time.Monday, time.Date(2024, 1, 2, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	expectedDate = time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	functionDate = temporal.WeekStartOn(time.Tuesday, time.Date(2024, 1, 3, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	now = time.Now()

	expectedDate = truccateSecond(time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday())+1, 0, 0, 0, 0, time.Local))
	functionDate = truccateSecond(temporal.WeekStartOn(time.Monday))

	assert.Equal(t, expectedDate, functionDate)

	// WeekEndOn
	expectedDate = time.Date(2024, 1, 6, 23, 59, 59, 999999999, time.UTC)
	functionDate = temporal.WeekEndOn(time.Sunday, time.Date(2024, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	expectedDate = time.Date(2024, 1, 7, 23, 59, 59, 999999999, time.UTC)
	functionDate = temporal.WeekEndOn(time.Monday, time.Date(2024, 1, 2, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	now = time.Now()

	expectedDate = time.Date(now.Year(), now.Month(), now.Day()-int(now.Weekday())+7, 23, 59, 59, 999999999, time.Local)
	functionDate = temporal.WeekEndOn(time.Monday)

	assert.Equal(t, expectedDate, functionDate)
}

func TestDay(t *testing.T) {
	// DayStart
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	functionDate := temporal.DayStart(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	now := time.Now()
	expectedDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	functionDate = temporal.DayStart()
	assert.Equal(t, expectedDate, functionDate)

	// DayEnd
	expectedDate = time.Date(2023, 1, 1, 23, 59, 59, 999999999, time.UTC)
	functionDate = temporal.DayEnd(time.Date(2023, 1, 1, 11, 2, 10, 0, time.UTC))

	assert.Equal(t, expectedDate, functionDate)

	expectedDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, time.Local)
	functionDate = temporal.DayEnd()
	assert.Equal(t, expectedDate, functionDate)

	// Yesterday
	expectedDate = time.Now().AddDate(0, 0, -1)
	functionDate = temporal.Yesterday()

	assert.Equal(t, expectedDate, functionDate)

	// Tomorrow
	expectedDate = time.Now().AddDate(0, 0, 1)
	functionDate = temporal.Tomorrow()

	assert.Equal(t, expectedDate, functionDate)

	// Days
	testRange := []int{-3, -2, -1, 1, 2, 3}
	fixedDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, day := range testRange {
		expectedDate := fixedDate.AddDate(0, 0, day)
		functionDate := temporal.Days(day, fixedDate)

		assert.Equal(t, expectedDate, functionDate)
	}

	expectedDate = truccateSecond(now.AddDate(0, 0, 1))
	functionDate = truccateSecond(temporal.Days(1))

	assert.Equal(t, expectedDate, functionDate)

	assert.Panics(t, func() { temporal.Days(0, fixedDate) })

}
