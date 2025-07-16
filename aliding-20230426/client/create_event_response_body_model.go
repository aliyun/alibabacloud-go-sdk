// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttendees(v []*CreateEventResponseBodyAttendees) *CreateEventResponseBody
	GetAttendees() []*CreateEventResponseBodyAttendees
	SetCreateTime(v string) *CreateEventResponseBody
	GetCreateTime() *string
	SetDescription(v string) *CreateEventResponseBody
	GetDescription() *string
	SetEnd(v *CreateEventResponseBodyEnd) *CreateEventResponseBody
	GetEnd() *CreateEventResponseBodyEnd
	SetId(v string) *CreateEventResponseBody
	GetId() *string
	SetIsAllDay(v bool) *CreateEventResponseBody
	GetIsAllDay() *bool
	SetLocation(v *CreateEventResponseBodyLocation) *CreateEventResponseBody
	GetLocation() *CreateEventResponseBodyLocation
	SetOnlineMeetingInfo(v *CreateEventResponseBodyOnlineMeetingInfo) *CreateEventResponseBody
	GetOnlineMeetingInfo() *CreateEventResponseBodyOnlineMeetingInfo
	SetOrganizer(v *CreateEventResponseBodyOrganizer) *CreateEventResponseBody
	GetOrganizer() *CreateEventResponseBodyOrganizer
	SetRecurrence(v *CreateEventResponseBodyRecurrence) *CreateEventResponseBody
	GetRecurrence() *CreateEventResponseBodyRecurrence
	SetReminders(v []*CreateEventResponseBodyReminders) *CreateEventResponseBody
	GetReminders() []*CreateEventResponseBodyReminders
	SetRequestId(v string) *CreateEventResponseBody
	GetRequestId() *string
	SetRichTextDescription(v *CreateEventResponseBodyRichTextDescription) *CreateEventResponseBody
	GetRichTextDescription() *CreateEventResponseBodyRichTextDescription
	SetStart(v *CreateEventResponseBodyStart) *CreateEventResponseBody
	GetStart() *CreateEventResponseBodyStart
	SetSummary(v string) *CreateEventResponseBody
	GetSummary() *string
	SetUiConfigs(v []*CreateEventResponseBodyUiConfigs) *CreateEventResponseBody
	GetUiConfigs() []*CreateEventResponseBodyUiConfigs
	SetUpdateTime(v string) *CreateEventResponseBody
	GetUpdateTime() *string
}

