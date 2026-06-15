// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlanMaintenanceWindowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribePlanMaintenanceWindowsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePlanMaintenanceWindowsResponseBody
	GetNextToken() *string
	SetPlanMaintenanceWindowList(v []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) *DescribePlanMaintenanceWindowsResponseBody
	GetPlanMaintenanceWindowList() []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList
	SetRequestId(v string) *DescribePlanMaintenanceWindowsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePlanMaintenanceWindowsResponseBody
	GetTotalCount() *int32
}

type DescribePlanMaintenanceWindowsResponseBody struct {
	// The number of entries to return per page. The default value is 10, and the maximum value is 100. If you omit this parameter or specify a value less than 10, the default value is used. If you specify a value greater than 100, the maximum value is used.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// A list of maintenance windows.
	PlanMaintenanceWindowList []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList `json:"PlanMaintenanceWindowList,omitempty" xml:"PlanMaintenanceWindowList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that match the query. This parameter is optional and not returned by default.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlanMaintenanceWindowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePlanMaintenanceWindowsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePlanMaintenanceWindowsResponseBody) GetPlanMaintenanceWindowList() []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	return s.PlanMaintenanceWindowList
}

func (s *DescribePlanMaintenanceWindowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlanMaintenanceWindowsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePlanMaintenanceWindowsResponseBody) SetMaxResults(v int32) *DescribePlanMaintenanceWindowsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBody) SetNextToken(v string) *DescribePlanMaintenanceWindowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBody) SetPlanMaintenanceWindowList(v []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) *DescribePlanMaintenanceWindowsResponseBody {
	s.PlanMaintenanceWindowList = v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBody) SetRequestId(v string) *DescribePlanMaintenanceWindowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBody) SetTotalCount(v int32) *DescribePlanMaintenanceWindowsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBody) Validate() error {
	if s.PlanMaintenanceWindowList != nil {
		for _, item := range s.PlanMaintenanceWindowList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList struct {
	// The creation time of the maintenance window.
	//
	// The time is in UTC and follows the ISO 8601 standard, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-04-11T02:20:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the maintenance window is enabled.
	//
	// example:
	//
	// true
	Enable                 *bool  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MinMaintenanceInterval *int32 `json:"MinMaintenanceInterval,omitempty" xml:"MinMaintenanceInterval,omitempty"`
	// The modification time of the maintenance window.
	//
	// The time is in UTC and follows the ISO 8601 standard, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-02-22 10:14:28 +0800
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the maintenance window.
	//
	// example:
	//
	// pw-bp1bqkbjb7h4j8zqzwvp
	PlanWindowId *string `json:"PlanWindowId,omitempty" xml:"PlanWindowId,omitempty"`
	// The name of the maintenance window.
	//
	// example:
	//
	// WindowName
	PlanWindowName *string `json:"PlanWindowName,omitempty" xml:"PlanWindowName,omitempty"`
	// The supported maintenance action.
	//
	// example:
	//
	// Reboot
	SupportMaintenanceAction *string `json:"SupportMaintenanceAction,omitempty" xml:"SupportMaintenanceAction,omitempty"`
	// The resources targeted by the maintenance window.
	TargetResource *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource `json:"TargetResource,omitempty" xml:"TargetResource,omitempty" type:"Struct"`
	// The recurrence schedule of the maintenance window.
	TimePeriod *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty" type:"Struct"`
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetEnable() *bool {
	return s.Enable
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetMinMaintenanceInterval() *int32 {
	return s.MinMaintenanceInterval
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetPlanWindowId() *string {
	return s.PlanWindowId
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetPlanWindowName() *string {
	return s.PlanWindowName
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetSupportMaintenanceAction() *string {
	return s.SupportMaintenanceAction
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetTargetResource() *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource {
	return s.TargetResource
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) GetTimePeriod() *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod {
	return s.TimePeriod
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetCreateTime(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.CreateTime = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetEnable(v bool) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.Enable = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetMinMaintenanceInterval(v int32) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.MinMaintenanceInterval = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetModifiedTime(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetPlanWindowId(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.PlanWindowId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetPlanWindowName(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.PlanWindowName = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetSupportMaintenanceAction(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.SupportMaintenanceAction = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetTargetResource(v *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.TargetResource = v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) SetTimePeriod(v *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList {
	s.TimePeriod = v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowList) Validate() error {
	if s.TargetResource != nil {
		if err := s.TargetResource.Validate(); err != nil {
			return err
		}
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource struct {
	// The ID of the target resource group.
	//
	// example:
	//
	// rg-aek2qxeteo7fr6y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scope of the target resources. Valid values: Tag, ResourceGroup, Instance, and AliUid.
	//
	// example:
	//
	// Tag
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The target tags.
	Tags []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) GetScope() *string {
	return s.Scope
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) GetTags() []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags {
	return s.Tags
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) SetResourceGroupId(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) SetScope(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource {
	s.Scope = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) SetTags(v []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource {
	s.Tags = v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResource) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags) GetKey() *string {
	return s.Key
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags) GetValue() *string {
	return s.Value
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags) SetKey(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags {
	s.Key = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags) SetValue(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags {
	s.Value = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTargetResourceTags) Validate() error {
	return dara.Validate(s)
}

type DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod struct {
	// The recurrence frequency. Valid values: Daily and Weekly.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The recurring UTC time ranges for the maintenance window.
	RangeList []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList `json:"RangeList,omitempty" xml:"RangeList,omitempty" type:"Repeated"`
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod) GetRangeList() []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList {
	return s.RangeList
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod) SetPeriodUnit(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod) SetRangeList(v []*DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod {
	s.RangeList = v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriod) Validate() error {
	if s.RangeList != nil {
		for _, item := range s.RangeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList struct {
	// The end time of the time range.
	//
	// example:
	//
	// Monday,22:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the time range.
	//
	// example:
	//
	// Monday,22:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList) SetEndTime(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList {
	s.EndTime = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList) SetStartTime(v string) *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList {
	s.StartTime = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsResponseBodyPlanMaintenanceWindowListTimePeriodRangeList) Validate() error {
	return dara.Validate(s)
}
