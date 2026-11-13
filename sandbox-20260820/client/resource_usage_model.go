// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceUsage interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ResourceUsage
	GetEndTime() *string
	SetGranularity(v string) *ResourceUsage
	GetGranularity() *string
	SetGroupBy(v []*string) *ResourceUsage
	GetGroupBy() []*string
	SetGroups(v []*ResourceUsageGroups) *ResourceUsage
	GetGroups() []*ResourceUsageGroups
	SetStartTime(v string) *ResourceUsage
	GetStartTime() *string
}

type ResourceUsage struct {
	// The end time of the requested interval, exclusive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-05T02:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The aggregation granularity used for this query. Valid values: All, Hour, and Day.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hour
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The grouping key used for this query. An empty array is returned when no grouping is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["teamID"]
	GroupBy []*string `json:"groupBy" xml:"groupBy" type:"Repeated"`
	// The usage for each group. If no grouping is applied, the usage for the entire query range is returned. An empty array is returned if there are no usage records within the query range.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"dimensions":{"teamID":"70d1c834-0333-58d8-97ac-5336eb9179eb"},"items":[{"usageName":"cpuActive","unit":"CoreSecond","total":5400,"buckets":[{"startTime":"2026-11-05T00:00:00Z","endTime":"2026-11-05T01:00:00Z","quantity":3600},{"startTime":"2026-11-05T01:00:00Z","endTime":"2026-11-05T02:00:00Z","quantity":1800}]}]}]
	Groups []*ResourceUsageGroups `json:"groups" xml:"groups" type:"Repeated"`
	// The start time of the requested interval, inclusive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-05T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ResourceUsage) String() string {
	return dara.Prettify(s)
}

func (s ResourceUsage) GoString() string {
	return s.String()
}

func (s *ResourceUsage) GetEndTime() *string {
	return s.EndTime
}

func (s *ResourceUsage) GetGranularity() *string {
	return s.Granularity
}

func (s *ResourceUsage) GetGroupBy() []*string {
	return s.GroupBy
}

func (s *ResourceUsage) GetGroups() []*ResourceUsageGroups {
	return s.Groups
}

func (s *ResourceUsage) GetStartTime() *string {
	return s.StartTime
}

func (s *ResourceUsage) SetEndTime(v string) *ResourceUsage {
	s.EndTime = &v
	return s
}

func (s *ResourceUsage) SetGranularity(v string) *ResourceUsage {
	s.Granularity = &v
	return s
}

func (s *ResourceUsage) SetGroupBy(v []*string) *ResourceUsage {
	s.GroupBy = v
	return s
}

func (s *ResourceUsage) SetGroups(v []*ResourceUsageGroups) *ResourceUsage {
	s.Groups = v
	return s
}

func (s *ResourceUsage) SetStartTime(v string) *ResourceUsage {
	s.StartTime = &v
	return s
}

func (s *ResourceUsage) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResourceUsageGroups struct {
	// The team, template, or sandbox version corresponding to the usage of this group.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"teamID":"70d1c834-0333-58d8-97ac-5336eb9179eb"}
	Dimensions map[string]*string `json:"dimensions" xml:"dimensions"`
	// The query result for each requested usage item in this group.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"usageName":"cpuActive","unit":"CoreSecond","total":5400,"buckets":[{"startTime":"2026-11-05T00:00:00Z","endTime":"2026-11-05T01:00:00Z","quantity":3600},{"startTime":"2026-11-05T01:00:00Z","endTime":"2026-11-05T02:00:00Z","quantity":1800}]}]
	Items []*ResourceUsageGroupsItems `json:"items" xml:"items" type:"Repeated"`
}

func (s ResourceUsageGroups) String() string {
	return dara.Prettify(s)
}

func (s ResourceUsageGroups) GoString() string {
	return s.String()
}

func (s *ResourceUsageGroups) GetDimensions() map[string]*string {
	return s.Dimensions
}

func (s *ResourceUsageGroups) GetItems() []*ResourceUsageGroupsItems {
	return s.Items
}

func (s *ResourceUsageGroups) SetDimensions(v map[string]*string) *ResourceUsageGroups {
	s.Dimensions = v
	return s
}

