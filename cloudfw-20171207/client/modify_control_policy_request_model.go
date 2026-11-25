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
	SetClientToken(v string) *ModifyControlPolicyRequest
	GetClientToken() *string
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
	SetDryRun(v bool) *ModifyControlPolicyRequest
	GetDryRun() *bool
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
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	//
	// > To modify an access control policy, you must provide its UUID. You can call the [DescribeControlPolicy](https://help.aliyun.com/document_detail/138866.html) operation to obtain the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221ad84c
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// Deprecated
	//
	// The application type supported by the access control policy. Valid values:
	//
	// - **ANY**
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
	// > **ANY*	- indicates that the policy applies to all application types.
	//
	// > You must specify either ApplicationNameList or ApplicationName. They cannot be both empty. If both are specified, ApplicationNameList takes precedence.
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The list of application names.
	//
	// > You must specify either ApplicationNameList or ApplicationName. They cannot be both empty. If both are specified, ApplicationNameList takes precedence.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The idempotency token.
	//
	// example:
	//
	// ddadxefexxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port for traffic in the access control policy.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book for traffic in the access control policy.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port for traffic in the access control policy. Valid values:
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
	// - If **DestinationType*	- is set to net, **Destination*	- is the destination CIDR block. Example: 1.2.XX.XX/24.
	//
	// - If **DestinationType*	- is set to group, **Destination*	- is the name of the destination address book. Example: db_group.
	//
	// - If **DestinationType*	- is set to domain, **Destination*	- is the destination domain name. Example: *.aliyuncs.com.
	//
	// - If **DestinationType*	- is set to location, **Destination*	- is the destination area (for the positional encoding of specific areas, see below). Example: ["BJ11", "ZB"\].
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
	// - **location**: destination region
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction of the access control policy. Valid values:
	//
	// - **in**: inbound traffic
	//
	// - **out**: outbound traffic
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution mode of the access control policy. Valid values:
	//
	// 	- **FQDN**: FQDN
	//
	// 	- **DNS**: DNS-based dynamic resolution
	//
	// 	- **FQDN_AND_DNS**: FQDN and DNS-based dynamic resolution
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// Specifies whether to perform a dry run.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The end time of the Policy Validity Period for the access control policy. The value is expressed in the format of a UNIX timestamp in seconds. The time must be on the hour or half-hour, and must be at least 30 minutes later than the start time.
	//
	// > When RepeatType is Permanent, EndTime is empty. When RepeatType is None, Daily, Weekly, or Monthly, EndTime must have a value, and you need to configure the end time Settings.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// The protocol type of traffic in the access control policy. Valid values:
	//
	// - **ANY**
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// > **ANY*	- indicates that the policy applies to all protocol types.
	//
	// > If the traffic direction is outbound and the destination address is a threat intelligence address book or cloud service address book of the domain name type, you can set this parameter to TCP or ANY. If you select TCP, you can set the application to HTTP, HTTPS, SMTP, SMTPS, or SSL. If you select ANY, you can set the application to ANY.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// - true: enabling status.
	//
	// - false: shutdown.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The collection of recurrence days for the Policy Validity Period of the access control policy.
	//
	// - If RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is an empty set. Example: [].
	//
	// - If RepeatType is set to Weekly, RepeatDays cannot be empty. Example: [0, 6].
	//
	// > When the RepeatType Settings is Weekly, the values in RepeatDays cannot be duplicate.
	//
	// - If RepeatType is set to `Monthly`, RepeatDays cannot be empty. Example: [1, 31].
	//
	// > When the RepeatType Settings is Monthly, the values in RepeatDays cannot be duplicate.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurrence end time when the policy takes effect. Use the 24-hour time format HH:mm, such as 23:00.
	//
	// > If RepeatType is set to Permanent or None, RepeatEndTime is empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime must be specified.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time when the policy takes effect. Use the 24-hour time format HH:mm, such as 08:00.
	//
	// > If RepeatType is set to Permanent or None, RepeatStartTime is empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatStartTime must be specified.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type of the Policy Validity Period for the access control policy. Valid values:
	//
	// - **Permanent*	- (default): always
	//
	// - **None**: single occurrence
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
	// The source address in the access control policy.
	//
	// - If **SourceType*	- is set to net, **Source*	- is the source CIDR block. Example: 1.2.XX.XX/24.
	//
	// - If **SourceType*	- is set to group, **Source*	- is the name of the source address book. Example: db_group.
	//
	// - If **SourceType*	- is set to location, **Source*	- is the source area (for the positional encoding of specific areas, see below). Example: ["BJ11", "ZB"\].
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: source CIDR block
	//
	// - **group**: source address book
	//
	// - **location**: source region
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the Policy Validity Period for the access control policy. The value is expressed in the format of a UNIX timestamp in seconds. The time must be on the hour or half-hour, and must be at least 30 minutes earlier than the end time.
	//
	// > When RepeatType is Permanent, StartTime is empty. When RepeatType is None, Daily, Weekly, or Monthly, StartTime must have a value, and you need to configure the start time Settings.
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

func (s *ModifyControlPolicyRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *ModifyControlPolicyRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *ModifyControlPolicyRequest) SetClientToken(v string) *ModifyControlPolicyRequest {
	s.ClientToken = &v
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

func (s *ModifyControlPolicyRequest) SetDryRun(v bool) *ModifyControlPolicyRequest {
	s.DryRun = &v
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
