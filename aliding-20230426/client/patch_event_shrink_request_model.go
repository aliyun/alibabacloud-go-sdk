// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchEventShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttendeesShrink(v string) *PatchEventShrinkRequest
	GetAttendeesShrink() *string
	SetCalendarId(v string) *PatchEventShrinkRequest
	GetCalendarId() *string
	SetCardInstancesShrink(v string) *PatchEventShrinkRequest
	GetCardInstancesShrink() *string
	SetDescription(v string) *PatchEventShrinkRequest
	GetDescription() *string
	SetEndShrink(v string) *PatchEventShrinkRequest
	GetEndShrink() *string
	SetEventId(v string) *PatchEventShrinkRequest
	GetEventId() *string
	SetExtraShrink(v string) *PatchEventShrinkRequest
	GetExtraShrink() *string
	SetIsAllDay(v bool) *PatchEventShrinkRequest
	GetIsAllDay() *bool
	SetLocationShrink(v string) *PatchEventShrinkRequest
	GetLocationShrink() *string
	SetRecurrenceShrink(v string) *PatchEventShrinkRequest
	GetRecurrenceShrink() *string
	SetRemindersShrink(v string) *PatchEventShrinkRequest
	GetRemindersShrink() *string
	SetStartShrink(v string) *PatchEventShrinkRequest
	GetStartShrink() *string
	SetSummary(v string) *PatchEventShrinkRequest
	GetSummary() *string
}

type PatchEventShrinkRequest struct {
	AttendeesShrink *string `json:"Attendees,omitempty" xml:"Attendees,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// primary
	CalendarId          *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
	CardInstancesShrink *string `json:"CardInstances,omitempty" xml:"CardInstances,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndShrink           *string `json:"End,omitempty" xml:"End,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iiiP35sJadba8aBSgjrwPRKgiEiF
	EventId     *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// true
	IsAllDay         *bool   `json:"IsAllDay,omitempty" xml:"IsAllDay,omitempty"`
	LocationShrink   *string `json:"Location,omitempty" xml:"Location,omitempty"`
	RecurrenceShrink *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	RemindersShrink  *string `json:"Reminders,omitempty" xml:"Reminders,omitempty"`
	StartShrink      *string `json:"Start,omitempty" xml:"Start,omitempty"`
	Summary          *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s PatchEventShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchEventShrinkRequest) GoString() string {
	return s.String()
}

func (s *PatchEventShrinkRequest) GetAttendeesShrink() *string {
	return s.AttendeesShrink
}

func (s *PatchEventShrinkRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *PatchEventShrinkRequest) GetCardInstancesShrink() *string {
	return s.CardInstancesShrink
}

func (s *PatchEventShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *PatchEventShrinkRequest) GetEndShrink() *string {
	return s.EndShrink
}

func (s *PatchEventShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *PatchEventShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *PatchEventShrinkRequest) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *PatchEventShrinkRequest) GetLocationShrink() *string {
	return s.LocationShrink
}

func (s *PatchEventShrinkRequest) GetRecurrenceShrink() *string {
	return s.RecurrenceShrink
}

func (s *PatchEventShrinkRequest) GetRemindersShrink() *string {
	return s.RemindersShrink
}

func (s *PatchEventShrinkRequest) GetStartShrink() *string {
	return s.StartShrink
}

func (s *PatchEventShrinkRequest) GetSummary() *string {
	return s.Summary
}

func (s *PatchEventShrinkRequest) SetAttendeesShrink(v string) *PatchEventShrinkRequest {
	s.AttendeesShrink = &v
	return s
}

func (s *PatchEventShrinkRequest) SetCalendarId(v string) *PatchEventShrinkRequest {
	s.CalendarId = &v
	return s
}

func (s *PatchEventShrinkRequest) SetCardInstancesShrink(v string) *PatchEventShrinkRequest {
	s.CardInstancesShrink = &v
	return s
}

func (s *PatchEventShrinkRequest) SetDescription(v string) *PatchEventShrinkRequest {
	s.Description = &v
	return s
}

func (s *PatchEventShrinkRequest) SetEndShrink(v string) *PatchEventShrinkRequest {
	s.EndShrink = &v
	return s
}

func (s *PatchEventShrinkRequest) SetEventId(v string) *PatchEventShrinkRequest {
	s.EventId = &v
	return s
}

func (s *PatchEventShrinkRequest) SetExtraShrink(v string) *PatchEventShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *PatchEventShrinkRequest) SetIsAllDay(v bool) *PatchEventShrinkRequest {
	s.IsAllDay = &v
	return s
}

func (s *PatchEventShrinkRequest) SetLocationShrink(v string) *PatchEventShrinkRequest {
	s.LocationShrink = &v
	return s
}

func (s *PatchEventShrinkRequest) SetRecurrenceShrink(v string) *PatchEventShrinkRequest {
	s.RecurrenceShrink = &v
	return s
}

func (s *PatchEventShrinkRequest) SetRemindersShrink(v string) *PatchEventShrinkRequest {
	s.RemindersShrink = &v
	return s
}

func (s *PatchEventShrinkRequest) SetStartShrink(v string) *PatchEventShrinkRequest {
	s.StartShrink = &v
	return s
}

func (s *PatchEventShrinkRequest) SetSummary(v string) *PatchEventShrinkRequest {
	s.Summary = &v
	return s
}

func (s *PatchEventShrinkRequest) Validate() error {
	return dara.Validate(s)
}
