// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEnterpriseSnapshotPolicyRequest
	GetClientToken() *string
	SetCrossRegionCopyInfo(v *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) *UpdateEnterpriseSnapshotPolicyRequest
	GetCrossRegionCopyInfo() *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo
	SetDesc(v string) *UpdateEnterpriseSnapshotPolicyRequest
	GetDesc() *string
	SetName(v string) *UpdateEnterpriseSnapshotPolicyRequest
	GetName() *string
	SetPolicyId(v string) *UpdateEnterpriseSnapshotPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *UpdateEnterpriseSnapshotPolicyRequest
	GetRegionId() *string
	SetRetainRule(v *UpdateEnterpriseSnapshotPolicyRequestRetainRule) *UpdateEnterpriseSnapshotPolicyRequest
	GetRetainRule() *UpdateEnterpriseSnapshotPolicyRequestRetainRule
	SetSchedule(v *UpdateEnterpriseSnapshotPolicyRequestSchedule) *UpdateEnterpriseSnapshotPolicyRequest
	GetSchedule() *UpdateEnterpriseSnapshotPolicyRequestSchedule
	SetSpecialRetainRules(v *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) *UpdateEnterpriseSnapshotPolicyRequest
	GetSpecialRetainRules() *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules
	SetState(v string) *UpdateEnterpriseSnapshotPolicyRequest
	GetState() *string
	SetStorageRule(v *UpdateEnterpriseSnapshotPolicyRequestStorageRule) *UpdateEnterpriseSnapshotPolicyRequest
	GetStorageRule() *UpdateEnterpriseSnapshotPolicyRequestStorageRule
}

type UpdateEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Snapshot replication destination information.
	CrossRegionCopyInfo *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty" type:"Struct"`
	// The description of the policy.
	//
	// example:
	//
	// xxx
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The id of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// esp-xxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID . You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which snapshot policy is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Snapshot retention rule.
	RetainRule *UpdateEnterpriseSnapshotPolicyRequestRetainRule `json:"RetainRule,omitempty" xml:"RetainRule,omitempty" type:"Struct"`
	// The rule for scheduling.
	Schedule *UpdateEnterpriseSnapshotPolicyRequestSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// The special snapshot retention rules.
	SpecialRetainRules *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty" type:"Struct"`
	// The status of the policy. Valid values:
	//
	// 	- **ENABLED**: Enable snapshot policy execution.
	//
	// 	- **DISABLED**: Disable snapshot policy execution.
	//
	// example:
	//
	// ENABLED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Advanced snapshot features.
	StorageRule *UpdateEnterpriseSnapshotPolicyRequestStorageRule `json:"StorageRule,omitempty" xml:"StorageRule,omitempty" type:"Struct"`
}

