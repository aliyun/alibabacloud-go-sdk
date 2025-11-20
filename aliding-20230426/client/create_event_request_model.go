// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttendees(v []*CreateEventRequestAttendees) *CreateEventRequest
	GetAttendees() []*CreateEventRequestAttendees
	SetCardInstances(v []*CreateEventRequestCardInstances) *CreateEventRequest
	GetCardInstances() []*CreateEventRequestCardInstances
	SetDescription(v string) *CreateEventRequest
	GetDescription() *string
	SetEnd(v *CreateEventRequestEnd) *CreateEventRequest
	GetEnd() *CreateEventRequestEnd
	SetExtra(v map[string]*string) *CreateEventRequest
	GetExtra() map[string]*string
	SetIsAllDay(v bool) *CreateEventRequest
	GetIsAllDay() *bool
	SetLocation(v *CreateEventRequestLocation) *CreateEventRequest
	GetLocation() *CreateEventRequestLocation
	SetOnlineMeetingInfo(v *CreateEventRequestOnlineMeetingInfo) *CreateEventRequest
	GetOnlineMeetingInfo() *CreateEventRequestOnlineMeetingInfo
	SetRecurrence(v *CreateEventRequestRecurrence) *CreateEventRequest
	GetRecurrence() *CreateEventRequestRecurrence
	SetReminders(v []*CreateEventRequestReminders) *CreateEventRequest
	GetReminders() []*CreateEventRequestReminders
	SetRichTextDescription(v *CreateEventRequestRichTextDescription) *CreateEventRequest
	GetRichTextDescription() *CreateEventRequestRichTextDescription
	SetSummary(v string) *CreateEventRequest
	GetSummary() *string
	SetUiConfigs(v []*CreateEventRequestUiConfigs) *CreateEventRequest
	GetUiConfigs() []*CreateEventRequestUiConfigs
	SetCalendarId(v string) *CreateEventRequest
	GetCalendarId() *string
	SetStart(v *CreateEventRequestStart) *CreateEventRequest
	GetStart() *CreateEventRequestStart
}

