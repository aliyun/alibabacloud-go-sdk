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
	// The information about the access control policies.
	Policys []*DescribeNatFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F183567D-8A52-5BAE-9472-F1C427378C28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The time when the access control policy was created.
	//
	// example:
	//
	// 1761062400
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	// The ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
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
	// The destination address in the access control policy. The value of this parameter varies based on the value of DestinationType. Valid values:
	//
	// 	- If the value of **DestinationType*	- is **net**, the value of this parameter is a CIDR block. Example: 192.0.XX.XX/24.
	//
	// 	- If the value of **DestinationType*	- is **domain**, the value of this parameter is a domain name. Example: aliyuncs.com.
	//
	// 	- If the value of **DestinationType*	- is **group**, the value of this parameter is the name of an address book. Example: db_group.
	//
	// 	- If the value of **DestinationType*	- is **location**, the value of this parameter is a location. For more information about location codes, see [AddControlPolicy](https://help.aliyun.com/document_detail/138867.html). Example: ["BJ11", "ZB"].
	//
	// example:
	//
	// x.x.x.x/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR blocks in the destination address book.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// 	- **ip**: an address book that includes one or more CIDR blocks
	//
	// 	- **domain**: an address book that includes one or more domain names
	//
	// example:
	//
	// ip
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
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
	// The DNS resolution result.
	//
	// example:
	//
	// 111.0.XX.XX,112.0.XX.XX
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// The time when the Domain Name System (DNS) resolution was performed. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1579261141
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// The domain name resolution method of the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// 	- **0**: fully qualified domain name (FQDN)-based resolution
	//
	// 	- **1**: DNS-based dynamic resolution
	//
	// 	- **2**: FQDN and DNS-based dynamic resolution
	//
	// example:
	//
	// 1
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The end time must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If RepeatType is set to Permanent, this parameter is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the access control policy was last hit. The value is a UNIX timestamp. Unit: seconds.
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
	// The time when the access control policy was modified.
	//
	// example:
	//
	// 1761062400
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-xxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A smaller priority value indicates a higher priority.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// 	- **ANY**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **ICMP**
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// 	- If RepeatType is set to `Permanent`, `None`, or `Daily`, the value of this parameter is an empty array. Example: [].
	//
	// 	- If RepeatType is set to Weekly, this parameter must be specified. Example: [0, 6].
	//
	// >  If RepeatType is set to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// 	- If RepeatType is set to `Monthly`, this parameter must be specified. Example: [1, 31].
	//
	// >  If RepeatType is set to Monthly, the fields in the value of this parameter cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The end time must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If RepeatType is set to Permanent or None, this parameter is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The start time must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If RepeatType is set to Permanent or None, this parameter is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
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
	// The source address in the access control policy. Valid values:
	//
	// 	- If the value of **SourceType*	- is `net`, the value of this parameter is a CIDR block. Example: 192.0.XX.XX/24.
	//
	// 	- If the value of **SourceType*	- is `group`, the value of this parameter is the name of an address book. Example: db_group.
	//
	// 	- If the value of **SourceType*	- is `location`, the value of this parameter is a location. For more information about location codes, see [AddControlPolicy](https://help.aliyun.com/document_detail/138867.html). Example: ["BJ11", "ZB"].
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The CIDR blocks in the source address book.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. The value is fixed as **ip**. The value indicates an address book that includes one or more CIDR blocks.
	//
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
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
	// The total quota consumed by the returned access control policies, which is the sum of the quota consumed by each policy. The quota that is consumed by an access control policy is calculated by using the following formula: Quota that is consumed by an access control policy = Number of source addresses (number of CIDR blocks or regions) × Number of destination addresses (number of CIDR blocks, regions, or domain names) × Number of port ranges × Number of applications.
	//
	// example:
	//
	// 10,000
	SpreadCnt *string `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The start time must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If RepeatType is set to Permanent, this parameter is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
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
