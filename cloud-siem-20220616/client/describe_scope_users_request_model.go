// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScopeUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeScopeUsersRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeScopeUsersRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeScopeUsersRequest
	GetRoleType() *int32
}

type DescribeScopeUsersRequest struct {
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
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

func (s DescribeScopeUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScopeUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScopeUsersRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeScopeUsersRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeScopeUsersRequest) SetRegionId(v string) *DescribeScopeUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScopeUsersRequest) SetRoleFor(v int64) *DescribeScopeUsersRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeScopeUsersRequest) SetRoleType(v int32) *DescribeScopeUsersRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeScopeUsersRequest) Validate() error {
	return dara.Validate(s)
}
