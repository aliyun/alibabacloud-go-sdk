// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttendees(v []*PatchEventResponseBodyAttendees) *PatchEventResponseBody
	GetAttendees() []*PatchEventResponseBodyAttendees
	SetCreateTime(v string) *PatchEventResponseBody
	GetCreateTime() *string
	SetDescription(v string) *PatchEventResponseBody
	GetDescription() *string
	SetEnd(v *PatchEventResponseBodyEnd) *PatchEventResponseBody
	GetEnd() *PatchEventResponseBodyEnd
	SetId(v string) *PatchEventResponseBody
	GetId() *string
	SetIsAllDay(v bool) *PatchEventResponseBody
	GetIsAllDay() *bool
	SetLocation(v *PatchEventResponseBodyLocation) *PatchEventResponseBody
	GetLocation() *PatchEventResponseBodyLocation
	SetOrganizer(v *PatchEventResponseBodyOrganizer) *PatchEventResponseBody
	GetOrganizer() *PatchEventResponseBodyOrganizer
	SetRecurrence(v *PatchEventResponseBodyRecurrence) *PatchEventResponseBody
	GetRecurrence() *PatchEventResponseBodyRecurrence
	SetReminders(v []*PatchEventResponseBodyReminders) *PatchEventResponseBody
	GetReminders() []*PatchEventResponseBodyReminders
	SetRequestId(v string) *PatchEventResponseBody
	GetRequestId() *string
	SetStart(v *PatchEventResponseBodyStart) *PatchEventResponseBody
	GetStart() *PatchEventResponseBodyStart
	SetSummary(v string) *PatchEventResponseBody
	GetSummary() *string
	SetUpdateTime(v string) *PatchEventResponseBody
	GetUpdateTime() *string
}

type PatchEventResponseBody struct {
	Attendees []*PatchEventResponseBodyAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-10-25T02:26:14Z
	CreateTime  *string                    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string                    `json:"description,omitempty" xml:"description,omitempty"`
	End         *PatchEventResponseBodyEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsAllDay   *bool                              `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	Location   *PatchEventResponseBodyLocation    `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	Organizer  *PatchEventResponseBodyOrganizer   `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	Recurrence *PatchEventResponseBodyRecurrence  `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders  []*PatchEventResponseBodyReminders `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 4248DCC9-785F-5A14-8BE0-830FD52E1261
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Start     *PatchEventResponseBodyStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	Summary   *string                      `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s PatchEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBody) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBody) GetAttendees() []*PatchEventResponseBodyAttendees {
	return s.Attendees
}

func (s *PatchEventResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PatchEventResponseBody) GetDescription() *string {
	return s.Description
}

func (s *PatchEventResponseBody) GetEnd() *PatchEventResponseBodyEnd {
	return s.End
}

func (s *PatchEventResponseBody) GetId() *string {
	return s.Id
}

func (s *PatchEventResponseBody) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *PatchEventResponseBody) GetLocation() *PatchEventResponseBodyLocation {
	return s.Location
}

func (s *PatchEventResponseBody) GetOrganizer() *PatchEventResponseBodyOrganizer {
	return s.Organizer
}

func (s *PatchEventResponseBody) GetRecurrence() *PatchEventResponseBodyRecurrence {
	return s.Recurrence
}

func (s *PatchEventResponseBody) GetReminders() []*PatchEventResponseBodyReminders {
	return s.Reminders
}

func (s *PatchEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PatchEventResponseBody) GetStart() *PatchEventResponseBodyStart {
	return s.Start
}

func (s *PatchEventResponseBody) GetSummary() *string {
	return s.Summary
}

func (s *PatchEventResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *PatchEventResponseBody) SetAttendees(v []*PatchEventResponseBodyAttendees) *PatchEventResponseBody {
	s.Attendees = v
	return s
}

func (s *PatchEventResponseBody) SetCreateTime(v string) *PatchEventResponseBody {
	s.CreateTime = &v
	return s
}

func (s *PatchEventResponseBody) SetDescription(v string) *PatchEventResponseBody {
	s.Description = &v
	return s
}

