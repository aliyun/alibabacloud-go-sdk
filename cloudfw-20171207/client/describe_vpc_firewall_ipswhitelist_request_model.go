// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallIPSWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVpcFirewallIPSWhitelistRequest
	GetLang() *string
	SetMemberUid(v int64) *DescribeVpcFirewallIPSWhitelistRequest
	GetMemberUid() *int64
	SetVpcFirewallId(v string) *DescribeVpcFirewallIPSWhitelistRequest
	GetVpcFirewallId() *string
}

type DescribeVpcFirewallIPSWhitelistRequest struct {
	// The language of the request and response messages.
	//
	// Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the Cloud Firewall member account.
	//
	// example:
	//
	// 176618****104675
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// - If the VPC firewall protects a network instance in a Cloud Enterprise Network (CEN) instance, set the value to the ID of the CEN instance. For a CEN instance of Basic Edition, you can call the [DescribeVpcFirewallCenList](https://help.aliyun.com/document_detail/345777.html) operation to query the instance ID. For a CEN instance of Enterprise Edition, you can call the [DescribeTrFirewallsV2List](https://help.aliyun.com/document_detail/2384695.html) operation to query the instance ID.
	//
	// - If the VPC firewall protects traffic between two VPCs that are connected by an Express Connect circuit, set the value to the instance ID of the VPC firewall. You can call the [DescribeVpcFirewallList](https://help.aliyun.com/document_detail/342932.html) operation to query the instance ID of the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-***
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallIPSWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallIPSWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) SetLang(v string) *DescribeVpcFirewallIPSWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) SetMemberUid(v int64) *DescribeVpcFirewallIPSWhitelistRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallIPSWhitelistRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
