// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CreateServiceLinkedRoleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateServiceLinkedRoleRequest
	GetRegionId() *string
}

type CreateServiceLinkedRoleRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateServiceLinkedRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceLinkedRoleRequest) SetOwnerId(v int64) *CreateServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
