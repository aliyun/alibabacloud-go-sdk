// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProdCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeProdCountRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeProdCountRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeProdCountRequest
	GetRoleType() *int32
}

type DescribeProdCountRequest struct {
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeProdCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProdCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeProdCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeProdCountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeProdCountRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeProdCountRequest) SetRegionId(v string) *DescribeProdCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProdCountRequest) SetRoleFor(v int64) *DescribeProdCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeProdCountRequest) SetRoleType(v int32) *DescribeProdCountRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeProdCountRequest) Validate() error {
	return dara.Validate(s)
}
