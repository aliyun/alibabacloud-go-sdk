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
	SetClientToken(v string) *AddControlPolicyRequest
	GetClientToken() *string
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
	SetDryRun(v bool) *AddControlPolicyRequest
	GetDryRun() *bool
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
	// The action that Cloud Firewall performs on traffic based on the access control policy. Valid values:
	//
	// - **accept**: allows the traffic.
	//
	// - **drop**: denies the traffic.
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
	// - **ANY*	- (all application types)
	//
	// > The supported application types depend on the protocol type specified by the Proto parameter. If Proto is set to TCP, ApplicationName can be set to any of the preceding application types. If Proto is set to UDP, ICMP, or ANY, ApplicationName can only be set to ANY. You must specify either ApplicationNameList or ApplicationName. You cannot leave both parameters empty.
	//
	// example:
	//
	// ANY
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The list of application types supported by the access control policy.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ddadxefexxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Release flow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port for traffic in the access control policy. Valid values:
	//
	// - If the protocol type is ICMP, leave DestPort empty.
	//
	// > If the protocol type is ICMP, access control on destination ports is not supported.
	//
	// - If the protocol type is TCP, UDP, or ANY, and the destination port type (DestPortType) is group, leave DestPort empty.
	//
	// > If you set the destination port type to group (destination port address book), you do not need to specify a destination port number. All ports managed by the access control policy are included in the destination port address book.
	//
	// - If the protocol type is TCP, UDP, or ANY, and the destination port type (DestPortType) is port, set DestPort to the destination port number.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
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
	// Valid values:
	//
	// - If DestinationType is set to net, Destination specifies the destination CIDR block.
	//
	//
	//
	//   Example: 1.2.XX.XX/24
	//
	// - If DestinationType is set to group, Destination specifies the name of the destination address book.
	//
	//   Example: db_group
	//
	// - If DestinationType is set to domain, Destination specifies the destination domain name.
	//
	//   Example: *.aliyuncs.com
	//
	// - If DestinationType is set to location, Destination specifies the destination location.
	//
	//   Example: ["BJ11", "ZB"\]
	//
	// > If Destination specifies a destination location, for more information, see [Location codes](https://help.aliyun.com/document_detail/2854161.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The destination address type in the access control policy. Valid values:
	//
	// - **net**: destination CIDR block
	//
	// - **group**: destination address book
	//
	// - **domain**: destination domain name
	//
	// - **location**: destination region
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction of the access control policy. Valid values:
	//
	// - **in**: Inbound traffic.
	//
	// - **out**: Outbound traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution mode of the access control policy. Valid values:
	//
	// 	- **FQDN**: FQDN-based resolution
	//
	// 	- **DNS**: DNS-based dynamic resolution
	//
	// 	- **FQDN_AND_DNS**: FQDN- and DNS-based dynamic resolution
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// Specifies whether to perform a dry run.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The end time of the Policy Validity Period for the access control policy. The time is expressed in the UNIX timestamp in seconds format. The time must be on the hour or half hour, and must be at least 30 minutes later than the start time.
	//
	// > When RepeatType is set to `Permanent`, EndTime is empty. When RepeatType is set to `None`, `Daily`, `Weekly`, or `Monthly`, EndTime is required, and you must specify the end time.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address version.
	//
	// Valid values:
	//
	// - **4**: IPv4
	//
	// - **6**: IPv6
	//
	// example:
	//
	// 6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The priority of the access control policy. The priority value starts from 1 and increases sequentially. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol type of traffic in the access control policy. Valid values:
	//
	// - **ANY**: all protocols
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// > If the traffic direction is outbound and the destination address is a domain name in a threat intelligence address book or a cloud service address book, you can only set Proto to TCP. In this case, you can set the application type to HTTP, HTTPS, SMTP, SMTPS, or SSL.
	//
	// This parameter is required.
	//
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, the policy is enabled after it is created. Valid values:
	//
	// - **true**: enables the access control policy.
	//
	// - **false**: does not enable the access control policy.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The set of repeat days for the Policy Validity Period of the access control policy.
	//
	// - When RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is an empty set. Example: []
	//
	// - When RepeatType is set to `Weekly`, RepeatDays cannot be empty. Example: [0, 6]
	//
	// > When RepeatType is set to Weekly, the values in RepeatDays must be unique.
	//
	// - When RepeatType is set to `Monthly`, RepeatDays cannot be empty. Example: [1, 31]
	//
	// > When RepeatType is set to Monthly, the values in RepeatDays must be unique.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The repeat end time for the Policy Validity Period of the access control policy. Example: 23:30. The time must be on the hour or half hour, and must be at least 30 minutes later than the repeat start time.
	//
	// > When RepeatType is set to `Permanent` or `None`, RepeatEndTime is empty. When RepeatType is set to `Daily`, `Weekly`, or `Monthly`, RepeatEndTime is required, and you must specify the repeat end time.
	//
	// > The time format is HH:mm (24-hour clock). Examples: 08:00 and 23:30.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The repeat start time for the Policy Validity Period of the access control policy. Example: 08:00. The time must be on the hour or half hour, and must be at least 30 minutes earlier than the repeat end time.
	//
	// > When RepeatType is set to `Permanent` or `None`, RepeatStartTime is empty. When RepeatType is set to `Daily`, `Weekly`, or `Monthly`, RepeatStartTime is required, and you must specify the repeat start time.
	//
	// > The time format is HH:mm (24-hour clock). Examples: 08:00 and 23:30.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The repeat type of the Policy Validity Period for the access control policy. Valid values:
	//
	// - **Permanent*	- (default): always
	//
	// - **None**: a single specified time
	//
	// - **Daily**: daily
	//
	// - **Weekly**: weekly
	//
	// - **Monthly**: monthly
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// - If SourceType is set to net, set Source to the source CIDR block.
	//
	//   Example: 1.1.XX.XX/24
	//
	// - If SourceType is set to group, set Source to the name of the source address book.
	//
	//   Example: db_group
	//
	// - If SourceType is set to location, set Source to the source region.
	//
	//   Example: ["BJ11", "ZB"\]
	//
	// > For more information about source regions, see [Region codes](https://help.aliyun.com/document_detail/2854161.html).
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
	// - **location**: source location
	//
	// This parameter is required.
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the Policy Validity Period for the access control policy. The time is expressed in the UNIX timestamp in seconds format. The time must be on the hour or half hour, and must be at least 30 minutes earlier than the end time.
	//
	// > When RepeatType is set to `Permanent`, StartTime is empty. When RepeatType is set to `None`, `Daily`, `Weekly`, or `Monthly`, StartTime is required, and you must specify the start time.
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

func (s *AddControlPolicyRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *AddControlPolicyRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *AddControlPolicyRequest) SetClientToken(v string) *AddControlPolicyRequest {
	s.ClientToken = &v
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

func (s *AddControlPolicyRequest) SetDryRun(v bool) *AddControlPolicyRequest {
	s.DryRun = &v
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
