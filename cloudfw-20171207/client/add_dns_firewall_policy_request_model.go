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
	// Specifies the action to take on traffic that matches the access control policy. Valid values:
	//
	// - **accept**: Allows the traffic.
	//
	// - **drop**: Denies the traffic.
	//
	// - **log**: Monitors the traffic.
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
	// - When **DestinationType*	- is `net`, this parameter specifies the destination CIDR block. Example: `1.2.3.4/24`.
	//
	// - When **DestinationType*	- is `group`, this parameter specifies the name of the destination address book. Example: `db_group`.
	//
	// - When **DestinationType*	- is `domain`, this parameter specifies the destination domain name. Example: `*.aliyuncs.com`.
	//
	// - When **DestinationType*	- is `location`, this parameter specifies the destination region. For more information about location codes, see the documentation. Example: `["BJ11", "ZB"]`.
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
	// - **net**: destination CIDR block
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
	// The traffic direction for the access control policy. Valid values:
	//
	// - **in**: inbound traffic
	//
	// - **out**: outbound traffic
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version supported by the policy.
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
	// The language of the request and response. Valid values:<br>-**zh**: Chinese<br>-**en**: English<br><br>
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The priority of the access control policy. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Specifies whether to enable the access control policy. Valid values:
	//
	// - **true**: Enables the access control policy.
	//
	// - **false**: Disables the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy.
	//
	// - When **SourceType*	- is `net`, this parameter specifies the source CIDR block. Example: `10.2.XX.XX/24`.
	//
	// - When **SourceType*	- is `group`, this parameter specifies the name of the source address book. Example: `db_group`.
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
