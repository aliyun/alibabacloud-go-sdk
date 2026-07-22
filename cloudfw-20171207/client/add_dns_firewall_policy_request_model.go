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
	// The method that is used by the access control policy to control traffic that passes through Cloud Firewall. Valid values:
	//
	// - **accept**: allows the traffic.
	//
	// - **drop**: deny the traffic.
	//
	// - **log**: monitors the traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The description of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy.
	//
	// - If **DestinationType*	- is set to net, **Destination*	- is a destination CIDR block. Example: 1.2.3.4/24.
	//
	// - If **DestinationType*	- is set to group, **Destination*	- is the name of a destination address book. Example: db_group.
	//
	// - If **DestinationType*	- is set to domain, **Destination*	- is a destination domain name. Example: *.aliyuncs.com.
	//
	// - If **DestinationType*	- is set to location, **Destination*	- is a destination area (for specific area positional encoding, see the following sections). Example: ["BJ11", "ZB"\\].
	//
	// This parameter is required.
	//
	// example:
	//
	// db_group
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy.
	//
	// Valid values:
	//
	// - **group**: destination address book
	//
	// - **domain**: destination domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// domain
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the DNS firewall policy. The backend fixes this value to out (internal-to-external). Set Direction to out.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP address version supported.
	//
	// Valid values:
	//
	// - **4**: IPv4
	//
	// - **6**: IPv6
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
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
	// The priority of the policy. A smaller value indicates a higher priority. Valid values: 1 to 20000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Specifies whether to enable the access control policy. The policy is enabled by default after it is created. Valid values:
	//
	// - **true**: enables the access control policy.
	//
	// - **false**: does not enable the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// - If **SourceType*	- is set to `net`, Source is a source CIDR block. Example: 10.2.XX.XX/24.
	//
	// - If **SourceType*	- is set to `group`, Source is the name of a source address book. Example: db_group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.223/32
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 140.205.118.97
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: source CIDR block
	//
	// - **group**: source address book
	//
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
