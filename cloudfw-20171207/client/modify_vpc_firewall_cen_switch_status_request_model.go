// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallCenSwitchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallSwitch(v string) *ModifyVpcFirewallCenSwitchStatusRequest
	GetFirewallSwitch() *string
	SetLang(v string) *ModifyVpcFirewallCenSwitchStatusRequest
	GetLang() *string
	SetMemberUid(v string) *ModifyVpcFirewallCenSwitchStatusRequest
	GetMemberUid() *string
	SetVpcFirewallId(v string) *ModifyVpcFirewallCenSwitchStatusRequest
	GetVpcFirewallId() *string
}

type ModifyVpcFirewallCenSwitchStatusRequest struct {
	// Specifies whether to enable the VPC firewall. Valid values:
	//
	// 	- **open**: yes
	//
	// 	- **close**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallCenList](https://help.aliyun.com/document_detail/345777.html) operation to query the instance IDs of VPC firewalls.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallCenSwitchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) GetFirewallSwitch() *string {
	return s.FirewallSwitch
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetFirewallSwitch(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetLang(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetMemberUid(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) Validate() error {
	return dara.Validate(s)
}
