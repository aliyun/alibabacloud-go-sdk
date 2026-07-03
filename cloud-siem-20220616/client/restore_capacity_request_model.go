// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *RestoreCapacityRequest
	GetRegionId() *string
	SetRoleFor(v int64) *RestoreCapacityRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *RestoreCapacityRequest
	GetRoleType() *int32
}

type RestoreCapacityRequest struct {
	// The region where the Data Management center is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the perspective of a specific member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of view. Valid values:
	//
	// - 0: The view for the current Alibaba Cloud account.
	//
	// - 1: The view for all accounts that belong to the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s RestoreCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreCapacityRequest) GoString() string {
	return s.String()
}

func (s *RestoreCapacityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestoreCapacityRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *RestoreCapacityRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *RestoreCapacityRequest) SetRegionId(v string) *RestoreCapacityRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreCapacityRequest) SetRoleFor(v int64) *RestoreCapacityRequest {
	s.RoleFor = &v
	return s
}

func (s *RestoreCapacityRequest) SetRoleType(v int32) *RestoreCapacityRequest {
	s.RoleType = &v
	return s
}

func (s *RestoreCapacityRequest) Validate() error {
	return dara.Validate(s)
}
