// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoney interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *Money
	GetAmount() *int64
	SetAmountAsString(v string) *Money
	GetAmountAsString() *string
	SetAmountString(v string) *Money
	GetAmountString() *string
	SetCent(v int64) *Money
	GetCent() *int64
	SetCurrency(v *MoneyCurrency) *Money
	GetCurrency() *MoneyCurrency
	SetCurrencyCode(v string) *Money
	GetCurrencyCode() *string
	SetPositive(v bool) *Money
	GetPositive() *bool
}

type Money struct {
	// amount
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// amountAsString
	AmountAsString *string `json:"amountAsString,omitempty" xml:"amountAsString,omitempty"`
	// amountString
	AmountString *string `json:"amountString,omitempty" xml:"amountString,omitempty"`
	// cent
	Cent *int64 `json:"cent,omitempty" xml:"cent,omitempty"`
	// currency
	Currency *MoneyCurrency `json:"currency,omitempty" xml:"currency,omitempty" type:"Struct"`
	// currencyCode
	CurrencyCode *string `json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	// positive
	Positive *bool `json:"positive,omitempty" xml:"positive,omitempty"`
}

func (s Money) String() string {
	return dara.Prettify(s)
}

func (s Money) GoString() string {
	return s.String()
}

func (s *Money) GetAmount() *int64 {
	return s.Amount
}

func (s *Money) GetAmountAsString() *string {
	return s.AmountAsString
}

func (s *Money) GetAmountString() *string {
	return s.AmountString
}

func (s *Money) GetCent() *int64 {
	return s.Cent
}

func (s *Money) GetCurrency() *MoneyCurrency {
	return s.Currency
}

func (s *Money) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *Money) GetPositive() *bool {
	return s.Positive
}

func (s *Money) SetAmount(v int64) *Money {
	s.Amount = &v
	return s
}

func (s *Money) SetAmountAsString(v string) *Money {
	s.AmountAsString = &v
	return s
}

func (s *Money) SetAmountString(v string) *Money {
	s.AmountString = &v
	return s
}

func (s *Money) SetCent(v int64) *Money {
	s.Cent = &v
	return s
}

func (s *Money) SetCurrency(v *MoneyCurrency) *Money {
	s.Currency = v
	return s
}

func (s *Money) SetCurrencyCode(v string) *Money {
	s.CurrencyCode = &v
	return s
}

func (s *Money) SetPositive(v bool) *Money {
	s.Positive = &v
	return s
}

func (s *Money) Validate() error {
	if s.Currency != nil {
		if err := s.Currency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MoneyCurrency struct {
	// currencyCode
	CurrencyCode *string `json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	// defaultFractionDigits
	DefaultFractionDigits *int32 `json:"defaultFractionDigits,omitempty" xml:"defaultFractionDigits,omitempty"`
	// displayName
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// numericCode
	NumericCode *int32 `json:"numericCode,omitempty" xml:"numericCode,omitempty"`
	// symbol
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
}

func (s MoneyCurrency) String() string {
	return dara.Prettify(s)
}

func (s MoneyCurrency) GoString() string {
	return s.String()
}

func (s *MoneyCurrency) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *MoneyCurrency) GetDefaultFractionDigits() *int32 {
	return s.DefaultFractionDigits
}

func (s *MoneyCurrency) GetDisplayName() *string {
	return s.DisplayName
}

func (s *MoneyCurrency) GetNumericCode() *int32 {
	return s.NumericCode
}

func (s *MoneyCurrency) GetSymbol() *string {
	return s.Symbol
}

func (s *MoneyCurrency) SetCurrencyCode(v string) *MoneyCurrency {
	s.CurrencyCode = &v
	return s
}

func (s *MoneyCurrency) SetDefaultFractionDigits(v int32) *MoneyCurrency {
	s.DefaultFractionDigits = &v
	return s
}

func (s *MoneyCurrency) SetDisplayName(v string) *MoneyCurrency {
	s.DisplayName = &v
	return s
}

func (s *MoneyCurrency) SetNumericCode(v int32) *MoneyCurrency {
	s.NumericCode = &v
	return s
}

func (s *MoneyCurrency) SetSymbol(v string) *MoneyCurrency {
	s.Symbol = &v
	return s
}

func (s *MoneyCurrency) Validate() error {
	return dara.Validate(s)
}
