// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanAllocateCreditAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFundAccountId(v int64) *GetFundAccountCanAllocateCreditAmountRequest
	GetFundAccountId() *int64
}

type GetFundAccountCanAllocateCreditAmountRequest struct {
	// example:
	//
	// 1233231
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountCanAllocateCreditAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanAllocateCreditAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanAllocateCreditAmountRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountCanAllocateCreditAmountRequest) SetFundAccountId(v int64) *GetFundAccountCanAllocateCreditAmountRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountRequest) Validate() error {
	return dara.Validate(s)
}
