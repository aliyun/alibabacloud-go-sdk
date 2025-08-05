// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *CreateVpcFirewallControlPolicyRequest
	GetAclAction() *string
	SetApplicationName(v string) *CreateVpcFirewallControlPolicyRequest
	GetApplicationName() *string
	SetApplicationNameList(v []*string) *CreateVpcFirewallControlPolicyRequest
	GetApplicationNameList() []*string
	SetDescription(v string) *CreateVpcFirewallControlPolicyRequest
	GetDescription() *string
	SetDestPort(v string) *CreateVpcFirewallControlPolicyRequest
	GetDestPort() *string
	SetDestPortGroup(v string) *CreateVpcFirewallControlPolicyRequest
	GetDestPortGroup() *string
	SetDestPortType(v string) *CreateVpcFirewallControlPolicyRequest
	GetDestPortType() *string
	SetDestination(v string) *CreateVpcFirewallControlPolicyRequest
	GetDestination() *string
	SetDestinationType(v string) *CreateVpcFirewallControlPolicyRequest
	GetDestinationType() *string
	SetDomainResolveType(v string) *CreateVpcFirewallControlPolicyRequest
	GetDomainResolveType() *string
	SetEndTime(v int64) *CreateVpcFirewallControlPolicyRequest
	GetEndTime() *int64
	SetLang(v string) *CreateVpcFirewallControlPolicyRequest
	GetLang() *string
	SetMemberUid(v string) *CreateVpcFirewallControlPolicyRequest
	GetMemberUid() *string
	SetNewOrder(v string) *CreateVpcFirewallControlPolicyRequest
	GetNewOrder() *string
	SetProto(v string) *CreateVpcFirewallControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *CreateVpcFirewallControlPolicyRequest
	GetRelease() *string
	SetRepeatDays(v []*int64) *CreateVpcFirewallControlPolicyRequest
	GetRepeatDays() []*int64
	SetRepeatEndTime(v string) *CreateVpcFirewallControlPolicyRequest
	GetRepeatEndTime() *string
	SetRepeatStartTime(v string) *CreateVpcFirewallControlPolicyRequest
	GetRepeatStartTime() *string
	SetRepeatType(v string) *CreateVpcFirewallControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *CreateVpcFirewallControlPolicyRequest
	GetSource() *string
	SetSourceType(v string) *CreateVpcFirewallControlPolicyRequest
	GetSourceType() *string
	SetStartTime(v int64) *CreateVpcFirewallControlPolicyRequest
	GetStartTime() *int64
	SetVpcFirewallId(v string) *CreateVpcFirewallControlPolicyRequest
	GetVpcFirewallId() *string
}

type CreateVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// - **accept**: allows the traffic.
	//
	// - **drop**: blocks the traffic.
	//
	// - **log**: monitors the traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// Deprecated
	//
	// The type of the applications that the access control policy supports. Valid values:
	//
	// - **FTP**
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// - **MySQL**
	//
	// - **SMTP**
	//
	// - **SMTPS**
	//
	// - **RDP**
	//
	// - **VNC**
	//
	// - **SSH**
	//
	// - **Redis**
	//
	// - **MQTT**
	//
	// - **MongoDB**
	//
	// - **Memcache**
	//
	// - **SSL**
	//
	// - **ANY**: all types of applications
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application types supported by the access control policy.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	//
	// >  If **DestPortType*	- is set to `port`, you must specify this parameter.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// >  If **DestPortType*	- is set to `group`, you must specify this parameter.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// - **port**: port
	//
	// - **group**: port address book
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. Valid values:
	//
	// - If **DestinationType*	- is set to `net`, the value of **Destination*	- must be a CIDR block.
	//
	// - If **DestinationType*	- is set to `group`, the value of **Destination*	- must be an address book.
	//
	// - If **DestinationType*	- is set to `domain`, the value of **Destination*	- must be a domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.2.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: CIDR block
	//
	// - **group**: address book
	//
	// - **domain**: domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
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
	// DNS
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// - **zh**: Chinese (default)
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A smaller priority value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// - **ANY*	- (If you are not sure about the protocol type, you can set this parameter to ANY.)
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// - **true**: enables the access control policy.
	//
	// - **false**: disables the access control policy.
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
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
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
	// - If SourceType is set to `net`, the value of Source must be a CIDR block.
	//
	// - If SourceType is set to `group`, the value of Source must be an address book.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.2.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: CIDR block
	//
	// - **group**: address book
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
	// The ID of the policy group in which you want to create the access control policy.
	//
	// - If a VPC firewall protects the traffic between two VPCs that are connected by using a CEN instance, the value of this parameter must be the ID of the CEN instance.
	//
	// - If a VPC firewall protects the traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallAclGroupList](https://www.alibabacloud.com/help/en/cloud-firewall/latest/describevpcfirewallaclgrouplist) operation to query the IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *CreateVpcFirewallControlPolicyRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateVpcFirewallControlPolicyRequest) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *CreateVpcFirewallControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVpcFirewallControlPolicyRequest) GetDestPort() *string {
	return s.DestPort
}

func (s *CreateVpcFirewallControlPolicyRequest) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *CreateVpcFirewallControlPolicyRequest) GetDestPortType() *string {
	return s.DestPortType
}

func (s *CreateVpcFirewallControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *CreateVpcFirewallControlPolicyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *CreateVpcFirewallControlPolicyRequest) GetDomainResolveType() *string {
	return s.DomainResolveType
}

func (s *CreateVpcFirewallControlPolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateVpcFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateVpcFirewallControlPolicyRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *CreateVpcFirewallControlPolicyRequest) GetNewOrder() *string {
	return s.NewOrder
}

func (s *CreateVpcFirewallControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *CreateVpcFirewallControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *CreateVpcFirewallControlPolicyRequest) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *CreateVpcFirewallControlPolicyRequest) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *CreateVpcFirewallControlPolicyRequest) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *CreateVpcFirewallControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *CreateVpcFirewallControlPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *CreateVpcFirewallControlPolicyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateVpcFirewallControlPolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateVpcFirewallControlPolicyRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *CreateVpcFirewallControlPolicyRequest) SetAclAction(v string) *CreateVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetApplicationName(v string) *CreateVpcFirewallControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetApplicationNameList(v []*string) *CreateVpcFirewallControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDescription(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPort(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPortGroup(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPortType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestination(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestinationType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDomainResolveType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DomainResolveType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetEndTime(v int64) *CreateVpcFirewallControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetLang(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetMemberUid(v string) *CreateVpcFirewallControlPolicyRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetNewOrder(v string) *CreateVpcFirewallControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetProto(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRelease(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRepeatDays(v []*int64) *CreateVpcFirewallControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRepeatEndTime(v string) *CreateVpcFirewallControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRepeatStartTime(v string) *CreateVpcFirewallControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRepeatType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetSource(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetSourceType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetStartTime(v int64) *CreateVpcFirewallControlPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *CreateVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
