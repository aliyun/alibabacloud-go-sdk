// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseSnapshotPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetClientToken() *string
	SetCrossRegionCopyInfoShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetCrossRegionCopyInfoShrink() *string
	SetDesc(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetDesc() *string
	SetName(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetName() *string
	SetPolicyId(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetPolicyId() *string
	SetRegionId(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetRegionId() *string
	SetRetainRuleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetRetainRuleShrink() *string
	SetScheduleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetScheduleShrink() *string
	SetSpecialRetainRulesShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetSpecialRetainRulesShrink() *string
	SetState(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetState() *string
	SetStorageRuleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest
	GetStorageRuleShrink() *string
}

type UpdateEnterpriseSnapshotPolicyShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Snapshot replication destination information.
	CrossRegionCopyInfoShrink *string `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty"`
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
	RetainRuleShrink *string `json:"RetainRule,omitempty" xml:"RetainRule,omitempty"`
	// The rule for scheduling.
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The special snapshot retention rules.
	SpecialRetainRulesShrink *string `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty"`
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
	StorageRuleShrink *string `json:"StorageRule,omitempty" xml:"StorageRule,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetCrossRegionCopyInfoShrink() *string {
	return s.CrossRegionCopyInfoShrink
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetDesc() *string {
	return s.Desc
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetRetainRuleShrink() *string {
	return s.RetainRuleShrink
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetScheduleShrink() *string {
	return s.ScheduleShrink
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetSpecialRetainRulesShrink() *string {
	return s.SpecialRetainRulesShrink
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetState() *string {
	return s.State
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) GetStorageRuleShrink() *string {
	return s.StorageRuleShrink
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetClientToken(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetCrossRegionCopyInfoShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.CrossRegionCopyInfoShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetDesc(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.Desc = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetName(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetPolicyId(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetRegionId(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetRetainRuleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.RetainRuleShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetScheduleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetSpecialRetainRulesShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.SpecialRetainRulesShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetState(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.State = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) SetStorageRuleShrink(v string) *UpdateEnterpriseSnapshotPolicyShrinkRequest {
	s.StorageRuleShrink = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
