// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanWithdrawAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFundAccountId(v int64) *GetFundAccountCanWithdrawAmountRequest
	GetFundAccountId() *int64
}

type GetFundAccountCanWithdrawAmountRequest struct {
	// example:
	//
	// 123212232
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountCanWithdrawAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanWithdrawAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanWithdrawAmountRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountCanWithdrawAmountRequest) SetFundAccountId(v int64) *GetFundAccountCanWithdrawAmountRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountRequest) Validate() error {
	return dara.Validate(s)
}
