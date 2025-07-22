// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountLowAvailableAmountAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFundAccountId(v int64) *GetFundAccountLowAvailableAmountAlarmRequest
	GetFundAccountId() *int64
}

type GetFundAccountLowAvailableAmountAlarmRequest struct {
	// example:
	//
	// 12332112
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountLowAvailableAmountAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountLowAvailableAmountAlarmRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountLowAvailableAmountAlarmRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountLowAvailableAmountAlarmRequest) SetFundAccountId(v int64) *GetFundAccountLowAvailableAmountAlarmRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmRequest) Validate() error {
	return dara.Validate(s)
}
