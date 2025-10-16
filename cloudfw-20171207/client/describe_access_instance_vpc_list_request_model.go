// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceVpcListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberUid(v int64) *DescribeAccessInstanceVpcListRequest
	GetMemberUid() *int64
	SetPageNo(v int32) *DescribeAccessInstanceVpcListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAccessInstanceVpcListRequest
	GetPageSize() *int32
	SetRegionNo(v string) *DescribeAccessInstanceVpcListRequest
	GetRegionNo() *string
	SetVpcId(v string) *DescribeAccessInstanceVpcListRequest
	GetVpcId() *string
}

type DescribeAccessInstanceVpcListRequest struct {
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
	// vpc-j6cvhdscntzuvr0x****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAccessInstanceVpcListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVpcListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVpcListRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeAccessInstanceVpcListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAccessInstanceVpcListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessInstanceVpcListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAccessInstanceVpcListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAccessInstanceVpcListRequest) SetMemberUid(v int64) *DescribeAccessInstanceVpcListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAccessInstanceVpcListRequest) SetPageNo(v int32) *DescribeAccessInstanceVpcListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAccessInstanceVpcListRequest) SetPageSize(v int32) *DescribeAccessInstanceVpcListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessInstanceVpcListRequest) SetRegionNo(v string) *DescribeAccessInstanceVpcListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAccessInstanceVpcListRequest) SetVpcId(v string) *DescribeAccessInstanceVpcListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeAccessInstanceVpcListRequest) Validate() error {
	return dara.Validate(s)
}
