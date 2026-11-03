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
	// The action that Cloud Firewall performs on traffic in the access control policy. Valid values:
	//
	// - **accept**: allows the traffic
	//
	// - **drop**: denies the traffic
	//
	// - **log**: monitors the traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The list of application types supported by the access control policy.
	//
	// This parameter is required.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Allow traffic
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy. Valid values:
	//
	// - If the protocol is ICMP, leave this parameter empty.
	//
	//
	//
	// > Access control based on destination ports is not supported when the protocol is ICMP.
	//
	// - If the protocol is TCP, UDP, or ANY and the destination Protocol Type (DestPortType) is group, leave this parameter empty.
	//
	// > If you set DestPortType to group, you do not need to specify a destination port. All ports managed by this policy are included in the destination port address book.
	//
	// - If the protocol is TCP, UDP, or ANY and the destination Protocol Type (DestPortType) is port, set this parameter to the destination port number.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// > If you set DestPortType to group, you must specify the name of the destination port address book.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// - **port**: a port
	//
	// - **group**: a port address book
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. Valid values:
	//
	// - If DestinationType is set to net, set Destination to the destination CIDR block.
	//
	//
	//
	//   Example: 1.2.XX.XX/24
	//
	// - If DestinationType is set to group, set Destination to the name of the destination address book.
	//
	//   Example: db_group
	//
	// - If DestinationType is set to domain, set Destination to the destination domain name.
	//
	//   Example: *.aliyuncs.com
	//
	// - If DestinationType is set to location, set Destination to the destination region.
	//
	//   Example: ["BJ11", "ZB"\\]
	//
	// This parameter is required.
	//
	// example:
	//
	// XX.XX.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: a destination CIDR block
	//
	// - **group**: a destination address book
	//
	// - **domain**: a destination domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction of the access control policy. Valid values:
	//
	// - **out**: outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name parse method of the access control policy. Valid values:
	//
	// - **0**: FQDN-based parse
	//
	// - **1**: DNS-based dynamic parse
	//
	// - **2**: FQDN-based and DNS-based dynamic parse
	//
	// > If the domain name identification mode includes FQDN, you can only configure the TCP protocol. The supported applications are HTTP, HTTPS, SMTP, SMTPS, SSL, IMAPS, and POPS.
	//
	// example:
	//
	// 0
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time for the Policy Validity Period of the access control policy. The value is in the format of a UNIX timestamp in seconds. The time must be on the hour or half hour, and must be at least 30 minutes later than the start time.
	//
	// > If RepeatType is set to Permanent, leave EndTime empty. If RepeatType is set to None, Daily, Weekly, or Monthly, EndTime must have a value, and you must specify the Settings for the end time.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address version. Valid values:
	//
	// - **4**: IPv4 (default)
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh**: Chinese (default)
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The instance ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-2vc2ustolqn6sr0******
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy. The priority value starts from 1 and increases sequentially. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol type of the traffic in the access control policy. Valid values:
	//
	// - ANY: all protocols
	//
	// - TCP
	//
	// - UDP
	//
	// - ICMP
	//
	// > If the destination address is a threat intelligence address book or a cloud service address book of the domain name type, you can only configure the TCP protocol. The supported applications are HTTP, HTTPS, SMTP, SMTPS, and SSL.
	//
	// This parameter is required.
	//
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, the policy is enabled after it is created. Valid values:
	//
	// - **true**: enables the access control policy
	//
	// - **false**: does not enable the access control policy
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The set of recurrence days for the Policy Validity Period of the access control policy.
	//
	// - If RepeatType is set to `Permanent`, `None`, or `Daily`, leave RepeatDays empty.
	//
	//   Example: []
	//
	// - If RepeatType is set to Weekly, RepeatDays cannot be empty.
	//
	//   Example: [0, 6]
	//
	// > When the RepeatType Settings are Weekly, the values in RepeatDays cannot be duplicate.
	//
	// - If RepeatType is set to `Monthly`, RepeatDays cannot be empty.
	//
	//   Example: [1, 31]
	//
	// > When the RepeatType Settings are Monthly, the values in RepeatDays cannot be duplicate.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurrence end time for the Policy Validity Period of the access control policy. Example: 23:30. The time must be on the hour or half hour, and must be at least 30 minutes later than the recurrence start time.
	//
	// > If RepeatType is set to Permanent or None, leave RepeatEndTime empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime must have a value, and you must specify the Settings for the recurrence end time.
	//
	// > The time format is HH:mm (24-hour clock), such as 08:00 and 23:30.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time for the Policy Validity Period of the access control policy. Example: 08:00. The time must be on the hour or half hour, and must be at least 30 minutes earlier than the recurrence end time.
	//
	// > If RepeatType is set to Permanent or None, leave RepeatStartTime empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatStartTime must have a value, and you must specify the Settings for the recurrence start time.
	//
	// > The time format is HH:mm (24-hour clock), such as 08:00 and 23:30.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type of the Policy Validity Period for the access control policy. Valid values:
	//
	// - **Permanent**: always (default)
	//
	// - **None**: a specific single time
	//
	// - **Daily**: every day
	//
	// - **Weekly**: every week
	//
	// - **Monthly**: every month
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// - If **SourceType*	- is set to `net`, set Source to the source CIDR block.
	//
	//   Example: 10.2.4.0/24
	//
	// - If **SourceType*	- is set to `group`, set Source to the name of the source address book.
	//
	//   Example: db_group
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.25/32
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: a source CIDR block
	//
	// - **group**: a source address book
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time for the Policy Validity Period of the access control policy. The value is in the format of a UNIX timestamp in seconds. The time must be on the hour or half hour, and must be at least 30 minutes earlier than the end time.
	//
	// > If RepeatType is set to Permanent, leave StartTime empty. If RepeatType is set to None, Daily, Weekly, or Monthly, StartTime must have a value, and you must specify the Settings for the start time.
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
