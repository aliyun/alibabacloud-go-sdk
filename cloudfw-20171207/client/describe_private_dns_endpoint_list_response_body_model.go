// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsEndpointListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceList(v []*DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) *DescribePrivateDnsEndpointListResponseBody
	GetAccessInstanceList() []*DescribePrivateDnsEndpointListResponseBodyAccessInstanceList
	SetPageNo(v int32) *DescribePrivateDnsEndpointListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribePrivateDnsEndpointListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePrivateDnsEndpointListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePrivateDnsEndpointListResponseBody
	GetTotalCount() *int64
}

type DescribePrivateDnsEndpointListResponseBody struct {
	AccessInstanceList []*DescribePrivateDnsEndpointListResponseBodyAccessInstanceList `json:"AccessInstanceList,omitempty" xml:"AccessInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePrivateDnsEndpointListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsEndpointListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsEndpointListResponseBody) GetAccessInstanceList() []*DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	return s.AccessInstanceList
}

func (s *DescribePrivateDnsEndpointListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribePrivateDnsEndpointListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePrivateDnsEndpointListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrivateDnsEndpointListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePrivateDnsEndpointListResponseBody) SetAccessInstanceList(v []*DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) *DescribePrivateDnsEndpointListResponseBody {
	s.AccessInstanceList = v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBody) SetPageNo(v int32) *DescribePrivateDnsEndpointListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBody) SetPageSize(v int32) *DescribePrivateDnsEndpointListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBody) SetRequestId(v string) *DescribePrivateDnsEndpointListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBody) SetTotalCount(v int64) *DescribePrivateDnsEndpointListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePrivateDnsEndpointListResponseBodyAccessInstanceList struct {
	// example:
	//
	// 123
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// example:
	//
	// test
	AccessInstanceName *string `json:"AccessInstanceName,omitempty" xml:"AccessInstanceName,omitempty"`
	// example:
	//
	// 1379490574415****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 3
	DomainNameCount *int64    `json:"DomainNameCount,omitempty" xml:"DomainNameCount,omitempty"`
	FirewallType    []*string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty" type:"Repeated"`
	// example:
	//
	// 1715075765
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// tcp
	IpProtocol *int32 `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// 1844802493****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 1.1.1.1
	PrimaryDns *string `json:"PrimaryDns,omitempty" xml:"PrimaryDns,omitempty"`
	// example:
	//
	// Custom
	PrivateDnsType *string `json:"PrivateDnsType,omitempty" xml:"PrivateDnsType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 1.1.1.2
	StandbyDns *string `json:"StandbyDns,omitempty" xml:"StandbyDns,omitempty"`
	// example:
	//
	// normal
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// vpc-2zelphbaourpun****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetAccessInstanceName() *string {
	return s.AccessInstanceName
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetDomainNameCount() *int64 {
	return s.DomainNameCount
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetFirewallType() []*string {
	return s.FirewallType
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetIpProtocol() *int32 {
	return s.IpProtocol
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetPort() *int32 {
	return s.Port
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetPrimaryDns() *string {
	return s.PrimaryDns
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetPrivateDnsType() *string {
	return s.PrivateDnsType
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetStandbyDns() *string {
	return s.StandbyDns
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetAccessInstanceId(v string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.AccessInstanceId = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetAccessInstanceName(v string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.AccessInstanceName = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetAliUid(v int64) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.AliUid = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetDomainNameCount(v int64) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.DomainNameCount = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetFirewallType(v []*string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.FirewallType = v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetGmtCreate(v int64) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetIpProtocol(v int32) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.IpProtocol = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetMemberUid(v int64) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.MemberUid = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetPort(v int32) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.Port = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetPrimaryDns(v string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.PrimaryDns = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetPrivateDnsType(v string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.PrivateDnsType = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetRegionNo(v string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.RegionNo = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetStandbyDns(v string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.StandbyDns = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetStatus(v int32) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.Status = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetTaskId(v string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.TaskId = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) SetVpcId(v string) *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList {
	s.VpcId = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponseBodyAccessInstanceList) Validate() error {
	return dara.Validate(s)
}
