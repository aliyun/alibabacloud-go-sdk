// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggTaskGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggTaskGroups(v []*ListAggTaskGroupsResponseBodyAggTaskGroups) *ListAggTaskGroupsResponseBody
	GetAggTaskGroups() []*ListAggTaskGroupsResponseBodyAggTaskGroups
	SetMaxResults(v int32) *ListAggTaskGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggTaskGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAggTaskGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAggTaskGroupsResponseBody
	GetTotalCount() *int32
}

type ListAggTaskGroupsResponseBody struct {
	// List of aggregation task groups.
	AggTaskGroups []*ListAggTaskGroupsResponseBodyAggTaskGroups `json:"aggTaskGroups,omitempty" xml:"aggTaskGroups,omitempty" type:"Repeated"`
	// The maximum number of records returned.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Token for the next query.
	//
	// example:
	//
	// aa9d0e569b880xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 7BF1F4D6-B9A8-5F0B-8C1D-4347FFCB798E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Total number of instances.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAggTaskGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggTaskGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggTaskGroupsResponseBody) GetAggTaskGroups() []*ListAggTaskGroupsResponseBodyAggTaskGroups {
	return s.AggTaskGroups
}

func (s *ListAggTaskGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggTaskGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggTaskGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggTaskGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAggTaskGroupsResponseBody) SetAggTaskGroups(v []*ListAggTaskGroupsResponseBodyAggTaskGroups) *ListAggTaskGroupsResponseBody {
	s.AggTaskGroups = v
	return s
}

