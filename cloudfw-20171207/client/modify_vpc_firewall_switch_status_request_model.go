// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallSwitchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallSwitch(v string) *ModifyVpcFirewallSwitchStatusRequest
	GetFirewallSwitch() *string
	SetLang(v string) *ModifyVpcFirewallSwitchStatusRequest
	GetLang() *string
	SetMemberUid(v string) *ModifyVpcFirewallSwitchStatusRequest
	GetMemberUid() *string
	SetVpcFirewallId(v string) *ModifyVpcFirewallSwitchStatusRequest
	GetVpcFirewallId() *string
}

type ModifyVpcFirewallSwitchStatusRequest struct {
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
	// > You can call the [DescribeVpcFirewallList](https://help.aliyun.com/document_detail/342932.html) operation to query the instance IDs of VPC firewalls.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallSwitchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusRequest) GetFirewallSwitch() *string {
	return s.FirewallSwitch
}

func (s *ModifyVpcFirewallSwitchStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyVpcFirewallSwitchStatusRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *ModifyVpcFirewallSwitchStatusRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetFirewallSwitch(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetLang(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetMemberUid(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) Validate() error {
	return dara.Validate(s)
}
