// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*DescribeGroupsResponseBodyGroups) *DescribeGroupsResponseBody
	GetGroups() []*DescribeGroupsResponseBodyGroups
	SetNextToken(v string) *DescribeGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeGroupsResponseBody
	GetRequestId() *string
}

type DescribeGroupsResponseBody struct {
	// The list of groups.
	Groups []*DescribeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The pagination token. Leave this parameter empty for the first request. For subsequent requests, use the `nextToken` value returned in the previous response.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68BD3312-53D8-123E-BB32-1A9F25E07A03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBody) GetGroups() []*DescribeGroupsResponseBodyGroups {
	return s.Groups
}

func (s *DescribeGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupsResponseBody) SetGroups(v []*DescribeGroupsResponseBodyGroups) *DescribeGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeGroupsResponseBody) SetNextToken(v string) *DescribeGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetRequestId(v string) *DescribeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupsResponseBody) Validate() error {
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

type DescribeGroupsResponseBodyGroups struct {
	// The group ID.
	//
	// example:
	//
	// group-0c5ambj9yaeqjpXXX
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 100
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The list of instance information.
	Instances []*DescribeGroupsResponseBodyGroupsInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The core-hour usage in the last week.
	//
	// example:
	//
	// 85949.99
	LastWeekUsedCoreTime *float32 `json:"LastWeekUsedCoreTime,omitempty" xml:"LastWeekUsedCoreTime,omitempty"`
	// The remaining core-hours.
	//
	// example:
	//
	// 92511.15
	LeftCoreTime *float32 `json:"LeftCoreTime,omitempty" xml:"LeftCoreTime,omitempty"`
	// The list of resource IDs.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The total duration with a precision factor.
	//
	// example:
	//
	// 2000.0000
	TotalTimeDecimal *string `json:"TotalTimeDecimal,omitempty" xml:"TotalTimeDecimal,omitempty"`
	// The used core-hours.
	//
	// example:
	//
	// 7488.85
	UsedCoreTime *float32 `json:"UsedCoreTime,omitempty" xml:"UsedCoreTime,omitempty"`
	// The usage duration with a precision factor.
	//
	// example:
	//
	// 11.0000
	UsedTimeDecimal *string `json:"UsedTimeDecimal,omitempty" xml:"UsedTimeDecimal,omitempty"`
}

func (s DescribeGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupsResponseBodyGroups) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeGroupsResponseBodyGroups) GetInstances() []*DescribeGroupsResponseBodyGroupsInstances {
	return s.Instances
}

func (s *DescribeGroupsResponseBodyGroups) GetLastWeekUsedCoreTime() *float32 {
	return s.LastWeekUsedCoreTime
}

func (s *DescribeGroupsResponseBodyGroups) GetLeftCoreTime() *float32 {
	return s.LeftCoreTime
}

func (s *DescribeGroupsResponseBodyGroups) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeGroupsResponseBodyGroups) GetTotalTimeDecimal() *string {
	return s.TotalTimeDecimal
}

func (s *DescribeGroupsResponseBodyGroups) GetUsedCoreTime() *float32 {
	return s.UsedCoreTime
}

func (s *DescribeGroupsResponseBodyGroups) GetUsedTimeDecimal() *string {
	return s.UsedTimeDecimal
}

