// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *DeleteVpcFirewallControlPolicyRequest
	GetAclUuid() *string
	SetLang(v string) *DeleteVpcFirewallControlPolicyRequest
	GetLang() *string
	SetVpcFirewallId(v string) *DeleteVpcFirewallControlPolicyRequest
	GetVpcFirewallId() *string
}

type DeleteVpcFirewallControlPolicyRequest struct {
	// The unique identity ID of the access control policy for virtual private cloud (VPC) firewalls at the border.
	//
	// When you delete an access control policy, you must provide the unique identity ID of the policy. You can invoke the [DescribeVpcFirewallControlPolicy](https://help.aliyun.com/document_detail/159758.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df2214****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The language of the request and response.
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the access control policy group for virtual private cloud (VPC) firewalls at the border. You can invoke the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to query the ID.
	//
	// Valid values:
	//
	// - If the virtual private cloud (VPC) firewalls at the border protect a CEN instance, the policy group ID is the instance ID of the CEN instance.
	//
	//    Example: cen-ervw0g12b5jbw*\*\*\*
	//
	// - If the virtual private cloud (VPC) firewalls at the border protect an Express Connect circuit, the policy group ID is the instance ID of the virtual private cloud (VPC) firewalls at the border.
	//
	//
	//
	//   Example: vfw-a42bbb7b887148c9*\*\*\*
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-a42bbb7b887148c91****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DeleteVpcFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DeleteVpcFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteVpcFirewallControlPolicyRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetAclUuid(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetLang(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
