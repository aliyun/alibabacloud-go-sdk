// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsFirewallPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *DeleteDnsFirewallPolicyRequest
	GetAclUuid() *string
	SetLang(v string) *DeleteDnsFirewallPolicyRequest
	GetLang() *string
	SetSourceIp(v string) *DeleteDnsFirewallPolicyRequest
	GetSourceIp() *string
}

type DeleteDnsFirewallPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 498946f4-c98a-45c0-9038-635c******
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 140.210.39.***
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteDnsFirewallPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsFirewallPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteDnsFirewallPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DeleteDnsFirewallPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDnsFirewallPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteDnsFirewallPolicyRequest) SetAclUuid(v string) *DeleteDnsFirewallPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteDnsFirewallPolicyRequest) SetLang(v string) *DeleteDnsFirewallPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDnsFirewallPolicyRequest) SetSourceIp(v string) *DeleteDnsFirewallPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteDnsFirewallPolicyRequest) Validate() error {
	return dara.Validate(s)
}
