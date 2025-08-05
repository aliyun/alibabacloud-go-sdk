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
	// The language of the content within the request and response.
	//
	// Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member in Cloud Firewall.
	//
	// example:
	//
	// 1766185894104675
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-3547deab1c9b4190a53f
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
