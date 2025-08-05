// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallCenConfigureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyVpcFirewallCenConfigureRequest
	GetLang() *string
	SetMemberUid(v string) *ModifyVpcFirewallCenConfigureRequest
	GetMemberUid() *string
	SetVpcFirewallId(v string) *ModifyVpcFirewallCenConfigureRequest
	GetVpcFirewallId() *string
	SetVpcFirewallName(v string) *ModifyVpcFirewallCenConfigureRequest
	GetVpcFirewallName() *string
}

type ModifyVpcFirewallCenConfigureRequest struct {
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
	// The instance name of the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test instance
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s ModifyVpcFirewallCenConfigureRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyVpcFirewallCenConfigureRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *ModifyVpcFirewallCenConfigureRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallCenConfigureRequest) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetLang(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetMemberUid(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetVpcFirewallName(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) Validate() error {
	return dara.Validate(s)
}
