// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFundAccountLowAvailableAmountAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CancelFundAccountLowAvailableAmountAlarmResponseBody
	GetData() *bool
	SetMetadata(v interface{}) *CancelFundAccountLowAvailableAmountAlarmResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *CancelFundAccountLowAvailableAmountAlarmResponseBody
	GetRequestId() *string
}

type CancelFundAccountLowAvailableAmountAlarmResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// DFC1F7F9-3BA9-BA4D-2F2E653
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelFundAccountLowAvailableAmountAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelFundAccountLowAvailableAmountAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) SetData(v bool) *CancelFundAccountLowAvailableAmountAlarmResponseBody {
	s.Data = &v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) SetMetadata(v interface{}) *CancelFundAccountLowAvailableAmountAlarmResponseBody {
	s.Metadata = v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) SetRequestId(v string) *CancelFundAccountLowAvailableAmountAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
