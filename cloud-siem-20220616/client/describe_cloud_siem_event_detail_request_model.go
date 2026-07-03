// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidentUuid(v string) *DescribeCloudSiemEventDetailRequest
	GetIncidentUuid() *string
	SetRegionId(v string) *DescribeCloudSiemEventDetailRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeCloudSiemEventDetailRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeCloudSiemEventDetailRequest
	GetRoleType() *int32
}

type DescribeCloudSiemEventDetailRequest struct {
	// The UUID of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region of the data management center for Threat Analysis. Select the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: assets in the Chinese mainland and China (Hong Kong)
	//
	// - ap-southeast-1: assets in regions outside the Chinese mainland
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member account. An administrator can use this parameter to query data from the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCloudSiemEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeCloudSiemEventDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudSiemEventDetailRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeCloudSiemEventDetailRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeCloudSiemEventDetailRequest) SetIncidentUuid(v string) *DescribeCloudSiemEventDetailRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailRequest) SetRegionId(v string) *DescribeCloudSiemEventDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailRequest) SetRoleFor(v int64) *DescribeCloudSiemEventDetailRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCloudSiemEventDetailRequest) SetRoleType(v int32) *DescribeCloudSiemEventDetailRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCloudSiemEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