type CreateEventResponseBody struct {
	Attendees []*CreateEventResponseBodyAttendees `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	CreateTime  *string                     `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string                     `json:"description,omitempty" xml:"description,omitempty"`
	End         *CreateEventResponseBodyEnd `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	// example:
	//
	// iiiP35sJadba8aBSgjrwPRKgiEiF
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsAllDay          *bool                                     `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	Location          *CreateEventResponseBodyLocation          `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	OnlineMeetingInfo *CreateEventResponseBodyOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty" xml:"onlineMeetingInfo,omitempty" type:"Struct"`
	Organizer         *CreateEventResponseBodyOrganizer         `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	Recurrence        *CreateEventResponseBodyRecurrence        `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders         []*CreateEventResponseBodyReminders       `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 4248DCC9-785F-5A14-8BE0-830FD52E1261
	RequestId           *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RichTextDescription *CreateEventResponseBodyRichTextDescription `json:"richTextDescription,omitempty" xml:"richTextDescription,omitempty" type:"Struct"`
	Start               *CreateEventResponseBodyStart               `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	Summary             *string                                     `json:"summary,omitempty" xml:"summary,omitempty"`
	UiConfigs           []*CreateEventResponseBodyUiConfigs         `json:"uiConfigs,omitempty" xml:"uiConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s CreateEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBody) GetAttendees() []*CreateEventResponseBodyAttendees {
	return s.Attendees
}

func (s *CreateEventResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateEventResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateEventResponseBody) GetEnd() *CreateEventResponseBodyEnd {
	return s.End
}

func (s *CreateEventResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateEventResponseBody) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *CreateEventResponseBody) GetLocation() *CreateEventResponseBodyLocation {
	return s.Location
}

func (s *CreateEventResponseBody) GetOnlineMeetingInfo() *CreateEventResponseBodyOnlineMeetingInfo {
	return s.OnlineMeetingInfo
}

func (s *CreateEventResponseBody) GetOrganizer() *CreateEventResponseBodyOrganizer {
	return s.Organizer
}

func (s *CreateEventResponseBody) GetRecurrence() *CreateEventResponseBodyRecurrence {
	return s.Recurrence
}

func (s *CreateEventResponseBody) GetReminders() []*CreateEventResponseBodyReminders {
	return s.Reminders
}

func (s *CreateEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventResponseBody) GetRichTextDescription() *CreateEventResponseBodyRichTextDescription {
	return s.RichTextDescription
}

func (s *CreateEventResponseBody) GetStart() *CreateEventResponseBodyStart {
	return s.Start
}

func (s *CreateEventResponseBody) GetSummary() *string {
	return s.Summary
}

func (s *CreateEventResponseBody) GetUiConfigs() []*CreateEventResponseBodyUiConfigs {
	return s.UiConfigs
}

func (s *CreateEventResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateEventResponseBody) SetAttendees(v []*CreateEventResponseBodyAttendees) *CreateEventResponseBody {
	s.Attendees = v
	return s
}

func (s *CreateEventResponseBody) SetCreateTime(v string) *CreateEventResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateEventResponseBody) SetDescription(v string) *CreateEventResponseBody {
	s.Description = &v
	return s
}

func (s *CreateEventResponseBody) SetEnd(v *CreateEventResponseBodyEnd) *CreateEventResponseBody {
	s.End = v
	return s
}

func (s *CreateEventResponseBody) SetId(v string) *CreateEventResponseBody {
	s.Id = &v
	return s
}

func (s *CreateEventResponseBody) SetIsAllDay(v bool) *CreateEventResponseBody {
	s.IsAllDay = &v
	return s
}

func (s *CreateEventResponseBody) SetLocation(v *CreateEventResponseBodyLocation) *CreateEventResponseBody {
	s.Location = v
	return s
}

func (s *CreateEventResponseBody) SetOnlineMeetingInfo(v *CreateEventResponseBodyOnlineMeetingInfo) *CreateEventResponseBody {
	s.OnlineMeetingInfo = v
	return s
}

func (s *CreateEventResponseBody) SetOrganizer(v *CreateEventResponseBodyOrganizer) *CreateEventResponseBody {
	s.Organizer = v
	return s
}

func (s *CreateEventResponseBody) SetRecurrence(v *CreateEventResponseBodyRecurrence) *CreateEventResponseBody {
	s.Recurrence = v
	return s
}

func (s *CreateEventResponseBody) SetReminders(v []*CreateEventResponseBodyReminders) *CreateEventResponseBody {
	s.Reminders = v
	return s
}

func (s *CreateEventResponseBody) SetRequestId(v string) *CreateEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventResponseBody) SetRichTextDescription(v *CreateEventResponseBodyRichTextDescription) *CreateEventResponseBody {
	s.RichTextDescription = v
	return s
}

func (s *CreateEventResponseBody) SetStart(v *CreateEventResponseBodyStart) *CreateEventResponseBody {
	s.Start = v
	return s
}

func (s *CreateEventResponseBody) SetSummary(v string) *CreateEventResponseBody {
	s.Summary = &v
	return s
}

func (s *CreateEventResponseBody) SetUiConfigs(v []*CreateEventResponseBodyUiConfigs) *CreateEventResponseBody {
	s.UiConfigs = v
	return s
}

func (s *CreateEventResponseBody) SetUpdateTime(v string) *CreateEventResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *CreateEventResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyAttendees struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// true
	IsOptional *bool `json:"IsOptional,omitempty" xml:"IsOptional,omitempty"`
	// example:
	//
	// needsAction
	ResponseStatus *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	// example:
	//
	// true
	Self *bool `json:"Self,omitempty" xml:"Self,omitempty"`
}

func (s CreateEventResponseBodyAttendees) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyAttendees) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyAttendees) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateEventResponseBodyAttendees) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *CreateEventResponseBodyAttendees) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *CreateEventResponseBodyAttendees) GetSelf() *bool {
	return s.Self
}

func (s *CreateEventResponseBodyAttendees) SetDisplayName(v string) *CreateEventResponseBodyAttendees {
	s.DisplayName = &v
	return s
}

func (s *CreateEventResponseBodyAttendees) SetIsOptional(v bool) *CreateEventResponseBodyAttendees {
	s.IsOptional = &v
	return s
}

func (s *CreateEventResponseBodyAttendees) SetResponseStatus(v string) *CreateEventResponseBodyAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *CreateEventResponseBodyAttendees) SetSelf(v bool) *CreateEventResponseBodyAttendees {
	s.Self = &v
	return s
}

func (s *CreateEventResponseBodyAttendees) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyEnd struct {
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

func (s CreateEventResponseBodyEnd) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyEnd) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyEnd) GetDate() *string {
	return s.Date
}

func (s *CreateEventResponseBodyEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *CreateEventResponseBodyEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateEventResponseBodyEnd) SetDate(v string) *CreateEventResponseBodyEnd {
	s.Date = &v
	return s
}

func (s *CreateEventResponseBodyEnd) SetDateTime(v string) *CreateEventResponseBodyEnd {
	s.DateTime = &v
	return s
}

func (s *CreateEventResponseBodyEnd) SetTimeZone(v string) *CreateEventResponseBodyEnd {
	s.TimeZone = &v
	return s
}

func (s *CreateEventResponseBodyEnd) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyLocation struct {
	// example:
	//
	// room 1-2-3
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s CreateEventResponseBodyLocation) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyLocation) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyLocation) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateEventResponseBodyLocation) SetDisplayName(v string) *CreateEventResponseBodyLocation {
	s.DisplayName = &v
	return s
}

func (s *CreateEventResponseBodyLocation) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyOnlineMeetingInfo struct {
	// example:
	//
	// 123
	ConferenceId *string                `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ExtraInfo    map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// dingtalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// http://meeting
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateEventResponseBodyOnlineMeetingInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) GetType() *string {
	return s.Type
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) GetUrl() *string {
	return s.Url
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) SetConferenceId(v string) *CreateEventResponseBodyOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) SetExtraInfo(v map[string]interface{}) *CreateEventResponseBodyOnlineMeetingInfo {
	s.ExtraInfo = v
	return s
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) SetType(v string) *CreateEventResponseBodyOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) SetUrl(v string) *CreateEventResponseBodyOnlineMeetingInfo {
	s.Url = &v
	return s
}

