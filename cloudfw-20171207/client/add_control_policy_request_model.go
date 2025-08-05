// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *AddControlPolicyRequest
	GetAclAction() *string
	SetApplicationName(v string) *AddControlPolicyRequest
	GetApplicationName() *string
	SetApplicationNameList(v []*string) *AddControlPolicyRequest
	GetApplicationNameList() []*string
	SetDescription(v string) *AddControlPolicyRequest
	GetDescription() *string
	SetDestPort(v string) *AddControlPolicyRequest
	GetDestPort() *string
	SetDestPortGroup(v string) *AddControlPolicyRequest
	GetDestPortGroup() *string
	SetDestPortType(v string) *AddControlPolicyRequest
	GetDestPortType() *string
	SetDestination(v string) *AddControlPolicyRequest
	GetDestination() *string
	SetDestinationType(v string) *AddControlPolicyRequest
	GetDestinationType() *string
	SetDirection(v string) *AddControlPolicyRequest
	GetDirection() *string
	SetDomainResolveType(v string) *AddControlPolicyRequest
	GetDomainResolveType() *string
	SetEndTime(v int64) *AddControlPolicyRequest
	GetEndTime() *int64
	SetIpVersion(v string) *AddControlPolicyRequest
	GetIpVersion() *string
	SetLang(v string) *AddControlPolicyRequest
	GetLang() *string
	SetNewOrder(v string) *AddControlPolicyRequest
	GetNewOrder() *string
	SetProto(v string) *AddControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *AddControlPolicyRequest
	GetRelease() *string
	SetRepeatDays(v []*int64) *AddControlPolicyRequest
	GetRepeatDays() []*int64
	SetRepeatEndTime(v string) *AddControlPolicyRequest
	GetRepeatEndTime() *string
	SetRepeatStartTime(v string) *AddControlPolicyRequest
	GetRepeatStartTime() *string
	SetRepeatType(v string) *AddControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *AddControlPolicyRequest
	GetSource() *string
	SetSourceIp(v string) *AddControlPolicyRequest
	GetSourceIp() *string
	SetSourceType(v string) *AddControlPolicyRequest
	GetSourceType() *string
	SetStartTime(v int64) *AddControlPolicyRequest
	GetStartTime() *int64
}

type AddControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
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
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// Deprecated
	//
	// The application type supported by the access control policy. Valid values:
	//
	// 	- **FTP**
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// 	- **Memcache**
	//
	// 	- **MongoDB**
	//
	// 	- **MQTT**
	//
	// 	- **MySQL**
	//
	// 	- **RDP**
	//
	// 	- **Redis**
	//
	// 	- **SMTP**
	//
	// 	- **SMTPS**
	//
	// 	- **SSH**
	//
	// 	- **SSL_No_Cert**
	//
	// 	- **SSL**
	//
	// 	- **VNC**
	//
	// 	- **ANY**
	//
	// > The value of this parameter is based on the value of Proto. If Proto is set to TCP, you can set ApplicationName to any valid value. If Proto is set to UDP, ICMP, or ANY, you can set ApplicationName only to ANY. You must specify at least one of the ApplicationNameList and ApplicationName parameters.
	//
	// example:
	//
	// ANY
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application types supported by the access control policy.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Allows traffic
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
	// The type of the destination port in the access control policy.
	//
	// Valid values:
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
	// 	- If DestinationType is set to group, the value of this parameter is an address book name.
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
	// 192.0.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// 	- **net**: CIDR block
	//
	// 	- **group**: address book
	//
	// 	- **domain**: domain name
	//
	// 	- **location**: location
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// 	- **in**: inbound traffic
	//
	// 	- **out**: outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// 	- **FQDN**: fully qualified domain name (FQDN)-based resolution
	//
	// 	- **DNS**: DNS-based dynamic resolution
	//
	// 	- **FQDN_AND_DNS**: FQDN and DNS-based dynamic resolution
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP version supported by the access control policy.
	//
	// Valid values:
	//
	// 	- **4**: IPv4
	//
	// 	- **6**: IPv6
	//
	// example:
	//
	// 6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol type supported by the access control policy. Valid values:
	//
	// 	- **ANY**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **ICMP**
	//
	// This parameter is required.
	//
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// 	- **true**: enables the access control policy.
	//
	// 	- **false**: disables the access control policy.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// 	- If you set RepeatType to `Permanent`, `None`, or `Daily`, leave this parameter empty. Example: [].
	//
	// 	- If you set RepeatType to Weekly, you must specify this parameter. Example: [0, 6].
	//
	// >  If you set RepeatType to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// 	- If you set RepeatType to `Monthly`, you must specify this parameter. Example: [1, 31].
	//
	// >  If you set RepeatType to Monthly, the fields in the value of this parameter cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The end time must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The start time must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
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
	// 	- If SourceType is set to net, the value of this parameter is a CIDR block.
	//
	//     Example: 1.1.XX.XX/24
	//
	// 	- If SourceType is set to group, the value of this parameter is an address book name.
	//
	//     Example: db_group
	//
	// 	- If SourceType is set to location, the value of this parameter is a location.
	//
	//     Example: ["BJ11", "ZB"]
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// 	- **net**: CIDR block
	//
	// 	- **group**: address book
	//
	// 	- **location**: location
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s AddControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *AddControlPolicyRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *AddControlPolicyRequest) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *AddControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *AddControlPolicyRequest) GetDestPort() *string {
	return s.DestPort
}

