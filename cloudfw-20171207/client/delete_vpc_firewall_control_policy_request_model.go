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
	// The UUID of the access control policy for the VPC firewall.
	//
	// When you delete an access control policy, you must provide the UUID of the policy. You can call the [DescribeVpcFirewallControlPolicy](https://help.aliyun.com/document_detail/159758.html) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df2214****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The language of the content within the request and response.
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the policy group for the VPC firewall. You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to query the ID.
	//
	// Valid values:
	//
	// - If the VPC firewall is used to protect a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance.
	//
	//   Example: cen-ervw0g12b5jbw\\*\\*\\*\\*
	//
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the ID of the VPC firewall instance.
	//
	//   Example: vfw-a42bbb7b887148c9\\*\\*\\*\\*
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
