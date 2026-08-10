// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryCalendarAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelQueryCalendarAvailabilityResponseBodyData) *GlobalHotelQueryCalendarAvailabilityResponseBody
	GetData() *GlobalHotelQueryCalendarAvailabilityResponseBodyData
	SetErrorCode(v string) *GlobalHotelQueryCalendarAvailabilityResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelQueryCalendarAvailabilityResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelQueryCalendarAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelQueryCalendarAvailabilityResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelQueryCalendarAvailabilityResponseBody
	GetTracerId() *string
}

type GlobalHotelQueryCalendarAvailabilityResponseBody struct {
	// The business data.
	Data *GlobalHotelQueryCalendarAvailabilityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GlobalHotelQueryCalendarAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryCalendarAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) GetData() *GlobalHotelQueryCalendarAvailabilityResponseBodyData {
	return s.Data
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) SetData(v *GlobalHotelQueryCalendarAvailabilityResponseBodyData) *GlobalHotelQueryCalendarAvailabilityResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) SetErrorCode(v string) *GlobalHotelQueryCalendarAvailabilityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) SetErrorMsg(v string) *GlobalHotelQueryCalendarAvailabilityResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) SetRequestId(v string) *GlobalHotelQueryCalendarAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) SetSuccess(v bool) *GlobalHotelQueryCalendarAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) SetTracerId(v string) *GlobalHotelQueryCalendarAvailabilityResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelQueryCalendarAvailabilityResponseBodyData struct {
	// The list of failed hotels (in partial success mode).
	FailedHotels []*GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels `json:"FailedHotels,omitempty" xml:"FailedHotels,omitempty" type:"Repeated"`
	// The calendar quotes grouped by standard hotel ID.
	Hotels map[string][]*DataHotelsValue `json:"Hotels,omitempty" xml:"Hotels,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryCalendarAvailabilityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryCalendarAvailabilityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyData) GetFailedHotels() []*GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels {
	return s.FailedHotels
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyData) GetHotels() map[string][]*DataHotelsValue {
	return s.Hotels
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyData) SetFailedHotels(v []*GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) *GlobalHotelQueryCalendarAvailabilityResponseBodyData {
	s.FailedHotels = v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyData) SetHotels(v map[string][]*DataHotelsValue) *GlobalHotelQueryCalendarAvailabilityResponseBodyData {
	s.Hotels = v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyData) SetTracerId(v string) *GlobalHotelQueryCalendarAvailabilityResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyData) Validate() error {
	if s.FailedHotels != nil {
		for _, item := range s.FailedHotels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels struct {
	// The error code.
	//
	// example:
	//
	// HOTEL_NOT_FOUND
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error description.
	//
	// example:
	//
	// The hotel does not exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The standard hotel ID.
	//
	// example:
	//
	// H001
	StandardHotelId *string `json:"StandardHotelId,omitempty" xml:"StandardHotelId,omitempty"`
}

func (s GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) GetStandardHotelId() *string {
	return s.StandardHotelId
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) SetErrorCode(v string) *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) SetErrorMessage(v string) *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels {
	s.ErrorMessage = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) SetStandardHotelId(v string) *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels {
	s.StandardHotelId = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponseBodyDataFailedHotels) Validate() error {
	return dara.Validate(s)
}
