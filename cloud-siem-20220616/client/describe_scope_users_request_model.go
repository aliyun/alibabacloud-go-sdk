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
	// The region of the Data Management center. Select the region based on where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. This parameter is used when an administrator switches to the view of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
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
