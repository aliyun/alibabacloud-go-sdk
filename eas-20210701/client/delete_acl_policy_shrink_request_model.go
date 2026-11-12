// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclPolicyListShrink(v string) *DeleteAclPolicyShrinkRequest
	GetAclPolicyListShrink() *string
	SetVpcId(v string) *DeleteAclPolicyShrinkRequest
	GetVpcId() *string
}

type DeleteAclPolicyShrinkRequest struct {
	// The whitelist of IP address ranges that are allowed to access the VPC. This parameter is required. If this parameter is not specified, an InvalidParameter error is returned, indicating that AclPolicyList must be set.
	AclPolicyListShrink *string `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty"`
	// The ID of the VPC. For information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteAclPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyShrinkRequest) GetAclPolicyListShrink() *string {
	return s.AclPolicyListShrink
}

func (s *DeleteAclPolicyShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteAclPolicyShrinkRequest) SetAclPolicyListShrink(v string) *DeleteAclPolicyShrinkRequest {
	s.AclPolicyListShrink = &v
	return s
}

func (s *DeleteAclPolicyShrinkRequest) SetVpcId(v string) *DeleteAclPolicyShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteAclPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