func (s *DescribeGroupsResponseBodyGroups) SetGroupId(v string) *DescribeGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetInstanceCount(v int32) *DescribeGroupsResponseBodyGroups {
	s.InstanceCount = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetInstances(v []*DescribeGroupsResponseBodyGroupsInstances) *DescribeGroupsResponseBodyGroups {
	s.Instances = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetLastWeekUsedCoreTime(v float32) *DescribeGroupsResponseBodyGroups {
	s.LastWeekUsedCoreTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetLeftCoreTime(v float32) *DescribeGroupsResponseBodyGroups {
	s.LeftCoreTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetResourceIds(v []*string) *DescribeGroupsResponseBodyGroups {
	s.ResourceIds = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetTotalTimeDecimal(v string) *DescribeGroupsResponseBodyGroups {
	s.TotalTimeDecimal = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetUsedCoreTime(v float32) *DescribeGroupsResponseBodyGroups {
	s.UsedCoreTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetUsedTimeDecimal(v string) *DescribeGroupsResponseBodyGroups {
	s.UsedTimeDecimal = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupsResponseBodyGroupsInstances struct {
	// The Alibaba Cloud user ID.
	//
	// example:
	//
	// 1712418364463000
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The expiration time. The time is in ISO 8601 format and the UTC time zone.
	//
	// example:
	//
	// 2026-09-05T15:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Specifies whether the instance is on a free trial.
	//
	// example:
	//
	// false
	FreeTrial *bool `json:"FreeTrial,omitempty" xml:"FreeTrial,omitempty"`
	// The group resource type.
	//
	// example:
	//
	// GROUP_CORE_PACKAGE
	GroupResourceType *string `json:"GroupResourceType,omitempty" xml:"GroupResourceType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// chp-0c1g77wmw50b6cXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The purchase period.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The end time of the current period. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-10-01 00:00:00
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" xml:"PeriodEndTime,omitempty"`
	// The start time of the current period. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-09-01 00:00:00
	PeriodStartTime *string `json:"PeriodStartTime,omitempty" xml:"PeriodStartTime,omitempty"`
	// The period unit.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Specifies whether priority deduction is enabled.
	//
	// example:
	//
	// true
	PriorityDeduction *bool `json:"PriorityDeduction,omitempty" xml:"PriorityDeduction,omitempty"`
	// The resource type.
	//
	// example:
	//
	// GROUP_CORE_PACKAGE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time. The time is in ISO 8601 format and the UTC time zone.
	//
	// example:
	//
	// 2026-08-05T02:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state.
	//
	// example:
	//
	// DEDUCTING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The total duration.
	//
	// example:
	//
	// 180000000
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The total duration with a precision factor.
	//
	// example:
	//
	// 2000.0000
	TotalTimeDecimal *string `json:"TotalTimeDecimal,omitempty" xml:"TotalTimeDecimal,omitempty"`
	// The usage Threshold Alerting rule.
	UsageAlertRule *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule `json:"UsageAlertRule,omitempty" xml:"UsageAlertRule,omitempty" type:"Struct"`
	// The usage duration.
	//
	// example:
	//
	// 26959861
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The usage duration with a precision factor.
	//
	// example:
	//
	// 11.0000
	UsedTimeDecimal *string `json:"UsedTimeDecimal,omitempty" xml:"UsedTimeDecimal,omitempty"`
}

func (s DescribeGroupsResponseBodyGroupsInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroupsInstances) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetFreeTrial() *bool {
	return s.FreeTrial
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetGroupResourceType() *string {
	return s.GroupResourceType
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetPeriodEndTime() *string {
	return s.PeriodEndTime
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetPeriodStartTime() *string {
	return s.PeriodStartTime
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetPriorityDeduction() *bool {
	return s.PriorityDeduction
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetState() *string {
	return s.State
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetTotalTimeDecimal() *string {
	return s.TotalTimeDecimal
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetUsageAlertRule() *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule {
	return s.UsageAlertRule
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *DescribeGroupsResponseBodyGroupsInstances) GetUsedTimeDecimal() *string {
	return s.UsedTimeDecimal
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetAliUid(v int64) *DescribeGroupsResponseBodyGroupsInstances {
	s.AliUid = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetExpiredTime(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetFreeTrial(v bool) *DescribeGroupsResponseBodyGroupsInstances {
	s.FreeTrial = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetGroupResourceType(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.GroupResourceType = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetInstanceId(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetPeriod(v int32) *DescribeGroupsResponseBodyGroupsInstances {
	s.Period = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetPeriodEndTime(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.PeriodEndTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetPeriodStartTime(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.PeriodStartTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetPeriodUnit(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetPriorityDeduction(v bool) *DescribeGroupsResponseBodyGroupsInstances {
	s.PriorityDeduction = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetResourceType(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.ResourceType = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetStartTime(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetState(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.State = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetTotalTime(v int64) *DescribeGroupsResponseBodyGroupsInstances {
	s.TotalTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetTotalTimeDecimal(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.TotalTimeDecimal = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetUsageAlertRule(v *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) *DescribeGroupsResponseBodyGroupsInstances {
	s.UsageAlertRule = v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetUsedTime(v int64) *DescribeGroupsResponseBodyGroupsInstances {
	s.UsedTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) SetUsedTimeDecimal(v string) *DescribeGroupsResponseBodyGroupsInstances {
	s.UsedTimeDecimal = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstances) Validate() error {
	if s.UsageAlertRule != nil {
		if err := s.UsageAlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule struct {
	// The list of contact group names.
	ContactGroupNames []*string `json:"ContactGroupNames,omitempty" xml:"ContactGroupNames,omitempty" type:"Repeated"`
	// The last trigger time.
	//
	// example:
	//
	// 2026-07-23 10:20:30
	LastTriggeredAt *string `json:"LastTriggeredAt,omitempty" xml:"LastTriggeredAt,omitempty"`
	// The percentage of the quota alert threshold.
	//
	// example:
	//
	// 80
	ThresholdPercent *int32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
}

func (s DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) GetContactGroupNames() []*string {
	return s.ContactGroupNames
}

func (s *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) GetLastTriggeredAt() *string {
	return s.LastTriggeredAt
}

func (s *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) GetThresholdPercent() *int32 {
	return s.ThresholdPercent
}

func (s *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) SetContactGroupNames(v []*string) *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule {
	s.ContactGroupNames = v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) SetLastTriggeredAt(v string) *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule {
	s.LastTriggeredAt = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) SetThresholdPercent(v int32) *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule {
	s.ThresholdPercent = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsInstancesUsageAlertRule) Validate() error {
	return dara.Validate(s)
}