func (s UpdateEnterpriseSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetCrossRegionCopyInfo() *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	return s.CrossRegionCopyInfo
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetDesc() *string {
	return s.Desc
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetRetainRule() *UpdateEnterpriseSnapshotPolicyRequestRetainRule {
	return s.RetainRule
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetSchedule() *UpdateEnterpriseSnapshotPolicyRequestSchedule {
	return s.Schedule
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetSpecialRetainRules() *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	return s.SpecialRetainRules
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetState() *string {
	return s.State
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) GetStorageRule() *UpdateEnterpriseSnapshotPolicyRequestStorageRule {
	return s.StorageRule
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetCrossRegionCopyInfo(v *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) *UpdateEnterpriseSnapshotPolicyRequest {
	s.CrossRegionCopyInfo = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetDesc(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.Desc = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetName(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetPolicyId(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetRetainRule(v *UpdateEnterpriseSnapshotPolicyRequestRetainRule) *UpdateEnterpriseSnapshotPolicyRequest {
	s.RetainRule = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetSchedule(v *UpdateEnterpriseSnapshotPolicyRequestSchedule) *UpdateEnterpriseSnapshotPolicyRequest {
	s.Schedule = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetSpecialRetainRules(v *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) *UpdateEnterpriseSnapshotPolicyRequest {
	s.SpecialRetainRules = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetState(v string) *UpdateEnterpriseSnapshotPolicyRequest {
	s.State = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) SetStorageRule(v *UpdateEnterpriseSnapshotPolicyRequestStorageRule) *UpdateEnterpriseSnapshotPolicyRequest {
	s.StorageRule = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequest) Validate() error {
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
	return nil
}

type UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo struct {
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
	// Destination region information.
	Regions []*UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) GetRegions() []*UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	return s.Regions
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) SetEnabled(v bool) *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	s.Enabled = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) SetRegions(v []*UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo {
	s.Regions = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo) Validate() error {
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

type UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions struct {
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which async replication is supported.
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

func (s UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) GetRetainDays() *int32 {
	return s.RetainDays
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) SetRegionId(v string) *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	s.RegionId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) SetRetainDays(v int32) *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions {
	s.RetainDays = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfoRegions) Validate() error {
	return dara.Validate(s)
}

type UpdateEnterpriseSnapshotPolicyRequestRetainRule struct {
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

func (s UpdateEnterpriseSnapshotPolicyRequestRetainRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestRetainRule) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) GetNumber() *int32 {
	return s.Number
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) SetNumber(v int32) *UpdateEnterpriseSnapshotPolicyRequestRetainRule {
	s.Number = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) SetTimeInterval(v int32) *UpdateEnterpriseSnapshotPolicyRequestRetainRule {
	s.TimeInterval = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) SetTimeUnit(v string) *UpdateEnterpriseSnapshotPolicyRequestRetainRule {
	s.TimeUnit = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestRetainRule) Validate() error {
	return dara.Validate(s)
}

type UpdateEnterpriseSnapshotPolicyRequestSchedule struct {
	// The time when the policy will to be scheduled. Valid values: Set the parameter in a cron expression.
	//
	// For example, you can use `0 0 4 1/1 	- ?` to specify 04:00:00 (UTC+8) on the first day of each month.
	//
	// This parameter is required.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestSchedule) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestSchedule) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSchedule) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSchedule) SetCronExpression(v string) *UpdateEnterpriseSnapshotPolicyRequestSchedule {
	s.CronExpression = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSchedule) Validate() error {
	return dara.Validate(s)
}

type UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules struct {
	// Indicates whether the special retention is enabled.
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The special retention rules.
	Rules []*UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) GetRules() []*UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	return s.Rules
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) SetEnabled(v bool) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	s.Enabled = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) SetRules(v []*UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules {
	s.Rules = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules) Validate() error {
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

type UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules struct {
	// The periodic unit for specially retained snapshots. If configured to WEEKS, it provides special retention for the first snapshot of each week. The retention period is determined by TimeUnit and TimeInterval. The range of values are:
	//
	// - WEEKS
	//
	// - MONTHS
	//
	// - YEARS"
	//
	// example:
	//
	// WEEKS
	SpecialPeriodUnit *string `json:"SpecialPeriodUnit,omitempty" xml:"SpecialPeriodUnit,omitempty"`
	// Retention Time Value. The range of values is greater than 1.
	//
	// example:
	//
	// 30
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

func (s UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GetSpecialPeriodUnit() *string {
	return s.SpecialPeriodUnit
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetSpecialPeriodUnit(v string) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.SpecialPeriodUnit = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetTimeInterval(v int32) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.TimeInterval = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) SetTimeUnit(v string) *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules {
	s.TimeUnit = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRulesRules) Validate() error {
	return dara.Validate(s)
}

type UpdateEnterpriseSnapshotPolicyRequestStorageRule struct {
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

func (s UpdateEnterpriseSnapshotPolicyRequestStorageRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyRequestStorageRule) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyRequestStorageRule) GetEnableImmediateAccess() *bool {
	return s.EnableImmediateAccess
}

func (s *UpdateEnterpriseSnapshotPolicyRequestStorageRule) SetEnableImmediateAccess(v bool) *UpdateEnterpriseSnapshotPolicyRequestStorageRule {
	s.EnableImmediateAccess = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyRequestStorageRule) Validate() error {
	return dara.Validate(s)
}
