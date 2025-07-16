// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclPolicyListShrink(v string) *CreateAclPolicyShrinkRequest
	GetAclPolicyListShrink() *string
	SetVpcId(v string) *CreateAclPolicyShrinkRequest
	GetVpcId() *string
}

type CreateAclPolicyShrinkRequest struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway.
	//
	// This parameter is required.
	AclPolicyListShrink *string `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty"`
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateAclPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAclPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyShrinkRequest) GetAclPolicyListShrink() *string {
	return s.AclPolicyListShrink
}

func (s *CreateAclPolicyShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateAclPolicyShrinkRequest) SetAclPolicyListShrink(v string) *CreateAclPolicyShrinkRequest {
	s.AclPolicyListShrink = &v
	return s
}

func (s *CreateAclPolicyShrinkRequest) SetVpcId(v string) *CreateAclPolicyShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateAclPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
