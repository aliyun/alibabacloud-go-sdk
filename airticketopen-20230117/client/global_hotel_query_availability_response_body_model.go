// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelQueryAvailabilityResponseBodyData) *GlobalHotelQueryAvailabilityResponseBody
	GetData() *GlobalHotelQueryAvailabilityResponseBodyData
	SetErrorCode(v string) *GlobalHotelQueryAvailabilityResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelQueryAvailabilityResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelQueryAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelQueryAvailabilityResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelQueryAvailabilityResponseBody
	GetTracerId() *string
}

type GlobalHotelQueryAvailabilityResponseBody struct {
	// The business data.
	Data *GlobalHotelQueryAvailabilityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// CityCodeRequired
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// City code cannot be empty
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryAvailabilityResponseBody) GetData() *GlobalHotelQueryAvailabilityResponseBodyData {
	return s.Data
}

func (s *GlobalHotelQueryAvailabilityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelQueryAvailabilityResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelQueryAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelQueryAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelQueryAvailabilityResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryAvailabilityResponseBody) SetData(v *GlobalHotelQueryAvailabilityResponseBodyData) *GlobalHotelQueryAvailabilityResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponseBody) SetErrorCode(v string) *GlobalHotelQueryAvailabilityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponseBody) SetErrorMsg(v string) *GlobalHotelQueryAvailabilityResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponseBody) SetRequestId(v string) *GlobalHotelQueryAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponseBody) SetSuccess(v bool) *GlobalHotelQueryAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponseBody) SetTracerId(v string) *GlobalHotelQueryAvailabilityResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelQueryAvailabilityResponseBodyData struct {
	// The room type offers grouped by standard hotel ID.
	Hotels map[string][]*DataHotelsValue `json:"Hotels,omitempty" xml:"Hotels,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryAvailabilityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryAvailabilityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryAvailabilityResponseBodyData) GetHotels() map[string][]*DataHotelsValue {
	return s.Hotels
}

func (s *GlobalHotelQueryAvailabilityResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryAvailabilityResponseBodyData) SetHotels(v map[string][]*DataHotelsValue) *GlobalHotelQueryAvailabilityResponseBodyData {
	s.Hotels = v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponseBodyData) SetTracerId(v string) *GlobalHotelQueryAvailabilityResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
