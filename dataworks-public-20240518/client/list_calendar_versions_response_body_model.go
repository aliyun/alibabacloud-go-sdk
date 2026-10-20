// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCalendarVersionsResponseBodyPagingInfo) *ListCalendarVersionsResponseBody
	GetPagingInfo() *ListCalendarVersionsResponseBodyPagingInfo
	SetRequestId(v string) *ListCalendarVersionsResponseBody
	GetRequestId() *string
}

type ListCalendarVersionsResponseBody struct {
	// The pagination information.
	PagingInfo *ListCalendarVersionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCalendarVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalendarVersionsResponseBody) GetPagingInfo() *ListCalendarVersionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCalendarVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCalendarVersionsResponseBody) SetPagingInfo(v *ListCalendarVersionsResponseBodyPagingInfo) *ListCalendarVersionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCalendarVersionsResponseBody) SetRequestId(v string) *ListCalendarVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalendarVersionsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCalendarVersionsResponseBodyPagingInfo struct {
	// The complete list of historical versions of the scheduling calendar, sorted in descending order by version number.
	CalendarVersions []*ListCalendarVersionsResponseBodyPagingInfoCalendarVersions `json:"CalendarVersions,omitempty" xml:"CalendarVersions,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCalendarVersionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarVersionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) GetCalendarVersions() []*ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	return s.CalendarVersions
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) SetCalendarVersions(v []*ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) *ListCalendarVersionsResponseBodyPagingInfo {
	s.CalendarVersions = v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListCalendarVersionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) SetPageSize(v int32) *ListCalendarVersionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListCalendarVersionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfo) Validate() error {
	if s.CalendarVersions != nil {
		for _, item := range s.CalendarVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCalendarVersionsResponseBodyPagingInfoCalendarVersions struct {
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
	// The default scheduling policy for dates that are not specified. This parameter is returned only for weakly constrained calendars. Valid values:
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
	SchedulePolicy *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy `json:"SchedulePolicy,omitempty" xml:"SchedulePolicy,omitempty" type:"Struct"`
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

func (s ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GoString() string {
	return s.String()
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetAllProjects() *bool {
	return s.AllProjects
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetDescription() *string {
	return s.Description
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetEndDate() *string {
	return s.EndDate
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetExpirationReminder() *string {
	return s.ExpirationReminder
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetFallbackPolicy() *string {
	return s.FallbackPolicy
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetId() *int64 {
	return s.Id
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetName() *string {
	return s.Name
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetOverduePolicy() *string {
	return s.OverduePolicy
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetOwner() *string {
	return s.Owner
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetSchedulePolicy() *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy {
	return s.SchedulePolicy
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetStartDate() *string {
	return s.StartDate
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetType() *string {
	return s.Type
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) GetVersion() *int32 {
	return s.Version
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetAllProjects(v bool) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.AllProjects = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetDescription(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.Description = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetEndDate(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.EndDate = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetExpirationReminder(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.ExpirationReminder = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetFallbackPolicy(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.FallbackPolicy = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetId(v int64) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.Id = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetModifyTime(v int64) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.ModifyTime = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetName(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.Name = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetOverduePolicy(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.OverduePolicy = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetOwner(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.Owner = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetProjectIds(v []*int64) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.ProjectIds = v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetSchedulePolicy(v *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.SchedulePolicy = v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetStartDate(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.StartDate = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetType(v string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.Type = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) SetVersion(v int32) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions {
	s.Version = &v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersions) Validate() error {
	if s.SchedulePolicy != nil {
		if err := s.SchedulePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy struct {
	// The list of normal scheduling dates.
	NormalDates []*string `json:"NormalDates,omitempty" xml:"NormalDates,omitempty" type:"Repeated"`
	// The list of non-scheduling dates. Tasks perform dry runs on these dates.
	SkipDates []*string `json:"SkipDates,omitempty" xml:"SkipDates,omitempty" type:"Repeated"`
}

func (s ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy) GoString() string {
	return s.String()
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy) GetNormalDates() []*string {
	return s.NormalDates
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy) GetSkipDates() []*string {
	return s.SkipDates
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy) SetNormalDates(v []*string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy {
	s.NormalDates = v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy) SetSkipDates(v []*string) *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy {
	s.SkipDates = v
	return s
}

func (s *ListCalendarVersionsResponseBodyPagingInfoCalendarVersionsSchedulePolicy) Validate() error {
	return dara.Validate(s)
}
