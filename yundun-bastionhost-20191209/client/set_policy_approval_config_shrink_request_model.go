// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyApprovalConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalConfigShrink(v string) *SetPolicyApprovalConfigShrinkRequest
	GetApprovalConfigShrink() *string
	SetInstanceId(v string) *SetPolicyApprovalConfigShrinkRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyApprovalConfigShrinkRequest
	GetPolicyId() *string
	SetRegionId(v string) *SetPolicyApprovalConfigShrinkRequest
	GetRegionId() *string
}

type SetPolicyApprovalConfigShrinkRequest struct {
	// The O&M approval setting in the control policy.
	//
	// This parameter is required.
	ApprovalConfigShrink *string `json:"ApprovalConfig,omitempty" xml:"ApprovalConfig,omitempty"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// >  You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyApprovalConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyApprovalConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyApprovalConfigShrinkRequest) GetApprovalConfigShrink() *string {
	return s.ApprovalConfigShrink
}

func (s *SetPolicyApprovalConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyApprovalConfigShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyApprovalConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyApprovalConfigShrinkRequest) SetApprovalConfigShrink(v string) *SetPolicyApprovalConfigShrinkRequest {
	s.ApprovalConfigShrink = &v
	return s
}

func (s *SetPolicyApprovalConfigShrinkRequest) SetInstanceId(v string) *SetPolicyApprovalConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyApprovalConfigShrinkRequest) SetPolicyId(v string) *SetPolicyApprovalConfigShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyApprovalConfigShrinkRequest) SetRegionId(v string) *SetPolicyApprovalConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyApprovalConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
