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
	// The information about the access control policies of the virtual private cloud (VPC) firewall.
	Policys []*DescribeVpcFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of access control policies for the virtual private cloud (VPC) firewall.
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
	// The action (settings) that Cloud Firewall performs on the traffic in the access control policy of the virtual private cloud (VPC) firewall. Valid values:
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique identity ID of the access control policy of the virtual private cloud (VPC) firewall.
	//
	// example:
	//
	// 4037fbf7-3e39-4634-92a4-d0155247****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the application with traffic settings in the access control policy of the virtual private cloud (VPC) firewall.
	//
	// example:
	//
	// 10**
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type supported by the access control policy of the virtual private cloud (VPC) firewall. Use ApplicationNameList instead. Valid values:
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The list of application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The time when the policy was created. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1761062400
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access control policy of the virtual private cloud (VPC) firewall.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port of the traffic in the access control policy of the virtual private cloud (VPC) firewall.
	//
	// example:
	//
	// 80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book for the traffic in the access control policy of the virtual private cloud (VPC) firewall.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The details of the destination port address book in the access control policy of the virtual private cloud (VPC) firewall.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The destination port type for the traffic in the access control policy of the virtual private cloud (VPC) firewall. Valid values:
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy of the virtual private cloud (VPC) firewall. Valid values:
	//
	// example:
	//
	// 192.0.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR block information in the destination address book of the access control policy of the virtual private cloud (VPC) firewall.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// example:
	//
	// ip
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The destination address type in the access control policy of the virtual private cloud (VPC) firewall. Valid values:
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy validity period for the access control policy. The value is a UNIX timestamp in seconds. The time must be on the hour or half hour and must be at least 30 minutes later than the start time.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The most recent time of hits. The value is a UNIX timestamp in seconds format.
	//
	// example:
	//
	// 1579261141
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of hits for the access control policy of the virtual private cloud (VPC) firewall.
	//
	// example:
	//
	// 100
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The UID of a member account of the current Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the policy was modified. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1761062400
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The priority of the access control policy of the virtual private cloud (VPC) firewall.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type of the traffic in the access control policy of the virtual private cloud (VPC) firewall. Valid values:
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The enabled status of the access control policy. The policy is enabled by default after creation. Valid values:
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The collection of recurrence days for the policy validity period of the access control policy.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurrence end time of the policy validity period. The value is in the HH:mm format using a 24-hour clock, such as 23:00.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time of the policy validity period. The value is in the HH:mm format using a 24-hour clock, such as 08:00.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type of the policy validity period for the access control policy. Valid values:
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy of the virtual private cloud (VPC) firewall. Valid values:
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The details of the source address book in the access control policy of the virtual private cloud (VPC) firewall.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. The only valid value is **ip**, which indicates an IP address book that contains one or more CIDR blocks.
	//
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The source address type in the access control policy of the virtual private cloud (VPC) firewall. Valid values:
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The number of access control policy specifications consumed, which is the cumulative number of specifications consumed by each policy.
	//
	// example:
	//
	// 10000
	SpreadCnt *int64 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The start time of the policy validity period for the access control policy. The value is a UNIX timestamp in seconds. The time must be on the hour or half hour and must be at least 30 minutes earlier than the end time.
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
