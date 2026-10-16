// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *MoveResourcesRequest
	GetResourceGroupId() *string
	SetResources(v []*MoveResourcesRequestResources) *MoveResourcesRequest
	GetResources() []*MoveResourcesRequestResources
}

type MoveResourcesRequest struct {
	// The ID of the destination resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzmeobk5w****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resources to move.
	//
	// > A maximum of 10 resources can be moved at a time. Move resources in batches if you have more than 10.
	//
	// This parameter is required.
	Resources []*MoveResourcesRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s MoveResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourcesRequest) GoString() string {
	return s.String()
}

func (s *MoveResourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *MoveResourcesRequest) GetResources() []*MoveResourcesRequestResources {
	return s.Resources
}

func (s *MoveResourcesRequest) SetResourceGroupId(v string) *MoveResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourcesRequest) SetResources(v []*MoveResourcesRequestResources) *MoveResourcesRequest {
	s.Resources = v
	return s
}

func (s *MoveResourcesRequest) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MoveResourcesRequestResources struct {
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-23v38****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// You can obtain the value from the **Resource Type*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The cloud service code.
	//
	// You can obtain the value from the **Cloud Service Code*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s MoveResourcesRequestResources) String() string {
	return dara.Prettify(s)
}

func (s MoveResourcesRequestResources) GoString() string {
	return s.String()
}

func (s *MoveResourcesRequestResources) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveResourcesRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *MoveResourcesRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *MoveResourcesRequestResources) GetService() *string {
	return s.Service
}

func (s *MoveResourcesRequestResources) SetRegionId(v string) *MoveResourcesRequestResources {
	s.RegionId = &v
	return s
}

func (s *MoveResourcesRequestResources) SetResourceId(v string) *MoveResourcesRequestResources {
	s.ResourceId = &v
	return s
}

func (s *MoveResourcesRequestResources) SetResourceType(v string) *MoveResourcesRequestResources {
	s.ResourceType = &v
	return s
}

func (s *MoveResourcesRequestResources) SetService(v string) *MoveResourcesRequestResources {
	s.Service = &v
	return s
}

func (s *MoveResourcesRequestResources) Validate() error {
	return dara.Validate(s)
}
