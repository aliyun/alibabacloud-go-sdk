// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListDeliveryRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDeliveryRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListDeliveryRequest
	GetRoleType() *int32
}

type ListDeliveryRequest struct {
	// The region of the data management center for threat analysis. Select the region based on where your assets are located. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Select this value if your assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud account ID of a member. An administrator can use this parameter to view data from the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of view. Valid values:
	//
	// - 0: The view for the current Alibaba Cloud account.
	//
	// - 1: The view for all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDeliveryRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDeliveryRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListDeliveryRequest) SetRegionId(v string) *ListDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *ListDeliveryRequest) SetRoleFor(v int64) *ListDeliveryRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDeliveryRequest) SetRoleType(v int32) *ListDeliveryRequest {
	s.RoleType = &v
	return s
}

func (s *ListDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
