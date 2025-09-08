// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportedLogCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeImportedLogCountRequest
	GetRegionId() *string
	SetRoleFor(v string) *DescribeImportedLogCountRequest
	GetRoleFor() *string
	SetRoleType(v string) *DescribeImportedLogCountRequest
	GetRoleType() *string
}

type DescribeImportedLogCountRequest struct {
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
	RoleFor  *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeImportedLogCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImportedLogCountRequest) GetRoleFor() *string {
	return s.RoleFor
}

func (s *DescribeImportedLogCountRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeImportedLogCountRequest) SetRegionId(v string) *DescribeImportedLogCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImportedLogCountRequest) SetRoleFor(v string) *DescribeImportedLogCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeImportedLogCountRequest) SetRoleType(v string) *DescribeImportedLogCountRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeImportedLogCountRequest) Validate() error {
	return dara.Validate(s)
}
