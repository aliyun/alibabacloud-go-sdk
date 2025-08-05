// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *ModifyVpcFirewallControlPolicyRequest
	GetAclAction() *string
	SetAclUuid(v string) *ModifyVpcFirewallControlPolicyRequest
	GetAclUuid() *string
	SetApplicationName(v string) *ModifyVpcFirewallControlPolicyRequest
	GetApplicationName() *string
	SetApplicationNameList(v []*string) *ModifyVpcFirewallControlPolicyRequest
	GetApplicationNameList() []*string
	SetDescription(v string) *ModifyVpcFirewallControlPolicyRequest
	GetDescription() *string
	SetDestPort(v string) *ModifyVpcFirewallControlPolicyRequest
	GetDestPort() *string
	SetDestPortGroup(v string) *ModifyVpcFirewallControlPolicyRequest
	GetDestPortGroup() *string
	SetDestPortType(v string) *ModifyVpcFirewallControlPolicyRequest
	GetDestPortType() *string
	SetDestination(v string) *ModifyVpcFirewallControlPolicyRequest
	GetDestination() *string
	SetDestinationType(v string) *ModifyVpcFirewallControlPolicyRequest
	GetDestinationType() *string
	SetDomainResolveType(v string) *ModifyVpcFirewallControlPolicyRequest
	GetDomainResolveType() *string
	SetEndTime(v int64) *ModifyVpcFirewallControlPolicyRequest
	GetEndTime() *int64
	SetLang(v string) *ModifyVpcFirewallControlPolicyRequest
	GetLang() *string
	SetProto(v string) *ModifyVpcFirewallControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *ModifyVpcFirewallControlPolicyRequest
	GetRelease() *string
	SetRepeatDays(v []*int64) *ModifyVpcFirewallControlPolicyRequest
	GetRepeatDays() []*int64
	SetRepeatEndTime(v string) *ModifyVpcFirewallControlPolicyRequest
	GetRepeatEndTime() *string
	SetRepeatStartTime(v string) *ModifyVpcFirewallControlPolicyRequest
	GetRepeatStartTime() *string
	SetRepeatType(v string) *ModifyVpcFirewallControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *ModifyVpcFirewallControlPolicyRequest
	GetSource() *string
	SetSourceType(v string) *ModifyVpcFirewallControlPolicyRequest
	GetSourceType() *string
	SetStartTime(v int64) *ModifyVpcFirewallControlPolicyRequest
	GetStartTime() *int64
	SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyRequest
	GetVpcFirewallId() *string
}

type ModifyVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic.
	//
	// Valid values:
	//
	// 	- **accept**: allows the traffic.
	//
	// 	- **drop**: blocks the traffic.
	//
	// 	- **log**: monitors the traffic.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// If you want to modify the configurations of an access control policy, you must provide the unique ID of the policy. You can call the [DescribeVpcFirewallControlPolicy](https://help.aliyun.com/document_detail/159758.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// Deprecated
	//
	// The type of the application that the access control policy supports.
	//
	// Valid values:
	//
	// 	- ANY: all application types
	//
	// 	- FTP
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// 	- MySQL
	//
	// 	- SMTP
	//
	// 	- SMTPS
	//
	// 	- RDP
	//
	// 	- VNC
	//
	// 	- SSH
	//
	// 	- Redis
	//
	// 	- MQTT
	//
	// 	- MongoDB
	//
	// 	- Memcache
	//
	// 	- SSL
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application names.
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
	// The type of the destination port in the access control policy.
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
	// 	- If **DestinationType*	- is set to `net`, the value of this parameter must be a CIDR block.
	//
	//     Example: 10.2.3.0/24
	//
	// 	- If **DestinationType*	- is set to `group`, the value of this parameter must be an address book name.
	//
	//     Example: db_group
	//
	// 	- If **DestinationType*	- is set to `domain`, the value of this parameter must be a domain name.
	//
	//     Example: \\*.aliyuncs.com
	//
	// example:
	//
	// 10.2.X.X/XX
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
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
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
	// The protocol type in the access control policy.
	//
	// Valid values:
	//
	// 	- ANY: all protocol types
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- ICMP
	//
	// example:
	//
	// TCP
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
	// 	- If you set RepeatType to `Permanent`, `None`, or `Daily`, the value of this parameter is an empty array. Example: [].
	//
	// 	- If you set RepeatType to Weekly, you must specify this parameter. Example: [0, 6].
	//
	// >  If you set RepeatType to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// 	- If you set RepeatType to `Monthly`, you must specify this parameter. Example: [1, 31].
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
	// Valid values:
	//
	// 	- If **SourceType*	- is set to `net`, the value of this parameter must be a CIDR block.
	//
	//     Example: 10.2.4.0/24
	//
	// 	- If **SourceType*	- is set to `group`, the value of this parameter must be an address book name.
	//
	//     Example: db_group
	//
	// example:
	//
	// 10.2.X.X/XX
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy.
	//
	// Valid values:
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
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The instance ID of the VPC firewall. You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to query the ID.
	//
	// 	- If the VPC firewall is used to protect a CEN instance, the value of this parameter must be the ID of the CEN instance.
	//
	//     Example: cen-ervw0g12b5jbw\\*\\*\\*\\*
	//
	// 	- If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	//     Example: vfw-a42bbb7b887148c9\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetDestPort() *string {
	return s.DestPort
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetDestPortType() *string {
	return s.DestPortType
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetDomainResolveType() *string {
	return s.DomainResolveType
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModifyVpcFirewallControlPolicyRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetAclAction(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetAclUuid(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetApplicationName(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetApplicationNameList(v []*string) *ModifyVpcFirewallControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDescription(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPort(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPortGroup(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPortType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestination(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestinationType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDomainResolveType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DomainResolveType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetEndTime(v int64) *ModifyVpcFirewallControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetLang(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetProto(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRelease(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRepeatDays(v []*int64) *ModifyVpcFirewallControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRepeatEndTime(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRepeatStartTime(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRepeatType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSource(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSourceType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetStartTime(v int64) *ModifyVpcFirewallControlPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
