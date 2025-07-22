// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFundAccountLowAvailableAmountAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFundAccountId(v int64) *SetFundAccountLowAvailableAmountAlarmRequest
	GetFundAccountId() *int64
	SetThresholdAmount(v string) *SetFundAccountLowAvailableAmountAlarmRequest
	GetThresholdAmount() *string
}

type SetFundAccountLowAvailableAmountAlarmRequest struct {
	// example:
	//
	// 12321213
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 100
	ThresholdAmount *string `json:"ThresholdAmount,omitempty" xml:"ThresholdAmount,omitempty"`
}

func (s SetFundAccountLowAvailableAmountAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s SetFundAccountLowAvailableAmountAlarmRequest) GoString() string {
	return s.String()
}

func (s *SetFundAccountLowAvailableAmountAlarmRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *SetFundAccountLowAvailableAmountAlarmRequest) GetThresholdAmount() *string {
	return s.ThresholdAmount
}

func (s *SetFundAccountLowAvailableAmountAlarmRequest) SetFundAccountId(v int64) *SetFundAccountLowAvailableAmountAlarmRequest {
	s.FundAccountId = &v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmRequest) SetThresholdAmount(v string) *SetFundAccountLowAvailableAmountAlarmRequest {
	s.ThresholdAmount = &v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmRequest) Validate() error {
	return dara.Validate(s)
}
