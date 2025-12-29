// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *CreateHotelAlarmRequest
	GetHotelId() *string
	SetMusicType(v string) *CreateHotelAlarmRequest
	GetMusicType() *string
	SetRooms(v []*string) *CreateHotelAlarmRequest
	GetRooms() []*string
	SetScheduleInfo(v *CreateHotelAlarmRequestScheduleInfo) *CreateHotelAlarmRequest
	GetScheduleInfo() *CreateHotelAlarmRequestScheduleInfo
}

type CreateHotelAlarmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cf2446fc9d144c85aaee4f9ae20a96e7
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// DOU_YIN
	MusicType *string `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// This parameter is required.
	Rooms []*string `json:"Rooms,omitempty" xml:"Rooms,omitempty" type:"Repeated"`
	// This parameter is required.
	ScheduleInfo *CreateHotelAlarmRequestScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
}

func (s CreateHotelAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *CreateHotelAlarmRequest) GetMusicType() *string {
	return s.MusicType
}

func (s *CreateHotelAlarmRequest) GetRooms() []*string {
	return s.Rooms
}

func (s *CreateHotelAlarmRequest) GetScheduleInfo() *CreateHotelAlarmRequestScheduleInfo {
	return s.ScheduleInfo
}

func (s *CreateHotelAlarmRequest) SetHotelId(v string) *CreateHotelAlarmRequest {
	s.HotelId = &v
	return s
}

func (s *CreateHotelAlarmRequest) SetMusicType(v string) *CreateHotelAlarmRequest {
	s.MusicType = &v
	return s
}

func (s *CreateHotelAlarmRequest) SetRooms(v []*string) *CreateHotelAlarmRequest {
	s.Rooms = v
	return s
}

func (s *CreateHotelAlarmRequest) SetScheduleInfo(v *CreateHotelAlarmRequestScheduleInfo) *CreateHotelAlarmRequest {
	s.ScheduleInfo = v
	return s
}

func (s *CreateHotelAlarmRequest) Validate() error {
	if s.ScheduleInfo != nil {
		if err := s.ScheduleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateHotelAlarmRequestScheduleInfo struct {
	Once *CreateHotelAlarmRequestScheduleInfoOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// ONCE, WEEKLY
	//
	// This parameter is required.
	Type   *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly *CreateHotelAlarmRequestScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s CreateHotelAlarmRequestScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmRequestScheduleInfo) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmRequestScheduleInfo) GetOnce() *CreateHotelAlarmRequestScheduleInfoOnce {
	return s.Once
}

func (s *CreateHotelAlarmRequestScheduleInfo) GetType() *string {
	return s.Type
}

func (s *CreateHotelAlarmRequestScheduleInfo) GetWeekly() *CreateHotelAlarmRequestScheduleInfoWeekly {
	return s.Weekly
}

func (s *CreateHotelAlarmRequestScheduleInfo) SetOnce(v *CreateHotelAlarmRequestScheduleInfoOnce) *CreateHotelAlarmRequestScheduleInfo {
	s.Once = v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfo) SetType(v string) *CreateHotelAlarmRequestScheduleInfo {
	s.Type = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfo) SetWeekly(v *CreateHotelAlarmRequestScheduleInfoWeekly) *CreateHotelAlarmRequestScheduleInfo {
	s.Weekly = v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfo) Validate() error {
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

type CreateHotelAlarmRequestScheduleInfoOnce struct {
	// example:
	//
	// 20
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// example:
	//
	// 19
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

func (s CreateHotelAlarmRequestScheduleInfoOnce) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmRequestScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) GetDay() *int32 {
	return s.Day
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) GetHour() *int32 {
	return s.Hour
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) GetMinute() *int32 {
	return s.Minute
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) GetMonth() *int32 {
	return s.Month
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) GetYear() *int32 {
	return s.Year
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetDay(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetHour(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetMinute(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetMonth(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) SetYear(v int32) *CreateHotelAlarmRequestScheduleInfoOnce {
	s.Year = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoOnce) Validate() error {
	return dara.Validate(s)
}

type CreateHotelAlarmRequestScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 30
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s CreateHotelAlarmRequestScheduleInfoWeekly) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmRequestScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) GetHour() *int32 {
	return s.Hour
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) GetMinute() *int32 {
	return s.Minute
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *CreateHotelAlarmRequestScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) SetHour(v int32) *CreateHotelAlarmRequestScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) SetMinute(v int32) *CreateHotelAlarmRequestScheduleInfoWeekly {
	s.Minute = &v
	return s
}

func (s *CreateHotelAlarmRequestScheduleInfoWeekly) Validate() error {
	return dara.Validate(s)
}
