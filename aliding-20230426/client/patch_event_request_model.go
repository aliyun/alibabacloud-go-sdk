// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttendees(v []*PatchEventRequestAttendees) *PatchEventRequest
	GetAttendees() []*PatchEventRequestAttendees
	SetCalendarId(v string) *PatchEventRequest
	GetCalendarId() *string
	SetCardInstances(v []*PatchEventRequestCardInstances) *PatchEventRequest
	GetCardInstances() []*PatchEventRequestCardInstances
	SetDescription(v string) *PatchEventRequest
	GetDescription() *string
	SetEnd(v *PatchEventRequestEnd) *PatchEventRequest
	GetEnd() *PatchEventRequestEnd
	SetEventId(v string) *PatchEventRequest
	GetEventId() *string
	SetExtra(v map[string]*string) *PatchEventRequest
	GetExtra() map[string]*string
	SetIsAllDay(v bool) *PatchEventRequest
	GetIsAllDay() *bool
	SetLocation(v *PatchEventRequestLocation) *PatchEventRequest
	GetLocation() *PatchEventRequestLocation
	SetRecurrence(v *PatchEventRequestRecurrence) *PatchEventRequest
	GetRecurrence() *PatchEventRequestRecurrence
	SetReminders(v []*PatchEventRequestReminders) *PatchEventRequest
	GetReminders() []*PatchEventRequestReminders
	SetStart(v *PatchEventRequestStart) *PatchEventRequest
	GetStart() *PatchEventRequestStart
	SetSummary(v string) *PatchEventRequest
	GetSummary() *string
}

type PatchEventRequest struct {
	Attendees []*PatchEventRequestAttendees `json:"Attendees,omitempty" xml:"Attendees,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// primary
	CalendarId    *string                           `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
	CardInstances []*PatchEventRequestCardInstances `json:"CardInstances,omitempty" xml:"CardInstances,omitempty" type:"Repeated"`
	Description   *string                           `json:"Description,omitempty" xml:"Description,omitempty"`
	End           *PatchEventRequestEnd             `json:"End,omitempty" xml:"End,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// iiiP35sJadba8aBSgjrwPRKgiEiF
	EventId *string            `json:"EventId,omitempty" xml:"EventId,omitempty"`
	Extra   map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// true
	IsAllDay   *bool                         `json:"IsAllDay,omitempty" xml:"IsAllDay,omitempty"`
	Location   *PatchEventRequestLocation    `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	Recurrence *PatchEventRequestRecurrence  `json:"Recurrence,omitempty" xml:"Recurrence,omitempty" type:"Struct"`
	Reminders  []*PatchEventRequestReminders `json:"Reminders,omitempty" xml:"Reminders,omitempty" type:"Repeated"`
	Start      *PatchEventRequestStart       `json:"Start,omitempty" xml:"Start,omitempty" type:"Struct"`
	Summary    *string                       `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s PatchEventRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequest) GoString() string {
	return s.String()
}

func (s *PatchEventRequest) GetAttendees() []*PatchEventRequestAttendees {
	return s.Attendees
}

func (s *PatchEventRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *PatchEventRequest) GetCardInstances() []*PatchEventRequestCardInstances {
	return s.CardInstances
}

func (s *PatchEventRequest) GetDescription() *string {
	return s.Description
}

func (s *PatchEventRequest) GetEnd() *PatchEventRequestEnd {
	return s.End
}

func (s *PatchEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *PatchEventRequest) GetExtra() map[string]*string {
	return s.Extra
}

func (s *PatchEventRequest) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *PatchEventRequest) GetLocation() *PatchEventRequestLocation {
	return s.Location
}

func (s *PatchEventRequest) GetRecurrence() *PatchEventRequestRecurrence {
	return s.Recurrence
}

func (s *PatchEventRequest) GetReminders() []*PatchEventRequestReminders {
	return s.Reminders
}

func (s *PatchEventRequest) GetStart() *PatchEventRequestStart {
	return s.Start
}

func (s *PatchEventRequest) GetSummary() *string {
	return s.Summary
}

func (s *PatchEventRequest) SetAttendees(v []*PatchEventRequestAttendees) *PatchEventRequest {
	s.Attendees = v
	return s
}

func (s *PatchEventRequest) SetCalendarId(v string) *PatchEventRequest {
	s.CalendarId = &v
	return s
}

func (s *PatchEventRequest) SetCardInstances(v []*PatchEventRequestCardInstances) *PatchEventRequest {
	s.CardInstances = v
	return s
}

func (s *PatchEventRequest) SetDescription(v string) *PatchEventRequest {
	s.Description = &v
	return s
}

func (s *PatchEventRequest) SetEnd(v *PatchEventRequestEnd) *PatchEventRequest {
	s.End = v
	return s
}

func (s *PatchEventRequest) SetEventId(v string) *PatchEventRequest {
	s.EventId = &v
	return s
}

func (s *PatchEventRequest) SetExtra(v map[string]*string) *PatchEventRequest {
	s.Extra = v
	return s
}

func (s *PatchEventRequest) SetIsAllDay(v bool) *PatchEventRequest {
	s.IsAllDay = &v
	return s
}

func (s *PatchEventRequest) SetLocation(v *PatchEventRequestLocation) *PatchEventRequest {
	s.Location = v
	return s
}

func (s *PatchEventRequest) SetRecurrence(v *PatchEventRequestRecurrence) *PatchEventRequest {
	s.Recurrence = v
	return s
}

func (s *PatchEventRequest) SetReminders(v []*PatchEventRequestReminders) *PatchEventRequest {
	s.Reminders = v
	return s
}

func (s *PatchEventRequest) SetStart(v *PatchEventRequestStart) *PatchEventRequest {
	s.Start = v
	return s
}

func (s *PatchEventRequest) SetSummary(v string) *PatchEventRequest {
	s.Summary = &v
	return s
}

func (s *PatchEventRequest) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestAttendees struct {
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// false
	IsOptional *bool `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
}

