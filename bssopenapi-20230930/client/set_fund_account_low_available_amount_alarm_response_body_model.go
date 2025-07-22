// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFundAccountLowAvailableAmountAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetFundAccountLowAvailableAmountAlarmResponseBody
	GetData() *bool
	SetMetadata(v interface{}) *SetFundAccountLowAvailableAmountAlarmResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *SetFundAccountLowAvailableAmountAlarmResponseBody
	GetRequestId() *string
}

type SetFundAccountLowAvailableAmountAlarmResponseBody struct {
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
	// 08108BF5-1AA3-518E-9986-95A3616E8DA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFundAccountLowAvailableAmountAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetFundAccountLowAvailableAmountAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) SetData(v bool) *SetFundAccountLowAvailableAmountAlarmResponseBody {
	s.Data = &v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) SetMetadata(v interface{}) *SetFundAccountLowAvailableAmountAlarmResponseBody {
	s.Metadata = v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) SetRequestId(v string) *SetFundAccountLowAvailableAmountAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
