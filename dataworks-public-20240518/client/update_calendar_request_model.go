// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllProjects(v bool) *UpdateCalendarRequest
	GetAllProjects() *bool
	SetDescription(v string) *UpdateCalendarRequest
	GetDescription() *string
	SetEndDate(v string) *UpdateCalendarRequest
	GetEndDate() *string
	SetExpirationReminder(v string) *UpdateCalendarRequest
	GetExpirationReminder() *string
	SetFallbackPolicy(v string) *UpdateCalendarRequest
	GetFallbackPolicy() *string
	SetId(v int64) *UpdateCalendarRequest
	GetId() *int64
	SetOverduePolicy(v string) *UpdateCalendarRequest
	GetOverduePolicy() *string
	SetOwner(v string) *UpdateCalendarRequest
	GetOwner() *string
	SetProjectIds(v []*int64) *UpdateCalendarRequest
	GetProjectIds() []*int64
	SetSchedulePolicy(v *UpdateCalendarRequestSchedulePolicy) *UpdateCalendarRequest
	GetSchedulePolicy() *UpdateCalendarRequestSchedulePolicy
	SetStartDate(v string) *UpdateCalendarRequest
	GetStartDate() *string
	SetType(v string) *UpdateCalendarRequest
	GetType() *string
}

type UpdateCalendarRequest struct {
	// Specifies whether the calendar applies to all workspaces. Default value: false. Valid values:
	//
	// - true: All workspaces.
	//
	// - false: The workspaces specified by ProjectIds.
	//
	// example:
	//
	// true
	AllProjects *bool `json:"AllProjects,omitempty" xml:"AllProjects,omitempty"`
	// The description of the scheduling calendar. The description must be 5 to 2,048 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Weekday scheduling calendar
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end date of the calendar validity period, in the `yyyy-mm-dd` format. Required when Type is set to StrongConstraint. The end date cannot be earlier than the start date, and the interval between the start date and the end date must be less than 5 years.
	//
	// example:
	//
	// 2026-10-09
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The expiration reminder. Required when Type is set to StrongConstraint. Valid values:
	//
	// - OneMonthBefore: One month in advance.
	//
	// - OneWeekBefore: One week in advance.
	//
	// - ThreeDaysBefore: Three days in advance.
	//
	// example:
	//
	// OneMonthBefore
	ExpirationReminder *string `json:"ExpirationReminder,omitempty" xml:"ExpirationReminder,omitempty"`
	// The default scheduling policy for dates that are not configured. Required when Type is set to WeakConstraint. Valid values:
	//
	// - Normal: Normal scheduling.
	//
	// - Pause: Pause scheduling.
	//
	// - Skip: Dry run.
	//
	// example:
	//
	// Normal
	FallbackPolicy *string `json:"FallbackPolicy,omitempty" xml:"FallbackPolicy,omitempty"`
	// The ID of the scheduling calendar.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The scheduling policy after the calendar expires. Required when Type is set to StrongConstraint. Valid values:
	//
	// - Normal: Normal scheduling.
	//
	// - Pause: Pause scheduling.
	//
	// - Skip: Dry run.
	//
	// example:
	//
	// Skip
	OverduePolicy *string `json:"OverduePolicy,omitempty" xml:"OverduePolicy,omitempty"`
	// The account ID of the owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The list of applicable workspace IDs.
	//
	// - Required and cannot be empty when AllProjects is set to false or left unspecified.
	//
	// - Leave empty or pass an empty array when AllProjects is set to true.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The date scheduling policy. For a strong constraint calendar, the normal scheduling dates and non-scheduling dates must cover every day within the validity period and cannot overlap.
	//
	// This parameter is required.
	SchedulePolicy *UpdateCalendarRequestSchedulePolicy `json:"SchedulePolicy,omitempty" xml:"SchedulePolicy,omitempty" type:"Struct"`
	// The start date of the calendar validity period, in the `yyyy-mm-dd` format. Required when Type is set to StrongConstraint.
	//
	// example:
	//
	// 2026-10-08
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The calendar type. Valid values:
	//
	// - StrongConstraint: Strong constraint calendar.
	//
	// - WeakConstraint: Weak constraint calendar.
	//
	// This parameter is required.
	//
	// example:
	//
	// StrongConstraint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCalendarRequest) GoString() string {
	return s.String()
}

