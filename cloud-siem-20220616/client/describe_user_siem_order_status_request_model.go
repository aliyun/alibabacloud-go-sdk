// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSiemOrderStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeUserSiemOrderStatusRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeUserSiemOrderStatusRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeUserSiemOrderStatusRequest
	GetRoleType() *int32
}

type DescribeUserSiemOrderStatusRequest struct {
	// The region of the data management center for threat detection and response. Select the data management center based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: assets in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: assets in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to another member\\"s perspective.
	//
	// example:
	//
	// 1234567890***
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the current Alibaba Cloud account view.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 0
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeUserSiemOrderStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSiemOrderStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserSiemOrderStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserSiemOrderStatusRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeUserSiemOrderStatusRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeUserSiemOrderStatusRequest) SetRegionId(v string) *DescribeUserSiemOrderStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserSiemOrderStatusRequest) SetRoleFor(v int64) *DescribeUserSiemOrderStatusRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeUserSiemOrderStatusRequest) SetRoleType(v int32) *DescribeUserSiemOrderStatusRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeUserSiemOrderStatusRequest) Validate() error {
	return dara.Validate(s)
}
