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
	// The ID of the resource group to which you want to move the resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzmeobk5w****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resources that you want to move.
	//
	// >  You can move a maximum of 10 resources at a time. If you want to move more than 10 resources, move them in batches.
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
	// The ID of the resource.
	//
	// example:
	//
	// vpc-bp1sig0mjktx5ewx1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// vpc
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// vpc
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
