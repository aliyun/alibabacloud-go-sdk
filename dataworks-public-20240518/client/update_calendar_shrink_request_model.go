// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCalendarShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllProjects(v bool) *UpdateCalendarShrinkRequest
	GetAllProjects() *bool
	SetDescription(v string) *UpdateCalendarShrinkRequest
	GetDescription() *string
	SetEndDate(v string) *UpdateCalendarShrinkRequest
	GetEndDate() *string
	SetExpirationReminder(v string) *UpdateCalendarShrinkRequest
	GetExpirationReminder() *string
	SetFallbackPolicy(v string) *UpdateCalendarShrinkRequest
	GetFallbackPolicy() *string
	SetId(v int64) *UpdateCalendarShrinkRequest
	GetId() *int64
	SetOverduePolicy(v string) *UpdateCalendarShrinkRequest
	GetOverduePolicy() *string
	SetOwner(v string) *UpdateCalendarShrinkRequest
	GetOwner() *string
	SetProjectIdsShrink(v string) *UpdateCalendarShrinkRequest
	GetProjectIdsShrink() *string
	SetSchedulePolicyShrink(v string) *UpdateCalendarShrinkRequest
	GetSchedulePolicyShrink() *string
	SetStartDate(v string) *UpdateCalendarShrinkRequest
	GetStartDate() *string
	SetType(v string) *UpdateCalendarShrinkRequest
	GetType() *string
}

type UpdateCalendarShrinkRequest struct {
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
	ProjectIdsShrink *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// The date scheduling policy. For a strong constraint calendar, the normal scheduling dates and non-scheduling dates must cover every day within the validity period and cannot overlap.
	//
	// This parameter is required.
	SchedulePolicyShrink *string `json:"SchedulePolicy,omitempty" xml:"SchedulePolicy,omitempty"`
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

func (s UpdateCalendarShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCalendarShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCalendarShrinkRequest) GetAllProjects() *bool {
	return s.AllProjects
}

func (s *UpdateCalendarShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCalendarShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateCalendarShrinkRequest) GetExpirationReminder() *string {
	return s.ExpirationReminder
}

func (s *UpdateCalendarShrinkRequest) GetFallbackPolicy() *string {
	return s.FallbackPolicy
}

func (s *UpdateCalendarShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateCalendarShrinkRequest) GetOverduePolicy() *string {
	return s.OverduePolicy
}

func (s *UpdateCalendarShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateCalendarShrinkRequest) GetProjectIdsShrink() *string {
	return s.ProjectIdsShrink
}

func (s *UpdateCalendarShrinkRequest) GetSchedulePolicyShrink() *string {
	return s.SchedulePolicyShrink
}

func (s *UpdateCalendarShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateCalendarShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdateCalendarShrinkRequest) SetAllProjects(v bool) *UpdateCalendarShrinkRequest {
	s.AllProjects = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetDescription(v string) *UpdateCalendarShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetEndDate(v string) *UpdateCalendarShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetExpirationReminder(v string) *UpdateCalendarShrinkRequest {
	s.ExpirationReminder = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetFallbackPolicy(v string) *UpdateCalendarShrinkRequest {
	s.FallbackPolicy = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetId(v int64) *UpdateCalendarShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetOverduePolicy(v string) *UpdateCalendarShrinkRequest {
	s.OverduePolicy = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetOwner(v string) *UpdateCalendarShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetProjectIdsShrink(v string) *UpdateCalendarShrinkRequest {
	s.ProjectIdsShrink = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetSchedulePolicyShrink(v string) *UpdateCalendarShrinkRequest {
	s.SchedulePolicyShrink = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetStartDate(v string) *UpdateCalendarShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) SetType(v string) *UpdateCalendarShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdateCalendarShrinkRequest) Validate() error {
	return dara.Validate(s)
}
