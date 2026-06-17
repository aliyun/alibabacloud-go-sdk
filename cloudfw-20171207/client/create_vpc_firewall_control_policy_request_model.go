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
	// - **accept**: Allows the traffic.
	//
	// - **drop**: Denies the traffic.
	//
	// - **log**: Monitors the traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// Deprecated
	//
	// The application type that the access control policy supports. Valid values:
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
	// - **ANY*	- (all application types)
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The list of application types that the access control policy supports.
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
	// > Set this parameter when **DestPortType*	- is set to `port`.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// > Set this parameter when **DestPortType*	- is set to `group`.
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
	// - If **DestinationType*	- is `net`, set this parameter to a destination CIDR block.
	//
	// - If **DestinationType*	- is `group`, set this parameter to the name of a destination address book.
	//
	// - If **DestinationType*	- is `domain`, set this parameter to a destination domain name.
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
	// The domain name resolution method for the access control policy. Valid values:
	//
	// - **FQDN**: FQDN-based
	//
	// - **DNS**: DNS-based dynamic resolution
	//
	// - **FQDN_AND_DNS**: FQDN-based and DNS-based dynamic resolution
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy validity period. This value is a UNIX timestamp. The time must be on the hour or half-hour and must be at least 30 minutes later than the start time.
	//
	// > If RepeatType is \\`Permanent\\`, leave this parameter empty. If RepeatType is \\`None\\`, \\`Daily\\`, \\`Weekly\\`, or \\`Monthly\\`, set this parameter.
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
	// The UID of the member account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The priority of the access control policy.
	//
	// The priority starts from 1. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// - **ANY*	- (Set this value if you are unsure of the protocol type.)
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
	// The status of the access control policy. The policy is enabled by default after it is created. Valid values:
	//
	// - **true**: Enables the access control policy.
	//
	// - **false**: Disables the access control policy.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of the week or month on which the policy is recurrently active.
	//
	// - If RepeatType is `Permanent`, `None`, or `Daily`, leave this parameter empty. Example: \\`[]\\`
	//
	// - If RepeatType is \\`Weekly\\`, set this parameter. Example: \\`[0, 6]\\`
	//
	// > If RepeatType is set to \\`Weekly\\`, the values in RepeatDays cannot be duplicates.
	//
	// - If **RepeatType*	- is \\`Monthly\\`, set this parameter. Example: \\`[1, 31]\\`
	//
	// > If RepeatType is set to \\`Monthly\\`, the values in RepeatDays cannot be duplicates.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurring end time of the policy validity period. For example: \\`23:30\\`. The time must be on the hour or half-hour and must be at least 30 minutes later than the recurring start time.
	//
	// > If RepeatType is \\`Permanent\\` or \\`None\\`, leave this parameter empty. If RepeatType is \\`Daily\\`, \\`Weekly\\`, or \\`Monthly\\`, set this parameter.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurring start time of the policy validity period. For example: \\`08:00\\`. The time must be on the hour or half-hour and must be at least 30 minutes earlier than the recurring end time.
	//
	// > If RepeatType is \\`Permanent\\` or \\`None\\`, leave this parameter empty. If RepeatType is \\`Daily\\`, \\`Weekly\\`, or \\`Monthly\\`, set this parameter.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the policy validity period. Valid values:
	//
	// - **Permanent*	- (default): always
	//
	// - **None**: one-time
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
	// - If SourceType is `net`, set this parameter to a source CIDR block.
	//
	// - If SourceType is `group`, set this parameter to the name of a source address book.
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
	// The start time of the policy validity period. This value is a UNIX timestamp. The time must be on the hour or half-hour and must be at least 30 minutes earlier than the end time.
	//
	// > If RepeatType is \\`Permanent\\`, leave this parameter empty. If RepeatType is \\`None\\`, \\`Daily\\`, \\`Weekly\\`, or \\`Monthly\\`, set this parameter.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the policy group for the VPC border firewall.
	//
	// - If the VPC border firewall protects traffic between two VPCs that are connected using a CEN instance, set this parameter to the ID of the CEN instance.
	//
	// - If the VPC border firewall protects traffic between two VPCs that are connected using an Express Connect circuit, set this parameter to the ID of the VPC border firewall instance.
	//
	// > Call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to get this ID.
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