func (s *PatchEventResponseBody) SetEnd(v *PatchEventResponseBodyEnd) *PatchEventResponseBody {
	s.End = v
	return s
}

func (s *PatchEventResponseBody) SetId(v string) *PatchEventResponseBody {
	s.Id = &v
	return s
}

func (s *PatchEventResponseBody) SetIsAllDay(v bool) *PatchEventResponseBody {
	s.IsAllDay = &v
	return s
}

func (s *PatchEventResponseBody) SetLocation(v *PatchEventResponseBodyLocation) *PatchEventResponseBody {
	s.Location = v
	return s
}

func (s *PatchEventResponseBody) SetOrganizer(v *PatchEventResponseBodyOrganizer) *PatchEventResponseBody {
	s.Organizer = v
	return s
}

func (s *PatchEventResponseBody) SetRecurrence(v *PatchEventResponseBodyRecurrence) *PatchEventResponseBody {
	s.Recurrence = v
	return s
}

func (s *PatchEventResponseBody) SetReminders(v []*PatchEventResponseBodyReminders) *PatchEventResponseBody {
	s.Reminders = v
	return s
}

func (s *PatchEventResponseBody) SetRequestId(v string) *PatchEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *PatchEventResponseBody) SetStart(v *PatchEventResponseBodyStart) *PatchEventResponseBody {
	s.Start = v
	return s
}

func (s *PatchEventResponseBody) SetSummary(v string) *PatchEventResponseBody {
	s.Summary = &v
	return s
}

func (s *PatchEventResponseBody) SetUpdateTime(v string) *PatchEventResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *PatchEventResponseBody) Validate() error {
	if s.Attendees != nil {
		for _, item := range s.Attendees {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.End != nil {
		if err := s.End.Validate(); err != nil {
			return err
		}
	}
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.Validate(); err != nil {
			return err
		}
	}
	if s.Recurrence != nil {
		if err := s.Recurrence.Validate(); err != nil {
			return err
		}
	}
	if s.Reminders != nil {
		for _, item := range s.Reminders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Start != nil {
		if err := s.Start.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PatchEventResponseBodyAttendees struct {
	// example:
	//
	// tony
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsOptional *bool `json:"IsOptional,omitempty" xml:"IsOptional,omitempty"`
	// example:
	//
	// accepted
	ResponseStatus *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	// example:
	//
	// true
	Self *bool `json:"Self,omitempty" xml:"Self,omitempty"`
}

func (s PatchEventResponseBodyAttendees) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyAttendees) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyAttendees) GetDisplayName() *string {
	return s.DisplayName
}

func (s *PatchEventResponseBodyAttendees) GetId() *string {
	return s.Id
}

func (s *PatchEventResponseBodyAttendees) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *PatchEventResponseBodyAttendees) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *PatchEventResponseBodyAttendees) GetSelf() *bool {
	return s.Self
}

func (s *PatchEventResponseBodyAttendees) SetDisplayName(v string) *PatchEventResponseBodyAttendees {
	s.DisplayName = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) SetId(v string) *PatchEventResponseBodyAttendees {
	s.Id = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) SetIsOptional(v bool) *PatchEventResponseBodyAttendees {
	s.IsOptional = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) SetResponseStatus(v string) *PatchEventResponseBodyAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) SetSelf(v bool) *PatchEventResponseBodyAttendees {
	s.Self = &v
	return s
}

func (s *PatchEventResponseBodyAttendees) Validate() error {
	return dara.Validate(s)
}

type PatchEventResponseBodyEnd struct {
	// example:
	//
	// 2020-01-01
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s PatchEventResponseBodyEnd) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyEnd) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyEnd) GetDate() *string {
	return s.Date
}

func (s *PatchEventResponseBodyEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *PatchEventResponseBodyEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *PatchEventResponseBodyEnd) SetDate(v string) *PatchEventResponseBodyEnd {
	s.Date = &v
	return s
}

func (s *PatchEventResponseBodyEnd) SetDateTime(v string) *PatchEventResponseBodyEnd {
	s.DateTime = &v
	return s
}

