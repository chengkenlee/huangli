package run

import (
	"time"
)

var (
	year  string
	month string
	day   string
)

func newReplaceVar(date string) (string, string) {
	ts, _ := time.Parse("2006-01-02", date)
	date = ts.Format("2006-1-2")
	year = ts.Format("2006")
	return year, date
}
