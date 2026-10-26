// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyIPAclConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIPAclConfigShrink(v string) *SetPolicyIPAclConfigShrinkRequest
	GetIPAclConfigShrink() *string
	SetInstanceId(v string) *SetPolicyIPAclConfigShrinkRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyIPAclConfigShrinkRequest
	GetPolicyId() *string
	SetProjectId(v int64) *SetPolicyIPAclConfigShrinkRequest
	GetProjectId() *int64
	SetRegionId(v string) *SetPolicyIPAclConfigShrinkRequest
	GetRegionId() *string
}

type SetPolicyIPAclConfigShrinkRequest struct {
	// The source IP address restriction configuration.
	//
	// This parameter is required.
	IPAclConfigShrink *string `json:"IPAclConfig,omitempty" xml:"IPAclConfig,omitempty"`
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the bastion host instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyIPAclConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyIPAclConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyIPAclConfigShrinkRequest) GetIPAclConfigShrink() *string {
	return s.IPAclConfigShrink
}

func (s *SetPolicyIPAclConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyIPAclConfigShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyIPAclConfigShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SetPolicyIPAclConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyIPAclConfigShrinkRequest) SetIPAclConfigShrink(v string) *SetPolicyIPAclConfigShrinkRequest {
	s.IPAclConfigShrink = &v
	return s
}

func (s *SetPolicyIPAclConfigShrinkRequest) SetInstanceId(v string) *SetPolicyIPAclConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyIPAclConfigShrinkRequest) SetPolicyId(v string) *SetPolicyIPAclConfigShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyIPAclConfigShrinkRequest) SetProjectId(v int64) *SetPolicyIPAclConfigShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *SetPolicyIPAclConfigShrinkRequest) SetRegionId(v string) *SetPolicyIPAclConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyIPAclConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
