// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewResourceGroupId(v string) *ChangeResourceGroupRequest
	GetNewResourceGroupId() *string
	SetRegionId(v string) *ChangeResourceGroupRequest
	GetRegionId() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *ChangeResourceGroupRequest
	GetResourceType() *string
}

type ChangeResourceGroupRequest struct {
	// The ID of the new resource group.
	//
	// You can view resource group IDs in the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups) .
	//
	// example:
	//
	// rg-acfmzmhzo******
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud resource that you want to move to a new resource group.
	//
	// example:
	//
	// si-5dc794a7fd254e******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- service: service
	//
	// 	- serviceinstance: service instance
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