func (s *PatchEventResponseBodyEnd) SetTimeZone(v string) *PatchEventResponseBodyEnd {
	s.TimeZone = &v
	return s
}

func (s *PatchEventResponseBodyEnd) Validate() error {
	return dara.Validate(s)
}

type PatchEventResponseBodyLocation struct {
	// example:
	//
	// true
	DisplayName  *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MeetingRooms []*string `json:"MeetingRooms,omitempty" xml:"MeetingRooms,omitempty" type:"Repeated"`
}

func (s PatchEventResponseBodyLocation) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyLocation) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyLocation) GetDisplayName() *string {
	return s.DisplayName
}

func (s *PatchEventResponseBodyLocation) GetMeetingRooms() []*string {
	return s.MeetingRooms
}

func (s *PatchEventResponseBodyLocation) SetDisplayName(v string) *PatchEventResponseBodyLocation {
	s.DisplayName = &v
	return s
}

func (s *PatchEventResponseBodyLocation) SetMeetingRooms(v []*string) *PatchEventResponseBodyLocation {
	s.MeetingRooms = v
	return s
}

func (s *PatchEventResponseBodyLocation) Validate() error {
	return dara.Validate(s)
}

type PatchEventResponseBodyOrganizer struct {
	// example:
	//
	// tony
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// accepted
	ResponseStatus *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	// example:
	//
	// true
	Self *bool `json:"Self,omitempty" xml:"Self,omitempty"`
}

func (s PatchEventResponseBodyOrganizer) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyOrganizer) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyOrganizer) GetDisplayName() *string {
	return s.DisplayName
}

func (s *PatchEventResponseBodyOrganizer) GetId() *string {
	return s.Id
}

func (s *PatchEventResponseBodyOrganizer) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *PatchEventResponseBodyOrganizer) GetSelf() *bool {
	return s.Self
}

func (s *PatchEventResponseBodyOrganizer) SetDisplayName(v string) *PatchEventResponseBodyOrganizer {
	s.DisplayName = &v
	return s
}

func (s *PatchEventResponseBodyOrganizer) SetId(v string) *PatchEventResponseBodyOrganizer {
	s.Id = &v
	return s
}

func (s *PatchEventResponseBodyOrganizer) SetResponseStatus(v string) *PatchEventResponseBodyOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *PatchEventResponseBodyOrganizer) SetSelf(v bool) *PatchEventResponseBodyOrganizer {
	s.Self = &v
	return s
}

func (s *PatchEventResponseBodyOrganizer) Validate() error {
	return dara.Validate(s)
}

type PatchEventResponseBodyRecurrence struct {
	Pattern *PatchEventResponseBodyRecurrencePattern `json:"Pattern,omitempty" xml:"Pattern,omitempty" type:"Struct"`
	Range   *PatchEventResponseBodyRecurrenceRange   `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
}

func (s PatchEventResponseBodyRecurrence) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyRecurrence) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyRecurrence) GetPattern() *PatchEventResponseBodyRecurrencePattern {
	return s.Pattern
}

func (s *PatchEventResponseBodyRecurrence) GetRange() *PatchEventResponseBodyRecurrenceRange {
	return s.Range
}

func (s *PatchEventResponseBodyRecurrence) SetPattern(v *PatchEventResponseBodyRecurrencePattern) *PatchEventResponseBodyRecurrence {
	s.Pattern = v
	return s
}

func (s *PatchEventResponseBodyRecurrence) SetRange(v *PatchEventResponseBodyRecurrenceRange) *PatchEventResponseBodyRecurrence {
	s.Range = v
	return s
}

func (s *PatchEventResponseBodyRecurrence) Validate() error {
	if s.Pattern != nil {
		if err := s.Pattern.Validate(); err != nil {
			return err
		}
	}
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PatchEventResponseBodyRecurrencePattern struct {
	// example:
	//
	// 14
	DayOfMonth *int32 `json:"DayOfMonth,omitempty" xml:"DayOfMonth,omitempty"`
	// example:
	//
	// sunday
	DaysOfWeek *string `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty"`
	// example:
	//
	// first
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// daily
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PatchEventResponseBodyRecurrencePattern) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyRecurrencePattern) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyRecurrencePattern) GetDayOfMonth() *int32 {
	return s.DayOfMonth
}

