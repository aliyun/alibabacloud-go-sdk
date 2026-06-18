// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferResourcesOutofGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *TransferResourcesOutofGroupRequest
	GetBusinessChannel() *string
	SetPlatform(v string) *TransferResourcesOutofGroupRequest
	GetPlatform() *string
	SetResourceGroupId(v string) *TransferResourcesOutofGroupRequest
	GetResourceGroupId() *string
	SetResources(v []*TransferResourcesOutofGroupRequestResources) *TransferResourcesOutofGroupRequest
	GetResources() []*TransferResourcesOutofGroupRequestResources
	SetTargetResourceGroupId(v string) *TransferResourcesOutofGroupRequest
	GetTargetResourceGroupId() *string
}

type TransferResourcesOutofGroupRequest struct {
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// example:
	//
	// AliyunConsole
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// rg-66konqltcpx43****
	ResourceGroupId *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Resources       []*TransferResourcesOutofGroupRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// rg-hcjo6dugpo1****
	TargetResourceGroupId *string `json:"TargetResourceGroupId,omitempty" xml:"TargetResourceGroupId,omitempty"`
}

func (s TransferResourcesOutofGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesOutofGroupRequest) GoString() string {
	return s.String()
}

func (s *TransferResourcesOutofGroupRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *TransferResourcesOutofGroupRequest) GetPlatform() *string {
	return s.Platform
}

func (s *TransferResourcesOutofGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TransferResourcesOutofGroupRequest) GetResources() []*TransferResourcesOutofGroupRequestResources {
	return s.Resources
}

func (s *TransferResourcesOutofGroupRequest) GetTargetResourceGroupId() *string {
	return s.TargetResourceGroupId
}

func (s *TransferResourcesOutofGroupRequest) SetBusinessChannel(v string) *TransferResourcesOutofGroupRequest {
	s.BusinessChannel = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetPlatform(v string) *TransferResourcesOutofGroupRequest {
	s.Platform = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetResourceGroupId(v string) *TransferResourcesOutofGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetResources(v []*TransferResourcesOutofGroupRequestResources) *TransferResourcesOutofGroupRequest {
	s.Resources = v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetTargetResourceGroupId(v string) *TransferResourcesOutofGroupRequest {
	s.TargetResourceGroupId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) Validate() error {
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

type TransferResourcesOutofGroupRequestResources struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ecd-1shc5sk0vo****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s TransferResourcesOutofGroupRequestResources) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesOutofGroupRequestResources) GoString() string {
	return s.String()
}

func (s *TransferResourcesOutofGroupRequestResources) GetRegionId() *string {
	return s.RegionId
}

func (s *TransferResourcesOutofGroupRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *TransferResourcesOutofGroupRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *TransferResourcesOutofGroupRequestResources) SetRegionId(v string) *TransferResourcesOutofGroupRequestResources {
	s.RegionId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequestResources) SetResourceId(v string) *TransferResourcesOutofGroupRequestResources {
	s.ResourceId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequestResources) SetResourceType(v string) *TransferResourcesOutofGroupRequestResources {
	s.ResourceType = &v
	return s
}

func (s *TransferResourcesOutofGroupRequestResources) Validate() error {
	return dara.Validate(s)
}
