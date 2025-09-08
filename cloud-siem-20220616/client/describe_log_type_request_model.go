// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeLogTypeRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeLogTypeRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeLogTypeRequest
	GetRoleType() *int32
}

type DescribeLogTypeRequest struct {
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

func (s DescribeLogTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogTypeRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeLogTypeRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeLogTypeRequest) SetRegionId(v string) *DescribeLogTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogTypeRequest) SetRoleFor(v int64) *DescribeLogTypeRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeLogTypeRequest) SetRoleType(v int32) *DescribeLogTypeRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeLogTypeRequest) Validate() error {
	return dara.Validate(s)
}
