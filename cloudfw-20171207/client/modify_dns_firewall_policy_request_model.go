// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDnsFirewallPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *ModifyDnsFirewallPolicyRequest
	GetAclAction() *string
	SetAclUuid(v string) *ModifyDnsFirewallPolicyRequest
	GetAclUuid() *string
	SetDescription(v string) *ModifyDnsFirewallPolicyRequest
	GetDescription() *string
	SetDestination(v string) *ModifyDnsFirewallPolicyRequest
	GetDestination() *string
	SetDestinationType(v string) *ModifyDnsFirewallPolicyRequest
	GetDestinationType() *string
	SetLang(v string) *ModifyDnsFirewallPolicyRequest
	GetLang() *string
	SetPriority(v string) *ModifyDnsFirewallPolicyRequest
	GetPriority() *string
	SetRelease(v string) *ModifyDnsFirewallPolicyRequest
	GetRelease() *string
	SetSource(v string) *ModifyDnsFirewallPolicyRequest
	GetSource() *string
	SetSourceIp(v string) *ModifyDnsFirewallPolicyRequest
	GetSourceIp() *string
	SetSourceType(v string) *ModifyDnsFirewallPolicyRequest
	GetSourceType() *string
}

type ModifyDnsFirewallPolicyRequest struct {
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df22****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// db_group
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	Source  *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ModifyDnsFirewallPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDnsFirewallPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDnsFirewallPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *ModifyDnsFirewallPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ModifyDnsFirewallPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDnsFirewallPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *ModifyDnsFirewallPolicyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ModifyDnsFirewallPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyDnsFirewallPolicyRequest) GetPriority() *string {
	return s.Priority
}

func (s *ModifyDnsFirewallPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *ModifyDnsFirewallPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *ModifyDnsFirewallPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyDnsFirewallPolicyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyDnsFirewallPolicyRequest) SetAclAction(v string) *ModifyDnsFirewallPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetAclUuid(v string) *ModifyDnsFirewallPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetDescription(v string) *ModifyDnsFirewallPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetDestination(v string) *ModifyDnsFirewallPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetDestinationType(v string) *ModifyDnsFirewallPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetLang(v string) *ModifyDnsFirewallPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetPriority(v string) *ModifyDnsFirewallPolicyRequest {
	s.Priority = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetRelease(v string) *ModifyDnsFirewallPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetSource(v string) *ModifyDnsFirewallPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetSourceIp(v string) *ModifyDnsFirewallPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) SetSourceType(v string) *ModifyDnsFirewallPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) Validate() error {
	return dara.Validate(s)
}