func (s *UpdateCalendarRequest) GetAllProjects() *bool {
	return s.AllProjects
}

func (s *UpdateCalendarRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCalendarRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateCalendarRequest) GetExpirationReminder() *string {
	return s.ExpirationReminder
}

func (s *UpdateCalendarRequest) GetFallbackPolicy() *string {
	return s.FallbackPolicy
}

func (s *UpdateCalendarRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateCalendarRequest) GetOverduePolicy() *string {
	return s.OverduePolicy
}

func (s *UpdateCalendarRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateCalendarRequest) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *UpdateCalendarRequest) GetSchedulePolicy() *UpdateCalendarRequestSchedulePolicy {
	return s.SchedulePolicy
}

func (s *UpdateCalendarRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateCalendarRequest) GetType() *string {
	return s.Type
}

func (s *UpdateCalendarRequest) SetAllProjects(v bool) *UpdateCalendarRequest {
	s.AllProjects = &v
	return s
}

func (s *UpdateCalendarRequest) SetDescription(v string) *UpdateCalendarRequest {
	s.Description = &v
	return s
}

func (s *UpdateCalendarRequest) SetEndDate(v string) *UpdateCalendarRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateCalendarRequest) SetExpirationReminder(v string) *UpdateCalendarRequest {
	s.ExpirationReminder = &v
	return s
}

func (s *UpdateCalendarRequest) SetFallbackPolicy(v string) *UpdateCalendarRequest {
	s.FallbackPolicy = &v
	return s
}

func (s *UpdateCalendarRequest) SetId(v int64) *UpdateCalendarRequest {
	s.Id = &v
	return s
}

func (s *UpdateCalendarRequest) SetOverduePolicy(v string) *UpdateCalendarRequest {
	s.OverduePolicy = &v
	return s
}

func (s *UpdateCalendarRequest) SetOwner(v string) *UpdateCalendarRequest {
	s.Owner = &v
	return s
}

func (s *UpdateCalendarRequest) SetProjectIds(v []*int64) *UpdateCalendarRequest {
	s.ProjectIds = v
	return s
}

func (s *UpdateCalendarRequest) SetSchedulePolicy(v *UpdateCalendarRequestSchedulePolicy) *UpdateCalendarRequest {
	s.SchedulePolicy = v
	return s
}

func (s *UpdateCalendarRequest) SetStartDate(v string) *UpdateCalendarRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateCalendarRequest) SetType(v string) *UpdateCalendarRequest {
	s.Type = &v
	return s
}

func (s *UpdateCalendarRequest) Validate() error {
	if s.SchedulePolicy != nil {
		if err := s.SchedulePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCalendarRequestSchedulePolicy struct {
	// The list of normal scheduling dates.
	NormalDates []*string `json:"NormalDates,omitempty" xml:"NormalDates,omitempty" type:"Repeated"`
	// The list of non-scheduling dates. Tasks perform dry runs on these dates. Used only for strong constraint calendars.
	SkipDates []*string `json:"SkipDates,omitempty" xml:"SkipDates,omitempty" type:"Repeated"`
}

func (s UpdateCalendarRequestSchedulePolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateCalendarRequestSchedulePolicy) GoString() string {
	return s.String()
}

func (s *UpdateCalendarRequestSchedulePolicy) GetNormalDates() []*string {
	return s.NormalDates
}

func (s *UpdateCalendarRequestSchedulePolicy) GetSkipDates() []*string {
	return s.SkipDates
}

func (s *UpdateCalendarRequestSchedulePolicy) SetNormalDates(v []*string) *UpdateCalendarRequestSchedulePolicy {
	s.NormalDates = v
	return s
}

func (s *UpdateCalendarRequestSchedulePolicy) SetSkipDates(v []*string) *UpdateCalendarRequestSchedulePolicy {
	s.SkipDates = v
	return s
}

func (s *UpdateCalendarRequestSchedulePolicy) Validate() error {
	return dara.Validate(s)
}