func (s *ResourceUsageGroups) SetItems(v []*ResourceUsageGroupsItems) *ResourceUsageGroups {
	s.Items = v
	return s
}

func (s *ResourceUsageGroups) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResourceUsageGroupsItems struct {
	// The hourly or daily usage arranged in chronological order. Only time periods with records are returned. This field is omitted when granularity is set to All.
	//
	// example:
	//
	// [{"startTime":"2026-11-05T00:00:00Z","endTime":"2026-11-05T01:00:00Z","quantity":3600},{"startTime":"2026-11-05T01:00:00Z","endTime":"2026-11-05T02:00:00Z","quantity":1800}]
	Buckets []*ResourceUsageGroupsItemsBuckets `json:"buckets" xml:"buckets" type:"Repeated"`
	// The cumulative usage of this usage item over the entire query interval. This field is omitted if there are no usage records for the item within the interval.
	//
	// example:
	//
	// 5400
	Total *float64 `json:"total,omitempty" xml:"total,omitempty"`
	// The unit of usage. The unit is CoreSecond for CPU, and GiBSecond for memory and cloud disks.
	//
	// This parameter is required.
	//
	// example:
	//
	// CoreSecond
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// The name of the usage item specified in the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpuActive
	UsageName *string `json:"usageName,omitempty" xml:"usageName,omitempty"`
}

func (s ResourceUsageGroupsItems) String() string {
	return dara.Prettify(s)
}

func (s ResourceUsageGroupsItems) GoString() string {
	return s.String()
}

func (s *ResourceUsageGroupsItems) GetBuckets() []*ResourceUsageGroupsItemsBuckets {
	return s.Buckets
}

func (s *ResourceUsageGroupsItems) GetTotal() *float64 {
	return s.Total
}

func (s *ResourceUsageGroupsItems) GetUnit() *string {
	return s.Unit
}

func (s *ResourceUsageGroupsItems) GetUsageName() *string {
	return s.UsageName
}

func (s *ResourceUsageGroupsItems) SetBuckets(v []*ResourceUsageGroupsItemsBuckets) *ResourceUsageGroupsItems {
	s.Buckets = v
	return s
}

func (s *ResourceUsageGroupsItems) SetTotal(v float64) *ResourceUsageGroupsItems {
	s.Total = &v
	return s
}

func (s *ResourceUsageGroupsItems) SetUnit(v string) *ResourceUsageGroupsItems {
	s.Unit = &v
	return s
}

func (s *ResourceUsageGroupsItems) SetUsageName(v string) *ResourceUsageGroupsItems {
	s.UsageName = &v
	return s
}

func (s *ResourceUsageGroupsItems) Validate() error {
	if s.Buckets != nil {
		for _, item := range s.Buckets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResourceUsageGroupsItemsBuckets struct {
	// The end time of this usage statistics period, exclusive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-05T01:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The cumulative usage within this time period. The unit is specified by unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3600
	Quantity *float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The start time of this usage statistics period, inclusive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-05T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ResourceUsageGroupsItemsBuckets) String() string {
	return dara.Prettify(s)
}

func (s ResourceUsageGroupsItemsBuckets) GoString() string {
	return s.String()
}

func (s *ResourceUsageGroupsItemsBuckets) GetEndTime() *string {
	return s.EndTime
}

func (s *ResourceUsageGroupsItemsBuckets) GetQuantity() *float64 {
	return s.Quantity
}

func (s *ResourceUsageGroupsItemsBuckets) GetStartTime() *string {
	return s.StartTime
}

func (s *ResourceUsageGroupsItemsBuckets) SetEndTime(v string) *ResourceUsageGroupsItemsBuckets {
	s.EndTime = &v
	return s
}

func (s *ResourceUsageGroupsItemsBuckets) SetQuantity(v float64) *ResourceUsageGroupsItemsBuckets {
	s.Quantity = &v
	return s
}

func (s *ResourceUsageGroupsItemsBuckets) SetStartTime(v string) *ResourceUsageGroupsItemsBuckets {
	s.StartTime = &v
	return s
}

func (s *ResourceUsageGroupsItemsBuckets) Validate() error {
	return dara.Validate(s)
}
