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
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// 	- **accept**: allows the traffic.
	//
	// 	- **drop**: denies the traffic.
	//
	// 	- **log**: monitors the traffic.
	//
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	//
	// To modify the configurations of an access control policy, you must provide the UUID of the policy. You can call the DescribeNatFirewallControlPolicy operation to query the UUIDs of access control policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61ab1c02-926a-4d1b-9ef7-595eed8c4226
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The name of the application.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy. Fuzzy match is supported.
	//
	// > If you do not specify this parameter, the descriptions of all policies are queried.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	//
	// > If **DestPortType*	- is set to `port`, you must specify this parameter.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// example:
	//
	// my_dest_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// 	- **port**: port
	//
	// 	- **group**: port address book
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// 	- If **DestinationType*	- is set to net, the value of **Destination*	- is a CIDR block. Example: 1.2.3.4/24
	//
	// 	- If **DestinationType*	- is set to group, the value of **Destination*	- is an address book. Example: db_group
	//
	// 	- If **DestinationType*	- is set to domain, the value of **Destination*	- is a domain name. Example: \\*.aliyuncs.com
	//
	// 	- If **DestinationType*	- is set to location, the value of **Destination*	- is a location. For more information about the location codes, see the "AddControlPolicy" topic. Example: ["BJ11", "ZB"]
	//
	// example:
	//
	// x.x.x.x/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// 	- **net**: CIDR block
	//
	// 	- **group**: address book
	//
	// 	- **domain**: domain name
	//
	// 	- **location**: destination location
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid value:
	//
	// 	- **out**: outbound.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// 	- **0**: Fully qualified domain name (FQDN)-based resolution
	//
	// 	- **1**: Domain Name System (DNS)-based dynamic resolution
	//
	// 	- **2**: FQDN and DNS-based dynamic resolution
	//
	// example:
	//
	// 0
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of StartTime.
	//
	// >  If RepeatType is set to Permanent, EndTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, EndTime must be specified.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
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
	// The protocol type in the access control policy. Valid values:
	//
	// 	- **ANY**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **ICMP**
	//
	// >  The value **ANY*	- indicates all types of applications.
	//
	// >  If the destination address type is a threat intelligence address book of the domain name type or a cloud service address book, you can set Proto to TCP. If you set Proto to TCP, you can set application types to HTTP, HTTPS, SMTP, SMTPS, and SSL.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// 	- If RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is left empty. Example: [].
	//
	// 	- If RepeatType is set to Weekly, RepeatDays must be specified. Example: [0, 6].
	//
	// >  If RepeatType is set to Weekly, the fields in the value of RepeatDays cannot be repeated.
	//
	// 	- If RepeatType is set to `Monthly`, RepeatDays must be specified. Example: [1, 31].
	//
	// >  If RepeatType is set to Monthly, the fields in the value of RepeatDays cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of RepeatStartTime.
	//
	// >  If RepeatType is set to Permanent or None, RepeatEndTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime must be specified.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of RepeatEndTime.
	//
	// >  If RepeatType is set to Permanent or None, RepeatStartTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// 	- **Permanent*	- (default): The policy always takes effect.
	//
	// 	- **None**: The policy takes effect for only once.
	//
	// 	- **Daily**: The policy takes effect on a daily basis.
	//
	// 	- **Weekly**: The policy takes effect on a weekly basis.
	//
	// 	- **Monthly**: The policy takes effect on a monthly basis.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// 	- If **SourceType*	- is set to `net`, the value of this parameter is a CIDR block. Example: 10.2.XX.XX/24.
	//
	// 	- If **SourceType*	- is set to `group`, the value of this parameter is an address book name. Example: db_group.
	//
	// example:
	//
	// 192.168.0.25/32
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// 	- **net**: CIDR block
	//
	// 	- **group**: address book
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of EndTime.
	//
	// >  If RepeatType is set to Permanent, StartTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, StartTime must be specified.
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