func (s PatchEventRequestAttendees) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestAttendees) GoString() string {
	return s.String()
}

func (s *PatchEventRequestAttendees) GetId() *string {
	return s.Id
}

func (s *PatchEventRequestAttendees) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *PatchEventRequestAttendees) SetId(v string) *PatchEventRequestAttendees {
	s.Id = &v
	return s
}

func (s *PatchEventRequestAttendees) SetIsOptional(v bool) *PatchEventRequestAttendees {
	s.IsOptional = &v
	return s
}

func (s *PatchEventRequestAttendees) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestCardInstances struct {
	OutTrackId *string `json:"OutTrackId,omitempty" xml:"OutTrackId,omitempty"`
	Scenario   *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
}

func (s PatchEventRequestCardInstances) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestCardInstances) GoString() string {
	return s.String()
}

func (s *PatchEventRequestCardInstances) GetOutTrackId() *string {
	return s.OutTrackId
}

func (s *PatchEventRequestCardInstances) GetScenario() *string {
	return s.Scenario
}

func (s *PatchEventRequestCardInstances) SetOutTrackId(v string) *PatchEventRequestCardInstances {
	s.OutTrackId = &v
	return s
}

func (s *PatchEventRequestCardInstances) SetScenario(v string) *PatchEventRequestCardInstances {
	s.Scenario = &v
	return s
}

func (s *PatchEventRequestCardInstances) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestEnd struct {
	// example:
	//
	// 2020-01-01
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s PatchEventRequestEnd) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestEnd) GoString() string {
	return s.String()
}

func (s *PatchEventRequestEnd) GetDate() *string {
	return s.Date
}

func (s *PatchEventRequestEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *PatchEventRequestEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *PatchEventRequestEnd) SetDate(v string) *PatchEventRequestEnd {
	s.Date = &v
	return s
}

func (s *PatchEventRequestEnd) SetDateTime(v string) *PatchEventRequestEnd {
	s.DateTime = &v
	return s
}

func (s *PatchEventRequestEnd) SetTimeZone(v string) *PatchEventRequestEnd {
	s.TimeZone = &v
	return s
}

func (s *PatchEventRequestEnd) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestLocation struct {
	// example:
	//
	// room 1-2-3
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s PatchEventRequestLocation) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestLocation) GoString() string {
	return s.String()
}

func (s *PatchEventRequestLocation) GetDisplayName() *string {
	return s.DisplayName
}

func (s *PatchEventRequestLocation) SetDisplayName(v string) *PatchEventRequestLocation {
	s.DisplayName = &v
	return s
}

