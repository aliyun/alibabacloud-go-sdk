// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAllProdsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListAllProdsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListAllProdsRequest
	GetRoleType() *int32
}

type ListAllProdsRequest struct {
	// The region of the Data Management hub for Threat Analysis. Select the region that corresponds to the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can specify this parameter to switch to the member\\"s view.
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

func (s ListAllProdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllProdsRequest) GoString() string {
	return s.String()
}

func (s *ListAllProdsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllProdsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListAllProdsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListAllProdsRequest) SetRegionId(v string) *ListAllProdsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllProdsRequest) SetRoleFor(v int64) *ListAllProdsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAllProdsRequest) SetRoleType(v int32) *ListAllProdsRequest {
	s.RoleType = &v
	return s
}

func (s *ListAllProdsRequest) Validate() error {
	return dara.Validate(s)
}
