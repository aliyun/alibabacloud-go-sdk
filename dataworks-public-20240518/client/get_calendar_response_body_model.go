// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCalendar(v *GetCalendarResponseBodyCalendar) *GetCalendarResponseBody
	GetCalendar() *GetCalendarResponseBodyCalendar
	SetRequestId(v string) *GetCalendarResponseBody
	GetRequestId() *string
}

type GetCalendarResponseBody struct {
	// The current complete configuration of the scheduling calendar.
	Calendar *GetCalendarResponseBodyCalendar `json:"Calendar,omitempty" xml:"Calendar,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *GetCalendarResponseBody) GetCalendar() *GetCalendarResponseBodyCalendar {
	return s.Calendar
}

func (s *GetCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCalendarResponseBody) SetCalendar(v *GetCalendarResponseBodyCalendar) *GetCalendarResponseBody {
	s.Calendar = v
	return s
}

func (s *GetCalendarResponseBody) SetRequestId(v string) *GetCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCalendarResponseBody) Validate() error {
	if s.Calendar != nil {
		if err := s.Calendar.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCalendarResponseBodyCalendar struct {
	// Specifies whether the calendar applies to all workspaces.
	//
	// example:
	//
	// true
	AllProjects *bool `json:"AllProjects,omitempty" xml:"AllProjects,omitempty"`
	// The description of the scheduling calendar.
	//
	// example:
	//
	// Weekday scheduling calendar
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end date of the calendar validity period, in the `yyyy-mm-dd` format. This parameter is returned only for strongly constrained calendars.
	//
	// example:
	//
	// 2026-10-09
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The expiration reminder. This parameter is returned only for strongly constrained calendars. Valid values:
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
	// The default scheduling policy for dates that are not configured. This parameter is returned only for weakly constrained calendars. Valid values:
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
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the calendar was modified. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1640000000000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the scheduling calendar.
	//
	// example:
	//
	// work_calendar
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The scheduling policy after the calendar expires. This parameter is returned only for strongly constrained calendars. Valid values:
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
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The list of applicable workspace IDs. An empty array is returned if AllProjects is set to true.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The date scheduling policy.
	SchedulePolicy *GetCalendarResponseBodyCalendarSchedulePolicy `json:"SchedulePolicy,omitempty" xml:"SchedulePolicy,omitempty" type:"Struct"`
	// The start date of the calendar validity period, in the `yyyy-mm-dd` format. This parameter is returned only for strongly constrained calendars.
	//
	// example:
	//
	// 2026-10-08
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The calendar type. Valid values:
	//
	// - StrongConstraint: Strongly constrained calendar.
	//
	// - WeakConstraint: Weakly constrained calendar.
	//
	// example:
	//
	// StrongConstraint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetCalendarResponseBodyCalendar) String() string {
	return dara.Prettify(s)
}

func (s GetCalendarResponseBodyCalendar) GoString() string {
	return s.String()
}

func (s *GetCalendarResponseBodyCalendar) GetAllProjects() *bool {
	return s.AllProjects
}

func (s *GetCalendarResponseBodyCalendar) GetDescription() *string {
	return s.Description
}

func (s *GetCalendarResponseBodyCalendar) GetEndDate() *string {
	return s.EndDate
}

func (s *GetCalendarResponseBodyCalendar) GetExpirationReminder() *string {
	return s.ExpirationReminder
}

func (s *GetCalendarResponseBodyCalendar) GetFallbackPolicy() *string {
	return s.FallbackPolicy
}

func (s *GetCalendarResponseBodyCalendar) GetId() *int64 {
	return s.Id
}

func (s *GetCalendarResponseBodyCalendar) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetCalendarResponseBodyCalendar) GetName() *string {
	return s.Name
}

func (s *GetCalendarResponseBodyCalendar) GetOverduePolicy() *string {
	return s.OverduePolicy
}

func (s *GetCalendarResponseBodyCalendar) GetOwner() *string {
	return s.Owner
}

func (s *GetCalendarResponseBodyCalendar) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *GetCalendarResponseBodyCalendar) GetSchedulePolicy() *GetCalendarResponseBodyCalendarSchedulePolicy {
	return s.SchedulePolicy
}

func (s *GetCalendarResponseBodyCalendar) GetStartDate() *string {
	return s.StartDate
}

func (s *GetCalendarResponseBodyCalendar) GetType() *string {
	return s.Type
}

func (s *GetCalendarResponseBodyCalendar) GetVersion() *int32 {
	return s.Version
}

func (s *GetCalendarResponseBodyCalendar) SetAllProjects(v bool) *GetCalendarResponseBodyCalendar {
	s.AllProjects = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetDescription(v string) *GetCalendarResponseBodyCalendar {
	s.Description = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetEndDate(v string) *GetCalendarResponseBodyCalendar {
	s.EndDate = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetExpirationReminder(v string) *GetCalendarResponseBodyCalendar {
	s.ExpirationReminder = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetFallbackPolicy(v string) *GetCalendarResponseBodyCalendar {
	s.FallbackPolicy = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetId(v int64) *GetCalendarResponseBodyCalendar {
	s.Id = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetModifyTime(v int64) *GetCalendarResponseBodyCalendar {
	s.ModifyTime = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetName(v string) *GetCalendarResponseBodyCalendar {
	s.Name = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetOverduePolicy(v string) *GetCalendarResponseBodyCalendar {
	s.OverduePolicy = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetOwner(v string) *GetCalendarResponseBodyCalendar {
	s.Owner = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetProjectIds(v []*int64) *GetCalendarResponseBodyCalendar {
	s.ProjectIds = v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetSchedulePolicy(v *GetCalendarResponseBodyCalendarSchedulePolicy) *GetCalendarResponseBodyCalendar {
	s.SchedulePolicy = v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetStartDate(v string) *GetCalendarResponseBodyCalendar {
	s.StartDate = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetType(v string) *GetCalendarResponseBodyCalendar {
	s.Type = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) SetVersion(v int32) *GetCalendarResponseBodyCalendar {
	s.Version = &v
	return s
}

func (s *GetCalendarResponseBodyCalendar) Validate() error {
	if s.SchedulePolicy != nil {
		if err := s.SchedulePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCalendarResponseBodyCalendarSchedulePolicy struct {
	// The list of normal scheduling dates.
	NormalDates []*string `json:"NormalDates,omitempty" xml:"NormalDates,omitempty" type:"Repeated"`
	// The list of non-scheduling dates. Tasks perform dry runs on these dates.
	SkipDates []*string `json:"SkipDates,omitempty" xml:"SkipDates,omitempty" type:"Repeated"`
}

func (s GetCalendarResponseBodyCalendarSchedulePolicy) String() string {
	return dara.Prettify(s)
}

func (s GetCalendarResponseBodyCalendarSchedulePolicy) GoString() string {
	return s.String()
}

func (s *GetCalendarResponseBodyCalendarSchedulePolicy) GetNormalDates() []*string {
	return s.NormalDates
}

func (s *GetCalendarResponseBodyCalendarSchedulePolicy) GetSkipDates() []*string {
	return s.SkipDates
}

func (s *GetCalendarResponseBodyCalendarSchedulePolicy) SetNormalDates(v []*string) *GetCalendarResponseBodyCalendarSchedulePolicy {
	s.NormalDates = v
	return s
}

func (s *GetCalendarResponseBodyCalendarSchedulePolicy) SetSkipDates(v []*string) *GetCalendarResponseBodyCalendarSchedulePolicy {
	s.SkipDates = v
	return s
}

func (s *GetCalendarResponseBodyCalendarSchedulePolicy) Validate() error {
	return dara.Validate(s)
}
