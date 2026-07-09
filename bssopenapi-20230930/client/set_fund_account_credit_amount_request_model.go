// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFundAccountCreditAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreditAmount(v string) *SetFundAccountCreditAmountRequest
	GetCreditAmount() *string
	SetCurrency(v string) *SetFundAccountCreditAmountRequest
	GetCurrency() *string
	SetFundAccountId(v int64) *SetFundAccountCreditAmountRequest
	GetFundAccountId() *int64
}

type SetFundAccountCreditAmountRequest struct {
	// Credit limit
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	CreditAmount *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// Currency for the credit control limit. Currently, only CNY is supported in mainland China, and only USD is supported for international use.
	//
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// Fund account ID. If not specified, the account owned by the current account (owner) is used by default.
	//
	// example:
	//
	// 1232312
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s SetFundAccountCreditAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s SetFundAccountCreditAmountRequest) GoString() string {
	return s.String()
}

func (s *SetFundAccountCreditAmountRequest) GetCreditAmount() *string {
	return s.CreditAmount
}

func (s *SetFundAccountCreditAmountRequest) GetCurrency() *string {
	return s.Currency
}

func (s *SetFundAccountCreditAmountRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *SetFundAccountCreditAmountRequest) SetCreditAmount(v string) *SetFundAccountCreditAmountRequest {
	s.CreditAmount = &v
	return s
}

func (s *SetFundAccountCreditAmountRequest) SetCurrency(v string) *SetFundAccountCreditAmountRequest {
	s.Currency = &v
	return s
}

func (s *SetFundAccountCreditAmountRequest) SetFundAccountId(v int64) *SetFundAccountCreditAmountRequest {
	s.FundAccountId = &v
	return s
}

func (s *SetFundAccountCreditAmountRequest) Validate() error {
	return dara.Validate(s)
}