func (s *ListAggTaskGroupsResponseBody) SetMaxResults(v int32) *ListAggTaskGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAggTaskGroupsResponseBody) SetNextToken(v string) *ListAggTaskGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAggTaskGroupsResponseBody) SetRequestId(v string) *ListAggTaskGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggTaskGroupsResponseBody) SetTotalCount(v int32) *ListAggTaskGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAggTaskGroupsResponseBody) Validate() error {
	if s.AggTaskGroups != nil {
		for _, item := range s.AggTaskGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggTaskGroupsResponseBodyAggTaskGroups struct {
	// Hash of the aggregation task group configuration.
	//
	// example:
	//
	// a54136014xxx
	AggTaskGroupConfigHash *string `json:"aggTaskGroupConfigHash,omitempty" xml:"aggTaskGroupConfigHash,omitempty"`
	// ID of the aggregation task group.
	//
	// example:
	//
	// aggTaskGroup-xxxx
	AggTaskGroupId *string `json:"aggTaskGroupId,omitempty" xml:"aggTaskGroupId,omitempty"`
	// Name of the aggregation task group.
	//
	// example:
	//
	// pipeline-aggtask-group
	AggTaskGroupName *string `json:"aggTaskGroupName,omitempty" xml:"aggTaskGroupName,omitempty"`
	// Cron expression for the aggregation task group when the scheduling mode is set to \\"Cron\\".
	//
	// example:
	//
	// 0 10 8 1 	- ? *
	CronExpr *string `json:"cronExpr,omitempty" xml:"cronExpr,omitempty"`
	// Fixed delay time (in seconds) for scheduling.
	//
	// example:
	//
	// 30
	Delay *int32 `json:"delay,omitempty" xml:"delay,omitempty"`
	// Description of the aggregation task group.
	//
	// example:
	//
	// workspace api monitor update test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Start time of the schedule in seconds since epoch.
	//
	// example:
	//
	// 1757409499000
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// Scheduling interval.
	//
	// example:
	//
	// 2025-04-24 00:00:00,2025-04-24 00:00:00
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
	// Maximum number of retries for the aggregation task.
	//
	// example:
	//
	// 2
	MaxRetries *int32 `json:"maxRetries,omitempty" xml:"maxRetries,omitempty"`
	// Maximum retry time (in seconds) for the aggregation task.
	//
	// example:
	//
	// 200
	MaxRunTimeInSeconds *int32 `json:"maxRunTimeInSeconds,omitempty" xml:"maxRunTimeInSeconds,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Scheduling mode.
	//
	// example:
	//
	// FixedRate
	ScheduleMode *string `json:"scheduleMode,omitempty" xml:"scheduleMode,omitempty"`
	// Scheduling time expression.
	//
	// example:
	//
	// @m
	ScheduleTimeExpr *string `json:"scheduleTimeExpr,omitempty" xml:"scheduleTimeExpr,omitempty"`
	// The source Prometheus instance ID of the aggregation task group.
	//
	// example:
	//
	// rw-xxx
	SourcePrometheusId *string `json:"sourcePrometheusId,omitempty" xml:"sourcePrometheusId,omitempty"`
	// Status of the aggregation task group.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Resource group tags
	Tags []*ListAggTaskGroupsResponseBodyAggTaskGroupsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The target Prometheus instance ID of the aggregation task group.
	//
	// example:
	//
	// rw-xxx
	TargetPrometheusId *string `json:"targetPrometheusId,omitempty" xml:"targetPrometheusId,omitempty"`
	// The second-level timestamp corresponding to the end time of scheduling.
	//
	// example:
	//
	// 0
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
	// Update time of the aggregation task group.
	//
	// example:
	//
	// 1757409499000
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListAggTaskGroupsResponseBodyAggTaskGroups) String() string {
	return dara.Prettify(s)
}

func (s ListAggTaskGroupsResponseBodyAggTaskGroups) GoString() string {
	return s.String()
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetAggTaskGroupConfigHash() *string {
	return s.AggTaskGroupConfigHash
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetAggTaskGroupId() *string {
	return s.AggTaskGroupId
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetAggTaskGroupName() *string {
	return s.AggTaskGroupName
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetCronExpr() *string {
	return s.CronExpr
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetDelay() *int32 {
	return s.Delay
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetDescription() *string {
	return s.Description
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetFromTime() *int64 {
	return s.FromTime
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetInterval() *string {
	return s.Interval
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetMaxRetries() *int32 {
	return s.MaxRetries
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetMaxRunTimeInSeconds() *int32 {
	return s.MaxRunTimeInSeconds
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetScheduleMode() *string {
	return s.ScheduleMode
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetScheduleTimeExpr() *string {
	return s.ScheduleTimeExpr
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetSourcePrometheusId() *string {
	return s.SourcePrometheusId
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetStatus() *string {
	return s.Status
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetTags() []*ListAggTaskGroupsResponseBodyAggTaskGroupsTags {
	return s.Tags
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetTargetPrometheusId() *string {
	return s.TargetPrometheusId
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetToTime() *int64 {
	return s.ToTime
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetAggTaskGroupConfigHash(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.AggTaskGroupConfigHash = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetAggTaskGroupId(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.AggTaskGroupId = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetAggTaskGroupName(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.AggTaskGroupName = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetCronExpr(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.CronExpr = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetDelay(v int32) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.Delay = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetDescription(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.Description = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetFromTime(v int64) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.FromTime = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetInterval(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.Interval = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetMaxRetries(v int32) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.MaxRetries = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetMaxRunTimeInSeconds(v int32) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.MaxRunTimeInSeconds = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetRegionId(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.RegionId = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetScheduleMode(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.ScheduleMode = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetScheduleTimeExpr(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.ScheduleTimeExpr = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetSourcePrometheusId(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.SourcePrometheusId = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetStatus(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.Status = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetTags(v []*ListAggTaskGroupsResponseBodyAggTaskGroupsTags) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.Tags = v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetTargetPrometheusId(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.TargetPrometheusId = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetToTime(v int64) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.ToTime = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) SetUpdateTime(v string) *ListAggTaskGroupsResponseBodyAggTaskGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroups) Validate() error {
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

type ListAggTaskGroupsResponseBodyAggTaskGroupsTags struct {
	// Key of the resource group tag.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Value of the resource group tag.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAggTaskGroupsResponseBodyAggTaskGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListAggTaskGroupsResponseBodyAggTaskGroupsTags) GoString() string {
	return s.String()
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroupsTags) GetKey() *string {
	return s.Key
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroupsTags) GetValue() *string {
	return s.Value
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroupsTags) SetKey(v string) *ListAggTaskGroupsResponseBodyAggTaskGroupsTags {
	s.Key = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroupsTags) SetValue(v string) *ListAggTaskGroupsResponseBodyAggTaskGroupsTags {
	s.Value = &v
	return s
}

func (s *ListAggTaskGroupsResponseBodyAggTaskGroupsTags) Validate() error {
	return dara.Validate(s)
}