func (s *PatchEventResponseBodyRecurrencePattern) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *PatchEventResponseBodyRecurrencePattern) GetIndex() *string {
	return s.Index
}

func (s *PatchEventResponseBodyRecurrencePattern) GetInterval() *int32 {
	return s.Interval
}

func (s *PatchEventResponseBodyRecurrencePattern) GetType() *string {
	return s.Type
}

func (s *PatchEventResponseBodyRecurrencePattern) SetDayOfMonth(v int32) *PatchEventResponseBodyRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) SetDaysOfWeek(v string) *PatchEventResponseBodyRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) SetIndex(v string) *PatchEventResponseBodyRecurrencePattern {
	s.Index = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) SetInterval(v int32) *PatchEventResponseBodyRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) SetType(v string) *PatchEventResponseBodyRecurrencePattern {
	s.Type = &v
	return s
}

func (s *PatchEventResponseBodyRecurrencePattern) Validate() error {
	return dara.Validate(s)
}

type PatchEventResponseBodyRecurrenceRange struct {
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 5
	NumberOfOccurrences *int32 `json:"NumberOfOccurrences,omitempty" xml:"NumberOfOccurrences,omitempty"`
	// example:
	//
	// noEnd
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PatchEventResponseBodyRecurrenceRange) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyRecurrenceRange) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyRecurrenceRange) GetEndDate() *string {
	return s.EndDate
}

func (s *PatchEventResponseBodyRecurrenceRange) GetNumberOfOccurrences() *int32 {
	return s.NumberOfOccurrences
}

func (s *PatchEventResponseBodyRecurrenceRange) GetType() *string {
	return s.Type
}

func (s *PatchEventResponseBodyRecurrenceRange) SetEndDate(v string) *PatchEventResponseBodyRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *PatchEventResponseBodyRecurrenceRange) SetNumberOfOccurrences(v int32) *PatchEventResponseBodyRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *PatchEventResponseBodyRecurrenceRange) SetType(v string) *PatchEventResponseBodyRecurrenceRange {
	s.Type = &v
	return s
}

func (s *PatchEventResponseBodyRecurrenceRange) Validate() error {
	return dara.Validate(s)
}

type PatchEventResponseBodyReminders struct {
	// example:
	//
	// dingtalk
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// 15
	Minutes *string `json:"Minutes,omitempty" xml:"Minutes,omitempty"`
}

func (s PatchEventResponseBodyReminders) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyReminders) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyReminders) GetMethod() *string {
	return s.Method
}

func (s *PatchEventResponseBodyReminders) GetMinutes() *string {
	return s.Minutes
}

func (s *PatchEventResponseBodyReminders) SetMethod(v string) *PatchEventResponseBodyReminders {
	s.Method = &v
	return s
}

func (s *PatchEventResponseBodyReminders) SetMinutes(v string) *PatchEventResponseBodyReminders {
	s.Minutes = &v
	return s
}

func (s *PatchEventResponseBodyReminders) Validate() error {
	return dara.Validate(s)
}

type PatchEventResponseBodyStart struct {
	// example:
	//
	// 2020-01-01
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s PatchEventResponseBodyStart) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponseBodyStart) GoString() string {
	return s.String()
}

func (s *PatchEventResponseBodyStart) GetDate() *string {
	return s.Date
}

func (s *PatchEventResponseBodyStart) GetDateTime() *string {
	return s.DateTime
}

func (s *PatchEventResponseBodyStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *PatchEventResponseBodyStart) SetDate(v string) *PatchEventResponseBodyStart {
	s.Date = &v
	return s
}

func (s *PatchEventResponseBodyStart) SetDateTime(v string) *PatchEventResponseBodyStart {
	s.DateTime = &v
	return s
}

func (s *PatchEventResponseBodyStart) SetTimeZone(v string) *PatchEventResponseBodyStart {
	s.TimeZone = &v
	return s
}

func (s *PatchEventResponseBodyStart) Validate() error {
	return dara.Validate(s)
}
