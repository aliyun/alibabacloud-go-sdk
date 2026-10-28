// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferResourcesIntoGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *TransferResourcesIntoGroupRequest
	GetBusinessChannel() *string
	SetPlatform(v string) *TransferResourcesIntoGroupRequest
	GetPlatform() *string
	SetResourceGroupId(v string) *TransferResourcesIntoGroupRequest
	GetResourceGroupId() *string
	SetResources(v []*TransferResourcesIntoGroupRequestResources) *TransferResourcesIntoGroupRequest
	GetResources() []*TransferResourcesIntoGroupRequestResources
}

type TransferResourcesIntoGroupRequest struct {
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The WUYING product type.
	//
	// example:
	//
	// AliyunConsole
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-ckf3cx7isinhk***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resource objects.
	Resources []*TransferResourcesIntoGroupRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s TransferResourcesIntoGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesIntoGroupRequest) GoString() string {
	return s.String()
}

func (s *TransferResourcesIntoGroupRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *TransferResourcesIntoGroupRequest) GetPlatform() *string {
	return s.Platform
}

func (s *TransferResourcesIntoGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TransferResourcesIntoGroupRequest) GetResources() []*TransferResourcesIntoGroupRequestResources {
	return s.Resources
}

func (s *TransferResourcesIntoGroupRequest) SetBusinessChannel(v string) *TransferResourcesIntoGroupRequest {
	s.BusinessChannel = &v
	return s
}

func (s *TransferResourcesIntoGroupRequest) SetPlatform(v string) *TransferResourcesIntoGroupRequest {
	s.Platform = &v
	return s
}

func (s *TransferResourcesIntoGroupRequest) SetResourceGroupId(v string) *TransferResourcesIntoGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TransferResourcesIntoGroupRequest) SetResources(v []*TransferResourcesIntoGroupRequestResources) *TransferResourcesIntoGroupRequest {
	s.Resources = v
	return s
}

func (s *TransferResourcesIntoGroupRequest) Validate() error {
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

type TransferResourcesIntoGroupRequestResources struct {
	// The region ID of the resource.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// ecd-0jl0dwz61hn59tto****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s TransferResourcesIntoGroupRequestResources) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesIntoGroupRequestResources) GoString() string {
	return s.String()
}

func (s *TransferResourcesIntoGroupRequestResources) GetRegionId() *string {
	return s.RegionId
}

func (s *TransferResourcesIntoGroupRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *TransferResourcesIntoGroupRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *TransferResourcesIntoGroupRequestResources) SetRegionId(v string) *TransferResourcesIntoGroupRequestResources {
	s.RegionId = &v
	return s
}

func (s *TransferResourcesIntoGroupRequestResources) SetResourceId(v string) *TransferResourcesIntoGroupRequestResources {
	s.ResourceId = &v
	return s
}

func (s *TransferResourcesIntoGroupRequestResources) SetResourceType(v string) *TransferResourcesIntoGroupRequestResources {
	s.ResourceType = &v
	return s
}

func (s *TransferResourcesIntoGroupRequestResources) Validate() error {
	return dara.Validate(s)
}