func (s *AddControlPolicyRequest) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *AddControlPolicyRequest) GetDestPortType() *string {
	return s.DestPortType
}

func (s *AddControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *AddControlPolicyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *AddControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *AddControlPolicyRequest) GetDomainResolveType() *string {
	return s.DomainResolveType
}

func (s *AddControlPolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AddControlPolicyRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *AddControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *AddControlPolicyRequest) GetNewOrder() *string {
	return s.NewOrder
}

func (s *AddControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *AddControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *AddControlPolicyRequest) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *AddControlPolicyRequest) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *AddControlPolicyRequest) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *AddControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *AddControlPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *AddControlPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *AddControlPolicyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *AddControlPolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *AddControlPolicyRequest) SetAclAction(v string) *AddControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *AddControlPolicyRequest) SetApplicationName(v string) *AddControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *AddControlPolicyRequest) SetApplicationNameList(v []*string) *AddControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *AddControlPolicyRequest) SetDescription(v string) *AddControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPort(v string) *AddControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPortGroup(v string) *AddControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPortType(v string) *AddControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestination(v string) *AddControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestinationType(v string) *AddControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *AddControlPolicyRequest) SetDirection(v string) *AddControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *AddControlPolicyRequest) SetDomainResolveType(v string) *AddControlPolicyRequest {
	s.DomainResolveType = &v
	return s
}

func (s *AddControlPolicyRequest) SetEndTime(v int64) *AddControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *AddControlPolicyRequest) SetIpVersion(v string) *AddControlPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *AddControlPolicyRequest) SetLang(v string) *AddControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *AddControlPolicyRequest) SetNewOrder(v string) *AddControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *AddControlPolicyRequest) SetProto(v string) *AddControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *AddControlPolicyRequest) SetRelease(v string) *AddControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *AddControlPolicyRequest) SetRepeatDays(v []*int64) *AddControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *AddControlPolicyRequest) SetRepeatEndTime(v string) *AddControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *AddControlPolicyRequest) SetRepeatStartTime(v string) *AddControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *AddControlPolicyRequest) SetRepeatType(v string) *AddControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *AddControlPolicyRequest) SetSource(v string) *AddControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *AddControlPolicyRequest) SetSourceIp(v string) *AddControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *AddControlPolicyRequest) SetSourceType(v string) *AddControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *AddControlPolicyRequest) SetStartTime(v int64) *AddControlPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *AddControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
