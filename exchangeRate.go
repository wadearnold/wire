// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// ExchangeRate is the ExchangeRate of the wire
type ExchangeRate struct {
	// tag
	tag string
	// ExchangeRate is the exchange rate
	// Must contain at least one numeric character and only one decimal comma marker (e.g., an exchange rate of 1.2345 should be entered as 1,2345).
	ExchangeRate string `json:"exchangeRate,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewExchangeRate returns a new ExchangeRate
func NewExchangeRate() *ExchangeRate {
	eRate := &ExchangeRate{
		tag: TagExchangeRate,
	}
	return eRate
}

// Parse takes the input string and parses the ExchangeRate values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (eRate *ExchangeRate) Parse(record string) error {
	if utf8.RuneCountInString(record) != 18 {
		return NewTagWrongLengthErr(18, len(record))
	}
	eRate.tag = record[:6]
	eRate.ExchangeRate = eRate.parseStringField(record[6:18])
	return nil
}

// String writes ExchangeRate
func (eRate *ExchangeRate) String() string {
	var buf strings.Builder
	buf.Grow(18)
	buf.WriteString(eRate.tag)
	buf.WriteString(eRate.ExchangeRateField())
	return buf.String()
}

// Validate performs WIRE format rule checks on ExchangeRate and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (eRate *ExchangeRate) Validate() error {
	if eRate.tag != TagExchangeRate {
		return fieldError("tag", ErrValidTagForType, eRate.tag)
	}
	if err := eRate.isAmount(eRate.ExchangeRate); err != nil {
		return fieldError("ExchangeRate", err, eRate.ExchangeRate)
	}
	return nil
}

// ExchangeRateField gets a string of the ExchangeRate field
func (eRate *ExchangeRate) ExchangeRateField() string {
	return eRate.alphaField(eRate.ExchangeRate, 12)
}
