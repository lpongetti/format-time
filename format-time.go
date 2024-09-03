package formattime

import (
	"strings"
	"time"
)

func Format(mask string, value time.Time) string {
	// https://day.js.org/docs/en/display/format
	layout := mask

	layout = strings.ReplaceAll(layout, "YYYY", "2006")
	layout = strings.ReplaceAll(layout, "YY", "06")
	layout = strings.ReplaceAll(layout, "HH", "15")
	layout = strings.ReplaceAll(layout, "hh", "03")
	layout = strings.ReplaceAll(layout, "h", "3")
	layout = strings.ReplaceAll(layout, "mm", "04")
	layout = strings.ReplaceAll(layout, "m", "4")
	layout = strings.ReplaceAll(layout, "ss", "05")
	layout = strings.ReplaceAll(layout, "s", "5")
	layout = strings.ReplaceAll(layout, "SSS", "000")
	layout = strings.ReplaceAll(layout, "ZZ", "-0700")
	layout = strings.ReplaceAll(layout, "Z", "-07:00")
	layout = strings.ReplaceAll(layout, "MMMM", "January")
	layout = strings.ReplaceAll(layout, "MMM", "Jan")
	layout = strings.ReplaceAll(layout, "MM", "01")
	layout = strings.ReplaceAll(layout, "M", "1")
	layout = strings.ReplaceAll(layout, "DDDD", "Monday")
	layout = strings.ReplaceAll(layout, "DDD", "Mon")
	layout = strings.ReplaceAll(layout, "DD", "02")
	layout = strings.ReplaceAll(layout, "D", "2")
	layout = strings.ReplaceAll(layout, "AA", "PM")
	layout = strings.ReplaceAll(layout, "aa", "pm")

	return value.Format(layout)
}
