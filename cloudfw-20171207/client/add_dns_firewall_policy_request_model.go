// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsFirewallPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *AddDnsFirewallPolicyRequest
	GetAclAction() *string
	SetDescription(v string) *AddDnsFirewallPolicyRequest
	GetDescription() *string
	SetDestination(v string) *AddDnsFirewallPolicyRequest
	GetDestination() *string
	SetDestinationType(v string) *AddDnsFirewallPolicyRequest
	GetDestinationType() *string
	SetDirection(v string) *AddDnsFirewallPolicyRequest
	GetDirection() *string
	SetIpVersion(v string) *AddDnsFirewallPolicyRequest
	GetIpVersion() *string
	SetLang(v string) *AddDnsFirewallPolicyRequest
	GetLang() *string
	SetPriority(v string) *AddDnsFirewallPolicyRequest
	GetPriority() *string
	SetRelease(v string) *AddDnsFirewallPolicyRequest
	GetRelease() *string
	SetSource(v string) *AddDnsFirewallPolicyRequest
	GetSource() *string
	SetSourceIp(v string) *AddDnsFirewallPolicyRequest
	GetSourceIp() *string
	SetSourceType(v string) *AddDnsFirewallPolicyRequest
	GetSourceType() *string
}

type AddDnsFirewallPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// db_group
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// domain
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.223/32
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 140.205.118.97
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s AddDnsFirewallPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDnsFirewallPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddDnsFirewallPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *AddDnsFirewallPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *AddDnsFirewallPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *AddDnsFirewallPolicyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *AddDnsFirewallPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *AddDnsFirewallPolicyRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *AddDnsFirewallPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDnsFirewallPolicyRequest) GetPriority() *string {
	return s.Priority
}

func (s *AddDnsFirewallPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *AddDnsFirewallPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *AddDnsFirewallPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *AddDnsFirewallPolicyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *AddDnsFirewallPolicyRequest) SetAclAction(v string) *AddDnsFirewallPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetDescription(v string) *AddDnsFirewallPolicyRequest {
	s.Description = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetDestination(v string) *AddDnsFirewallPolicyRequest {
	s.Destination = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetDestinationType(v string) *AddDnsFirewallPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetDirection(v string) *AddDnsFirewallPolicyRequest {
	s.Direction = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetIpVersion(v string) *AddDnsFirewallPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetLang(v string) *AddDnsFirewallPolicyRequest {
	s.Lang = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetPriority(v string) *AddDnsFirewallPolicyRequest {
	s.Priority = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetRelease(v string) *AddDnsFirewallPolicyRequest {
	s.Release = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetSource(v string) *AddDnsFirewallPolicyRequest {
	s.Source = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetSourceIp(v string) *AddDnsFirewallPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) SetSourceType(v string) *AddDnsFirewallPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *AddDnsFirewallPolicyRequest) Validate() error {
	return dara.Validate(s)
}
