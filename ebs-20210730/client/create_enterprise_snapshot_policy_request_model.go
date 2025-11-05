// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEnterpriseSnapshotPolicyRequest
	GetClientToken() *string
	SetCrossRegionCopyInfo(v *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) *CreateEnterpriseSnapshotPolicyRequest
	GetCrossRegionCopyInfo() *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo
	SetDesc(v string) *CreateEnterpriseSnapshotPolicyRequest
	GetDesc() *string
	SetName(v string) *CreateEnterpriseSnapshotPolicyRequest
	GetName() *string
	SetRegionId(v string) *CreateEnterpriseSnapshotPolicyRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEnterpriseSnapshotPolicyRequest
	GetResourceGroupId() *string
	SetRetainRule(v *CreateEnterpriseSnapshotPolicyRequestRetainRule) *CreateEnterpriseSnapshotPolicyRequest
	GetRetainRule() *CreateEnterpriseSnapshotPolicyRequestRetainRule
	SetSchedule(v *CreateEnterpriseSnapshotPolicyRequestSchedule) *CreateEnterpriseSnapshotPolicyRequest
	GetSchedule() *CreateEnterpriseSnapshotPolicyRequestSchedule
	SetSpecialRetainRules(v *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) *CreateEnterpriseSnapshotPolicyRequest
	GetSpecialRetainRules() *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules
	SetState(v string) *CreateEnterpriseSnapshotPolicyRequest
	GetState() *string
	SetStorageRule(v *CreateEnterpriseSnapshotPolicyRequestStorageRule) *CreateEnterpriseSnapshotPolicyRequest
	GetStorageRule() *CreateEnterpriseSnapshotPolicyRequestStorageRule
	SetTag(v []*CreateEnterpriseSnapshotPolicyRequestTag) *CreateEnterpriseSnapshotPolicyRequest
	GetTag() []*CreateEnterpriseSnapshotPolicyRequestTag
	SetTargetType(v string) *CreateEnterpriseSnapshotPolicyRequest
	GetTargetType() *string
}

type CreateEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Snapshot replication destination information.
	CrossRegionCopyInfo *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty" type:"Struct"`
	// The description of the policy.
	//
	// example:
	//
	// xxx
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID . You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which snapshot policy is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the snapshot policy.
	//
	// example:
	//
	// xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot retention rule.
	//
	// This parameter is required.
	RetainRule *CreateEnterpriseSnapshotPolicyRequestRetainRule `json:"RetainRule,omitempty" xml:"RetainRule,omitempty" type:"Struct"`
	// The rule for scheduling.
	//
	// This parameter is required.
	Schedule *CreateEnterpriseSnapshotPolicyRequestSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// The special snapshot retention rules.
	SpecialRetainRules *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty" type:"Struct"`
	// The status of the policy. Valid values:
	//
	// - ENABLED: Enable snapshot policy execution.
	//
	// - DISABLED: Disable snapshot policy execution.
	//
	// example:
	//
	// ENABLED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Advanced snapshot features.
	StorageRule *CreateEnterpriseSnapshotPolicyRequestStorageRule `json:"StorageRule,omitempty" xml:"StorageRule,omitempty" type:"Struct"`
	// The list of tags.
	Tag []*CreateEnterpriseSnapshotPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Binding target type, valid value:
	//
	// - DISK
	//
	// This parameter is required.
	//
	// example:
	//
	// DISK
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetCrossRegionCopyInfo() *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	return s.CrossRegionCopyInfo
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetDesc() *string {
	return s.Desc
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetRetainRule() *CreateEnterpriseSnapshotPolicyRequestRetainRule {
	return s.RetainRule
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetSchedule() *CreateEnterpriseSnapshotPolicyRequestSchedule {
	return s.Schedule
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetSpecialRetainRules() *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	return s.SpecialRetainRules
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetState() *string {
	return s.State
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetStorageRule() *CreateEnterpriseSnapshotPolicyRequestStorageRule {
	return s.StorageRule
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetTag() []*CreateEnterpriseSnapshotPolicyRequestTag {
	return s.Tag
}

func (s *CreateEnterpriseSnapshotPolicyRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetCrossRegionCopyInfo(v *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) *CreateEnterpriseSnapshotPolicyRequest {
	s.CrossRegionCopyInfo = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetDesc(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.Desc = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetName(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetResourceGroupId(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetRetainRule(v *CreateEnterpriseSnapshotPolicyRequestRetainRule) *CreateEnterpriseSnapshotPolicyRequest {
	s.RetainRule = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetSchedule(v *CreateEnterpriseSnapshotPolicyRequestSchedule) *CreateEnterpriseSnapshotPolicyRequest {
	s.Schedule = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetSpecialRetainRules(v *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) *CreateEnterpriseSnapshotPolicyRequest {
	s.SpecialRetainRules = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetState(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.State = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetStorageRule(v *CreateEnterpriseSnapshotPolicyRequestStorageRule) *CreateEnterpriseSnapshotPolicyRequest {
	s.StorageRule = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetTag(v []*CreateEnterpriseSnapshotPolicyRequestTag) *CreateEnterpriseSnapshotPolicyRequest {
	s.Tag = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) SetTargetType(v string) *CreateEnterpriseSnapshotPolicyRequest {
	s.TargetType = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequest) Validate() error {
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo struct {
	// Whether cross-region replication is enabled. The range of values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of destination regions.
	Regions []*CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) GetRegions() []*CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	return s.Regions
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) SetEnabled(v bool) *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	s.Enabled = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) SetRegions(v []*CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	s.Regions = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) Validate() error {
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

type CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions struct {
	// The region ID of the destination. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of days to retain the destination snapshot. The range of values is greater than 1.
	//
	// example:
	//
	// 7
	RetainDays *int32 `json:"RetainDays,omitempty" xml:"RetainDays,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) GetRetainDays() *int32 {
	return s.RetainDays
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) SetRegionId(v string) *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	s.RegionId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) SetRetainDays(v int32) *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	s.RetainDays = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) Validate() error {
	return dara.Validate(s)
}

type CreateEnterpriseSnapshotPolicyRequestRetainRule struct {
	// Maximum number of retained snapshots.
	//
	// example:
	//
	// 10
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// The time interval , valid value greater than 1.
	//
	// example:
	//
	// 14
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The unit of time, valid values:
	//
	// - DAYS
	//
	// - WEEKS
	//
	// example:
	//
	// DAYS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestRetainRule) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestRetainRule) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) GetNumber() *int32 {
	return s.Number
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) SetNumber(v int32) *CreateEnterpriseSnapshotPolicyRequestRetainRule {
	s.Number = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) SetTimeInterval(v int32) *CreateEnterpriseSnapshotPolicyRequestRetainRule {
	s.TimeInterval = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) SetTimeUnit(v string) *CreateEnterpriseSnapshotPolicyRequestRetainRule {
	s.TimeUnit = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestRetainRule) Validate() error {
	return dara.Validate(s)
}

type CreateEnterpriseSnapshotPolicyRequestSchedule struct {
	// The time when the policy will to be scheduled. Valid values: Set the parameter in a cron expression.
	//
	// For example, you can use 0 0 4 1/1 	- ? to specify 04:00:00 (UTC+8) on the first day of each month.
	//
	// This parameter is required.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestSchedule) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestSchedule) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestSchedule) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateEnterpriseSnapshotPolicyRequestSchedule) SetCronExpression(v string) *CreateEnterpriseSnapshotPolicyRequestSchedule {
	s.CronExpression = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSchedule) Validate() error {
	return dara.Validate(s)
}

type CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules struct {
	// Indicates whether the special retention is enabled.
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The special retention rules.
	Rules []*CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) GetRules() []*CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	return s.Rules
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) SetEnabled(v bool) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	s.Enabled = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) SetRules(v []*CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	s.Rules = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRules) Validate() error {
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

type CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules struct {
	// The periodic unit for specially retained snapshots. If configured to WEEKS, it provides special retention for the first snapshot of each week. The retention period is determined by TimeUnit and TimeInterval. The range of values are:
	//
	// - WEEKS
	//
	// - MONTHS
	//
	// - YEARS
	//
	// example:
	//
	// WEEKS
	SpecialPeriodUnit *string `json:"SpecialPeriodUnit,omitempty" xml:"SpecialPeriodUnit,omitempty"`
	// Retention Time Value. The range of values is greater than 1.
	//
	// example:
	//
	// 14
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// Retention time unit for special snapshots. The range of values:
	//
	// - DAYS
	//
	// - WEEKS
	//
	// example:
	//
	// WEEKS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GetSpecialPeriodUnit() *string {
	return s.SpecialPeriodUnit
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetSpecialPeriodUnit(v string) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.SpecialPeriodUnit = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetTimeInterval(v int32) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.TimeInterval = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetTimeUnit(v string) *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.TimeUnit = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) Validate() error {
	return dara.Validate(s)
}

type CreateEnterpriseSnapshotPolicyRequestStorageRule struct {
	// Whether to enable the rapid availability of snapshots. The range of values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	EnableImmediateAccess *bool `json:"EnableImmediateAccess,omitempty" xml:"EnableImmediateAccess,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestStorageRule) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestStorageRule) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestStorageRule) GetEnableImmediateAccess() *bool {
	return s.EnableImmediateAccess
}

func (s *CreateEnterpriseSnapshotPolicyRequestStorageRule) SetEnableImmediateAccess(v bool) *CreateEnterpriseSnapshotPolicyRequestStorageRule {
	s.EnableImmediateAccess = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestStorageRule) Validate() error {
	return dara.Validate(s)
}

type CreateEnterpriseSnapshotPolicyRequestTag struct {
	// The key of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEnterpriseSnapshotPolicyRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEnterpriseSnapshotPolicyRequestTag) SetKey(v string) *CreateEnterpriseSnapshotPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestTag) SetValue(v string) *CreateEnterpriseSnapshotPolicyRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyRequestTag) Validate() error {
	return dara.Validate(s)
}
