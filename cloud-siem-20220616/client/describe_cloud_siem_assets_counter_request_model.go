// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemAssetsCounterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidentUuid(v string) *DescribeCloudSiemAssetsCounterRequest
	GetIncidentUuid() *string
	SetRegionId(v string) *DescribeCloudSiemAssetsCounterRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeCloudSiemAssetsCounterRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeCloudSiemAssetsCounterRequest
	GetRoleType() *int32
}

type DescribeCloudSiemAssetsCounterRequest struct {
	// The UUID of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeCloudSiemAssetsCounterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudSiemAssetsCounterRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeCloudSiemAssetsCounterRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetIncidentUuid(v string) *DescribeCloudSiemAssetsCounterRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetRegionId(v string) *DescribeCloudSiemAssetsCounterRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetRoleFor(v int64) *DescribeCloudSiemAssetsCounterRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetRoleType(v int32) *DescribeCloudSiemAssetsCounterRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterRequest) Validate() error {
	return dara.Validate(s)
}
