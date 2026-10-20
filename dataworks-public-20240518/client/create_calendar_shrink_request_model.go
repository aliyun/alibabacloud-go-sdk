// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalendarShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllProjects(v bool) *CreateCalendarShrinkRequest
	GetAllProjects() *bool
	SetDescription(v string) *CreateCalendarShrinkRequest
	GetDescription() *string
	SetEndDate(v string) *CreateCalendarShrinkRequest
	GetEndDate() *string
	SetExpirationReminder(v string) *CreateCalendarShrinkRequest
	GetExpirationReminder() *string
	SetFallbackPolicy(v string) *CreateCalendarShrinkRequest
	GetFallbackPolicy() *string
	SetName(v string) *CreateCalendarShrinkRequest
	GetName() *string
	SetOverduePolicy(v string) *CreateCalendarShrinkRequest
	GetOverduePolicy() *string
	SetOwner(v string) *CreateCalendarShrinkRequest
	GetOwner() *string
	SetProjectIdsShrink(v string) *CreateCalendarShrinkRequest
	GetProjectIdsShrink() *string
	SetSchedulePolicyShrink(v string) *CreateCalendarShrinkRequest
	GetSchedulePolicyShrink() *string
	SetStartDate(v string) *CreateCalendarShrinkRequest
	GetStartDate() *string
	SetType(v string) *CreateCalendarShrinkRequest
	GetType() *string
}

type CreateCalendarShrinkRequest struct {
	// Specifies whether the calendar applies to all workspaces. Default value: false.
	//
	// - true: all workspaces.
	//
	// - false: the workspaces specified by ProjectIds.
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
	// The end date of the calendar validity period. Format: `yyyy-mm-dd`. This parameter is required when Type is set to StrongConstraint. The end date cannot be earlier than the start date, and the interval between the start date and the end date must be less than 5 years.
	//
	// example:
	//
	// 2026-10-09
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The expiration reminder. This parameter is required when Type is set to StrongConstraint.
	//
	// - OneMonthBefore: one month in advance.
	//
	// - OneWeekBefore: one week in advance.
	//
	// - ThreeDaysBefore: three days in advance.
	//
	// example:
	//
	// OneMonthBefore
	ExpirationReminder *string `json:"ExpirationReminder,omitempty" xml:"ExpirationReminder,omitempty"`
	// The default scheduling policy for dates that are not configured. This parameter is required when Type is set to WeakConstraint.
	//
	// - Normal: normal scheduling.
	//
	// - Pause: pause scheduling.
	//
	// - Skip: dry run.
	//
	// example:
	//
	// Normal
	FallbackPolicy *string `json:"FallbackPolicy,omitempty" xml:"FallbackPolicy,omitempty"`
	// The name of the scheduling calendar. The name must be unique within the current tenant and can be up to 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// work_calendar
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The scheduling policy after the calendar expires. This parameter is required when Type is set to StrongConstraint.
	//
	// - Normal: normal scheduling.
	//
	// - Pause: pause scheduling.
	//
	// - Skip: dry run.
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
	// - If AllProjects is set to false or left empty: This parameter is required and cannot be empty.
	//
	// - If AllProjects is set to true: Leave this parameter empty or pass an empty array.
	ProjectIdsShrink *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// The date scheduling policy. For a strong constraint calendar, the normal scheduling dates and non-scheduling dates must cover every day within the validity period and cannot overlap.
	//
	// This parameter is required.
	SchedulePolicyShrink *string `json:"SchedulePolicy,omitempty" xml:"SchedulePolicy,omitempty"`
	// The start date of the calendar validity period. Format: `yyyy-mm-dd`. This parameter is required when Type is set to StrongConstraint.
	//
	// example:
	//
	// 2026-10-08
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The type of the calendar.
	//
	// - StrongConstraint: strong constraint calendar.
	//
	// - WeakConstraint: weak constraint calendar.
	//
	// This parameter is required.
	//
	// example:
	//
	// StrongConstraint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCalendarShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCalendarShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCalendarShrinkRequest) GetAllProjects() *bool {
	return s.AllProjects
}

func (s *CreateCalendarShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCalendarShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateCalendarShrinkRequest) GetExpirationReminder() *string {
	return s.ExpirationReminder
}

func (s *CreateCalendarShrinkRequest) GetFallbackPolicy() *string {
	return s.FallbackPolicy
}

func (s *CreateCalendarShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCalendarShrinkRequest) GetOverduePolicy() *string {
	return s.OverduePolicy
}

func (s *CreateCalendarShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateCalendarShrinkRequest) GetProjectIdsShrink() *string {
	return s.ProjectIdsShrink
}

func (s *CreateCalendarShrinkRequest) GetSchedulePolicyShrink() *string {
	return s.SchedulePolicyShrink
}

func (s *CreateCalendarShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateCalendarShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateCalendarShrinkRequest) SetAllProjects(v bool) *CreateCalendarShrinkRequest {
	s.AllProjects = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetDescription(v string) *CreateCalendarShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetEndDate(v string) *CreateCalendarShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetExpirationReminder(v string) *CreateCalendarShrinkRequest {
	s.ExpirationReminder = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetFallbackPolicy(v string) *CreateCalendarShrinkRequest {
	s.FallbackPolicy = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetName(v string) *CreateCalendarShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetOverduePolicy(v string) *CreateCalendarShrinkRequest {
	s.OverduePolicy = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetOwner(v string) *CreateCalendarShrinkRequest {
	s.Owner = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetProjectIdsShrink(v string) *CreateCalendarShrinkRequest {
	s.ProjectIdsShrink = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetSchedulePolicyShrink(v string) *CreateCalendarShrinkRequest {
	s.SchedulePolicyShrink = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetStartDate(v string) *CreateCalendarShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *CreateCalendarShrinkRequest) SetType(v string) *CreateCalendarShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateCalendarShrinkRequest) Validate() error {
	return dara.Validate(s)
}