func (s *PatchEventRequestLocation) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestRecurrence struct {
	Pattern *PatchEventRequestRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	Range   *PatchEventRequestRecurrenceRange   `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s PatchEventRequestRecurrence) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestRecurrence) GoString() string {
	return s.String()
}

func (s *PatchEventRequestRecurrence) GetPattern() *PatchEventRequestRecurrencePattern {
	return s.Pattern
}

func (s *PatchEventRequestRecurrence) GetRange() *PatchEventRequestRecurrenceRange {
	return s.Range
}

func (s *PatchEventRequestRecurrence) SetPattern(v *PatchEventRequestRecurrencePattern) *PatchEventRequestRecurrence {
	s.Pattern = v
	return s
}

func (s *PatchEventRequestRecurrence) SetRange(v *PatchEventRequestRecurrenceRange) *PatchEventRequestRecurrence {
	s.Range = v
	return s
}

func (s *PatchEventRequestRecurrence) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestRecurrencePattern struct {
	// example:
	//
	// 1
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	// example:
	//
	// sunday
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	// example:
	//
	// last
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// 1
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// daily
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PatchEventRequestRecurrencePattern) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestRecurrencePattern) GoString() string {
	return s.String()
}

func (s *PatchEventRequestRecurrencePattern) GetDayOfMonth() *int32 {
	return s.DayOfMonth
}

func (s *PatchEventRequestRecurrencePattern) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *PatchEventRequestRecurrencePattern) GetIndex() *string {
	return s.Index
}

func (s *PatchEventRequestRecurrencePattern) GetInterval() *int32 {
	return s.Interval
}

func (s *PatchEventRequestRecurrencePattern) GetType() *string {
	return s.Type
}

func (s *PatchEventRequestRecurrencePattern) SetDayOfMonth(v int32) *PatchEventRequestRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) SetDaysOfWeek(v string) *PatchEventRequestRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) SetIndex(v string) *PatchEventRequestRecurrencePattern {
	s.Index = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) SetInterval(v int32) *PatchEventRequestRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) SetType(v string) *PatchEventRequestRecurrencePattern {
	s.Type = &v
	return s
}

func (s *PatchEventRequestRecurrencePattern) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestRecurrenceRange struct {
	// example:
	//
	// 2021-12-31T10:15:30+08:00
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 1
	NumberOfOccurrences *int32 `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	// example:
	//
	// endDate
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PatchEventRequestRecurrenceRange) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestRecurrenceRange) GoString() string {
	return s.String()
}

func (s *PatchEventRequestRecurrenceRange) GetEndDate() *string {
	return s.EndDate
}

func (s *PatchEventRequestRecurrenceRange) GetNumberOfOccurrences() *int32 {
	return s.NumberOfOccurrences
}

func (s *PatchEventRequestRecurrenceRange) GetType() *string {
	return s.Type
}

func (s *PatchEventRequestRecurrenceRange) SetEndDate(v string) *PatchEventRequestRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *PatchEventRequestRecurrenceRange) SetNumberOfOccurrences(v int32) *PatchEventRequestRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *PatchEventRequestRecurrenceRange) SetType(v string) *PatchEventRequestRecurrenceRange {
	s.Type = &v
	return s
}

func (s *PatchEventRequestRecurrenceRange) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestReminders struct {
	// example:
	//
	// dingtalk
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// 15
	Minutes *int32 `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s PatchEventRequestReminders) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestReminders) GoString() string {
	return s.String()
}

func (s *PatchEventRequestReminders) GetMethod() *string {
	return s.Method
}

func (s *PatchEventRequestReminders) GetMinutes() *int32 {
	return s.Minutes
}

func (s *PatchEventRequestReminders) SetMethod(v string) *PatchEventRequestReminders {
	s.Method = &v
	return s
}

func (s *PatchEventRequestReminders) SetMinutes(v int32) *PatchEventRequestReminders {
	s.Minutes = &v
	return s
}

func (s *PatchEventRequestReminders) Validate() error {
	return dara.Validate(s)
}

type PatchEventRequestStart struct {
	// example:
	//
	// 2020-01-01
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s PatchEventRequestStart) String() string {
	return dara.Prettify(s)
}

func (s PatchEventRequestStart) GoString() string {
	return s.String()
}

func (s *PatchEventRequestStart) GetDate() *string {
	return s.Date
}

func (s *PatchEventRequestStart) GetDateTime() *string {
	return s.DateTime
}

func (s *PatchEventRequestStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *PatchEventRequestStart) SetDate(v string) *PatchEventRequestStart {
	s.Date = &v
	return s
}

func (s *PatchEventRequestStart) SetDateTime(v string) *PatchEventRequestStart {
	s.DateTime = &v
	return s
}

func (s *PatchEventRequestStart) SetTimeZone(v string) *PatchEventRequestStart {
	s.TimeZone = &v
	return s
}

func (s *PatchEventRequestStart) Validate() error {
	return dara.Validate(s)
}
