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
	// The region where the Data Management hub for threat analysis is located. Select a region for the management hub based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are in the Chinese mainland or the China (Hong Kong) region.
	//
	// - ap-southeast-1: Select this value if your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. This parameter is used by an administrator to switch to the perspective of a member.
	//
	// example:
	//
	// 127XXXX
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
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
