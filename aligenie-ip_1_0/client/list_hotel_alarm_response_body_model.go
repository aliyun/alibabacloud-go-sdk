// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *ListHotelAlarmResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *ListHotelAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelAlarmResponseBody
	GetRequestId() *string
	SetResult(v []*ListHotelAlarmResponseBodyResult) *ListHotelAlarmResponseBody
	GetResult() []*ListHotelAlarmResponseBodyResult
	SetStatusCode(v int32) *ListHotelAlarmResponseBody
	GetStatusCode() *int32
}

type ListHotelAlarmResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***881
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelAlarmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListHotelAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *ListHotelAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelAlarmResponseBody) GetResult() []*ListHotelAlarmResponseBodyResult {
	return s.Result
}

func (s *ListHotelAlarmResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelAlarmResponseBody) SetExtentions(v map[string]interface{}) *ListHotelAlarmResponseBody {
	s.Extentions = v
	return s
}

func (s *ListHotelAlarmResponseBody) SetMessage(v string) *ListHotelAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelAlarmResponseBody) SetRequestId(v string) *ListHotelAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelAlarmResponseBody) SetResult(v []*ListHotelAlarmResponseBodyResult) *ListHotelAlarmResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelAlarmResponseBody) SetStatusCode(v int32) *ListHotelAlarmResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListHotelAlarmResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelAlarmResponseBodyResult struct {
	// example:
	//
	// 5039
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// example:
	//
	// PvkB****VVTA==
	DeviceOpenId *string                                       `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	ScheduleInfo *ListHotelAlarmResponseBodyResultScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	// example:
	//
	// mgw/k***HQd
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s ListHotelAlarmResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBodyResult) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *ListHotelAlarmResponseBodyResult) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *ListHotelAlarmResponseBodyResult) GetScheduleInfo() *ListHotelAlarmResponseBodyResultScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListHotelAlarmResponseBodyResult) GetUserOpenId() *string {
	return s.UserOpenId
}

func (s *ListHotelAlarmResponseBodyResult) SetAlarmId(v int64) *ListHotelAlarmResponseBodyResult {
	s.AlarmId = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResult) SetDeviceOpenId(v string) *ListHotelAlarmResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResult) SetScheduleInfo(v *ListHotelAlarmResponseBodyResultScheduleInfo) *ListHotelAlarmResponseBodyResult {
	s.ScheduleInfo = v
	return s
}

func (s *ListHotelAlarmResponseBodyResult) SetUserOpenId(v string) *ListHotelAlarmResponseBodyResult {
	s.UserOpenId = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResult) Validate() error {
	if s.ScheduleInfo != nil {
		if err := s.ScheduleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelAlarmResponseBodyResultScheduleInfo struct {
	Once *ListHotelAlarmResponseBodyResultScheduleInfoOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// ONCE, WEEKLY
	//
	// example:
	//
	// ONCE
	Type   *string                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly *ListHotelAlarmResponseBodyResultScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s ListHotelAlarmResponseBodyResultScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmResponseBodyResultScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) GetOnce() *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	return s.Once
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) GetType() *string {
	return s.Type
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) GetWeekly() *ListHotelAlarmResponseBodyResultScheduleInfoWeekly {
	return s.Weekly
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) SetOnce(v *ListHotelAlarmResponseBodyResultScheduleInfoOnce) *ListHotelAlarmResponseBodyResultScheduleInfo {
	s.Once = v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) SetType(v string) *ListHotelAlarmResponseBodyResultScheduleInfo {
	s.Type = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) SetWeekly(v *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) *ListHotelAlarmResponseBodyResultScheduleInfo {
	s.Weekly = v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfo) Validate() error {
	if s.Once != nil {
		if err := s.Once.Validate(); err != nil {
			return err
		}
	}
	if s.Weekly != nil {
		if err := s.Weekly.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelAlarmResponseBodyResultScheduleInfoOnce struct {
	// example:
	//
	// 20
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 30
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 9
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 2022
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ListHotelAlarmResponseBodyResultScheduleInfoOnce) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmResponseBodyResultScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) GetDay() *int32 {
	return s.Day
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) GetHour() *int32 {
	return s.Hour
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) GetMinute() *int32 {
	return s.Minute
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) GetMonth() *int32 {
	return s.Month
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) GetYear() *int32 {
	return s.Year
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetDay(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetHour(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetMinute(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetMonth(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) SetYear(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoOnce {
	s.Year = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoOnce) Validate() error {
	return dara.Validate(s)
}

type ListHotelAlarmResponseBodyResultScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s ListHotelAlarmResponseBodyResultScheduleInfoWeekly) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmResponseBodyResultScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) GetHour() *int32 {
	return s.Hour
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) GetMinute() *int32 {
	return s.Minute
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *ListHotelAlarmResponseBodyResultScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) SetHour(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) SetMinute(v int32) *ListHotelAlarmResponseBodyResultScheduleInfoWeekly {
	s.Minute = &v
	return s
}

func (s *ListHotelAlarmResponseBodyResultScheduleInfoWeekly) Validate() error {
	return dara.Validate(s)
}
