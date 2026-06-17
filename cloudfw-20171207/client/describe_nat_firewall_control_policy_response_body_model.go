// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicys(v []*DescribeNatFirewallControlPolicyResponseBodyPolicys) *DescribeNatFirewallControlPolicyResponseBody
	GetPolicys() []*DescribeNatFirewallControlPolicyResponseBodyPolicys
	SetRequestId(v string) *DescribeNatFirewallControlPolicyResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeNatFirewallControlPolicyResponseBody
	GetTotalCount() *string
}

type DescribeNatFirewallControlPolicyResponseBody struct {
	// The information about the access control policies for the NAT firewall.
	Policys []*DescribeNatFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F283567D-8A52-5BAE-9472-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 28
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNatFirewallControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallControlPolicyResponseBody) GetPolicys() []*DescribeNatFirewallControlPolicyResponseBodyPolicys {
	return s.Policys
}

func (s *DescribeNatFirewallControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallControlPolicyResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeNatFirewallControlPolicyResponseBody) SetPolicys(v []*DescribeNatFirewallControlPolicyResponseBodyPolicys) *DescribeNatFirewallControlPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBody) SetRequestId(v string) *DescribeNatFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBody) SetTotalCount(v string) *DescribeNatFirewallControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBody) Validate() error {
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

type DescribeNatFirewallControlPolicyResponseBodyPolicys struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// - **accept**: Allow
	//
	// - **drop**: Deny
	//
	// - **log**: Monitor
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// example:
	//
	// 01281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application names. Multiple applications are supported.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The time when the policy was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1761062400
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// test-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port for the traffic in the access control policy.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book for the traffic in the access control policy.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The list of ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The destination port type for the traffic in the access control policy. Valid values:
	//
	// - **port**: port
	//
	// - **group**: port address book
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. The value of this parameter varies based on the value of the DestinationType parameter. Valid values:
	//
	// - If **DestinationType*	- is **net**, the value of this parameter is a CIDR block. Example: 192.0.XX.XX/24.
	//
	// - If **DestinationType*	- is **domain**, the value of this parameter is a domain name. Example: aliyuncs.com.
	//
	// - If **DestinationType*	- is **group**, the value of this parameter is the name of an address book. Example: db_group.
	//
	// - If **DestinationType*	- is **location**, the value of this parameter is a region name. For more information, see [AddControlPolicy](https://help.aliyun.com/document_detail/138867.html). Example: ["BJ11", "ZB"].
	//
	// example:
	//
	// x.x.x.x/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The list of CIDR blocks in the destination address book of the access control policy.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// - **ip**: an IP address book that contains one or more IP address CIDR blocks.
	//
	// - **domain**: a domain name address book that contains one or more domain names.
	//
	// example:
	//
	// ip
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
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
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The DNS resolution result.
	//
	// example:
	//
	// 111.0.XX.XX,112.0.XX.XX
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// The timestamp of the DNS resolution. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1579261141
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// The domain name resolution method of the access control policy. By default, this feature is enabled. Valid values:
	//
	// - **0**: FQDN-based
	//
	// - **1**: DNS-based dynamic resolution
	//
	// - **2**: FQDN- and DNS-based dynamic resolution
	//
	// example:
	//
	// 0
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy validity period. The value is a UNIX timestamp. Unit: seconds. The time must be on the hour or half hour, and at least 30 minutes later than the start time.
	//
	// > If RepeatType is set to Permanent, this parameter is empty. If RepeatType is set to None, Daily, Weekly, or Monthly, you must set this parameter.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The timestamp of the last hit. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1579261141
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of hits for the access control policy.
	//
	// example:
	//
	// 100
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The time when the policy was last modified. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1761062400
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the NAT Gateway.
	//
	// example:
	//
	// ngw-xxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy.
	//
	// The priority starts from 1. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type of the traffic in the access control policy. Valid values:
	//
	// - **ANY**
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// - **true**: enabled
	//
	// - **false**: disabled
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of the week or month for the policy to repeat.
	//
	// - If RepeatType is set to `Permanent`, `None`, or `Daily`, this parameter is an empty set.
	//
	//   Example: []
	//
	// - If RepeatType is set to Weekly, this parameter cannot be empty.
	//
	//   Example: [0, 6]
	//
	// > If RepeatType is set to Weekly, the days in RepeatDays cannot be repeated.
	//
	// - If RepeatType is set to `Monthly`, this parameter cannot be empty.
	//
	//   Example: [1, 31]
	//
	// > If RepeatType is set to Monthly, the days in RepeatDays cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The end time of the recurrence. The time is in the HH:mm format, based on a 24-hour clock. Example: 23:00.
	//
	// > If RepeatType is set to Permanent or None, this parameter is empty. If RepeatType is set to Daily, Weekly, or Monthly, you must set this parameter.
	//
	// > The time is in the HH:mm format, based on a 24-hour clock. Examples: 08:00 and 23:30.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The start time of the recurrence. The time is in the HH:mm format, based on a 24-hour clock. Example: 08:00.
	//
	// > If RepeatType is set to Permanent or None, this parameter is empty. If RepeatType is set to Daily, Weekly, or Monthly, you must set this parameter.
	//
	// > The time is in the HH:mm format, based on a 24-hour clock. Examples: 08:00 and 23:30.
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
	// The source address in the access control policy. Valid values:
	//
	// - If **SourceType*	- is `net`, the value of this parameter is a CIDR block. Example: 192.0.XX.XX/24.
	//
	// - If **SourceType*	- is `group`, the value of this parameter is the name of a source address book. Example: db_group.
	//
	// - If **SourceType*	- is `location`, the value of this parameter is a region. For more information, see [AddControlPolicy](https://help.aliyun.com/document_detail/138867.html). Example: ["BJ11", "ZB"].
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of CIDR blocks in the source address book of the access control policy.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. The value is fixed at **ip**. This indicates an IP address book that contains one or more IP address CIDR blocks.
	//
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The source address type in the access control policy. Valid values:
	//
	// - **net**: CIDR block
	//
	// - **group**: source address book
	//
	// - **location**: source region
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The number of policy specifications that are occupied. This is the cumulative value of the number of specifications occupied by each policy.
	//
	// The number of specifications occupied by a single policy = Number of source CIDR blocks × Number of destination addresses (IP address CIDR blocks, regions, or domain names) × Number of applications × Number of port ranges.
	//
	// example:
	//
	// 10,000
	SpreadCnt *string `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The start time of the policy validity period. The value is a UNIX timestamp. Unit: seconds. The time must be on the hour or half hour, and at least 30 minutes earlier than the end time.
	//
	// > If RepeatType is set to Permanent, this parameter is empty. If RepeatType is set to None, Daily, Weekly, or Monthly, you must set this parameter.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeNatFirewallControlPolicyResponseBodyPolicys) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDescription() *string {
	return s.Description
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDestPort() *string {
	return s.DestPort
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDestPortGroupPorts() []*string {
	return s.DestPortGroupPorts
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDestPortType() *string {
	return s.DestPortType
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDestination() *string {
	return s.Destination
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDestinationGroupCidrs() []*string {
	return s.DestinationGroupCidrs
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDestinationGroupType() *string {
	return s.DestinationGroupType
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDestinationType() *string {
	return s.DestinationType
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDnsResult() *string {
	return s.DnsResult
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDnsResultTime() *int64 {
	return s.DnsResultTime
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetDomainResolveType() *int32 {
	return s.DomainResolveType
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetHitLastTime() *int64 {
	return s.HitLastTime
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetHitTimes() *int64 {
	return s.HitTimes
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetOrder() *int32 {
	return s.Order
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetProto() *string {
	return s.Proto
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetRelease() *string {
	return s.Release
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetRepeatType() *string {
	return s.RepeatType
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetSource() *string {
	return s.Source
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetSourceGroupCidrs() []*string {
	return s.SourceGroupCidrs
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetSourceGroupType() *string {
	return s.SourceGroupType
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetSpreadCnt() *string {
	return s.SpreadCnt
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetApplicationNameList(v []*string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetCreateTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.CreateTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestPortGroupPorts(v []*string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestPortType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestPortType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDnsResult(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DnsResult = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDnsResultTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DnsResultTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDomainResolveType(v int32) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DomainResolveType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetEndTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.EndTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetHitLastTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.HitLastTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetHitTimes(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetModifyTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.ModifyTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetNatGatewayId(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRelease(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRepeatDays(v []*int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.RepeatDays = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRepeatEndTime(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.RepeatEndTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRepeatStartTime(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.RepeatStartTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRepeatType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.RepeatType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSpreadCnt(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.SpreadCnt = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetStartTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.StartTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) Validate() error {
	return dara.Validate(s)
}
