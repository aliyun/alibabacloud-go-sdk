// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetCapacityRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetCapacityRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *GetCapacityRequest
	GetRoleType() *int32
}

type GetCapacityRequest struct {
	// Region where the Data Management Center for threat analysis is located. Choose a region based on where your assets are located.
	//
	// - cn-hangzhou: Select this if your assets are in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Select this if your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// User ID of the member whose view the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// View type.
	//
	// - 0: View for the current Alibaba Cloud account.
	//
	// - 1: View for all accounts under your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s GetCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCapacityRequest) GoString() string {
	return s.String()
}

func (s *GetCapacityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCapacityRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetCapacityRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *GetCapacityRequest) SetRegionId(v string) *GetCapacityRequest {
	s.RegionId = &v
	return s
}

func (s *GetCapacityRequest) SetRoleFor(v int64) *GetCapacityRequest {
	s.RoleFor = &v
	return s
}

func (s *GetCapacityRequest) SetRoleType(v int32) *GetCapacityRequest {
	s.RoleType = &v
	return s
}

func (s *GetCapacityRequest) Validate() error {
	return dara.Validate(s)
}
