// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnterpriseSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeEnterpriseSnapshotPolicyResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeEnterpriseSnapshotPolicyResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnterpriseSnapshotPolicyResponseBody
	GetPageSize() *int32
	SetPolicies(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) *DescribeEnterpriseSnapshotPolicyResponseBody
	GetPolicies() []*DescribeEnterpriseSnapshotPolicyResponseBodyPolicies
	SetRequestId(v string) *DescribeEnterpriseSnapshotPolicyResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeEnterpriseSnapshotPolicyResponseBody
	GetTotalCount() *int64
}

type DescribeEnterpriseSnapshotPolicyResponseBody struct {
	// The pagination token (Token) returned for the next query.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number for paged queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paged queries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of policies.
	Policies []*DescribeEnterpriseSnapshotPolicyResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 549BB71F-6778-53CE-AA73-E94ACE802AB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) GetPolicies() []*DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	return s.Policies
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetNextToken(v string) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetPageNumber(v int32) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetPageSize(v int32) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetPolicies(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.Policies = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetRequestId(v string) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) SetTotalCount(v int64) *DescribeEnterpriseSnapshotPolicyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPolicies struct {
	// The creation time in UTC ISO 8601 format.
	//
	// example:
	//
	// 2023-06-24T06:03:35Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The snapshot cross-region copy information.
	CrossRegionCopyInfo *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty" type:"Struct"`
	// The snapshot policy description.
	//
	// example:
	//
	// xxx
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The list of bound cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// The snapshot managed status.
	//
	// example:
	//
	// false
	ManagedForEcs *bool `json:"ManagedForEcs,omitempty" xml:"ManagedForEcs,omitempty"`
	// The snapshot policy name.
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The snapshot policy ID.
	//
	// example:
	//
	// esp-xxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// the resource group
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot policy retention rule.
	RetainRule *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule `json:"RetainRule,omitempty" xml:"RetainRule,omitempty" type:"Struct"`
	// The snapshot policy schedule rule.
	Schedule *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// The special retention rules for the snapshot policy.
	SpecialRetainRules *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty" type:"Struct"`
	// The status. Valid values:
	//
	// - DISABLED
	//
	// - ENABLED
	//
	// example:
	//
	// DISABLED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The snapshot policy storage rule.
	StorageRule *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule `json:"StorageRule,omitempty" xml:"StorageRule,omitempty" type:"Struct"`
	// the pair tags
	Tags []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The number of bound targets.
	//
	// example:
	//
	// 10
	TargetCount *int32 `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
	// The type. Valid values:
	//
	// - DISK
	//
	// example:
	//
	// DISK
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetCrossRegionCopyInfo() *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo {
	return s.CrossRegionCopyInfo
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetDesc() *string {
	return s.Desc
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetDiskIds() []*string {
	return s.DiskIds
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetManagedForEcs() *bool {
	return s.ManagedForEcs
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetName() *string {
	return s.Name
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetRetainRule() *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule {
	return s.RetainRule
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetSchedule() *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule {
	return s.Schedule
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetSpecialRetainRules() *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules {
	return s.SpecialRetainRules
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetState() *string {
	return s.State
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetStorageRule() *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule {
	return s.StorageRule
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetTags() []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags {
	return s.Tags
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetTargetCount() *int32 {
	return s.TargetCount
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetCreateTime(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.CreateTime = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetCrossRegionCopyInfo(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.CrossRegionCopyInfo = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetDesc(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.Desc = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetDiskIds(v []*string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.DiskIds = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetManagedForEcs(v bool) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.ManagedForEcs = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetName(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetPolicyId(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetResourceGroupId(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetRetainRule(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.RetainRule = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetSchedule(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.Schedule = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetSpecialRetainRules(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.SpecialRetainRules = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetState(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.State = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetStorageRule(v *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.StorageRule = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetTags(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.Tags = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetTargetCount(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.TargetCount = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) SetTargetType(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies {
	s.TargetType = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPolicies) Validate() error {
	if s.CrossRegionCopyInfo != nil {
		if err := s.CrossRegionCopyInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RetainRule != nil {
		if err := s.RetainRule.Validate(); err != nil {
			return err
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	if s.SpecialRetainRules != nil {
		if err := s.SpecialRetainRules.Validate(); err != nil {
			return err
		}
	}
	if s.StorageRule != nil {
		if err := s.StorageRule.Validate(); err != nil {
			return err
		}
	}
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

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo struct {
	// Indicates whether cross-region copy is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The destination region information.
	Regions []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) GetRegions() []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions {
	return s.Regions
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) SetEnabled(v bool) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo {
	s.Enabled = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) SetRegions(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo {
	s.Regions = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfo) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions struct {
	// The snapshot copy destination region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of days to retain snapshots at the destination region.
	//
	// example:
	//
	// 7
	RetainDays *int32 `json:"RetainDays,omitempty" xml:"RetainDays,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) GetRetainDays() *int32 {
	return s.RetainDays
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) SetRegionId(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) SetRetainDays(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions {
	s.RetainDays = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesCrossRegionCopyInfoRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule struct {
	// The retention count.
	//
	// example:
	//
	// 10
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// The time interval of the retention rule. The unit is specified by the TimeUnit parameter. The value must be greater than 1.
	//
	// example:
	//
	// 14
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The retention time unit.
	//
	// example:
	//
	// DAYS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) GetNumber() *int32 {
	return s.Number
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) SetNumber(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule {
	s.Number = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) SetTimeInterval(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule {
	s.TimeInterval = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) SetTimeUnit(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule {
	s.TimeUnit = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesRetainRule) Validate() error {
	return dara.Validate(s)
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule struct {
	// The cron expression.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) SetCronExpression(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule {
	s.CronExpression = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSchedule) Validate() error {
	return dara.Validate(s)
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules struct {
	// Indicates whether special retention is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The special retention rules.
	Rules []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) GetRules() []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules {
	return s.Rules
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) SetEnabled(v bool) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules {
	s.Enabled = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) SetRules(v []*DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules {
	s.Rules = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRules) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules struct {
	// The special period unit.
	//
	// example:
	//
	// WEEKS
	SpecialPeriodUnit *string `json:"SpecialPeriodUnit,omitempty" xml:"SpecialPeriodUnit,omitempty"`
	// The time interval of the retention rule. The unit is specified by the TimeUnit parameter. The value must be greater than 1.
	//
	// example:
	//
	// 1
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The retention time unit.
	//
	// example:
	//
	// WEEKS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) GetSpecialPeriodUnit() *string {
	return s.SpecialPeriodUnit
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) SetSpecialPeriodUnit(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules {
	s.SpecialPeriodUnit = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) SetTimeInterval(v int32) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules {
	s.TimeInterval = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) SetTimeUnit(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules {
	s.TimeUnit = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesSpecialRetainRulesRules) Validate() error {
	return dara.Validate(s)
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule struct {
	// Indicates whether instant access is enabled for snapshots.
	//
	// example:
	//
	// false
	EnableImmediateAccess *bool `json:"EnableImmediateAccess,omitempty" xml:"EnableImmediateAccess,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) GetEnableImmediateAccess() *bool {
	return s.EnableImmediateAccess
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) SetEnableImmediateAccess(v bool) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule {
	s.EnableImmediateAccess = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesStorageRule) Validate() error {
	return dara.Validate(s)
}

type DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags struct {
	// The tag key of the resource.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) SetTagKey(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags {
	s.TagKey = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) SetTagValue(v string) *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags {
	s.TagValue = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponseBodyPoliciesTags) Validate() error {
	return dara.Validate(s)
}
