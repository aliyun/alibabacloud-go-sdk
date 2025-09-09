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
	SetRegionId(v string) *SetPolicyIPAclConfigShrinkRequest
	GetRegionId() *string
}

type SetPolicyIPAclConfigShrinkRequest struct {
	// The access control settings for source IP addresses.
	//
	// This parameter is required.
	IPAclConfigShrink *string `json:"IPAclConfig,omitempty" xml:"IPAclConfig,omitempty"`
	// The bastion host ID.
	//
	// > You can call the DescribeInstances operation to query the bastion host ID.[](~~153281~~)
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
	// 3
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *SetPolicyIPAclConfigShrinkRequest) SetRegionId(v string) *SetPolicyIPAclConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyIPAclConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
