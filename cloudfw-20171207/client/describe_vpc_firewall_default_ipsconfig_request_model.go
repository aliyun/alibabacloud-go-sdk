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
	// The UID of the member account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// - If the VPC firewall protects a network instance in a Cloud Enterprise Network (CEN) instance, set this parameter to the ID of the CEN instance. To query the ID of a CEN instance of Basic Edition, call the [DescribeVpcFirewallCenList](https://help.aliyun.com/document_detail/345777.html) operation. To query the ID of a CEN instance of Enterprise Edition, call the [DescribeTrFirewallsV2List](https://help.aliyun.com/document_detail/2384695.html) operation.
	//
	// - If the VPC firewall protects traffic between two VPCs connected by an Express Connect circuit, set this parameter to the ID of the VPC firewall instance. To query the ID of the VPC firewall instance, call the [DescribeVpcFirewallList](https://help.aliyun.com/document_detail/342932.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-****
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
