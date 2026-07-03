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
	// The region where the data management center of Threat Analysis is deployed. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or the China (Hong Kong) region.
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UID of the member. An administrator can use this parameter to switch to the member\\"s view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of view.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that belong to the enterprise.
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
