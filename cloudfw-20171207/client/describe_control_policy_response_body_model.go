// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v string) *DescribeControlPolicyResponseBody
	GetPageNo() *string
	SetPageSize(v string) *DescribeControlPolicyResponseBody
	GetPageSize() *string
	SetPolicys(v []*DescribeControlPolicyResponseBodyPolicys) *DescribeControlPolicyResponseBody
	GetPolicys() []*DescribeControlPolicyResponseBodyPolicys
	SetRequestId(v string) *DescribeControlPolicyResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeControlPolicyResponseBody
	GetTotalCount() *string
}

type DescribeControlPolicyResponseBody struct {
	// The page number of the current page displayed in a paging query.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of entries per page displayed in a paging query.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the access control policies.
	Policys []*DescribeControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the access control policies.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponseBody) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeControlPolicyResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeControlPolicyResponseBody) GetPolicys() []*DescribeControlPolicyResponseBodyPolicys {
	return s.Policys
}

func (s *DescribeControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeControlPolicyResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeControlPolicyResponseBody) SetPageNo(v string) *DescribeControlPolicyResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetPageSize(v string) *DescribeControlPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetPolicys(v []*DescribeControlPolicyResponseBodyPolicys) *DescribeControlPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetRequestId(v string) *DescribeControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetTotalCount(v string) *DescribeControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) Validate() error {
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

type DescribeControlPolicyResponseBodyPolicys struct {
	// The action that Cloud Firewall performs on the traffic in the access control policy. Valid values:
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application ID of the traffic in the access control policy.
	//
	// example:
	//
	// 10***
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type supported by the access control policy. We recommend that you use ApplicationNameList. Valid values:
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The list of application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The time when the policy was created. The value is a UNIX timestamp in seconds, which is the number of seconds that have elapsed since January 1, 1970 (UTC).
	//
	// example:
	//
	// 1761062400
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// Allow access to office network segment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port of the traffic in the access control policy.
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
	// The list of ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. The value varies depending on the DestinationType (destination type). Valid values:
	//
	// example:
	//
	// 192.0.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The list of CIDR blocks in the destination address book of the access control policy.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// example:
	//
	// ip
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction of the access control policy. Valid values:
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Deprecated
	//
	// The DNS resolution result.
	//
	// example:
	//
	// 192.0.XX.XX,192.0.XX.XX
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// The timestamp of the DNS resolution. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1579261141
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// example:
	//
	// FQDN
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy validity period for the access control policy. The value is a UNIX timestamp in seconds. The value must be on the hour or half hour and must be at least 30 minutes later than the start time.
	//
	// example:
	//
	// 1694764800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The most recent time of hits. The value is in the format of a UNIX timestamp in seconds.
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
	// The supported IP address version. Valid values:
	//
	// example:
	//
	// 6
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The time when the policy was last modified. The value is a UNIX timestamp in seconds, which is the number of seconds that have elapsed since January 1, 1970 (UTC).
	//
	// example:
	//
	// 1761062400
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The priority of the access control policy.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The security protocol type of the traffic in the access control policy. Valid values:
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
	// The recurrence end time of the policy validity period for the access control policy. Example: 23:30. The value must be on the hour or half hour and must be at least 30 minutes later than the recurrence start time.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time of the policy validity period for the access control policy. Example: 08:00. The value must be on the hour or half hour and must be at least 30 minutes earlier than the recurrence end time.
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
	// The source address in the access control policy. Valid values:
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of CIDR blocks in the source address book of the access control policy.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. Valid values:
	//
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The number of quota units consumed by the access control policy, which is the cumulative number of quota units consumed by each policy.
	//
	// example:
	//
	// 10000
	SpreadCnt *int32 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The start time of the policy validity period for the access control policy. The value is a UNIX timestamp in seconds. The value must be on the hour or half hour and must be at least 30 minutes earlier than the end time.
	//
	// example:
	//
	// 1694761200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeControlPolicyResponseBodyPolicys) String() string {
	return dara.Prettify(s)
}

func (s DescribeControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDescription() *string {
	return s.Description
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDestPort() *string {
	return s.DestPort
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDestPortGroupPorts() []*string {
	return s.DestPortGroupPorts
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDestPortType() *string {
	return s.DestPortType
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDestination() *string {
	return s.Destination
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDestinationGroupCidrs() []*string {
	return s.DestinationGroupCidrs
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDestinationGroupType() *string {
	return s.DestinationGroupType
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDestinationType() *string {
	return s.DestinationType
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDirection() *string {
	return s.Direction
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDnsResult() *string {
	return s.DnsResult
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDnsResultTime() *int64 {
	return s.DnsResultTime
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetDomainResolveType() *string {
	return s.DomainResolveType
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetHitLastTime() *int64 {
	return s.HitLastTime
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetHitTimes() *int64 {
	return s.HitTimes
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetOrder() *int32 {
	return s.Order
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetProto() *string {
	return s.Proto
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetRelease() *string {
	return s.Release
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetRepeatType() *string {
	return s.RepeatType
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetSource() *string {
	return s.Source
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetSourceGroupCidrs() []*string {
	return s.SourceGroupCidrs
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetSourceGroupType() *string {
	return s.SourceGroupType
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetSpreadCnt() *int32 {
	return s.SpreadCnt
}

func (s *DescribeControlPolicyResponseBodyPolicys) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationId(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationId = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationName(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationName = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationNameList(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetCreateTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.CreateTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortGroupPorts(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDirection(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Direction = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDnsResult(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DnsResult = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDnsResultTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.DnsResultTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDomainResolveType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DomainResolveType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetEndTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.EndTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetHitLastTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.HitLastTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetHitTimes(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetIpVersion(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.IpVersion = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetModifyTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.ModifyTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRelease(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRepeatDays(v []*int64) *DescribeControlPolicyResponseBodyPolicys {
	s.RepeatDays = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRepeatEndTime(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.RepeatEndTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRepeatStartTime(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.RepeatStartTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRepeatType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.RepeatType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSpreadCnt(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.SpreadCnt = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetStartTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.StartTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) Validate() error {
	return dara.Validate(s)
}
