// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcFirewallConfigureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteVpcFirewallConfigureRequest
	GetLang() *string
	SetMemberUid(v string) *DeleteVpcFirewallConfigureRequest
	GetMemberUid() *string
	SetVpcFirewallIdList(v []*string) *DeleteVpcFirewallConfigureRequest
	GetVpcFirewallIdList() []*string
}

type DeleteVpcFirewallConfigureRequest struct {
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
	// The instance IDs of VPC firewalls.
	//
	// This parameter is required.
	VpcFirewallIdList []*string `json:"VpcFirewallIdList,omitempty" xml:"VpcFirewallIdList,omitempty" type:"Repeated"`
}

func (s DeleteVpcFirewallConfigureRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteVpcFirewallConfigureRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DeleteVpcFirewallConfigureRequest) GetVpcFirewallIdList() []*string {
	return s.VpcFirewallIdList
}

func (s *DeleteVpcFirewallConfigureRequest) SetLang(v string) *DeleteVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallConfigureRequest) SetMemberUid(v string) *DeleteVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *DeleteVpcFirewallConfigureRequest) SetVpcFirewallIdList(v []*string) *DeleteVpcFirewallConfigureRequest {
	s.VpcFirewallIdList = v
	return s
}

func (s *DeleteVpcFirewallConfigureRequest) Validate() error {
	return dara.Validate(s)
}
