// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *JoinResourceGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *JoinResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *JoinResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *JoinResourceGroupRequest
	GetResourceType() *string
}

type JoinResourceGroupRequest struct {
	// The ID of the region in which you want to create the instance.
	//
	// This parameter is required.
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
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-c95f0a39d8ff****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- cluster: cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s JoinResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *JoinResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *JoinResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *JoinResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *JoinResourceGroupRequest) SetRegionId(v string) *JoinResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceGroupId(v string) *JoinResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceId(v string) *JoinResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceType(v string) *JoinResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *JoinResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
