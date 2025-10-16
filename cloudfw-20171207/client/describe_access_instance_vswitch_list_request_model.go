// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceVSwitchListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberUid(v int64) *DescribeAccessInstanceVSwitchListRequest
	GetMemberUid() *int64
	SetPageNo(v int32) *DescribeAccessInstanceVSwitchListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAccessInstanceVSwitchListRequest
	GetPageSize() *int32
	SetRegionNo(v string) *DescribeAccessInstanceVSwitchListRequest
	GetRegionNo() *string
	SetVSwitchId(v string) *DescribeAccessInstanceVSwitchListRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeAccessInstanceVSwitchListRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeAccessInstanceVSwitchListRequest
	GetZoneId() *string
}

type DescribeAccessInstanceVSwitchListRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vsw-qzeaol304m***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-uf6b5lyul0x******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAccessInstanceVSwitchListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVSwitchListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVSwitchListRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeAccessInstanceVSwitchListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAccessInstanceVSwitchListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessInstanceVSwitchListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAccessInstanceVSwitchListRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAccessInstanceVSwitchListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAccessInstanceVSwitchListRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAccessInstanceVSwitchListRequest) SetMemberUid(v int64) *DescribeAccessInstanceVSwitchListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListRequest) SetPageNo(v int32) *DescribeAccessInstanceVSwitchListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListRequest) SetPageSize(v int32) *DescribeAccessInstanceVSwitchListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListRequest) SetRegionNo(v string) *DescribeAccessInstanceVSwitchListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListRequest) SetVSwitchId(v string) *DescribeAccessInstanceVSwitchListRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListRequest) SetVpcId(v string) *DescribeAccessInstanceVSwitchListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListRequest) SetZoneId(v string) *DescribeAccessInstanceVSwitchListRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListRequest) Validate() error {
	return dara.Validate(s)
}
