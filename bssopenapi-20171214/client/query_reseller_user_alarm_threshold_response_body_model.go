// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResellerUserAlarmThresholdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryResellerUserAlarmThresholdResponseBody
	GetCode() *string
	SetCount(v int32) *QueryResellerUserAlarmThresholdResponseBody
	GetCount() *int32
	SetData(v []*QueryResellerUserAlarmThresholdResponseBodyData) *QueryResellerUserAlarmThresholdResponseBody
	GetData() []*QueryResellerUserAlarmThresholdResponseBodyData
	SetMessage(v string) *QueryResellerUserAlarmThresholdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryResellerUserAlarmThresholdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryResellerUserAlarmThresholdResponseBody
	GetSuccess() *bool
}

type QueryResellerUserAlarmThresholdResponseBody struct {
	// The status code of the request. If the call fails, an error code is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The data returned.
	Data []*QueryResellerUserAlarmThresholdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EAE08A27-386C-579E-966D-8853EC3C5D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryResellerUserAlarmThresholdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryResellerUserAlarmThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResellerUserAlarmThresholdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryResellerUserAlarmThresholdResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QueryResellerUserAlarmThresholdResponseBody) GetData() []*QueryResellerUserAlarmThresholdResponseBodyData {
	return s.Data
}

func (s *QueryResellerUserAlarmThresholdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryResellerUserAlarmThresholdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryResellerUserAlarmThresholdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryResellerUserAlarmThresholdResponseBody) SetCode(v string) *QueryResellerUserAlarmThresholdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBody) SetCount(v int32) *QueryResellerUserAlarmThresholdResponseBody {
	s.Count = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBody) SetData(v []*QueryResellerUserAlarmThresholdResponseBodyData) *QueryResellerUserAlarmThresholdResponseBody {
	s.Data = v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBody) SetMessage(v string) *QueryResellerUserAlarmThresholdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBody) SetRequestId(v string) *QueryResellerUserAlarmThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBody) SetSuccess(v bool) *QueryResellerUserAlarmThresholdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryResellerUserAlarmThresholdResponseBodyData struct {
	// The denominator of the alert percentage. This parameter is valid only when ThresholdType is set to 1.
	//
	// example:
	//
	// 100
	Denominator *int32 `json:"Denominator,omitempty" xml:"Denominator,omitempty"`
	// The numerator of the alert percentage. This parameter is valid only when ThresholdType is set to 1.
	//
	// example:
	//
	// 10
	Numerator *int32 `json:"Numerator,omitempty" xml:"Numerator,omitempty"`
	// The alert threshold in amount. This parameter is valid only when ThresholdType is set to 0.
	//
	// example:
	//
	// 100
	ThresholdAmount *string `json:"ThresholdAmount,omitempty" xml:"ThresholdAmount,omitempty"`
	// The type of the alert threshold. Valid values: 0 (amount) and 1 (percentage).
	//
	// example:
	//
	// 0
	ThresholdType *int32 `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
}

func (s QueryResellerUserAlarmThresholdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryResellerUserAlarmThresholdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) GetDenominator() *int32 {
	return s.Denominator
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) GetNumerator() *int32 {
	return s.Numerator
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) GetThresholdAmount() *string {
	return s.ThresholdAmount
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) GetThresholdType() *int32 {
	return s.ThresholdType
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) SetDenominator(v int32) *QueryResellerUserAlarmThresholdResponseBodyData {
	s.Denominator = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) SetNumerator(v int32) *QueryResellerUserAlarmThresholdResponseBodyData {
	s.Numerator = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) SetThresholdAmount(v string) *QueryResellerUserAlarmThresholdResponseBodyData {
	s.ThresholdAmount = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) SetThresholdType(v int32) *QueryResellerUserAlarmThresholdResponseBodyData {
	s.ThresholdType = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
