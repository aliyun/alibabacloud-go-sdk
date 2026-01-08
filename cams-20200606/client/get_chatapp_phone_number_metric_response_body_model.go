// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappPhoneNumberMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetChatappPhoneNumberMetricResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetChatappPhoneNumberMetricResponseBody
	GetCode() *string
	SetData(v []*GetChatappPhoneNumberMetricResponseBodyData) *GetChatappPhoneNumberMetricResponseBody
	GetData() []*GetChatappPhoneNumberMetricResponseBodyData
	SetMessage(v string) *GetChatappPhoneNumberMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatappPhoneNumberMetricResponseBody
	GetRequestId() *string
}

type GetChatappPhoneNumberMetricResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// NONE
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The value OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*GetChatappPhoneNumberMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1612C226-E271-4CFE-9F18-4066D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappPhoneNumberMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatappPhoneNumberMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberMetricResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetChatappPhoneNumberMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatappPhoneNumberMetricResponseBody) GetData() []*GetChatappPhoneNumberMetricResponseBodyData {
	return s.Data
}

func (s *GetChatappPhoneNumberMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatappPhoneNumberMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetAccessDeniedDetail(v string) *GetChatappPhoneNumberMetricResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetCode(v string) *GetChatappPhoneNumberMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetData(v []*GetChatappPhoneNumberMetricResponseBodyData) *GetChatappPhoneNumberMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetMessage(v string) *GetChatappPhoneNumberMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) SetRequestId(v string) *GetChatappPhoneNumberMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBody) Validate() error {
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

type GetChatappPhoneNumberMetricResponseBodyData struct {
	// The number of delivered messages.
	//
	// example:
	//
	// 5
	DeliveredCount *int32 `json:"DeliveredCount,omitempty" xml:"DeliveredCount,omitempty"`
	// The end of the time range that you queried.
	//
	// example:
	//
	// 1667196043904
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The granularity of the metric.
	//
	// Valid values:
	//
	// 	- DAILY
	//
	// 	- HALF_HOUR
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The business phone number.
	//
	// example:
	//
	// 861380000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The number of sent messages.
	//
	// example:
	//
	// 10
	SentCount *int32 `json:"SentCount,omitempty" xml:"SentCount,omitempty"`
	// The beginning of the time range that you queried.
	//
	// example:
	//
	// 1669619491000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetChatappPhoneNumberMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatappPhoneNumberMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) GetDeliveredCount() *int32 {
	return s.DeliveredCount
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) GetEnd() *int64 {
	return s.End
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) GetGranularity() *string {
	return s.Granularity
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) GetSentCount() *int32 {
	return s.SentCount
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) GetStart() *int64 {
	return s.Start
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetDeliveredCount(v int32) *GetChatappPhoneNumberMetricResponseBodyData {
	s.DeliveredCount = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetEnd(v int64) *GetChatappPhoneNumberMetricResponseBodyData {
	s.End = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetGranularity(v string) *GetChatappPhoneNumberMetricResponseBodyData {
	s.Granularity = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetPhoneNumber(v string) *GetChatappPhoneNumberMetricResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetSentCount(v int32) *GetChatappPhoneNumberMetricResponseBodyData {
	s.SentCount = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) SetStart(v int64) *GetChatappPhoneNumberMetricResponseBodyData {
	s.Start = &v
	return s
}

func (s *GetChatappPhoneNumberMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
