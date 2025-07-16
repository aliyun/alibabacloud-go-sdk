// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttendeesShrink(v string) *CreateEventShrinkRequest
	GetAttendeesShrink() *string
	SetCardInstancesShrink(v string) *CreateEventShrinkRequest
	GetCardInstancesShrink() *string
	SetDescription(v string) *CreateEventShrinkRequest
	GetDescription() *string
	SetEndShrink(v string) *CreateEventShrinkRequest
	GetEndShrink() *string
	SetExtraShrink(v string) *CreateEventShrinkRequest
	GetExtraShrink() *string
	SetIsAllDay(v bool) *CreateEventShrinkRequest
	GetIsAllDay() *bool
	SetLocationShrink(v string) *CreateEventShrinkRequest
	GetLocationShrink() *string
	SetOnlineMeetingInfoShrink(v string) *CreateEventShrinkRequest
	GetOnlineMeetingInfoShrink() *string
	SetRecurrenceShrink(v string) *CreateEventShrinkRequest
	GetRecurrenceShrink() *string
	SetRemindersShrink(v string) *CreateEventShrinkRequest
	GetRemindersShrink() *string
	SetRichTextDescriptionShrink(v string) *CreateEventShrinkRequest
	GetRichTextDescriptionShrink() *string
	SetSummary(v string) *CreateEventShrinkRequest
	GetSummary() *string
	SetUiConfigsShrink(v string) *CreateEventShrinkRequest
	GetUiConfigsShrink() *string
	SetCalendarId(v string) *CreateEventShrinkRequest
	GetCalendarId() *string
	SetStartShrink(v string) *CreateEventShrinkRequest
	GetStartShrink() *string
}

type CreateEventShrinkRequest struct {
	AttendeesShrink     *string `json:"Attendees,omitempty" xml:"Attendees,omitempty"`
	CardInstancesShrink *string `json:"CardInstances,omitempty" xml:"CardInstances,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndShrink           *string `json:"End,omitempty" xml:"End,omitempty"`
	ExtraShrink         *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// true
	IsAllDay                *bool   `json:"IsAllDay,omitempty" xml:"IsAllDay,omitempty"`
	LocationShrink          *string `json:"Location,omitempty" xml:"Location,omitempty"`
	OnlineMeetingInfoShrink *string `json:"OnlineMeetingInfo,omitempty" xml:"OnlineMeetingInfo,omitempty"`
	RecurrenceShrink        *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// if can be null:
	// false
	RemindersShrink           *string `json:"Reminders,omitempty" xml:"Reminders,omitempty"`
	RichTextDescriptionShrink *string `json:"RichTextDescription,omitempty" xml:"RichTextDescription,omitempty"`
	// This parameter is required.
	Summary         *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	UiConfigsShrink *string `json:"UiConfigs,omitempty" xml:"UiConfigs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// primary
	CalendarId *string `json:"calendarId,omitempty" xml:"calendarId,omitempty"`
	// This parameter is required.
	StartShrink *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s CreateEventShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEventShrinkRequest) GetAttendeesShrink() *string {
	return s.AttendeesShrink
}

func (s *CreateEventShrinkRequest) GetCardInstancesShrink() *string {
	return s.CardInstancesShrink
}

func (s *CreateEventShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEventShrinkRequest) GetEndShrink() *string {
	return s.EndShrink
}

func (s *CreateEventShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *CreateEventShrinkRequest) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *CreateEventShrinkRequest) GetLocationShrink() *string {
	return s.LocationShrink
}

func (s *CreateEventShrinkRequest) GetOnlineMeetingInfoShrink() *string {
	return s.OnlineMeetingInfoShrink
}

func (s *CreateEventShrinkRequest) GetRecurrenceShrink() *string {
	return s.RecurrenceShrink
}

func (s *CreateEventShrinkRequest) GetRemindersShrink() *string {
	return s.RemindersShrink
}

func (s *CreateEventShrinkRequest) GetRichTextDescriptionShrink() *string {
	return s.RichTextDescriptionShrink
}

func (s *CreateEventShrinkRequest) GetSummary() *string {
	return s.Summary
}

func (s *CreateEventShrinkRequest) GetUiConfigsShrink() *string {
	return s.UiConfigsShrink
}

func (s *CreateEventShrinkRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *CreateEventShrinkRequest) GetStartShrink() *string {
	return s.StartShrink
}

func (s *CreateEventShrinkRequest) SetAttendeesShrink(v string) *CreateEventShrinkRequest {
	s.AttendeesShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetCardInstancesShrink(v string) *CreateEventShrinkRequest {
	s.CardInstancesShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetDescription(v string) *CreateEventShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateEventShrinkRequest) SetEndShrink(v string) *CreateEventShrinkRequest {
	s.EndShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetExtraShrink(v string) *CreateEventShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetIsAllDay(v bool) *CreateEventShrinkRequest {
	s.IsAllDay = &v
	return s
}

func (s *CreateEventShrinkRequest) SetLocationShrink(v string) *CreateEventShrinkRequest {
	s.LocationShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetOnlineMeetingInfoShrink(v string) *CreateEventShrinkRequest {
	s.OnlineMeetingInfoShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetRecurrenceShrink(v string) *CreateEventShrinkRequest {
	s.RecurrenceShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetRemindersShrink(v string) *CreateEventShrinkRequest {
	s.RemindersShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetRichTextDescriptionShrink(v string) *CreateEventShrinkRequest {
	s.RichTextDescriptionShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetSummary(v string) *CreateEventShrinkRequest {
	s.Summary = &v
	return s
}

func (s *CreateEventShrinkRequest) SetUiConfigsShrink(v string) *CreateEventShrinkRequest {
	s.UiConfigsShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) SetCalendarId(v string) *CreateEventShrinkRequest {
	s.CalendarId = &v
	return s
}

func (s *CreateEventShrinkRequest) SetStartShrink(v string) *CreateEventShrinkRequest {
	s.StartShrink = &v
	return s
}

func (s *CreateEventShrinkRequest) Validate() error {
	return dara.Validate(s)
}