type CreateEventRequest struct {
	Attendees     []*CreateEventRequestAttendees     `json:"Attendees,omitempty" xml:"Attendees,omitempty" type:"Repeated"`
	CardInstances []*CreateEventRequestCardInstances `json:"CardInstances,omitempty" xml:"CardInstances,omitempty" type:"Repeated"`
	Description   *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	End           *CreateEventRequestEnd             `json:"End,omitempty" xml:"End,omitempty" type:"Struct"`
	Extra         map[string]*string                 `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// true
	IsAllDay          *bool                                `json:"IsAllDay,omitempty" xml:"IsAllDay,omitempty"`
	Location          *CreateEventRequestLocation          `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	OnlineMeetingInfo *CreateEventRequestOnlineMeetingInfo `json:"OnlineMeetingInfo,omitempty" xml:"OnlineMeetingInfo,omitempty" type:"Struct"`
	Recurrence        *CreateEventRequestRecurrence        `json:"Recurrence,omitempty" xml:"Recurrence,omitempty" type:"Struct"`
	// if can be null:
	// false
	Reminders           []*CreateEventRequestReminders         `json:"Reminders,omitempty" xml:"Reminders,omitempty" type:"Repeated"`
	RichTextDescription *CreateEventRequestRichTextDescription `json:"RichTextDescription,omitempty" xml:"RichTextDescription,omitempty" type:"Struct"`
	// This parameter is required.
	Summary   *string                        `json:"Summary,omitempty" xml:"Summary,omitempty"`
	UiConfigs []*CreateEventRequestUiConfigs `json:"UiConfigs,omitempty" xml:"UiConfigs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// primary
	CalendarId *string `json:"calendarId,omitempty" xml:"calendarId,omitempty"`
	// This parameter is required.
	Start *CreateEventRequestStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s CreateEventRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRequest) GetAttendees() []*CreateEventRequestAttendees {
	return s.Attendees
}

func (s *CreateEventRequest) GetCardInstances() []*CreateEventRequestCardInstances {
	return s.CardInstances
}

func (s *CreateEventRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEventRequest) GetEnd() *CreateEventRequestEnd {
	return s.End
}

func (s *CreateEventRequest) GetExtra() map[string]*string {
	return s.Extra
}

func (s *CreateEventRequest) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *CreateEventRequest) GetLocation() *CreateEventRequestLocation {
	return s.Location
}

func (s *CreateEventRequest) GetOnlineMeetingInfo() *CreateEventRequestOnlineMeetingInfo {
	return s.OnlineMeetingInfo
}

func (s *CreateEventRequest) GetRecurrence() *CreateEventRequestRecurrence {
	return s.Recurrence
}

func (s *CreateEventRequest) GetReminders() []*CreateEventRequestReminders {
	return s.Reminders
}

func (s *CreateEventRequest) GetRichTextDescription() *CreateEventRequestRichTextDescription {
	return s.RichTextDescription
}

func (s *CreateEventRequest) GetSummary() *string {
	return s.Summary
}

func (s *CreateEventRequest) GetUiConfigs() []*CreateEventRequestUiConfigs {
	return s.UiConfigs
}

func (s *CreateEventRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *CreateEventRequest) GetStart() *CreateEventRequestStart {
	return s.Start
}

func (s *CreateEventRequest) SetAttendees(v []*CreateEventRequestAttendees) *CreateEventRequest {
	s.Attendees = v
	return s
}

func (s *CreateEventRequest) SetCardInstances(v []*CreateEventRequestCardInstances) *CreateEventRequest {
	s.CardInstances = v
	return s
}

func (s *CreateEventRequest) SetDescription(v string) *CreateEventRequest {
	s.Description = &v
	return s
}

func (s *CreateEventRequest) SetEnd(v *CreateEventRequestEnd) *CreateEventRequest {
	s.End = v
	return s
}

func (s *CreateEventRequest) SetExtra(v map[string]*string) *CreateEventRequest {
	s.Extra = v
	return s
}

func (s *CreateEventRequest) SetIsAllDay(v bool) *CreateEventRequest {
	s.IsAllDay = &v
	return s
}

func (s *CreateEventRequest) SetLocation(v *CreateEventRequestLocation) *CreateEventRequest {
	s.Location = v
	return s
}

func (s *CreateEventRequest) SetOnlineMeetingInfo(v *CreateEventRequestOnlineMeetingInfo) *CreateEventRequest {
	s.OnlineMeetingInfo = v
	return s
}

func (s *CreateEventRequest) SetRecurrence(v *CreateEventRequestRecurrence) *CreateEventRequest {
	s.Recurrence = v
	return s
}

func (s *CreateEventRequest) SetReminders(v []*CreateEventRequestReminders) *CreateEventRequest {
	s.Reminders = v
	return s
}

func (s *CreateEventRequest) SetRichTextDescription(v *CreateEventRequestRichTextDescription) *CreateEventRequest {
	s.RichTextDescription = v
	return s
}

func (s *CreateEventRequest) SetSummary(v string) *CreateEventRequest {
	s.Summary = &v
	return s
}

func (s *CreateEventRequest) SetUiConfigs(v []*CreateEventRequestUiConfigs) *CreateEventRequest {
	s.UiConfigs = v
	return s
}

func (s *CreateEventRequest) SetCalendarId(v string) *CreateEventRequest {
	s.CalendarId = &v
	return s
}

func (s *CreateEventRequest) SetStart(v *CreateEventRequestStart) *CreateEventRequest {
	s.Start = v
	return s
}

func (s *CreateEventRequest) Validate() error {
	if s.Attendees != nil {
		for _, item := range s.Attendees {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CardInstances != nil {
		for _, item := range s.CardInstances {
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
	if s.OnlineMeetingInfo != nil {
		if err := s.OnlineMeetingInfo.Validate(); err != nil {
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
	if s.RichTextDescription != nil {
		if err := s.RichTextDescription.Validate(); err != nil {
			return err
		}
	}
	if s.UiConfigs != nil {
		for _, item := range s.UiConfigs {
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

type CreateEventRequestAttendees struct {
	// example:
	//
	// 7845
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsOptional *bool `json:"isOptional,omitempty" xml:"isOptional,omitempty"`
}

func (s CreateEventRequestAttendees) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestAttendees) GoString() string {
	return s.String()
}

func (s *CreateEventRequestAttendees) GetId() *string {
	return s.Id
}

func (s *CreateEventRequestAttendees) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *CreateEventRequestAttendees) SetId(v string) *CreateEventRequestAttendees {
	s.Id = &v
	return s
}

func (s *CreateEventRequestAttendees) SetIsOptional(v bool) *CreateEventRequestAttendees {
	s.IsOptional = &v
	return s
}

func (s *CreateEventRequestAttendees) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestCardInstances struct {
	OutTrackId *string `json:"OutTrackId,omitempty" xml:"OutTrackId,omitempty"`
	Scenario   *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
}

func (s CreateEventRequestCardInstances) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestCardInstances) GoString() string {
	return s.String()
}

func (s *CreateEventRequestCardInstances) GetOutTrackId() *string {
	return s.OutTrackId
}

func (s *CreateEventRequestCardInstances) GetScenario() *string {
	return s.Scenario
}

func (s *CreateEventRequestCardInstances) SetOutTrackId(v string) *CreateEventRequestCardInstances {
	s.OutTrackId = &v
	return s
}

func (s *CreateEventRequestCardInstances) SetScenario(v string) *CreateEventRequestCardInstances {
	s.Scenario = &v
	return s
}

func (s *CreateEventRequestCardInstances) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestEnd struct {
	// example:
	//
	// "2020-09-21"
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// "2021-09-20T10:15:30+08:00"
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// example:
	//
	// "Asia/Shanghai"
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s CreateEventRequestEnd) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestEnd) GoString() string {
	return s.String()
}

func (s *CreateEventRequestEnd) GetDate() *string {
	return s.Date
}

func (s *CreateEventRequestEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *CreateEventRequestEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateEventRequestEnd) SetDate(v string) *CreateEventRequestEnd {
	s.Date = &v
	return s
}

func (s *CreateEventRequestEnd) SetDateTime(v string) *CreateEventRequestEnd {
	s.DateTime = &v
	return s
}

func (s *CreateEventRequestEnd) SetTimeZone(v string) *CreateEventRequestEnd {
	s.TimeZone = &v
	return s
}

func (s *CreateEventRequestEnd) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestLocation struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s CreateEventRequestLocation) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestLocation) GoString() string {
	return s.String()
}

func (s *CreateEventRequestLocation) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateEventRequestLocation) SetDisplayName(v string) *CreateEventRequestLocation {
	s.DisplayName = &v
	return s
}

func (s *CreateEventRequestLocation) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestOnlineMeetingInfo struct {
	// example:
	//
	// dingtalk
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventRequestOnlineMeetingInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *CreateEventRequestOnlineMeetingInfo) GetType() *string {
	return s.Type
}

func (s *CreateEventRequestOnlineMeetingInfo) SetType(v string) *CreateEventRequestOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *CreateEventRequestOnlineMeetingInfo) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestRecurrence struct {
	Pattern *CreateEventRequestRecurrencePattern `json:"pattern,omitempty" xml:"pattern,omitempty" type:"Struct"`
	Range   *CreateEventRequestRecurrenceRange   `json:"range,omitempty" xml:"range,omitempty" type:"Struct"`
}

func (s CreateEventRequestRecurrence) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestRecurrence) GoString() string {
	return s.String()
}

func (s *CreateEventRequestRecurrence) GetPattern() *CreateEventRequestRecurrencePattern {
	return s.Pattern
}

func (s *CreateEventRequestRecurrence) GetRange() *CreateEventRequestRecurrenceRange {
	return s.Range
}

func (s *CreateEventRequestRecurrence) SetPattern(v *CreateEventRequestRecurrencePattern) *CreateEventRequestRecurrence {
	s.Pattern = v
	return s
}

func (s *CreateEventRequestRecurrence) SetRange(v *CreateEventRequestRecurrenceRange) *CreateEventRequestRecurrence {
	s.Range = v
	return s
}

func (s *CreateEventRequestRecurrence) Validate() error {
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

type CreateEventRequestRecurrencePattern struct {
	// example:
	//
	// 1
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	// example:
	//
	// "monday"
	DaysOfWeek *string `json:"daysOfWeek,omitempty" xml:"daysOfWeek,omitempty"`
	// example:
	//
	// "last"
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// 1
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// "daily"
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventRequestRecurrencePattern) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestRecurrencePattern) GoString() string {
	return s.String()
}

func (s *CreateEventRequestRecurrencePattern) GetDayOfMonth() *int32 {
	return s.DayOfMonth
}

func (s *CreateEventRequestRecurrencePattern) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *CreateEventRequestRecurrencePattern) GetIndex() *string {
	return s.Index
}

func (s *CreateEventRequestRecurrencePattern) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateEventRequestRecurrencePattern) GetType() *string {
	return s.Type
}

func (s *CreateEventRequestRecurrencePattern) SetDayOfMonth(v int32) *CreateEventRequestRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) SetDaysOfWeek(v string) *CreateEventRequestRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) SetIndex(v string) *CreateEventRequestRecurrencePattern {
	s.Index = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) SetInterval(v int32) *CreateEventRequestRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) SetType(v string) *CreateEventRequestRecurrencePattern {
	s.Type = &v
	return s
}

func (s *CreateEventRequestRecurrencePattern) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestRecurrenceRange struct {
	// example:
	//
	// "2021-12-31T10:15:30+08:00"
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 5
	NumberOfOccurrences *int32 `json:"numberOfOccurrences,omitempty" xml:"numberOfOccurrences,omitempty"`
	// example:
	//
	// "endDate"
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventRequestRecurrenceRange) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestRecurrenceRange) GoString() string {
	return s.String()
}

func (s *CreateEventRequestRecurrenceRange) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateEventRequestRecurrenceRange) GetNumberOfOccurrences() *int32 {
	return s.NumberOfOccurrences
}

func (s *CreateEventRequestRecurrenceRange) GetType() *string {
	return s.Type
}

func (s *CreateEventRequestRecurrenceRange) SetEndDate(v string) *CreateEventRequestRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *CreateEventRequestRecurrenceRange) SetNumberOfOccurrences(v int32) *CreateEventRequestRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *CreateEventRequestRecurrenceRange) SetType(v string) *CreateEventRequestRecurrenceRange {
	s.Type = &v
	return s
}

func (s *CreateEventRequestRecurrenceRange) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestReminders struct {
	// example:
	//
	// dingtalk
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// 15
	Minutes *int32 `json:"minutes,omitempty" xml:"minutes,omitempty"`
}

func (s CreateEventRequestReminders) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestReminders) GoString() string {
	return s.String()
}

func (s *CreateEventRequestReminders) GetMethod() *string {
	return s.Method
}

func (s *CreateEventRequestReminders) GetMinutes() *int32 {
	return s.Minutes
}

func (s *CreateEventRequestReminders) SetMethod(v string) *CreateEventRequestReminders {
	s.Method = &v
	return s
}

func (s *CreateEventRequestReminders) SetMinutes(v int32) *CreateEventRequestReminders {
	s.Minutes = &v
	return s
}

func (s *CreateEventRequestReminders) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestRichTextDescription struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CreateEventRequestRichTextDescription) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestRichTextDescription) GoString() string {
	return s.String()
}

func (s *CreateEventRequestRichTextDescription) GetText() *string {
	return s.Text
}

func (s *CreateEventRequestRichTextDescription) SetText(v string) *CreateEventRequestRichTextDescription {
	s.Text = &v
	return s
}

func (s *CreateEventRequestRichTextDescription) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestUiConfigs struct {
	// example:
	//
	// "updateEventButton"
	UiName *string `json:"uiName,omitempty" xml:"uiName,omitempty"`
	// example:
	//
	// "hide"
	UiStatus *string `json:"uiStatus,omitempty" xml:"uiStatus,omitempty"`
}

func (s CreateEventRequestUiConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestUiConfigs) GoString() string {
	return s.String()
}

func (s *CreateEventRequestUiConfigs) GetUiName() *string {
	return s.UiName
}

func (s *CreateEventRequestUiConfigs) GetUiStatus() *string {
	return s.UiStatus
}

func (s *CreateEventRequestUiConfigs) SetUiName(v string) *CreateEventRequestUiConfigs {
	s.UiName = &v
	return s
}

func (s *CreateEventRequestUiConfigs) SetUiStatus(v string) *CreateEventRequestUiConfigs {
	s.UiStatus = &v
	return s
}

func (s *CreateEventRequestUiConfigs) Validate() error {
	return dara.Validate(s)
}

type CreateEventRequestStart struct {
	// example:
	//
	// "2021-09-20"
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// "2021-09-20T10:15:30+08:00"
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// example:
	//
	// "Asia/Shanghai"
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s CreateEventRequestStart) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequestStart) GoString() string {
	return s.String()
}

func (s *CreateEventRequestStart) GetDate() *string {
	return s.Date
}

func (s *CreateEventRequestStart) GetDateTime() *string {
	return s.DateTime
}

func (s *CreateEventRequestStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateEventRequestStart) SetDate(v string) *CreateEventRequestStart {
	s.Date = &v
	return s
}

func (s *CreateEventRequestStart) SetDateTime(v string) *CreateEventRequestStart {
	s.DateTime = &v
	return s
}

func (s *CreateEventRequestStart) SetTimeZone(v string) *CreateEventRequestStart {
	s.TimeZone = &v
	return s
}

func (s *CreateEventRequestStart) Validate() error {
	return dara.Validate(s)
}
