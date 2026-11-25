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
	SetSourceType(v string) *ModifyDnsFirewallPolicyRequest
	GetSourceType() *string
}

type ModifyDnsFirewallPolicyRequest struct {
	// The action that Cloud Firewall performs on traffic that matches the access control policy. Valid values:
	//
	// - **accept**: Allows the traffic.
	//
	// - **drop**: Denies the traffic.
	//
	// - **log**: Monitors the traffic.
	//
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df22****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy.
	//
	// - If **DestinationType*	- is set to net, **Destination*	- specifies the destination CIDR block. Example: 1.2.3.4/24.
	//
	// - If **DestinationType*	- is set to group, **Destination*	- specifies the name of the destination address book. Example: db_group.
	//
	// - If **DestinationType*	- is set to domain, **Destination*	- specifies the destination domain name. Example: *.aliyuncs.com.
	//
	// - If **DestinationType*	- is set to location, **Destination*	- specifies the destination area (for specific area positional encoding, see below). Example: ["BJ11", "ZB"\].
	//
	// example:
	//
	// db_group
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: a destination CIDR block.
	//
	// - **group**: a destination address book.
	//
	// - **domain**: a destination domain name.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Policy priority of the access control policy before policy priority is modified.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Specifies whether the access control policy is enabled. By default, the policy is enabled after it is created. Valid values:
	//
	// - **true**: enables the access control policy.
	//
	// - **false**: disables the access control policy.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	Source  *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: a source CIDR block.
	//
	// - **group**: a source address book.
	//
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

func (s *ModifyDnsFirewallPolicyRequest) SetSourceType(v string) *ModifyDnsFirewallPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyDnsFirewallPolicyRequest) Validate() error {
	return dara.Validate(s)
}
