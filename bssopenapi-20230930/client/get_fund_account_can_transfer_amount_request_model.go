// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanTransferAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *GetFundAccountCanTransferAmountRequest
	GetCurrency() *string
	SetFundAccountId(v string) *GetFundAccountCanTransferAmountRequest
	GetFundAccountId() *string
}

type GetFundAccountCanTransferAmountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 123212
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountCanTransferAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanTransferAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanTransferAmountRequest) GetCurrency() *string {
	return s.Currency
}

func (s *GetFundAccountCanTransferAmountRequest) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *GetFundAccountCanTransferAmountRequest) SetCurrency(v string) *GetFundAccountCanTransferAmountRequest {
	s.Currency = &v
	return s
}

func (s *GetFundAccountCanTransferAmountRequest) SetFundAccountId(v string) *GetFundAccountCanTransferAmountRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountCanTransferAmountRequest) Validate() error {
	return dara.Validate(s)
}
