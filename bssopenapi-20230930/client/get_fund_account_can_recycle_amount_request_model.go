// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanRecycleAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *GetFundAccountCanRecycleAmountRequest
	GetCurrency() *string
	SetRecycleFromFundAccountId(v string) *GetFundAccountCanRecycleAmountRequest
	GetRecycleFromFundAccountId() *string
}

type GetFundAccountCanRecycleAmountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 122321223
	RecycleFromFundAccountId *string `json:"RecycleFromFundAccountId,omitempty" xml:"RecycleFromFundAccountId,omitempty"`
}

func (s GetFundAccountCanRecycleAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanRecycleAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanRecycleAmountRequest) GetCurrency() *string {
	return s.Currency
}

func (s *GetFundAccountCanRecycleAmountRequest) GetRecycleFromFundAccountId() *string {
	return s.RecycleFromFundAccountId
}

func (s *GetFundAccountCanRecycleAmountRequest) SetCurrency(v string) *GetFundAccountCanRecycleAmountRequest {
	s.Currency = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountRequest) SetRecycleFromFundAccountId(v string) *GetFundAccountCanRecycleAmountRequest {
	s.RecycleFromFundAccountId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountRequest) Validate() error {
	return dara.Validate(s)
}
