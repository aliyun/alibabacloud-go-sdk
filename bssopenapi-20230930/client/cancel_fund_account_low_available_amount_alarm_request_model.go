// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFundAccountLowAvailableAmountAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFundAccountId(v int64) *CancelFundAccountLowAvailableAmountAlarmRequest
	GetFundAccountId() *int64
}

type CancelFundAccountLowAvailableAmountAlarmRequest struct {
	// example:
	//
	// 123321123
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s CancelFundAccountLowAvailableAmountAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelFundAccountLowAvailableAmountAlarmRequest) GoString() string {
	return s.String()
}

func (s *CancelFundAccountLowAvailableAmountAlarmRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *CancelFundAccountLowAvailableAmountAlarmRequest) SetFundAccountId(v int64) *CancelFundAccountLowAvailableAmountAlarmRequest {
	s.FundAccountId = &v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmRequest) Validate() error {
	return dara.Validate(s)
}
