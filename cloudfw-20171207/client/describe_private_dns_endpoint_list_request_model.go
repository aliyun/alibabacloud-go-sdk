// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsEndpointListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *DescribePrivateDnsEndpointListRequest
	GetAccessInstanceId() *string
	SetAccessInstanceName(v string) *DescribePrivateDnsEndpointListRequest
	GetAccessInstanceName() *string
	SetFirewallType(v string) *DescribePrivateDnsEndpointListRequest
	GetFirewallType() *string
	SetMemberUid(v int64) *DescribePrivateDnsEndpointListRequest
	GetMemberUid() *int64
	SetPageNo(v int32) *DescribePrivateDnsEndpointListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribePrivateDnsEndpointListRequest
	GetPageSize() *int32
	SetRegionNo(v string) *DescribePrivateDnsEndpointListRequest
	GetRegionNo() *string
	SetStatus(v string) *DescribePrivateDnsEndpointListRequest
	GetStatus() *string
	SetVpcId(v string) *DescribePrivateDnsEndpointListRequest
	GetVpcId() *string
}

type DescribePrivateDnsEndpointListRequest struct {
	// example:
	//
	// pd-12345
	AccessInstanceId   *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	AccessInstanceName *string `json:"AccessInstanceName,omitempty" xml:"AccessInstanceName,omitempty"`
	// example:
	//
	// vpc
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// example:
	//
	// 135809047715****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
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
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribePrivateDnsEndpointListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsEndpointListRequest) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsEndpointListRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DescribePrivateDnsEndpointListRequest) GetAccessInstanceName() *string {
	return s.AccessInstanceName
}

func (s *DescribePrivateDnsEndpointListRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribePrivateDnsEndpointListRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribePrivateDnsEndpointListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribePrivateDnsEndpointListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePrivateDnsEndpointListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePrivateDnsEndpointListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribePrivateDnsEndpointListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribePrivateDnsEndpointListRequest) SetAccessInstanceId(v string) *DescribePrivateDnsEndpointListRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) SetAccessInstanceName(v string) *DescribePrivateDnsEndpointListRequest {
	s.AccessInstanceName = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) SetFirewallType(v string) *DescribePrivateDnsEndpointListRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) SetMemberUid(v int64) *DescribePrivateDnsEndpointListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) SetPageNo(v int32) *DescribePrivateDnsEndpointListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) SetPageSize(v int32) *DescribePrivateDnsEndpointListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) SetRegionNo(v string) *DescribePrivateDnsEndpointListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) SetStatus(v string) *DescribePrivateDnsEndpointListRequest {
	s.Status = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) SetVpcId(v string) *DescribePrivateDnsEndpointListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribePrivateDnsEndpointListRequest) Validate() error {
	return dara.Validate(s)
}
