// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryAvailabilityResponseBodyData) *QueryAvailabilityResponseBody
	GetData() *QueryAvailabilityResponseBodyData
	SetErrorCode(v string) *QueryAvailabilityResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryAvailabilityResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAvailabilityResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *QueryAvailabilityResponseBody
	GetTracerId() *string
}

type QueryAvailabilityResponseBody struct {
	Data *QueryAvailabilityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CityCodeRequired
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 城市编码不能为空
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvailabilityResponseBody) GetData() *QueryAvailabilityResponseBodyData {
	return s.Data
}

func (s *QueryAvailabilityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryAvailabilityResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAvailabilityResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryAvailabilityResponseBody) SetData(v *QueryAvailabilityResponseBodyData) *QueryAvailabilityResponseBody {
	s.Data = v
	return s
}

func (s *QueryAvailabilityResponseBody) SetErrorCode(v string) *QueryAvailabilityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryAvailabilityResponseBody) SetErrorMsg(v string) *QueryAvailabilityResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryAvailabilityResponseBody) SetRequestId(v string) *QueryAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvailabilityResponseBody) SetSuccess(v bool) *QueryAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAvailabilityResponseBody) SetTracerId(v string) *QueryAvailabilityResponseBody {
	s.TracerId = &v
	return s
}

func (s *QueryAvailabilityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAvailabilityResponseBodyData struct {
	Hotels map[string][]*DataHotelsValue `json:"Hotels,omitempty" xml:"Hotels,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryAvailabilityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailabilityResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAvailabilityResponseBodyData) GetHotels() map[string][]*DataHotelsValue {
	return s.Hotels
}

func (s *QueryAvailabilityResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryAvailabilityResponseBodyData) SetHotels(v map[string][]*DataHotelsValue) *QueryAvailabilityResponseBodyData {
	s.Hotels = v
	return s
}

func (s *QueryAvailabilityResponseBodyData) SetTracerId(v string) *QueryAvailabilityResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *QueryAvailabilityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
