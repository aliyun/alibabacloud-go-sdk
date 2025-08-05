// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcFirewallCenConfigureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteVpcFirewallCenConfigureRequest
	GetLang() *string
	SetMemberUid(v string) *DeleteVpcFirewallCenConfigureRequest
	GetMemberUid() *string
	SetVpcFirewallIdList(v []*string) *DeleteVpcFirewallCenConfigureRequest
	GetVpcFirewallIdList() []*string
}

type DeleteVpcFirewallCenConfigureRequest struct {
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

func (s DeleteVpcFirewallCenConfigureRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteVpcFirewallCenConfigureRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DeleteVpcFirewallCenConfigureRequest) GetVpcFirewallIdList() []*string {
	return s.VpcFirewallIdList
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetLang(v string) *DeleteVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetMemberUid(v string) *DeleteVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetVpcFirewallIdList(v []*string) *DeleteVpcFirewallCenConfigureRequest {
	s.VpcFirewallIdList = v
	return s
}

func (s *DeleteVpcFirewallCenConfigureRequest) Validate() error {
	return dara.Validate(s)
}
