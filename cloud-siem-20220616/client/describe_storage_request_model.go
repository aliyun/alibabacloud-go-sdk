// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeStorageRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeStorageRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeStorageRequest
	GetRoleType() *int32
}

type DescribeStorageRequest struct {
	// The region of the Data Management center for threat analysis. Select a region for the center based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are in the Chinese mainland or the China (Hong Kong) region.
	//
	// - ap-southeast-1: Select this value if your assets are in regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose view you want to use. This parameter is available only for administrators.
	//
	// example:
	//
	// 137820528780****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of view. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeStorageRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeStorageRequest) SetRegionId(v string) *DescribeStorageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageRequest) SetRoleFor(v int64) *DescribeStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeStorageRequest) SetRoleType(v int32) *DescribeStorageRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeStorageRequest) Validate() error {
	return dara.Validate(s)
}
