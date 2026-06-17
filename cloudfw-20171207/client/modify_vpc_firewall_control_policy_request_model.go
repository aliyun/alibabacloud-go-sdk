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
	// The action that is performed on traffic that hits the access control policy.
	//
	// Valid values:
	//
	// - **accept**: allows the traffic.
	//
	// - **drop**: denies the traffic.
	//
	// - **log**: monitors the traffic.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// To modify an access control policy, you must provide the unique ID of the policy. You can call the [DescribeVpcFirewallControlPolicy](https://help.aliyun.com/document_detail/159758.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// Deprecated
	//
	// The application type supported by the access control policy.
	//
	// Valid values:
	//
	// - ANY (all application types)
	//
	// - FTP
	//
	// - HTTP
	//
	// - HTTPS
	//
	// - MySQL
	//
	// - SMTP
	//
	// - SMTPS
	//
	// - RDP
	//
	// - VNC
	//
	// - SSH
	//
	// - Redis
	//
	// - MQTT
	//
	// - MongoDB
	//
	// - Memcache
	//
	// - SSL
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The list of application names.
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
	// The name of the destination port address book for the traffic to which the access control policy applies.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port for the traffic to which the access control policy applies.
	//
	// - **port**: port
	//
	// - **group**: port address book
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// - If **DestinationType*	- is set to `net`, specify a destination CIDR block.
	//
	//   Example: 10.2.3.0/24
	//
	// - If **DestinationType*	- is set to `group`, specify the name of a destination address book.
	//
	//   Example: db_group
	//
	// - If **DestinationType*	- is set to `domain`, specify a destination domain name.
	//
	//   Example: \\*.aliyuncs.com
	//
	// example:
	//
	// 10.2.X.X/XX
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
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The domain name resolution method for the access control policy. Valid values:
	//
	// - **FQDN**: based on Fully Qualified Domain Name (FQDN)
	//
	// - **DNS**: based on dynamic DNS resolution
	//
	// - **FQDN_AND_DNS**: based on FQDN and dynamic DNS resolution
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy validity period. The value is a UNIX timestamp. The time is on the hour or on the half hour, and is at least 30 minutes later than the start time.
	//
	// > If RepeatType is set to \\`Permanent\\`, this parameter is empty. If RepeatType is set to \\`None\\`, \\`Daily\\`, \\`Weekly\\`, or \\`Monthly\\`, you must specify this parameter.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response message.
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
	// The protocol type for the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// - ANY (all protocol types)
	//
	// - TCP
	//
	// - UDP
	//
	// - ICMP
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// - **true**: enables the access control policy.
	//
	// - **false**: disables the access control policy.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or month on which the policy takes effect.
	//
	// - If RepeatType is set to \\`Permanent\\`, \\`None\\`, or \\`Daily\\`, this parameter is empty.
	//
	//   Example: []
	//
	// - If RepeatType is set to \\`Weekly\\`, this parameter cannot be empty.
	//
	//   Example: [0, 6]
	//
	// > If RepeatType is set to \\`Weekly\\`, the elements in the array cannot be repeated.
	//
	// - If RepeatType is set to \\`Monthly\\`, this parameter cannot be empty.
	//
	//   Example: [1, 31]
	//
	// > If RepeatType is set to \\`Monthly\\`, the elements in the array cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The end time of the recurrence. The time is in the HH:mm format. The time is on the hour or on the half hour, and is at least 30 minutes later than the start time. Example: 23:30.
	//
	// > If RepeatType is set to \\`Permanent\\` or \\`None\\`, this parameter is empty. If RepeatType is set to \\`Daily\\`, \\`Weekly\\`, or \\`Monthly\\`, you must specify this parameter.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The start time of the recurrence. The time is in the HH:mm format. The time is on the hour or on the half hour, and is at least 30 minutes earlier than the end time. Example: 08:00.
	//
	// > If RepeatType is set to \\`Permanent\\` or \\`None\\`, this parameter is empty. If RepeatType is set to \\`Daily\\`, \\`Weekly\\`, or \\`Monthly\\`, you must specify this parameter.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the policy to take effect. Valid values:
	//
	// - **Permanent*	- (default): The policy is always in effect.
	//
	// - **None**: The policy takes effect for a specific period of time.
	//
	// - **Daily**: The policy takes effect daily.
	//
	// - **Weekly**: The policy takes effect weekly.
	//
	// - **Monthly**: The policy takes effect monthly.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy.
	//
	// Valid values:
	//
	// - If **SourceType*	- is set to `net`, specify a source CIDR block.
	//
	//   Example: 10.2.4.0/24
	//
	// - If **SourceType*	- is set to `group`, specify the name of a source address book.
	//
	//   Example: db_group
	//
	// example:
	//
	// 10.2.X.X/XX
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy.
	//
	// Valid values:
	//
	// - **net**: source CIDR block
	//
	// - **group**: source address book
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the policy validity period. The value is a UNIX timestamp. The time is on the hour or on the half hour, and is at least 30 minutes earlier than the end time.
	//
	// > If RepeatType is set to \\`Permanent\\`, this parameter is empty. If RepeatType is set to \\`None\\`, \\`Daily\\`, \\`Weekly\\`, or \\`Monthly\\`, you must specify this parameter.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the VPC firewall instance. You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to query the ID.
	//
	// - If the VPC firewall protects a CEN instance, the value of this parameter is the ID of the CEN instance.
	//
	//   Example: cen-ervw0g12b5jbw\\*\\*\\*\\*
	//
	// - If the VPC firewall protects an Express Connect circuit, the value of this parameter is the ID of the VPC firewall instance.
	//
	//   Example: vfw-a42bbb7b887148c9\\*\\*\\*\\*
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
