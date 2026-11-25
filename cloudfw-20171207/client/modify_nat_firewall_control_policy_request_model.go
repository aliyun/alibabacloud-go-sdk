// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *ModifyNatFirewallControlPolicyRequest
	GetAclAction() *string
	SetAclUuid(v string) *ModifyNatFirewallControlPolicyRequest
	GetAclUuid() *string
	SetApplicationNameList(v []*string) *ModifyNatFirewallControlPolicyRequest
	GetApplicationNameList() []*string
	SetDescription(v string) *ModifyNatFirewallControlPolicyRequest
	GetDescription() *string
	SetDestPort(v string) *ModifyNatFirewallControlPolicyRequest
	GetDestPort() *string
	SetDestPortGroup(v string) *ModifyNatFirewallControlPolicyRequest
	GetDestPortGroup() *string
	SetDestPortType(v string) *ModifyNatFirewallControlPolicyRequest
	GetDestPortType() *string
	SetDestination(v string) *ModifyNatFirewallControlPolicyRequest
	GetDestination() *string
	SetDestinationType(v string) *ModifyNatFirewallControlPolicyRequest
	GetDestinationType() *string
	SetDirection(v string) *ModifyNatFirewallControlPolicyRequest
	GetDirection() *string
	SetDomainResolveType(v string) *ModifyNatFirewallControlPolicyRequest
	GetDomainResolveType() *string
	SetEndTime(v int64) *ModifyNatFirewallControlPolicyRequest
	GetEndTime() *int64
	SetLang(v string) *ModifyNatFirewallControlPolicyRequest
	GetLang() *string
	SetNatGatewayId(v string) *ModifyNatFirewallControlPolicyRequest
	GetNatGatewayId() *string
	SetProto(v string) *ModifyNatFirewallControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *ModifyNatFirewallControlPolicyRequest
	GetRelease() *string
	SetRepeatDays(v []*int64) *ModifyNatFirewallControlPolicyRequest
	GetRepeatDays() []*int64
	SetRepeatEndTime(v string) *ModifyNatFirewallControlPolicyRequest
	GetRepeatEndTime() *string
	SetRepeatStartTime(v string) *ModifyNatFirewallControlPolicyRequest
	GetRepeatStartTime() *string
	SetRepeatType(v string) *ModifyNatFirewallControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *ModifyNatFirewallControlPolicyRequest
	GetSource() *string
	SetSourceType(v string) *ModifyNatFirewallControlPolicyRequest
	GetSourceType() *string
	SetStartTime(v int64) *ModifyNatFirewallControlPolicyRequest
	GetStartTime() *int64
}

type ModifyNatFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on traffic in the access control policy. Valid values:
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
	// The UUID of the access control policy.
	//
	// To delete an access control policy, you must provide the UUID of the policy. You can call the DescribeNatFirewallControlPolicy operation to obtain the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 63ab1c02-926a-4d1b-9ef7-*****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The names of the applications.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy. Fuzzy match is supported.
	//
	// > If you do not specify this parameter, the descriptions of all policies are queried.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port for traffic in the access control policy of the virtual private cloud (VPC) firewalls.
	//
	// > Specify this parameter when **DestPortType*	- is set to `port`.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book for traffic in the access control policy.
	//
	// example:
	//
	// my_dest_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port for traffic in the access control policy. Valid values:
	//
	// - **port**: port.
	//
	// - **group**: port address book.
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// - If **DestinationType*	- is set to net, **Destination*	- is the destination CIDR block. Example: 1.2.3.4/24.
	//
	// - If **DestinationType*	- is set to group, **Destination*	- is the name of the destination address book. Example: db_group.
	//
	// - If **DestinationType*	- is set to domain, **Destination*	- is the destination domain name. Example: *.aliyuncs.com.
	//
	// - If **DestinationType*	- is set to location, **Destination*	- is the destination region. For the region codes, see the following section. Example: ["BJ11", "ZB"\].
	//
	// example:
	//
	// x.x.x.x/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: destination CIDR block.
	//
	// - **group**: destination address book.
	//
	// - **domain**: destination domain name.
	//
	// - **location**: destination region.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction of the access control policy. Valid values:
	//
	// - **out**: outbound traffic.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// - **0**: FQDN-based resolution.
	//
	// - **1**: DNS-based dynamic resolution.
	//
	// - **2**: FQDN and DNS-based dynamic resolution.
	//
	// > If the domain name identification mode includes FQDN, only the TCP protocol is allowed. You can select HTTP, HTTPS, SMTP, SMTPS, SSL, IMAPS, or POPS as the application.
	//
	// example:
	//
	// 0
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy validity period for the access control policy. Specify the time in the UNIX timestamp format (seconds). The time must be on the hour or half-hour, and must be at least 30 minutes later than the start time.
	//
	// > If RepeatType is set to Permanent, EndTime is empty. If RepeatType is set to None, Daily, Weekly, or Monthly, EndTime must be specified.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-xxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The type of the security protocol for traffic in the access control policy. Supported protocol types:
	//
	// - **ANY**
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// > **ANY*	- indicates that the policy applies to all types of applications.
	//
	// > If the destination address is a threat intelligence address book or a cloud service address book of the domain name type, you can set this parameter to TCP. If you select TCP, you can select HTTP, HTTPS, SMTP, SMTPS, or SSL as the application.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// - **true**: The policy is enabled.
	//
	// - **false**: The policy is disabled.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The set of recurrence days for the policy validity period of the access control policy.
	//
	// - If RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is an empty set. Example: [].
	//
	// - If RepeatType is set to Weekly, RepeatDays cannot be empty. Example: [0, 6].
	//
	// > If RepeatType is set to Weekly, the values in RepeatDays must be unique.
	//
	// - If RepeatType is set to `Monthly`, RepeatDays cannot be empty. Example: [1, 31].
	//
	// > If RepeatType is set to Monthly, the values in RepeatDays must be unique.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurrence end time when the policy takes effect. Use the 24-hour time format HH:mm, such as 23:30. The time must be on the hour or half-hour, and must be at least 30 minutes later than the recurrence start time.
	//
	// > If RepeatType is set to Permanent or None, RepeatEndTime is empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime must be specified.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time when the policy takes effect. Use the 24-hour time format HH:mm, such as 08:00. The time must be on the hour or half-hour, and must be at least 30 minutes earlier than the recurrence end time.
	//
	// > If RepeatType is set to Permanent or None, RepeatStartTime is empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatStartTime must be specified.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type of the policy validity period for the access control policy. Valid values:
	//
	// - **Permanent*	- (default): Always.
	//
	// - **None**: Specified single time.
	//
	// - **Daily**: Every day.
	//
	// - **Weekly**: Every week.
	//
	// - **Monthly**: Every month.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// - If **SourceType*	- is set to `net`, Source is the source CIDR block. Example: 10.2.XX.XX/24.
	//
	// - If **SourceType*	- is set to `group`, Source is the name of the source address book. Example: db_group.
	//
	// example:
	//
	// 10.2.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: source CIDR block.
	//
	// - **group**: source address book.
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the policy validity period for the access control policy. Specify the time in the UNIX timestamp format (seconds). The time must be on the hour or half-hour, and must be at least 30 minutes earlier than the end time.
	//
	// > If RepeatType is set to Permanent, StartTime is empty. If RepeatType is set to None, Daily, Weekly, or Monthly, StartTime must be specified.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyNatFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *ModifyNatFirewallControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ModifyNatFirewallControlPolicyRequest) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *ModifyNatFirewallControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyNatFirewallControlPolicyRequest) GetDestPort() *string {
	return s.DestPort
}

func (s *ModifyNatFirewallControlPolicyRequest) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *ModifyNatFirewallControlPolicyRequest) GetDestPortType() *string {
	return s.DestPortType
}

func (s *ModifyNatFirewallControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *ModifyNatFirewallControlPolicyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ModifyNatFirewallControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyNatFirewallControlPolicyRequest) GetDomainResolveType() *string {
	return s.DomainResolveType
}

func (s *ModifyNatFirewallControlPolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModifyNatFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyNatFirewallControlPolicyRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatFirewallControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *ModifyNatFirewallControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *ModifyNatFirewallControlPolicyRequest) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *ModifyNatFirewallControlPolicyRequest) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *ModifyNatFirewallControlPolicyRequest) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *ModifyNatFirewallControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *ModifyNatFirewallControlPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *ModifyNatFirewallControlPolicyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyNatFirewallControlPolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModifyNatFirewallControlPolicyRequest) SetAclAction(v string) *ModifyNatFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetAclUuid(v string) *ModifyNatFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetApplicationNameList(v []*string) *ModifyNatFirewallControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDescription(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestPort(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestPortGroup(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestPortType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestination(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestinationType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDirection(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDomainResolveType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DomainResolveType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetEndTime(v int64) *ModifyNatFirewallControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetLang(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetNatGatewayId(v string) *ModifyNatFirewallControlPolicyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetProto(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRelease(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRepeatDays(v []*int64) *ModifyNatFirewallControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRepeatEndTime(v string) *ModifyNatFirewallControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRepeatStartTime(v string) *ModifyNatFirewallControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRepeatType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetSource(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetSourceType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetStartTime(v int64) *ModifyNatFirewallControlPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
