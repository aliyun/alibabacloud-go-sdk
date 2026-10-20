// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllProjects(v bool) *CreateCalendarRequest
	GetAllProjects() *bool
	SetDescription(v string) *CreateCalendarRequest
	GetDescription() *string
	SetEndDate(v string) *CreateCalendarRequest
	GetEndDate() *string
	SetExpirationReminder(v string) *CreateCalendarRequest
	GetExpirationReminder() *string
	SetFallbackPolicy(v string) *CreateCalendarRequest
	GetFallbackPolicy() *string
	SetName(v string) *CreateCalendarRequest
	GetName() *string
	SetOverduePolicy(v string) *CreateCalendarRequest
	GetOverduePolicy() *string
	SetOwner(v string) *CreateCalendarRequest
	GetOwner() *string
	SetProjectIds(v []*int64) *CreateCalendarRequest
	GetProjectIds() []*int64
	SetSchedulePolicy(v *CreateCalendarRequestSchedulePolicy) *CreateCalendarRequest
	GetSchedulePolicy() *CreateCalendarRequestSchedulePolicy
	SetStartDate(v string) *CreateCalendarRequest
	GetStartDate() *string
	SetType(v string) *CreateCalendarRequest
	GetType() *string
}

type CreateCalendarRequest struct {
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
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The date scheduling policy. For a strong constraint calendar, the normal scheduling dates and non-scheduling dates must cover every day within the validity period and cannot overlap.
	//
	// This parameter is required.
	SchedulePolicy *CreateCalendarRequestSchedulePolicy `json:"SchedulePolicy,omitempty" xml:"SchedulePolicy,omitempty" type:"Struct"`
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

func (s CreateCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCalendarRequest) GoString() string {
	return s.String()
}

func (s *CreateCalendarRequest) GetAllProjects() *bool {
	return s.AllProjects
}

func (s *CreateCalendarRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCalendarRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateCalendarRequest) GetExpirationReminder() *string {
	return s.ExpirationReminder
}

func (s *CreateCalendarRequest) GetFallbackPolicy() *string {
	return s.FallbackPolicy
}

func (s *CreateCalendarRequest) GetName() *string {
	return s.Name
}

func (s *CreateCalendarRequest) GetOverduePolicy() *string {
	return s.OverduePolicy
}

func (s *CreateCalendarRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateCalendarRequest) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *CreateCalendarRequest) GetSchedulePolicy() *CreateCalendarRequestSchedulePolicy {
	return s.SchedulePolicy
}

func (s *CreateCalendarRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateCalendarRequest) GetType() *string {
	return s.Type
}

func (s *CreateCalendarRequest) SetAllProjects(v bool) *CreateCalendarRequest {
	s.AllProjects = &v
	return s
}

func (s *CreateCalendarRequest) SetDescription(v string) *CreateCalendarRequest {
	s.Description = &v
	return s
}

func (s *CreateCalendarRequest) SetEndDate(v string) *CreateCalendarRequest {
	s.EndDate = &v
	return s
}

func (s *CreateCalendarRequest) SetExpirationReminder(v string) *CreateCalendarRequest {
	s.ExpirationReminder = &v
	return s
}

func (s *CreateCalendarRequest) SetFallbackPolicy(v string) *CreateCalendarRequest {
	s.FallbackPolicy = &v
	return s
}

func (s *CreateCalendarRequest) SetName(v string) *CreateCalendarRequest {
	s.Name = &v
	return s
}

func (s *CreateCalendarRequest) SetOverduePolicy(v string) *CreateCalendarRequest {
	s.OverduePolicy = &v
	return s
}

func (s *CreateCalendarRequest) SetOwner(v string) *CreateCalendarRequest {
	s.Owner = &v
	return s
}

func (s *CreateCalendarRequest) SetProjectIds(v []*int64) *CreateCalendarRequest {
	s.ProjectIds = v
	return s
}

func (s *CreateCalendarRequest) SetSchedulePolicy(v *CreateCalendarRequestSchedulePolicy) *CreateCalendarRequest {
	s.SchedulePolicy = v
	return s
}

func (s *CreateCalendarRequest) SetStartDate(v string) *CreateCalendarRequest {
	s.StartDate = &v
	return s
}

func (s *CreateCalendarRequest) SetType(v string) *CreateCalendarRequest {
	s.Type = &v
	return s
}

func (s *CreateCalendarRequest) Validate() error {
	if s.SchedulePolicy != nil {
		if err := s.SchedulePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCalendarRequestSchedulePolicy struct {
	// The list of normal scheduling dates.
	NormalDates []*string `json:"NormalDates,omitempty" xml:"NormalDates,omitempty" type:"Repeated"`
	// The list of non-scheduling dates. Tasks perform dry runs on these dates. This parameter is used only for strong constraint calendars.
	SkipDates []*string `json:"SkipDates,omitempty" xml:"SkipDates,omitempty" type:"Repeated"`
}

func (s CreateCalendarRequestSchedulePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateCalendarRequestSchedulePolicy) GoString() string {
	return s.String()
}

func (s *CreateCalendarRequestSchedulePolicy) GetNormalDates() []*string {
	return s.NormalDates
}

func (s *CreateCalendarRequestSchedulePolicy) GetSkipDates() []*string {
	return s.SkipDates
}

func (s *CreateCalendarRequestSchedulePolicy) SetNormalDates(v []*string) *CreateCalendarRequestSchedulePolicy {
	s.NormalDates = v
	return s
}

func (s *CreateCalendarRequestSchedulePolicy) SetSkipDates(v []*string) *CreateCalendarRequestSchedulePolicy {
	s.SkipDates = v
	return s
}

func (s *CreateCalendarRequestSchedulePolicy) Validate() error {
	return dara.Validate(s)
}
