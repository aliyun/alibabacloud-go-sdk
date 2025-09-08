// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetStorageRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetStorageRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *GetStorageRequest
	GetRoleType() *int32
}

type GetStorageRequest struct {
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
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
	// 127XXXX
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

func (s GetStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageRequest) GoString() string {
	return s.String()
}

func (s *GetStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStorageRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetStorageRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *GetStorageRequest) SetRegionId(v string) *GetStorageRequest {
	s.RegionId = &v
	return s
}

func (s *GetStorageRequest) SetRoleFor(v int64) *GetStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *GetStorageRequest) SetRoleType(v int32) *GetStorageRequest {
	s.RoleType = &v
	return s
}

func (s *GetStorageRequest) Validate() error {
	return dara.Validate(s)
}