func (s *CreateEventResponseBodyOnlineMeetingInfo) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyOrganizer struct {
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

func (s CreateEventResponseBodyOrganizer) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyOrganizer) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyOrganizer) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateEventResponseBodyOrganizer) GetId() *string {
	return s.Id
}

func (s *CreateEventResponseBodyOrganizer) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *CreateEventResponseBodyOrganizer) GetSelf() *bool {
	return s.Self
}

func (s *CreateEventResponseBodyOrganizer) SetDisplayName(v string) *CreateEventResponseBodyOrganizer {
	s.DisplayName = &v
	return s
}

func (s *CreateEventResponseBodyOrganizer) SetId(v string) *CreateEventResponseBodyOrganizer {
	s.Id = &v
	return s
}

func (s *CreateEventResponseBodyOrganizer) SetResponseStatus(v string) *CreateEventResponseBodyOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *CreateEventResponseBodyOrganizer) SetSelf(v bool) *CreateEventResponseBodyOrganizer {
	s.Self = &v
	return s
}

func (s *CreateEventResponseBodyOrganizer) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyRecurrence struct {
	Pattern *CreateEventResponseBodyRecurrencePattern `json:"Pattern,omitempty" xml:"Pattern,omitempty" type:"Struct"`
	Range   *CreateEventResponseBodyRecurrenceRange   `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
}

func (s CreateEventResponseBodyRecurrence) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyRecurrence) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyRecurrence) GetPattern() *CreateEventResponseBodyRecurrencePattern {
	return s.Pattern
}

func (s *CreateEventResponseBodyRecurrence) GetRange() *CreateEventResponseBodyRecurrenceRange {
	return s.Range
}

func (s *CreateEventResponseBodyRecurrence) SetPattern(v *CreateEventResponseBodyRecurrencePattern) *CreateEventResponseBodyRecurrence {
	s.Pattern = v
	return s
}

func (s *CreateEventResponseBodyRecurrence) SetRange(v *CreateEventResponseBodyRecurrenceRange) *CreateEventResponseBodyRecurrence {
	s.Range = v
	return s
}

func (s *CreateEventResponseBodyRecurrence) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyRecurrencePattern struct {
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

func (s CreateEventResponseBodyRecurrencePattern) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyRecurrencePattern) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyRecurrencePattern) GetDayOfMonth() *int32 {
	return s.DayOfMonth
}

func (s *CreateEventResponseBodyRecurrencePattern) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *CreateEventResponseBodyRecurrencePattern) GetIndex() *string {
	return s.Index
}

func (s *CreateEventResponseBodyRecurrencePattern) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateEventResponseBodyRecurrencePattern) GetType() *string {
	return s.Type
}

func (s *CreateEventResponseBodyRecurrencePattern) SetDayOfMonth(v int32) *CreateEventResponseBodyRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) SetDaysOfWeek(v string) *CreateEventResponseBodyRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) SetIndex(v string) *CreateEventResponseBodyRecurrencePattern {
	s.Index = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) SetInterval(v int32) *CreateEventResponseBodyRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) SetType(v string) *CreateEventResponseBodyRecurrencePattern {
	s.Type = &v
	return s
}

func (s *CreateEventResponseBodyRecurrencePattern) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyRecurrenceRange struct {
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

func (s CreateEventResponseBodyRecurrenceRange) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyRecurrenceRange) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyRecurrenceRange) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateEventResponseBodyRecurrenceRange) GetNumberOfOccurrences() *int32 {
	return s.NumberOfOccurrences
}

func (s *CreateEventResponseBodyRecurrenceRange) GetType() *string {
	return s.Type
}

func (s *CreateEventResponseBodyRecurrenceRange) SetEndDate(v string) *CreateEventResponseBodyRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *CreateEventResponseBodyRecurrenceRange) SetNumberOfOccurrences(v int32) *CreateEventResponseBodyRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *CreateEventResponseBodyRecurrenceRange) SetType(v string) *CreateEventResponseBodyRecurrenceRange {
	s.Type = &v
	return s
}

func (s *CreateEventResponseBodyRecurrenceRange) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyReminders struct {
	// example:
	//
	// dingtalk
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// 15
	Minutes *string `json:"Minutes,omitempty" xml:"Minutes,omitempty"`
}

func (s CreateEventResponseBodyReminders) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyReminders) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyReminders) GetMethod() *string {
	return s.Method
}

func (s *CreateEventResponseBodyReminders) GetMinutes() *string {
	return s.Minutes
}

func (s *CreateEventResponseBodyReminders) SetMethod(v string) *CreateEventResponseBodyReminders {
	s.Method = &v
	return s
}

func (s *CreateEventResponseBodyReminders) SetMinutes(v string) *CreateEventResponseBodyReminders {
	s.Minutes = &v
	return s
}

func (s *CreateEventResponseBodyReminders) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyRichTextDescription struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CreateEventResponseBodyRichTextDescription) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyRichTextDescription) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyRichTextDescription) GetText() *string {
	return s.Text
}

func (s *CreateEventResponseBodyRichTextDescription) SetText(v string) *CreateEventResponseBodyRichTextDescription {
	s.Text = &v
	return s
}

func (s *CreateEventResponseBodyRichTextDescription) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyStart struct {
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

func (s CreateEventResponseBodyStart) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyStart) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyStart) GetDate() *string {
	return s.Date
}

func (s *CreateEventResponseBodyStart) GetDateTime() *string {
	return s.DateTime
}

func (s *CreateEventResponseBodyStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateEventResponseBodyStart) SetDate(v string) *CreateEventResponseBodyStart {
	s.Date = &v
	return s
}

func (s *CreateEventResponseBodyStart) SetDateTime(v string) *CreateEventResponseBodyStart {
	s.DateTime = &v
	return s
}

func (s *CreateEventResponseBodyStart) SetTimeZone(v string) *CreateEventResponseBodyStart {
	s.TimeZone = &v
	return s
}

func (s *CreateEventResponseBodyStart) Validate() error {
	return dara.Validate(s)
}

type CreateEventResponseBodyUiConfigs struct {
	// example:
	//
	// updateEventButton
	UiName *string `json:"UiName,omitempty" xml:"UiName,omitempty"`
	// example:
	//
	// hide
	UiStatus *string `json:"UiStatus,omitempty" xml:"UiStatus,omitempty"`
}

func (s CreateEventResponseBodyUiConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBodyUiConfigs) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBodyUiConfigs) GetUiName() *string {
	return s.UiName
}

func (s *CreateEventResponseBodyUiConfigs) GetUiStatus() *string {
	return s.UiStatus
}

func (s *CreateEventResponseBodyUiConfigs) SetUiName(v string) *CreateEventResponseBodyUiConfigs {
	s.UiName = &v
	return s
}

func (s *CreateEventResponseBodyUiConfigs) SetUiStatus(v string) *CreateEventResponseBodyUiConfigs {
	s.UiStatus = &v
	return s
}

func (s *CreateEventResponseBodyUiConfigs) Validate() error {
	return dara.Validate(s)
}
