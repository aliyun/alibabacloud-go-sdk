// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVpcId(v string) *ListAclPolicyRequest
	GetVpcId() *string
}

type ListAclPolicyRequest struct {
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAclPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListAclPolicyRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAclPolicyRequest) SetVpcId(v string) *ListAclPolicyRequest {
	s.VpcId = &v
	return s
}

func (s *ListAclPolicyRequest) Validate() error {
	return dara.Validate(s)
}
