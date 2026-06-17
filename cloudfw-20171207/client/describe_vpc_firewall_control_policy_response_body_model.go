// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicys(v []*DescribeVpcFirewallControlPolicyResponseBodyPolicys) *DescribeVpcFirewallControlPolicyResponseBody
	GetPolicys() []*DescribeVpcFirewallControlPolicyResponseBodyPolicys
	SetRequestId(v string) *DescribeVpcFirewallControlPolicyResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeVpcFirewallControlPolicyResponseBody
	GetTotalCount() *string
}

type DescribeVpcFirewallControlPolicyResponseBody struct {
	// The VPC firewall access control policies.
	Policys []*DescribeVpcFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VPC firewall access control policies.
	//
	// example:
	//
	// 20
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) GetPolicys() []*DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	return s.Policys
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetPolicys(v []*DescribeVpcFirewallControlPolicyResponseBodyPolicys) *DescribeVpcFirewallControlPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *DescribeVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetTotalCount(v string) *DescribeVpcFirewallControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) Validate() error {
	if s.Policys != nil {
		for _, item := range s.Policys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallControlPolicyResponseBodyPolicys struct {
	// The action to perform on traffic that matches the access control policy. Valid values:
	//
	// - **accept**: allows the traffic.
	//
	// - **drop**: denies the traffic.
	//
	// - **log**: logs the traffic.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique identifier of the access control policy.
	//
	// example:
	//
	// 4037fbf7-3e39-4634-92a4-d0155247****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 10**
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type. We recommend that you use `ApplicationNameList` instead. Valid values:
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
	// The list of application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The UNIX timestamp, in seconds, of when the policy was created.
	//
	// example:
	//
	// 1761062400
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The policy description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port. Valid values:
	//
	// - **port**: a single port
	//
	// - **group**: a port address book
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address for the access control policy. The value depends on `DestinationType`.
	//
	// - If `DestinationType` is `net`, the value is a destination CIDR block.
	//
	// - If `DestinationType` is `domain`, the value is a destination domain name.
	//
	// - If `DestinationType` is `group`, the value is the name of a destination address book.
	//
	// example:
	//
	// 192.0.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR blocks in the destination address book.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book. Valid values:
	//
	// - **ip**: an address book of IP addresses or CIDR blocks.
	//
	// - **domain**: an address book of domain names.
	//
	// example:
	//
	// ip
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address. Valid values:
	//
	// - **net**: a destination CIDR block
	//
	// - **group**: a destination address book
	//
	// - **domain**: a destination domain name
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The domain name resolution mode. Valid values:
	//
	// - **FQDN**: FQDN-based resolution
	//
	// - **DNS**: DNS-based dynamic resolution
	//
	// - **FQDN_AND_DNS**: FQDN-based and DNS-based dynamic resolution
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The UNIX timestamp, in seconds, for the end of the policy\\"s effective period. The time must be on the hour or half-hour and at least 30 minutes after the start time.
	//
	// > This parameter is not used if `RepeatType` is `Permanent`. It is required for `None`, `Daily`, `Weekly`, or `Monthly` recurrence.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The UNIX timestamp, in seconds, of the last policy hit.
	//
	// example:
	//
	// 1579261141
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of policy hits.
	//
	// example:
	//
	// 100
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The UID of the member account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The UNIX timestamp, in seconds, of when the policy was last modified.
	//
	// example:
	//
	// 1761062400
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The priority of the access control policy, starting from 1. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type. Valid values:
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// - **ANY*	- (all protocol types)
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The enabled status of the access control policy. A policy is enabled by default after it is created. Valid values:
	//
	// - **true**: The policy is enabled.
	//
	// - **false**: The policy is disabled.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of the week or month on which the policy recurs.
	//
	// - If `RepeatType` is set to `Permanent`, `None`, or `Daily`, this parameter is empty. Example: `[]`
	//
	// - If `RepeatType` is set to `Weekly`, this parameter is required. Example: `[0, 6]`
	//
	// > If `RepeatType` is set to `Weekly`, do not specify duplicate values for this parameter.
	//
	// - If `RepeatType` is set to `Monthly`, this parameter is required. Example: `[1, 31]`
	//
	// > If `RepeatType` is set to `Monthly`, do not specify duplicate values for this parameter.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurrence end time. The time is in the `HH:mm` 24-hour format, such as `23:30`.
	//
	// > This parameter is not used if `RepeatType` is `Permanent` or `None`. It is required for `Daily`, `Weekly`, or `Monthly` recurrence.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time. The time is in the `HH:mm` 24-hour format, such as `08:00`.
	//
	// > This parameter is not used if `RepeatType` is `Permanent` or `None`. It is required for `Daily`, `Weekly`, or `Monthly` recurrence.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the policy\\"s effective period. Valid values:
	//
	// - **Permanent*	- (default): The policy is always active.
	//
	// - **None**: The policy applies only once.
	//
	// - **Daily**: The policy recurs daily.
	//
	// - **Weekly**: The policy recurs weekly.
	//
	// - **Monthly**: The policy recurs monthly.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address for the access control policy. The value depends on `SourceType`.
	//
	// - If `SourceType` is `net`, the value is a source CIDR block.
	//
	// - If `SourceType` is `group`, the value is the name of a source address book.
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The CIDR blocks in the source address book.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book. The value is always **ip**, which indicates an address book that contains IP addresses or CIDR blocks.
	//
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address. Valid values:
	//
	// - **net**: a source CIDR block
	//
	// - **group**: a source address book
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The number of rule capacity units that the access control policy consumes. This is calculated as: Number of source addresses × Number of destination addresses × Number of applications × Number of port ranges.
	//
	// example:
	//
	// 10000
	SpreadCnt *int64 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The UNIX timestamp, in seconds, for the start of the policy\\"s effective period. The time must be on the hour or half-hour and at least 30 minutes before the end time.
	//
	// > This parameter is not used if `RepeatType` is `Permanent`. It is required for `None`, `Daily`, `Weekly`, or `Monthly` recurrence.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponseBodyPolicys) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDestPort() *string {
	return s.DestPort
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDestPortGroupPorts() []*string {
	return s.DestPortGroupPorts
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDestPortType() *string {
	return s.DestPortType
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDestination() *string {
	return s.Destination
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDestinationGroupCidrs() []*string {
	return s.DestinationGroupCidrs
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDestinationGroupType() *string {
	return s.DestinationGroupType
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDestinationType() *string {
	return s.DestinationType
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetDomainResolveType() *string {
	return s.DomainResolveType
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetHitLastTime() *int64 {
	return s.HitLastTime
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetHitTimes() *int64 {
	return s.HitTimes
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetOrder() *int32 {
	return s.Order
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetProto() *string {
	return s.Proto
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetRelease() *string {
	return s.Release
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetRepeatType() *string {
	return s.RepeatType
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetSource() *string {
	return s.Source
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetSourceGroupCidrs() []*string {
	return s.SourceGroupCidrs
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetSourceGroupType() *string {
	return s.SourceGroupType
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetSpreadCnt() *int64 {
	return s.SpreadCnt
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationId(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationName(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationName = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationNameList(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetCreateTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortGroupPorts(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDomainResolveType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DomainResolveType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetEndTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetHitLastTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.HitLastTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetHitTimes(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetMemberUid(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetModifyTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ModifyTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRelease(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRepeatDays(v []*int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.RepeatDays = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRepeatEndTime(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.RepeatEndTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRepeatStartTime(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.RepeatStartTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRepeatType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.RepeatType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSpreadCnt(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SpreadCnt = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetStartTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.StartTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) Validate() error {
	return dara.Validate(s)
}
