// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePickUpWaybillPreQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreatePickUpWaybillPreQueryResponseBodyData) *CreatePickUpWaybillPreQueryResponseBody
	GetData() *CreatePickUpWaybillPreQueryResponseBodyData
	SetHttpStatusCode(v int32) *CreatePickUpWaybillPreQueryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePickUpWaybillPreQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePickUpWaybillPreQueryResponseBody
	GetRequestId() *string
}

type CreatePickUpWaybillPreQueryResponseBody struct {
	// The result set.
	Data *CreatePickUpWaybillPreQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FC30594-3841-43AD-9008-03393BCB5CD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBody) GetData() *CreatePickUpWaybillPreQueryResponseBodyData {
	return s.Data
}

func (s *CreatePickUpWaybillPreQueryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePickUpWaybillPreQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePickUpWaybillPreQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePickUpWaybillPreQueryResponseBody) SetData(v *CreatePickUpWaybillPreQueryResponseBodyData) *CreatePickUpWaybillPreQueryResponseBody {
	s.Data = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBody) SetHttpStatusCode(v int32) *CreatePickUpWaybillPreQueryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBody) SetMessage(v string) *CreatePickUpWaybillPreQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBody) SetRequestId(v string) *CreatePickUpWaybillPreQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePickUpWaybillPreQueryResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of information about whether courier companies are available for placing orders.
	CpTimeSelectList []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList `json:"CpTimeSelectList,omitempty" xml:"CpTimeSelectList,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// None
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The response content.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) GetCpTimeSelectList() []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList {
	return s.CpTimeSelectList
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetCode(v string) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetCpTimeSelectList(v []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.CpTimeSelectList = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetErrorCode(v string) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetErrorMsg(v string) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetMessage(v string) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) SetSuccess(v bool) *CreatePickUpWaybillPreQueryResponseBodyData {
	s.Success = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyData) Validate() error {
	if s.CpTimeSelectList != nil {
		for _, item := range s.CpTimeSelectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList struct {
	// The available time options for scheduled orders. If the courier company does not support scheduled orders, this field is empty.
	AppointTimes []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes `json:"AppointTimes,omitempty" xml:"AppointTimes,omitempty" type:"Repeated"`
	// The estimated price, in CNY, rounded to two decimal places. This field has a value only when the estimated weight is specified in the request parameters.
	//
	// example:
	//
	// 12.50
	PrePrice *string `json:"PrePrice,omitempty" xml:"PrePrice,omitempty"`
	// The available options for real-time orders.
	RealTime *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime `json:"RealTime,omitempty" xml:"RealTime,omitempty" type:"Struct"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) GetAppointTimes() []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes {
	return s.AppointTimes
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) GetPrePrice() *string {
	return s.PrePrice
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) GetRealTime() *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime {
	return s.RealTime
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) SetAppointTimes(v []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList {
	s.AppointTimes = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) SetPrePrice(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList {
	s.PrePrice = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) SetRealTime(v *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList {
	s.RealTime = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectList) Validate() error {
	if s.AppointTimes != nil {
		for _, item := range s.AppointTimes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RealTime != nil {
		if err := s.RealTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes struct {
	// The date string in YYYY-MM-DD format.
	//
	// example:
	//
	// 2022-04-28
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Indicates whether the date is selectable.
	//
	// example:
	//
	// true
	DateSelectable *bool `json:"DateSelectable,omitempty" xml:"DateSelectable,omitempty"`
	// The list of available time slots for the day.
	TimeList []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) GetDate() *string {
	return s.Date
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) GetDateSelectable() *bool {
	return s.DateSelectable
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) GetTimeList() []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	return s.TimeList
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) SetDate(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes {
	s.Date = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) SetDateSelectable(v bool) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes {
	s.DateSelectable = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) SetTimeList(v []*CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes {
	s.TimeList = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimes) Validate() error {
	if s.TimeList != nil {
		for _, item := range s.TimeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList struct {
	// The end time of the time slot, in HH:mm:ss format.
	//
	// example:
	//
	// 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The tip message displayed when the option is not selectable.
	//
	// example:
	//
	// Fully booked
	SelectDisableTip *string `json:"SelectDisableTip,omitempty" xml:"SelectDisableTip,omitempty"`
	// Indicates whether the scheduled time slot is selectable.
	//
	// example:
	//
	// true
	Selectable *bool `json:"Selectable,omitempty" xml:"Selectable,omitempty"`
	// The start time of the time slot, in HH:mm:ss format.
	//
	// example:
	//
	// 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) GetEndTime() *string {
	return s.EndTime
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) GetSelectDisableTip() *string {
	return s.SelectDisableTip
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) GetSelectable() *bool {
	return s.Selectable
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) GetStartTime() *string {
	return s.StartTime
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) SetEndTime(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	s.EndTime = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) SetSelectDisableTip(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	s.SelectDisableTip = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) SetSelectable(v bool) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	s.Selectable = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) SetStartTime(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList {
	s.StartTime = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListAppointTimesTimeList) Validate() error {
	return dara.Validate(s)
}

type CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime struct {
	// The service type name for real-time orders.
	//
	// example:
	//
	// 2-hour pickup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tip message displayed when the option is not selectable.
	//
	// example:
	//
	// The real-time order placement time range has been exceeded
	SelectDisableTip *string `json:"SelectDisableTip,omitempty" xml:"SelectDisableTip,omitempty"`
	// Indicates whether a real-time order can still be placed after the real-time order placement time has passed.
	//
	// example:
	//
	// false
	Selectable *bool `json:"Selectable,omitempty" xml:"Selectable,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) GetName() *string {
	return s.Name
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) GetSelectDisableTip() *string {
	return s.SelectDisableTip
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) GetSelectable() *bool {
	return s.Selectable
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) SetName(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime {
	s.Name = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) SetSelectDisableTip(v string) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime {
	s.SelectDisableTip = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) SetSelectable(v bool) *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime {
	s.Selectable = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponseBodyDataCpTimeSelectListRealTime) Validate() error {
	return dara.Validate(s)
}
