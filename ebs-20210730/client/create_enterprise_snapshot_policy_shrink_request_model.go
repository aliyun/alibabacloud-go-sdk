// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseSnapshotPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetClientToken() *string
	SetCrossRegionCopyInfoShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetCrossRegionCopyInfoShrink() *string
	SetDesc(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetDesc() *string
	SetName(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetName() *string
	SetRegionId(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetResourceGroupId() *string
	SetRetainRuleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetRetainRuleShrink() *string
	SetScheduleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetScheduleShrink() *string
	SetSpecialRetainRulesShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetSpecialRetainRulesShrink() *string
	SetState(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetState() *string
	SetStorageRuleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetStorageRuleShrink() *string
	SetTag(v []*CreateEnterpriseSnapshotPolicyShrinkRequestTag) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetTag() []*CreateEnterpriseSnapshotPolicyShrinkRequestTag
	SetTargetType(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest
	GetTargetType() *string
}

type CreateEnterpriseSnapshotPolicyShrinkRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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
	RetainRuleShrink *string `json:"RetainRule,omitempty" xml:"RetainRule,omitempty"`
	// The rule for scheduling.
	//
	// This parameter is required.
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The special snapshot retention rules.
	SpecialRetainRulesShrink *string `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty"`
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
	StorageRuleShrink *string `json:"StorageRule,omitempty" xml:"StorageRule,omitempty"`
	// The list of tags.
	Tag []*CreateEnterpriseSnapshotPolicyShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s CreateEnterpriseSnapshotPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetCrossRegionCopyInfoShrink() *string {
	return s.CrossRegionCopyInfoShrink
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetDesc() *string {
	return s.Desc
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetRetainRuleShrink() *string {
	return s.RetainRuleShrink
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetScheduleShrink() *string {
	return s.ScheduleShrink
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetSpecialRetainRulesShrink() *string {
	return s.SpecialRetainRulesShrink
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetState() *string {
	return s.State
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetStorageRuleShrink() *string {
	return s.StorageRuleShrink
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetTag() []*CreateEnterpriseSnapshotPolicyShrinkRequestTag {
	return s.Tag
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetClientToken(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetCrossRegionCopyInfoShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.CrossRegionCopyInfoShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetDesc(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.Desc = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetName(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetRegionId(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetResourceGroupId(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetRetainRuleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.RetainRuleShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetScheduleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetSpecialRetainRulesShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.SpecialRetainRulesShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetState(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.State = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetStorageRuleShrink(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.StorageRuleShrink = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetTag(v []*CreateEnterpriseSnapshotPolicyShrinkRequestTag) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) SetTargetType(v string) *CreateEnterpriseSnapshotPolicyShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequest) Validate() error {
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

type CreateEnterpriseSnapshotPolicyShrinkRequestTag struct {
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

func (s CreateEnterpriseSnapshotPolicyShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequestTag) SetKey(v string) *CreateEnterpriseSnapshotPolicyShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequestTag) SetValue(v string) *CreateEnterpriseSnapshotPolicyShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
