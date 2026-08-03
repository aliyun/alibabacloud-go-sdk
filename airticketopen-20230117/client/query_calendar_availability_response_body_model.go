// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCalendarAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryCalendarAvailabilityResponseBodyData) *QueryCalendarAvailabilityResponseBody
	GetData() *QueryCalendarAvailabilityResponseBodyData
	SetErrorCode(v string) *QueryCalendarAvailabilityResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryCalendarAvailabilityResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryCalendarAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCalendarAvailabilityResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *QueryCalendarAvailabilityResponseBody
	GetTracerId() *string
}

type QueryCalendarAvailabilityResponseBody struct {
	Data *QueryCalendarAvailabilityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryCalendarAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCalendarAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCalendarAvailabilityResponseBody) GetData() *QueryCalendarAvailabilityResponseBodyData {
	return s.Data
}

func (s *QueryCalendarAvailabilityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryCalendarAvailabilityResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryCalendarAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCalendarAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCalendarAvailabilityResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryCalendarAvailabilityResponseBody) SetData(v *QueryCalendarAvailabilityResponseBodyData) *QueryCalendarAvailabilityResponseBody {
	s.Data = v
	return s
}

func (s *QueryCalendarAvailabilityResponseBody) SetErrorCode(v string) *QueryCalendarAvailabilityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBody) SetErrorMsg(v string) *QueryCalendarAvailabilityResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBody) SetRequestId(v string) *QueryCalendarAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBody) SetSuccess(v bool) *QueryCalendarAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBody) SetTracerId(v string) *QueryCalendarAvailabilityResponseBody {
	s.TracerId = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCalendarAvailabilityResponseBodyData struct {
	FailedHotels []*QueryCalendarAvailabilityResponseBodyDataFailedHotels `json:"FailedHotels,omitempty" xml:"FailedHotels,omitempty" type:"Repeated"`
	Hotels       map[string][]*DataHotelsValue                            `json:"Hotels,omitempty" xml:"Hotels,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryCalendarAvailabilityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCalendarAvailabilityResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCalendarAvailabilityResponseBodyData) GetFailedHotels() []*QueryCalendarAvailabilityResponseBodyDataFailedHotels {
	return s.FailedHotels
}

func (s *QueryCalendarAvailabilityResponseBodyData) GetHotels() map[string][]*DataHotelsValue {
	return s.Hotels
}

func (s *QueryCalendarAvailabilityResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryCalendarAvailabilityResponseBodyData) SetFailedHotels(v []*QueryCalendarAvailabilityResponseBodyDataFailedHotels) *QueryCalendarAvailabilityResponseBodyData {
	s.FailedHotels = v
	return s
}

func (s *QueryCalendarAvailabilityResponseBodyData) SetHotels(v map[string][]*DataHotelsValue) *QueryCalendarAvailabilityResponseBodyData {
	s.Hotels = v
	return s
}

func (s *QueryCalendarAvailabilityResponseBodyData) SetTracerId(v string) *QueryCalendarAvailabilityResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBodyData) Validate() error {
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

type QueryCalendarAvailabilityResponseBodyDataFailedHotels struct {
	// example:
	//
	// HOTEL_NOT_FOUND
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 酒店不存在
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// H001
	StandardHotelId *string `json:"StandardHotelId,omitempty" xml:"StandardHotelId,omitempty"`
}

func (s QueryCalendarAvailabilityResponseBodyDataFailedHotels) String() string {
	return dara.Prettify(s)
}

func (s QueryCalendarAvailabilityResponseBodyDataFailedHotels) GoString() string {
	return s.String()
}

func (s *QueryCalendarAvailabilityResponseBodyDataFailedHotels) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryCalendarAvailabilityResponseBodyDataFailedHotels) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryCalendarAvailabilityResponseBodyDataFailedHotels) GetStandardHotelId() *string {
	return s.StandardHotelId
}

func (s *QueryCalendarAvailabilityResponseBodyDataFailedHotels) SetErrorCode(v string) *QueryCalendarAvailabilityResponseBodyDataFailedHotels {
	s.ErrorCode = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBodyDataFailedHotels) SetErrorMessage(v string) *QueryCalendarAvailabilityResponseBodyDataFailedHotels {
	s.ErrorMessage = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBodyDataFailedHotels) SetStandardHotelId(v string) *QueryCalendarAvailabilityResponseBodyDataFailedHotels {
	s.StandardHotelId = &v
	return s
}

func (s *QueryCalendarAvailabilityResponseBodyDataFailedHotels) Validate() error {
	return dara.Validate(s)
}
