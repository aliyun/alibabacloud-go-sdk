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
	// The action that is set in the access control policy. Settings the method in which traffic passes through Cloud Firewall. Valid values:
	//
	// - **accept**: allows the access.
	//
	// - **drop**: deny the access.
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
	// The application type supported by the access control policy. Valid values:
	//
	// - **FTP**
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// - **Memcache**
	//
	// - **MongoDB**
	//
	// - **MQTT**
	//
	// - **MySQL**
	//
	// - **RDP**
	//
	// - **Redis**
	//
	// - **SMTP**
	//
	// - **SMTPS**
	//
	// - **SSH**
	//
	// - **SSL_No_Cert**
	//
	// - **SSL**
	//
	// - **VNC**
	//
	// - **ANY**: all application types
	//
	// > The valid values of ApplicationName depend on the value of the protocol type (Proto). If Proto is set to TCP, ApplicationName can be set to any of the preceding application types. If Proto is set to UDP, ICMP, or ANY, ApplicationName can be set only to ANY. You must specify either ApplicationNameList or ApplicationName. You cannot leave both of them empty.
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
	// Release flow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy. Valid values:
	//
	// - If the protocol type is ICMP, the value of DestPort is empty.
	//
	//
	//
	// > If the protocol type is ICMP, access control on the destination port is not supported.
	//
	// - If the protocol type is TCP, UDP, or ANY, and the destination port type (DestPortType) is group, the value of DestPort is empty.
	//
	// > If the destination port type of the access control policy is set to group (port address book), you do not need to specify a destination port number. All ports that the access control policy manages are included in the port address book.
	//
	// - If the protocol type is TCP, UDP, or ANY, and the destination port type (DestPortType) is port, the value of DestPort is the destination port number.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	//
	// > If DestPortType is set to group, you must specify the destination port address book name.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy.
	//
	// Valid values:
	//
	// - **port**: port
	//
	// - **group**: port address book.
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// Valid values:
	//
	// - If DestinationType is set to net, the value of Destination is a destination CIDR block.
	//
	//
	//
	//   Example: 1.2.XX.XX/24
	//
	// - If DestinationType is set to group, the value of Destination is a destination address book name.
	//
	//   Example: db_group
	//
	// - If DestinationType is set to domain, the value of Destination is a destination domain name.
	//
	//   Example: *.aliyuncs.com
	//
	// - If DestinationType is set to location, the value of Destination is a destination region.
	//
	//   Example: ["BJ11", "ZB"\\]
	//
	// > If Destination is set to a destination region, for more information, see [Region codes](https://help.aliyun.com/document_detail/2854161.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: destination CIDR block
	//
	// - **group**: destination address book
	//
	// - **domain**: destination domain name
	//
	// - **location**: destination region.
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction of the access control policy. Valid values:
	//
	// - **in**: inbound traffic
	//
	// - **out**: outbound traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// 	- **FQDN**: FQDN-based resolution
	//
	// 	- **DNS**: DNS-based dynamic resolution
	//
	// 	- **FQDN_AND_DNS**: FQDN-based and DNS-based dynamic resolution.
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy validity period for the access control policy. The value is a UNIX timestamp in seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// > If RepeatType is set to Permanent, this parameter is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter is required.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address version supported.
	//
	// Valid values:
	//
	// - **4**: IPv4
	//
	// - **6**: IPv6.
	//
	// example:
	//
	// 6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The priority of the access control policy. The priority value starts from 1. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// - **ANY**: any protocol
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// > If the traffic direction is outbound and the destination address is a threat intelligence address book or cloud service address book of the domain name type, only TCP is supported. The application type can be set to HTTP, HTTPS, SMTP, SMTPS, or SSL.
	//
	// This parameter is required.
	//
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. The policy is enabled by default after it is created. Valid values:
	//
	// - **true**: enables the access control policy.
	//
	// - **false**: disables the access control policy.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of the recurrence for the policy validity period of the access control policy.
	//
	// - If RepeatType is set to `Permanent`, `None`, or `Daily`, the value of RepeatDays is an empty array.
	//
	//   Example: []
	//
	// - If RepeatType is set to Weekly, the value of RepeatDays must not be empty.
	//
	//   Example: [0, 6]
	//
	// > If RepeatType is set to Weekly, the values in RepeatDays cannot be repeated.
	//
	// - If RepeatType is set to `Monthly`, the value of RepeatDays must not be empty.
	//
	//   Example: [1, 31]
	//
	// > If RepeatType is set to Monthly, the values in RepeatDays cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurrence end time of the policy validity period for the access control policy. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the recurrence start time.
	//
	// > If RepeatType is set to Permanent or None, this parameter is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter is required.
	//
	// > The time is in the HH:mm format (24-hour clock). Example: 08:00 or 23:30.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time of the policy validity period for the access control policy. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the recurrence end time.
	//
	// > If RepeatType is set to Permanent or None, this parameter is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter is required.
	//
	// > The time is in the HH:mm format (24-hour clock). Example: 08:00 or 23:30.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type of the policy validity period for the access control policy. Valid values:
	//
	// - **Permanent*	- (default): The policy is always valid.
	//
	// - **None**: The policy is valid for a specified single time period.
	//
	// - **Daily**: The policy is valid on a daily basis.
	//
	// - **Weekly**: The policy is valid on a weekly basis.
	//
	// - **Monthly**: The policy is valid on a monthly basis.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// - If SourceType is set to net, the value of Source is a source CIDR block.
	//
	//   Example: 1.1.XX.XX/24
	//
	// - If SourceType is set to group, the value of Source is a source address book name.
	//
	//   Example: db_group
	//
	// - If SourceType is set to location, the value of Source is a source region.
	//
	//   Example: ["BJ11", "ZB"\\]
	//
	// > If Source is set to a source region, for more information, see [Region codes](https://help.aliyun.com/document_detail/2854161.html).
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
	// - **net**: source CIDR block
	//
	// - **group**: source address book
	//
	// - **location**: source region.
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the policy validity period for the access control policy. The value is a UNIX timestamp in seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// > If RepeatType is set to Permanent, this parameter is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter is required.
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
