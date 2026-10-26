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
	SetProjectId(v int64) *SetPolicyApprovalConfigShrinkRequest
	GetProjectId() *int64
	SetRegionId(v string) *SetPolicyApprovalConfigShrinkRequest
	GetRegionId() *string
}

type SetPolicyApprovalConfigShrinkRequest struct {
	// The O&M approval configuration in the control policy.
	//
	// This parameter is required.
	ApprovalConfigShrink *string `json:"ApprovalConfig,omitempty" xml:"ApprovalConfig,omitempty"`
	// The Bastionhost instance ID.
	//
	// > You can call the DescribeInstances operation to obtain the Bastionhost instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *SetPolicyApprovalConfigShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *SetPolicyApprovalConfigShrinkRequest) SetProjectId(v int64) *SetPolicyApprovalConfigShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *SetPolicyApprovalConfigShrinkRequest) SetRegionId(v string) *SetPolicyApprovalConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyApprovalConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
