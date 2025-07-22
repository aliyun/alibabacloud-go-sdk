// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountAvailableAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFundAccountId(v string) *GetFundAccountAvailableAmountRequest
	GetFundAccountId() *string
}

type GetFundAccountAvailableAmountRequest struct {
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountAvailableAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountAvailableAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountRequest) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *GetFundAccountAvailableAmountRequest) SetFundAccountId(v string) *GetFundAccountAvailableAmountRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountAvailableAmountRequest) Validate() error {
	return dara.Validate(s)
}
