// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarms(v []*UpdateHotelAlarmRequestAlarms) *UpdateHotelAlarmRequest
	GetAlarms() []*UpdateHotelAlarmRequestAlarms
	SetHotelId(v string) *UpdateHotelAlarmRequest
	GetHotelId() *string
	SetScheduleInfo(v *UpdateHotelAlarmRequestScheduleInfo) *UpdateHotelAlarmRequest
	GetScheduleInfo() *UpdateHotelAlarmRequestScheduleInfo
}

type UpdateHotelAlarmRequest struct {
	// This parameter is required.
	Alarms []*UpdateHotelAlarmRequestAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7a381a668bc485980bed3876a75e013
	HotelId      *string                              `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	ScheduleInfo *UpdateHotelAlarmRequestScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
}

func (s UpdateHotelAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequest) GetAlarms() []*UpdateHotelAlarmRequestAlarms {
	return s.Alarms
}

func (s *UpdateHotelAlarmRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateHotelAlarmRequest) GetScheduleInfo() *UpdateHotelAlarmRequestScheduleInfo {
	return s.ScheduleInfo
}

func (s *UpdateHotelAlarmRequest) SetAlarms(v []*UpdateHotelAlarmRequestAlarms) *UpdateHotelAlarmRequest {
	s.Alarms = v
	return s
}

func (s *UpdateHotelAlarmRequest) SetHotelId(v string) *UpdateHotelAlarmRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelAlarmRequest) SetScheduleInfo(v *UpdateHotelAlarmRequestScheduleInfo) *UpdateHotelAlarmRequest {
	s.ScheduleInfo = v
	return s
}

func (s *UpdateHotelAlarmRequest) Validate() error {
	if s.Alarms != nil {
		for _, item := range s.Alarms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduleInfo != nil {
		if err := s.ScheduleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHotelAlarmRequestAlarms struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pvk***VTA==
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	// example:
	//
	// 101
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mgw/***dHQd
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s UpdateHotelAlarmRequestAlarms) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmRequestAlarms) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequestAlarms) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *UpdateHotelAlarmRequestAlarms) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *UpdateHotelAlarmRequestAlarms) GetRoomNo() *string {
	return s.RoomNo
}

func (s *UpdateHotelAlarmRequestAlarms) GetUserOpenId() *string {
	return s.UserOpenId
}

func (s *UpdateHotelAlarmRequestAlarms) SetAlarmId(v int64) *UpdateHotelAlarmRequestAlarms {
	s.AlarmId = &v
	return s
}

func (s *UpdateHotelAlarmRequestAlarms) SetDeviceOpenId(v string) *UpdateHotelAlarmRequestAlarms {
	s.DeviceOpenId = &v
	return s
}

func (s *UpdateHotelAlarmRequestAlarms) SetRoomNo(v string) *UpdateHotelAlarmRequestAlarms {
	s.RoomNo = &v
	return s
}

func (s *UpdateHotelAlarmRequestAlarms) SetUserOpenId(v string) *UpdateHotelAlarmRequestAlarms {
	s.UserOpenId = &v
	return s
}

func (s *UpdateHotelAlarmRequestAlarms) Validate() error {
	return dara.Validate(s)
}

type UpdateHotelAlarmRequestScheduleInfo struct {
	Once *UpdateHotelAlarmRequestScheduleInfoOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// ONCE, WEEKLY
	//
	// example:
	//
	// ONCE
	Type   *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly *UpdateHotelAlarmRequestScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s UpdateHotelAlarmRequestScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmRequestScheduleInfo) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequestScheduleInfo) GetOnce() *UpdateHotelAlarmRequestScheduleInfoOnce {
	return s.Once
}

func (s *UpdateHotelAlarmRequestScheduleInfo) GetType() *string {
	return s.Type
}

func (s *UpdateHotelAlarmRequestScheduleInfo) GetWeekly() *UpdateHotelAlarmRequestScheduleInfoWeekly {
	return s.Weekly
}

func (s *UpdateHotelAlarmRequestScheduleInfo) SetOnce(v *UpdateHotelAlarmRequestScheduleInfoOnce) *UpdateHotelAlarmRequestScheduleInfo {
	s.Once = v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfo) SetType(v string) *UpdateHotelAlarmRequestScheduleInfo {
	s.Type = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfo) SetWeekly(v *UpdateHotelAlarmRequestScheduleInfoWeekly) *UpdateHotelAlarmRequestScheduleInfo {
	s.Weekly = v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfo) Validate() error {
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

type UpdateHotelAlarmRequestScheduleInfoOnce struct {
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
	// 0
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

func (s UpdateHotelAlarmRequestScheduleInfoOnce) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmRequestScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) GetDay() *int32 {
	return s.Day
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) GetHour() *int32 {
	return s.Hour
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) GetMinute() *int32 {
	return s.Minute
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) GetMonth() *int32 {
	return s.Month
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) GetYear() *int32 {
	return s.Year
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetDay(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetHour(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetMinute(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetMonth(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) SetYear(v int32) *UpdateHotelAlarmRequestScheduleInfoOnce {
	s.Year = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoOnce) Validate() error {
	return dara.Validate(s)
}

type UpdateHotelAlarmRequestScheduleInfoWeekly struct {
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

func (s UpdateHotelAlarmRequestScheduleInfoWeekly) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmRequestScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) GetHour() *int32 {
	return s.Hour
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) GetMinute() *int32 {
	return s.Minute
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *UpdateHotelAlarmRequestScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) SetHour(v int32) *UpdateHotelAlarmRequestScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) SetMinute(v int32) *UpdateHotelAlarmRequestScheduleInfoWeekly {
	s.Minute = &v
	return s
}

func (s *UpdateHotelAlarmRequestScheduleInfoWeekly) Validate() error {
	return dara.Validate(s)
}
