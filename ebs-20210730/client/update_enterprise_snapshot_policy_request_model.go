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
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique across different requests. The ClientToken value supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cross-region copy destination information.
	CrossRegionCopyInfo *UpdateEnterpriseSnapshotPolicyRequestCrossRegionCopyInfo `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty" type:"Struct"`
	// The description of the snapshot policy.
	//
	// example:
	//
	// xxx
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The ID of the policy to modify.
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The snapshot policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esp-xxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The retention rule.
	RetainRule *UpdateEnterpriseSnapshotPolicyRequestRetainRule `json:"RetainRule,omitempty" xml:"RetainRule,omitempty" type:"Struct"`
	// The schedule rule.
	Schedule *UpdateEnterpriseSnapshotPolicyRequestSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Struct"`
	// The special retention rules.
	SpecialRetainRules *UpdateEnterpriseSnapshotPolicyRequestSpecialRetainRules `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty" type:"Struct"`
	// The status of the snapshot policy. Valid values:
	//
	// - ENABLED
	//
	// - DISABLED
	//
	// example:
	//
	// ENABLED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The advanced snapshot feature.
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
	// Specifies whether to enable cross-region replication. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The destination region information.
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
	// The destination region ID for snapshot replication. You can call [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) to query the region information of existing asynchronous replication relationships.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of days to retain snapshots in the destination region. The value must be greater than 1.
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
	// The number of snapshots to retain. Valid values: 1 to 256.
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
	// The unit of the retention time. Valid values:
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
	// The execution cycle and time of the policy. A cron expression is used.
	//
	// For example, `0 0 4 1/1 	- ?` specifies that the snapshot operation is performed at 04:00 every day, starting from the first day of each month.
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
	// Specifies whether to enable special retention. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of special retention rules. Multiple rules are supported.
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
	// The period unit for special retention snapshots. For example, if this parameter is set to WEEKS, the first snapshot of each week is given special retention. The retention duration is determined by the TimeUnit and TimeInterval parameters. Valid values:
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
	// The time interval of the retention rule. The unit is specified by the TimeUnit parameter. The value must be greater than 1.
	//
	// example:
	//
	// 30
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The unit of the retention time for special snapshots. Valid values:
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
	// Specifies whether to enable instant access for snapshots. Valid values:
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
