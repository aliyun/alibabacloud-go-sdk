// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallAclEngineModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberUid(v string) *ModifyVpcFirewallAclEngineModeRequest
	GetMemberUid() *string
	SetStrictMode(v string) *ModifyVpcFirewallAclEngineModeRequest
	GetStrictMode() *string
	SetVpcFirewallId(v string) *ModifyVpcFirewallAclEngineModeRequest
	GetVpcFirewallId() *string
}

type ModifyVpcFirewallAclEngineModeRequest struct {
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	StrictMode *string `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallAclEngineModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallAclEngineModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallAclEngineModeRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *ModifyVpcFirewallAclEngineModeRequest) GetStrictMode() *string {
	return s.StrictMode
}

func (s *ModifyVpcFirewallAclEngineModeRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallAclEngineModeRequest) SetMemberUid(v string) *ModifyVpcFirewallAclEngineModeRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallAclEngineModeRequest) SetStrictMode(v string) *ModifyVpcFirewallAclEngineModeRequest {
	s.StrictMode = &v
	return s
}

func (s *ModifyVpcFirewallAclEngineModeRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallAclEngineModeRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallAclEngineModeRequest) Validate() error {
	return dara.Validate(s)
}
