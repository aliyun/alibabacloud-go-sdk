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
	// The action that Cloud Firewall performs on the traffic.
	//
	// Valid values:
	//
	// 	- **accept**: allows the traffic.
	//
	// 	- **drop**: denies the traffic.
	//
	// 	- **log**: monitors the traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The application types supported by the access control policy.
	//
	// This parameter is required.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy. Valid values:
	//
	// 	- If Proto is set to ICMP, DestPort is automatically left empty.
	//
	// > If Proto is set to ICMP, access control does not take effect on the destination port.
	//
	// 	- If Proto is set to TCP, UDP, or ANY and DestPortType is set to group, DestPort is empty.
	//
	// > If DestPortType is set to group, you do not need to specify the destination port number. All ports on which the access control policy takes effect are included in the destination port address book.
	//
	// 	- If Proto is set to TCP, UDP, or ANY and DestPortType is set to port, the value of DestPort is the destination port number.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// > If DestPortType is set to group, you must specify the name of the destination port address book.
	//
	// example:
	//
	// my_port_group
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
	// Valid values:
	//
	// 	- If DestinationType is set to net, the value of this parameter is a CIDR block.
	//
	//     Example: 1.2.XX.XX/24
	//
	// 	- If DestinationType is set to group, the value of this parameter is an address book.
	//
	//     Example: db_group
	//
	// 	- If DestinationType is set to domain, the value of this parameter is a domain name.
	//
	//     Example: \\*.aliyuncs.com
	//
	// 	- If DestinationType is set to location, the value of this parameter is a location.
	//
	//     Example: ["BJ11", "ZB"]
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
	// 	- **net**: CIDR block
	//
	// 	- **group**: address book
	//
	// 	- **domain**: domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid value:
	//
	// 	- **out**: outbound.
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// 	- **0**: fully qualified domain name (FQDN)-based resolution
	//
	// 	- **1**: Domain Name System (DNS)-based dynamic resolution
	//
	// 	- **2**: FQDN and DNS-based dynamic resolution
	//
	// example:
	//
	// 0
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of StartTime.
	//
	// >  If RepeatType is set to Permanent, EndTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP version supported by the access control policy. Valid values:
	//
	// 	- **4**: IPv4 (default)
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response.
	//
	// Valid values:
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
	// ngx-xxxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol type in the access control policy.
	//
	// Valid values:
	//
	// 	- ANY: all types of protocols.
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- ICMP
	//
	// >  If the destination address is a threat intelligence address book of the domain name type or a cloud service address book, you can set Proto only to TCP and set ApplicationNameList to HTTP, HTTPS, SMTP, SMTPS, or SSL.
	//
	// This parameter is required.
	//
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	// >  If RepeatType is set to Permanent or None, RepeatEndTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
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
	// The source address in the access control policy.
	//
	// Valid values:
	//
	// 	- If **SourceType*	- is set to `net`, the value of Source is a CIDR block.
	//
	//     Example: 10.2.4.0/24
	//
	// 	- If **SourceType*	- is set to `group`, the value of this parameter must be an address book name.
	//
	//     Example: db_group
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
	// 	- **net**: source CIDR block
	//
	// 	- **group**: source address book
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of EndTime.
	//
	// >  If RepeatType is set to Permanent, StartTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
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
