// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallIPSWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyVpcFirewallIPSWhitelistRequest
	GetLang() *string
	SetListType(v int64) *ModifyVpcFirewallIPSWhitelistRequest
	GetListType() *int64
	SetListValue(v string) *ModifyVpcFirewallIPSWhitelistRequest
	GetListValue() *string
	SetMemberUid(v int64) *ModifyVpcFirewallIPSWhitelistRequest
	GetMemberUid() *int64
	SetVpcFirewallId(v string) *ModifyVpcFirewallIPSWhitelistRequest
	GetVpcFirewallId() *string
	SetWhiteType(v int64) *ModifyVpcFirewallIPSWhitelistRequest
	GetWhiteType() *int64
}

type ModifyVpcFirewallIPSWhitelistRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the list. Valid values:
	//
	// 	- **1**: user-defined
	//
	// 	- **2**: address book
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ListType *int64 `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The entry in the list.
	//
	// example:
	//
	// 10.130.0.0/20,10.130.17.11/32
	ListValue *string `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 1415189284827022
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-d7b8ce273791475b9b0b
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The type of the whitelist. Valid values:
	//
	// 	- **1**: destination
	//
	// 	- **2**: source
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WhiteType *int64 `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s ModifyVpcFirewallIPSWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallIPSWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) GetListType() *int64 {
	return s.ListType
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) GetListValue() *string {
	return s.ListValue
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) GetWhiteType() *int64 {
	return s.WhiteType
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetLang(v string) *ModifyVpcFirewallIPSWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetListType(v int64) *ModifyVpcFirewallIPSWhitelistRequest {
	s.ListType = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetListValue(v string) *ModifyVpcFirewallIPSWhitelistRequest {
	s.ListValue = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetMemberUid(v int64) *ModifyVpcFirewallIPSWhitelistRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallIPSWhitelistRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetWhiteType(v int64) *ModifyVpcFirewallIPSWhitelistRequest {
	s.WhiteType = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
