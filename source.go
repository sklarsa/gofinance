package gofinance

import (
	"fmt"
	"time"

	"github.com/HT808s/gofinance/models"
)

type Source interface {
	Quote(symbols []string) (models.Quotes, error)

	Hist(symbols []string) (models.HistMap, error)
	HistLimit(symbols []string, start time.Time, end time.Time) (models.HistMap, error)

	DividendHist(symbols []string) (models.DividendHistMap, error)
	DividendHistLimit(symbols []string, start time.Time, end time.Time) (models.DividendHistMap, error)

	fmt.Stringer
}
