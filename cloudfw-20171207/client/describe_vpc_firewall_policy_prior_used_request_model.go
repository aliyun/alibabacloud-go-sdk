// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallPolicyPriorUsedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVpcFirewallPolicyPriorUsedRequest
	GetLang() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallPolicyPriorUsedRequest
	GetVpcFirewallId() *string
}

type DescribeVpcFirewallPolicyPriorUsedRequest struct {
	// The language of the content within the request and response.
	//
	// Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the access control policy group. You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to query the ID.
	//
	// 	- If the VPC firewall is used to protect a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance.
	//
	//     Example: cen-ervw0g12b5jbw\\*\\*\\*\\*.
	//
	// 	- If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the ID of the VPC firewall.
	//
	//     Example: vfw-a42bbb7b887148c9\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) SetLang(v string) *DescribeVpcFirewallPolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallPolicyPriorUsedRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) Validate() error {
	return dara.Validate(s)
}
