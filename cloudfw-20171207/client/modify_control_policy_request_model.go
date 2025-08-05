// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *ModifyControlPolicyRequest
	GetAclAction() *string
	SetAclUuid(v string) *ModifyControlPolicyRequest
	GetAclUuid() *string
	SetApplicationName(v string) *ModifyControlPolicyRequest
	GetApplicationName() *string
	SetApplicationNameList(v []*string) *ModifyControlPolicyRequest
	GetApplicationNameList() []*string
	SetDescription(v string) *ModifyControlPolicyRequest
	GetDescription() *string
	SetDestPort(v string) *ModifyControlPolicyRequest
	GetDestPort() *string
	SetDestPortGroup(v string) *ModifyControlPolicyRequest
	GetDestPortGroup() *string
	SetDestPortType(v string) *ModifyControlPolicyRequest
	GetDestPortType() *string
	SetDestination(v string) *ModifyControlPolicyRequest
	GetDestination() *string
	SetDestinationType(v string) *ModifyControlPolicyRequest
	GetDestinationType() *string
	SetDirection(v string) *ModifyControlPolicyRequest
	GetDirection() *string
	SetDomainResolveType(v string) *ModifyControlPolicyRequest
	GetDomainResolveType() *string
	SetEndTime(v int64) *ModifyControlPolicyRequest
	GetEndTime() *int64
	SetLang(v string) *ModifyControlPolicyRequest
	GetLang() *string
	SetProto(v string) *ModifyControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *ModifyControlPolicyRequest
	GetRelease() *string
	SetRepeatDays(v []*int64) *ModifyControlPolicyRequest
	GetRepeatDays() []*int64
	SetRepeatEndTime(v string) *ModifyControlPolicyRequest
	GetRepeatEndTime() *string
	SetRepeatStartTime(v string) *ModifyControlPolicyRequest
	GetRepeatStartTime() *string
	SetRepeatType(v string) *ModifyControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *ModifyControlPolicyRequest
	GetSource() *string
	SetSourceType(v string) *ModifyControlPolicyRequest
	GetSourceType() *string
	SetStartTime(v int64) *ModifyControlPolicyRequest
	GetStartTime() *int64
}

type ModifyControlPolicyRequest struct {
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
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	//
	// >  To modify an access control policy, you must specify the UUID of the policy. You can call the [DescribeControlPolicy](https://help.aliyun.com/document_detail/138866.html) interface to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221ad84c
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// Deprecated
	//
	// The type of the application that the access control policy supports. Valid values:
	//
	// 	- **ANY**
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// 	- **MySQL**
	//
	// 	- **SMTP**
	//
	// 	- **SMTPS**
	//
	// 	- **RDP**
	//
	// 	- **VNC**
	//
	// 	- **SSH**
	//
	// 	- **Redis**
	//
	// 	- **MQTT**
	//
	// 	- **MongoDB**
	//
	// 	- **Memcache**
	//
	// 	- **SSL**
	//
	// >  The value **ANY*	- indicates all types of applications.
	//
	// >  You must specify one of the ApplicationNameList and ApplicationName parameters. If you configure both ApplicationNameList and ApplicationName, only the value of ApplicationNameList is used.
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application names.
	//
	// >  You must specify one of the ApplicationNameList and ApplicationName parameters. If you configure both ApplicationNameList and ApplicationName, only the value of ApplicationNameList is used.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
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
	// 	- If **DestinationType*	- is set to net, the value of **Destination*	- is a CIDR block. Example: 1.2.XX.XX/24.
	//
	// 	- If **DestinationType*	- is set to group, the value of **Destination*	- is an address book. Example: db_group.
	//
	// 	- If **DestinationType*	- is set to domain, the value of **Destination*	- is a domain name. Example: \\*.aliyuncs.com.
	//
	// 	- If **DestinationType*	- is set to location, the value of **Destination*	- is a location. For more information about the location codes, see the "AddControlPolicy" topic. Example: ["BJ11", "ZB"].
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
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method of the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
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
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of StartTime.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
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
	// The protocol type that the access control policy supports. Valid values:
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
	// >  If the traffic direction is outbound and the destination address is a threat intelligence address book of the domain name type or a cloud service address book, you can set Proto to TCP or ANY. If you set Proto to TCP, you can set ApplicationName to HTTP, HTTPS, SMTP, SMTPS, and SSL. If you set Proto to ANY, you can set ApplicationName to ANY.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// 	- If you set RepeatType to `Permanent`, `None`, or `Daily`, the value of this parameter is an empty array. Example: []
	//
	// 	- If you set RepeatType to Weekly, you must specify this parameter. Example: [0, 6]
	//
	// >  If you set RepeatType to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// 	- If you set RepeatType to `Monthly`, you must specify this parameter. Example: [1, 31]
	//
	// >  If you set RepeatType to Monthly, the fields in the value of this parameter cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of RepeatStartTime.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of RepeatEndTime.
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
	// The source address in the access control policy.
	//
	// 	- If **SourceType*	- is set to net, the value of **Source*	- is a CIDR block. Example: 1.2.XX.XX/24.
	//
	// 	- If **SourceType*	- is set to group, the value of **Source*	- is an address book. Example: db_group.
	//
	// 	- If **SourceType*	- is set to location, the value of **Source*	- is a location. For more information about the location codes, see the "AddControlPolicy" topic. Example: ["BJ11", "ZB"]
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// 	- **net**: CIDR block
	//
	// 	- **group**: address book
	//
	// 	- **location**: location
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of EndTime.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *ModifyControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ModifyControlPolicyRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ModifyControlPolicyRequest) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *ModifyControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyControlPolicyRequest) GetDestPort() *string {
	return s.DestPort
}

func (s *ModifyControlPolicyRequest) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *ModifyControlPolicyRequest) GetDestPortType() *string {
	return s.DestPortType
}

func (s *ModifyControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *ModifyControlPolicyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ModifyControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyControlPolicyRequest) GetDomainResolveType() *string {
	return s.DomainResolveType
}

func (s *ModifyControlPolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModifyControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *ModifyControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *ModifyControlPolicyRequest) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *ModifyControlPolicyRequest) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *ModifyControlPolicyRequest) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *ModifyControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *ModifyControlPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *ModifyControlPolicyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyControlPolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModifyControlPolicyRequest) SetAclAction(v string) *ModifyControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetAclUuid(v string) *ModifyControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetApplicationName(v string) *ModifyControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetApplicationNameList(v []*string) *ModifyControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *ModifyControlPolicyRequest) SetDescription(v string) *ModifyControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPort(v string) *ModifyControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPortGroup(v string) *ModifyControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPortType(v string) *ModifyControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestination(v string) *ModifyControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestinationType(v string) *ModifyControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDirection(v string) *ModifyControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDomainResolveType(v string) *ModifyControlPolicyRequest {
	s.DomainResolveType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetEndTime(v int64) *ModifyControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetLang(v string) *ModifyControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetProto(v string) *ModifyControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRelease(v string) *ModifyControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRepeatDays(v []*int64) *ModifyControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *ModifyControlPolicyRequest) SetRepeatEndTime(v string) *ModifyControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRepeatStartTime(v string) *ModifyControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRepeatType(v string) *ModifyControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSource(v string) *ModifyControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSourceType(v string) *ModifyControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetStartTime(v int64) *ModifyControlPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
