// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *CreateNatFirewallControlPolicyRequest
	GetAclAction() *string
	SetApplicationNameList(v []*string) *CreateNatFirewallControlPolicyRequest
	GetApplicationNameList() []*string
	SetDescription(v string) *CreateNatFirewallControlPolicyRequest
	GetDescription() *string
	SetDestPort(v string) *CreateNatFirewallControlPolicyRequest
	GetDestPort() *string
	SetDestPortGroup(v string) *CreateNatFirewallControlPolicyRequest
	GetDestPortGroup() *string
	SetDestPortType(v string) *CreateNatFirewallControlPolicyRequest
	GetDestPortType() *string
	SetDestination(v string) *CreateNatFirewallControlPolicyRequest
	GetDestination() *string
	SetDestinationType(v string) *CreateNatFirewallControlPolicyRequest
	GetDestinationType() *string
	SetDirection(v string) *CreateNatFirewallControlPolicyRequest
	GetDirection() *string
	SetDomainResolveType(v int32) *CreateNatFirewallControlPolicyRequest
	GetDomainResolveType() *int32
	SetEndTime(v int64) *CreateNatFirewallControlPolicyRequest
	GetEndTime() *int64
	SetIpVersion(v string) *CreateNatFirewallControlPolicyRequest
	GetIpVersion() *string
	SetLang(v string) *CreateNatFirewallControlPolicyRequest
	GetLang() *string
	SetNatGatewayId(v string) *CreateNatFirewallControlPolicyRequest
	GetNatGatewayId() *string
	SetNewOrder(v string) *CreateNatFirewallControlPolicyRequest
	GetNewOrder() *string
	SetProto(v string) *CreateNatFirewallControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *CreateNatFirewallControlPolicyRequest
	GetRelease() *string
	SetRepeatDays(v []*int64) *CreateNatFirewallControlPolicyRequest
	GetRepeatDays() []*int64
	SetRepeatEndTime(v string) *CreateNatFirewallControlPolicyRequest
	GetRepeatEndTime() *string
	SetRepeatStartTime(v string) *CreateNatFirewallControlPolicyRequest
	GetRepeatStartTime() *string
	SetRepeatType(v string) *CreateNatFirewallControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *CreateNatFirewallControlPolicyRequest
	GetSource() *string
	SetSourceType(v string) *CreateNatFirewallControlPolicyRequest
	GetSourceType() *string
	SetStartTime(v int64) *CreateNatFirewallControlPolicyRequest
	GetStartTime() *int64
}

type CreateNatFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on traffic that matches the access control policy.
	//
	// Valid values:
	//
	// - **accept**: Allows the traffic.
	//
	// - **drop**: Drops the traffic.
	//
	// - **log**: Logs the traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The list of applications to which the access control policy applies.
	//
	// This parameter is required.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 放行流量
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port for the traffic. The value of this parameter depends on the `Proto` and `DestPortType` parameters.
	//
	// - If `Proto` is `ICMP`, leave this parameter empty.
	//
	// > Access control cannot be configured based on the destination port for ICMP traffic.
	//
	// - If the destination port type (`DestPortType`) is `group`, leave this parameter empty.
	//
	// > If `DestPortType` is set to `group`, you do not need to specify a destination port because the required ports are defined in the selected port address book.
	//
	// - If the protocol is TCP, UDP, or ANY and the destination port type (`DestPortType`) is `port`, specify the destination port number.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book.
	//
	// > This parameter is required if `DestPortType` is set to `group`.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port.
	//
	// - **port**: Port or port range
	//
	// - **group**: Port address book
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// The value of this parameter varies based on the value of `DestinationType`:
	//
	// - If `DestinationType` is `net`, set this parameter to the destination CIDR.
	//
	//   Example: `1.2.XX.XX/24`
	//
	// - If `DestinationType` is `group`, set this parameter to the name of the destination address book.
	//
	//   Example: `db_group`
	//
	// - If `DestinationType` is `domain`, set this parameter to the destination domain.
	//
	//   Example: \\*.aliyuncs.com
	//
	// - If `DestinationType` is `location`, set this parameter to the destination location.
	//
	//   Example: ["BJ11", "ZB"]
	//
	// This parameter is required.
	//
	// example:
	//
	// XX.XX.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy.
	//
	// Valid values:
	//
	// - **net**: Destination CIDR
	//
	// - **group**: Destination address book
	//
	// - **domain**: Destination domain
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction for the access control policy. Valid values:
	//
	// - **out**: outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method. Valid values:
	//
	// - **0**: FQDN-based resolution
	//
	// - **1**: Dynamic DNS-based resolution
	//
	// - **2**: FQDN-based and dynamic DNS-based resolution
	//
	// > If the resolution method includes FQDN, you can set the protocol only to TCP. The supported applications are HTTP, HTTPS, SMTP, SMTPS, and SSL.
	//
	// example:
	//
	// 0
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy\\"s validity period, specified as a Unix timestamp. The time must be on the hour or half-hour and be at least 30 minutes after the start time.
	//
	// > This parameter is required if `RepeatType` is `None`, `Daily`, `Weekly`, or `Monthly`. If `RepeatType` is `Permanent`, leave this parameter empty.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP version for the access control policy. Valid values:
	//
	// - **4*	- (default): IPv4
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the response messages.
	//
	// Valid values:
	//
	// - **zh**: Chinese (default)
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The instance ID of the NAT Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-2vc2ustolqn6sr0******
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy. Values start from 1. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol for the traffic in the access control policy.
	//
	// Valid values:
	//
	// - ANY: all protocols
	//
	// - TCP
	//
	// - UDP
	//
	// - ICMP
	//
	// > If the destination is a domain, a threat intelligence address book, or a cloud service address book, you can only set this parameter to `TCP`. The supported application types are HTTP, HTTPS, SMTP, SMTPS, and SSL.
	//
	// This parameter is required.
	//
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether the access control policy is enabled. By default, policies are enabled upon creation. Valid values:
	//
	// - **true**: Enables the policy.
	//
	// - **false**: Disables the policy.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of the week or month on which the policy recurs.
	//
	// - If `RepeatType` is set to `Permanent`, `None`, or `Daily`, leave this parameter empty. Example: `[]`
	//
	// - If `RepeatType` is `Weekly`, this parameter is required. Example: `[0, 6]`
	//
	// > If `RepeatType` is `Weekly`, the values in `RepeatDays` must be unique.
	//
	// - If `RepeatType` is `Monthly`, this parameter is required. Example: `[1, 31]`
	//
	// > If `RepeatType` is `Monthly`, the values in `RepeatDays` must be unique.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The end time of the recurrence. The time must be on the hour or half-hour, and must be at least 30 minutes later than the start time.
	//
	// > This parameter is required if `RepeatType` is set to `Daily`, `Weekly`, or `Monthly`. If `RepeatType` is `Permanent` or `None`, leave this parameter empty.
	//
	// > The time is in the HH:mm format (24-hour). For example, `08:00` or `23:30`.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The start time of the recurrence. The time must be on the hour or half-hour, and must be at least 30 minutes earlier than the end time.
	//
	// > This parameter is required if `RepeatType` is set to `Daily`, `Weekly`, or `Monthly`. If `RepeatType` is `Permanent` or `None`, leave this parameter empty.
	//
	// > The time is in the HH:mm format (24-hour). For example, `08:00` or `23:30`.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the policy validity period. Valid values:
	//
	// - **Permanent*	- (default): The policy is always active.
	//
	// - **None**: The policy runs once for a specified duration.
	//
	// - **Daily**: The policy recurs daily.
	//
	// - **Weekly**: The policy recurs weekly within a specified time range.
	//
	// - **Monthly**: The policy recurs monthly within a specified time range.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy.
	//
	// The value of this parameter varies based on the value of `SourceType`:
	//
	// - If **SourceType*	- is `net`, set this parameter to the source CIDR.
	//
	//   Example: 10.2.4.0/24
	//
	// - If **SourceType*	- is `group`, set this parameter to the name of the source address book.
	//
	//   Example: db_group
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.25/32
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy.
	//
	// Valid values:
	//
	// - **net**: Source CIDR
	//
	// - **group**: Source address book
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the policy\\"s validity period, specified as a Unix timestamp. The time must be on the hour or half-hour and be at least 30 minutes before the end time.
	//
	// > This parameter is required if `RepeatType` is `None`, `Daily`, `Weekly`, or `Monthly`. If `RepeatType` is `Permanent`, leave this parameter empty.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateNatFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *CreateNatFirewallControlPolicyRequest) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *CreateNatFirewallControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNatFirewallControlPolicyRequest) GetDestPort() *string {
	return s.DestPort
}

func (s *CreateNatFirewallControlPolicyRequest) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *CreateNatFirewallControlPolicyRequest) GetDestPortType() *string {
	return s.DestPortType
}

func (s *CreateNatFirewallControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *CreateNatFirewallControlPolicyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *CreateNatFirewallControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *CreateNatFirewallControlPolicyRequest) GetDomainResolveType() *int32 {
	return s.DomainResolveType
}

func (s *CreateNatFirewallControlPolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateNatFirewallControlPolicyRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateNatFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateNatFirewallControlPolicyRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateNatFirewallControlPolicyRequest) GetNewOrder() *string {
	return s.NewOrder
}

func (s *CreateNatFirewallControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *CreateNatFirewallControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *CreateNatFirewallControlPolicyRequest) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *CreateNatFirewallControlPolicyRequest) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *CreateNatFirewallControlPolicyRequest) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *CreateNatFirewallControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *CreateNatFirewallControlPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *CreateNatFirewallControlPolicyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateNatFirewallControlPolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateNatFirewallControlPolicyRequest) SetAclAction(v string) *CreateNatFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetApplicationNameList(v []*string) *CreateNatFirewallControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDescription(v string) *CreateNatFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestPort(v string) *CreateNatFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestPortGroup(v string) *CreateNatFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestPortType(v string) *CreateNatFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestination(v string) *CreateNatFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestinationType(v string) *CreateNatFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDirection(v string) *CreateNatFirewallControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDomainResolveType(v int32) *CreateNatFirewallControlPolicyRequest {
	s.DomainResolveType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetEndTime(v int64) *CreateNatFirewallControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetIpVersion(v string) *CreateNatFirewallControlPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetLang(v string) *CreateNatFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetNatGatewayId(v string) *CreateNatFirewallControlPolicyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetNewOrder(v string) *CreateNatFirewallControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetProto(v string) *CreateNatFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRelease(v string) *CreateNatFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRepeatDays(v []*int64) *CreateNatFirewallControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRepeatEndTime(v string) *CreateNatFirewallControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRepeatStartTime(v string) *CreateNatFirewallControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRepeatType(v string) *CreateNatFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetSource(v string) *CreateNatFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetSourceType(v string) *CreateNatFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetStartTime(v int64) *CreateNatFirewallControlPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
