// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsFirewallPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v string) *DescribeDnsFirewallPolicyResponseBody
	GetPageNo() *string
	SetPageSize(v string) *DescribeDnsFirewallPolicyResponseBody
	GetPageSize() *string
	SetPolicys(v []*DescribeDnsFirewallPolicyResponseBodyPolicys) *DescribeDnsFirewallPolicyResponseBody
	GetPolicys() []*DescribeDnsFirewallPolicyResponseBodyPolicys
	SetRequestId(v string) *DescribeDnsFirewallPolicyResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeDnsFirewallPolicyResponseBody
	GetTotalCount() *string
}

type DescribeDnsFirewallPolicyResponseBody struct {
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *string                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Policys  []*DescribeDnsFirewallPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// example:
	//
	// 0A4ACDE9-9F9F-56C1-B3B7-60971BA1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDnsFirewallPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsFirewallPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsFirewallPolicyResponseBody) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeDnsFirewallPolicyResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDnsFirewallPolicyResponseBody) GetPolicys() []*DescribeDnsFirewallPolicyResponseBodyPolicys {
	return s.Policys
}

func (s *DescribeDnsFirewallPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsFirewallPolicyResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeDnsFirewallPolicyResponseBody) SetPageNo(v string) *DescribeDnsFirewallPolicyResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBody) SetPageSize(v string) *DescribeDnsFirewallPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBody) SetPolicys(v []*DescribeDnsFirewallPolicyResponseBodyPolicys) *DescribeDnsFirewallPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBody) SetRequestId(v string) *DescribeDnsFirewallPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBody) SetTotalCount(v string) *DescribeDnsFirewallPolicyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBody) Validate() error {
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

type DescribeDnsFirewallPolicyResponseBodyPolicys struct {
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// example:
	//
	// 01281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// x.x.x.x/32
	Destination      *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationAddrs []*string `json:"DestinationAddrs,omitempty" xml:"DestinationAddrs,omitempty" type:"Repeated"`
	// example:
	//
	// ip
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// example:
	//
	// net
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 1579261141
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// example:
	//
	// 100
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// example:
	//
	// 6
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// 110
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// example:
	//
	// 192.0.XX.XX/24
	Source      *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceAddrs []*string `json:"SourceAddrs,omitempty" xml:"SourceAddrs,omitempty" type:"Repeated"`
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeDnsFirewallPolicyResponseBodyPolicys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsFirewallPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetDescription() *string {
	return s.Description
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetDestination() *string {
	return s.Destination
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetDestinationAddrs() []*string {
	return s.DestinationAddrs
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetDestinationGroupType() *string {
	return s.DestinationGroupType
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetDestinationType() *string {
	return s.DestinationType
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetDirection() *string {
	return s.Direction
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetHitLastTime() *int64 {
	return s.HitLastTime
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetHitTimes() *int64 {
	return s.HitTimes
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetRelease() *string {
	return s.Release
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetSource() *string {
	return s.Source
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetSourceAddrs() []*string {
	return s.SourceAddrs
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetSourceGroupType() *string {
	return s.SourceGroupType
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetDescription(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetDestination(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetDestinationAddrs(v []*string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.DestinationAddrs = v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetDirection(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.Direction = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetHitLastTime(v int64) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.HitLastTime = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetHitTimes(v int64) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetIpVersion(v int32) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.IpVersion = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetPriority(v int32) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.Priority = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetRelease(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetSource(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetSourceAddrs(v []*string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.SourceAddrs = v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeDnsFirewallPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeDnsFirewallPolicyResponseBodyPolicys) Validate() error {
	return dara.Validate(s)
}
