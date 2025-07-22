// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountLowAvailableAmountAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmEnabled(v bool) *GetFundAccountLowAvailableAmountAlarmResponseBody
	GetAlarmEnabled() *bool
	SetMetadata(v interface{}) *GetFundAccountLowAvailableAmountAlarmResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *GetFundAccountLowAvailableAmountAlarmResponseBody
	GetRequestId() *string
	SetThresholdAmount(v string) *GetFundAccountLowAvailableAmountAlarmResponseBody
	GetThresholdAmount() *string
}

type GetFundAccountLowAvailableAmountAlarmResponseBody struct {
	AlarmEnabled *bool `json:"AlarmEnabled,omitempty" xml:"AlarmEnabled,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ThresholdAmount *string `json:"ThresholdAmount,omitempty" xml:"ThresholdAmount,omitempty"`
}

func (s GetFundAccountLowAvailableAmountAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountLowAvailableAmountAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) GetAlarmEnabled() *bool {
	return s.AlarmEnabled
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) GetThresholdAmount() *string {
	return s.ThresholdAmount
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) SetAlarmEnabled(v bool) *GetFundAccountLowAvailableAmountAlarmResponseBody {
	s.AlarmEnabled = &v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) SetMetadata(v interface{}) *GetFundAccountLowAvailableAmountAlarmResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) SetRequestId(v string) *GetFundAccountLowAvailableAmountAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) SetThresholdAmount(v string) *GetFundAccountLowAvailableAmountAlarmResponseBody {
	s.ThresholdAmount = &v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
