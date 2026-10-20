// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCalendarsResponseBodyPagingInfo) *ListCalendarsResponseBody
	GetPagingInfo() *ListCalendarsResponseBodyPagingInfo
	SetRequestId(v string) *ListCalendarsResponseBody
	GetRequestId() *string
}

type ListCalendarsResponseBody struct {
	// The pagination information.
	PagingInfo *ListCalendarsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCalendarsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBody) GetPagingInfo() *ListCalendarsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCalendarsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCalendarsResponseBody) SetPagingInfo(v *ListCalendarsResponseBodyPagingInfo) *ListCalendarsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCalendarsResponseBody) SetRequestId(v string) *ListCalendarsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalendarsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCalendarsResponseBodyPagingInfo struct {
	// The list of scheduling calendars.
	Calendars []*ListCalendarsResponseBodyPagingInfoCalendars `json:"Calendars,omitempty" xml:"Calendars,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCalendarsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyPagingInfo) GetCalendars() []*ListCalendarsResponseBodyPagingInfoCalendars {
	return s.Calendars
}

func (s *ListCalendarsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCalendarsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCalendarsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCalendarsResponseBodyPagingInfo) SetCalendars(v []*ListCalendarsResponseBodyPagingInfoCalendars) *ListCalendarsResponseBodyPagingInfo {
	s.Calendars = v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfo) SetPageNumber(v int32) *ListCalendarsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfo) SetPageSize(v int32) *ListCalendarsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfo) SetTotalCount(v int32) *ListCalendarsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfo) Validate() error {
	if s.Calendars != nil {
		for _, item := range s.Calendars {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCalendarsResponseBodyPagingInfoCalendars struct {
	// Indicates whether the calendar is applicable to all workspaces.
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
	// The end date of the calendar validity period. Format: `yyyy-mm-dd`. This parameter is returned only for strongly constrained calendars.
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
	SchedulePolicy *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy `json:"SchedulePolicy,omitempty" xml:"SchedulePolicy,omitempty" type:"Struct"`
	// The start date of the calendar validity period. Format: `yyyy-mm-dd`. This parameter is returned only for strongly constrained calendars.
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

func (s ListCalendarsResponseBodyPagingInfoCalendars) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBodyPagingInfoCalendars) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetAllProjects() *bool {
	return s.AllProjects
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetDescription() *string {
	return s.Description
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetEndDate() *string {
	return s.EndDate
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetExpirationReminder() *string {
	return s.ExpirationReminder
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetFallbackPolicy() *string {
	return s.FallbackPolicy
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetId() *int64 {
	return s.Id
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetName() *string {
	return s.Name
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetOverduePolicy() *string {
	return s.OverduePolicy
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetOwner() *string {
	return s.Owner
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetSchedulePolicy() *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy {
	return s.SchedulePolicy
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetStartDate() *string {
	return s.StartDate
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetType() *string {
	return s.Type
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) GetVersion() *int32 {
	return s.Version
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetAllProjects(v bool) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.AllProjects = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetDescription(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.Description = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetEndDate(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.EndDate = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetExpirationReminder(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.ExpirationReminder = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetFallbackPolicy(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.FallbackPolicy = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetId(v int64) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.Id = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetModifyTime(v int64) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.ModifyTime = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetName(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.Name = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetOverduePolicy(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.OverduePolicy = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetOwner(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.Owner = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetProjectIds(v []*int64) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.ProjectIds = v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetSchedulePolicy(v *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.SchedulePolicy = v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetStartDate(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.StartDate = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetType(v string) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.Type = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) SetVersion(v int32) *ListCalendarsResponseBodyPagingInfoCalendars {
	s.Version = &v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendars) Validate() error {
	if s.SchedulePolicy != nil {
		if err := s.SchedulePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy struct {
	// The list of normal scheduling dates.
	NormalDates []*string `json:"NormalDates,omitempty" xml:"NormalDates,omitempty" type:"Repeated"`
	// The list of non-scheduling dates. Tasks perform dry runs on these dates.
	SkipDates []*string `json:"SkipDates,omitempty" xml:"SkipDates,omitempty" type:"Repeated"`
}

func (s ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy) GetNormalDates() []*string {
	return s.NormalDates
}

func (s *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy) GetSkipDates() []*string {
	return s.SkipDates
}

func (s *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy) SetNormalDates(v []*string) *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy {
	s.NormalDates = v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy) SetSkipDates(v []*string) *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy {
	s.SkipDates = v
	return s
}

func (s *ListCalendarsResponseBodyPagingInfoCalendarsSchedulePolicy) Validate() error {
	return dara.Validate(s)
}
