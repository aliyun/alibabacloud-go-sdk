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
	// 137820528780****
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
