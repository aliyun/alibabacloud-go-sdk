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
	// Ensures the idempotence of the request. Generate a parameter value from your client that is unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The snapshot replication information.
	CrossRegionCopyInfoShrink *string `json:"CrossRegionCopyInfo,omitempty" xml:"CrossRegionCopyInfo,omitempty"`
	// The description.
	//
	// example:
	//
	// xxx
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The Policy Name.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. You can call DescribeRegions to query the regions that support asynchronous replication.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The retention rule.
	//
	// This parameter is required.
	RetainRuleShrink *string `json:"RetainRule,omitempty" xml:"RetainRule,omitempty"`
	// The schedule rule.
	//
	// This parameter is required.
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The special retention rules.
	SpecialRetainRulesShrink *string `json:"SpecialRetainRules,omitempty" xml:"SpecialRetainRules,omitempty"`
	// The status. Valid values:
	//
	// - DISABLED
	//
	// - ENABLED
	//
	// example:
	//
	// ENABLED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The advanced snapshot feature.
	StorageRuleShrink *string `json:"StorageRule,omitempty" xml:"StorageRule,omitempty"`
	// The tag key-value pairs. Valid values of n: 1 to 20.
	Tag []*CreateEnterpriseSnapshotPolicyShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type. Valid values:
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
	// The tag key of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
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
