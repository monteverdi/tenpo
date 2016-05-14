package tenpo

import (
	"fmt"
	"time"
)

// monthOrders is an order array of the months to determine
// timezone offsets for daylight savings evaluation.
var monthOrders = [11]time.Month{
	time.April,
	time.December,
	time.May,
	time.November,
	time.June,
	time.October,
	time.July,
	time.September,
	time.February,
	time.August,
	time.March,
}

// localCache will store whether or not the timezone in
// a given year has daylight saving adjustments
var localCache = make(map[string]bool)

// IsDaylightSavingTZ determines whether or not the given timezone
// in the current year has daylight saving adjustments
func IsTZDaylightSaving(timezone string) (bool, error) {
	now := time.Now()
	return IsTZDaylightSavingInYear(now.Year(), timezone)
}

// IsTZDaylightSavingInYear determines whether or not the given timezone
// in the given year has daylight saving adjustments
func IsTZDaylightSavingInYear(year int, timezone string) (bool, error) {
	cacheKey := fmt.Sprintf("%s/%d", timezone, year)
	if b, ok := localCache[cacheKey]; ok {
		return b, nil
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return false, err
	}
	var isDaylightSaving bool
	defer func() {
		localCache[cacheKey] = isDaylightSaving
	}()
	_, offset := time.Date(year, time.January, 1, 0, 0, 0, 0, location).Zone()
	for _, month := range monthOrders {
		_, monthOffset := time.Date(year, month, 1, 0, 0, 0, 0, location).Zone()
		if monthOffset != offset {
			isDaylightSaving = true
			break
		}
	}
	return isDaylightSaving, nil
}
