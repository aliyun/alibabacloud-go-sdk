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
	// The region where the data management center of threat detection and response is located. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets belong to the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets belong to regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
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
