// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *MoveResourceGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *MoveResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *MoveResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *MoveResourceGroupRequest
	GetResourceType() *string
}

type MoveResourceGroupRequest struct {
	// The region of the organization to which the owner of the certificate belongs. Valid values: ap-southeast-1 and cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmykgxu5d46ey
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas-cn-4591d3xa****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.\\
	//
	// Default value: **instance**
	//
	// Valid values:
	//
	// 	- instance: certificate order
	//
	// 	- Certificate: certificate
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *MoveResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *MoveResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
