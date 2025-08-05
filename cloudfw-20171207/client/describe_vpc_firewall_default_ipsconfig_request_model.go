// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDefaultIPSConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberUid(v string) *DescribeVpcFirewallDefaultIPSConfigRequest
	GetMemberUid() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallDefaultIPSConfigRequest
	GetVpcFirewallId() *string
}

type DescribeVpcFirewallDefaultIPSConfigRequest struct {
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// 	- If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. You can call the [DescribeVpcFirewallCenList](https://help.aliyun.com/document_detail/345777.html) operation to query the IDs of CEN instances.
	//
	// 	- If the VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall. You can call the [DescribeVpcFirewallList](https://help.aliyun.com/document_detail/342932.html) operation to query the instance IDs of VPC firewalls.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallDefaultIPSConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) SetMemberUid(v string) *DescribeVpcFirewallDefaultIPSConfigRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallDefaultIPSConfigRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) Validate() error {
	return dara.Validate(s)
}
