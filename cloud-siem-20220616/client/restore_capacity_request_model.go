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
