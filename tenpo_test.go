package tenpo

import (
	"fmt"
	"testing"
)

func TestIsTZDaylightSaving(t *testing.T) {
	tz := "Asia/Hong_Kong"
	// we know that Hong_Kong does not have daylight savings time
	dls, err := IsTZDaylightSaving(tz)
	if err != nil || dls {
		t.Fatalf("Timezone %s should not be daylight saving", tz)
	}
}

func TestIsTZDaylightSaving_badTZ(t *testing.T) {
	tz := "Dummy"
	_, err := IsTZDaylightSaving(tz)
	if err == nil {
		t.Fatalf("Expected an Error with bad Timezone %s", tz)
	}
}

func TestIsTZDaylightSavingInYear(t *testing.T) {
	tz := "America/New_York"
	year := 2000
	dls, err := IsTZDaylightSavingInYear(year, tz)
	if err != nil || !dls {
		t.Fatalf("Expected Timezone %s in Year %d to be a Daylight Saving",
			tz, year)
	}
	// now, before the US law is passed
	year = 1900
	dls, err = IsTZDaylightSavingInYear(year, tz)
	if err != nil || dls {
		t.Fatalf("Expected Timezone %s in Year %s is not a Daylight Saving",
			tz, year)
	}
}

func TestIsTZDaylightSaving_checkCache(t *testing.T) {
	tz := "Asia/Taipei"
	year := 2016
	dls, err := IsTZDaylightSavingInYear(year, tz)
	if err != nil || dls {
		t.Fatalf("Timezone %s should not be daylight saving", tz)
	}
	if b, ok := localCache[fmt.Sprintf("%s/%d", tz, year)]; !ok || b {
		t.Fatalf("Expected Timezone %s in Year %d to be in the local cache %v",
			tz, year, localCache)
	}
	dls, err = IsTZDaylightSavingInYear(year, tz) // should get this from cache
	if err != nil || dls {
		t.Fatalf("Timezone %s should not be daylight saving", tz)
	}
}

func ExampleIsTZDaylightSaving() {
	tz := "Asia/Hong_Kong"
	dls, err := IsTZDaylightSaving(tz)
	if err != nil {
		// handle the error where the timezone is not valid
		fmt.Errorf("Timezone %s is not a valid timezone", tz)
	}
	fmt.Printf("%t", dls)
	// Output: false
}

func ExampleIsTZDaylightSavingInYear() {
	tz := "America/Los_Angeles"
	year := 1970
	dls1, err := IsTZDaylightSavingInYear(year, tz)
	year = 1900
	dls2, _ := IsTZDaylightSavingInYear(year, tz) // this is before the law passed
	fmt.Printf("%t\n", dls1)
	fmt.Printf("%t\n", dls2)
	// Output:
	// true
	// false
}
