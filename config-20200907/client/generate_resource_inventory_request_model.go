// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResourceInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v string) *GenerateResourceInventoryRequest
	GetRegions() *string
	SetResourceDeleted(v int32) *GenerateResourceInventoryRequest
	GetResourceDeleted() *int32
	SetResourceTypes(v string) *GenerateResourceInventoryRequest
	GetResourceTypes() *string
}

type GenerateResourceInventoryRequest struct {
	// The region IDs of the resources. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// Indicates whether the resource is deleted. Valid values:
	//
	// 	- 1 (default): The resource is retained.
	//
	// 	- 0: The resource is deleted.
	//
	// example:
	//
	// 1
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The resource types. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s GenerateResourceInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateResourceInventoryRequest) GoString() string {
	return s.String()
}

func (s *GenerateResourceInventoryRequest) GetRegions() *string {
	return s.Regions
}

func (s *GenerateResourceInventoryRequest) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *GenerateResourceInventoryRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *GenerateResourceInventoryRequest) SetRegions(v string) *GenerateResourceInventoryRequest {
	s.Regions = &v
	return s
}

func (s *GenerateResourceInventoryRequest) SetResourceDeleted(v int32) *GenerateResourceInventoryRequest {
	s.ResourceDeleted = &v
	return s
}

func (s *GenerateResourceInventoryRequest) SetResourceTypes(v string) *GenerateResourceInventoryRequest {
	s.ResourceTypes = &v
	return s
}

func (s *GenerateResourceInventoryRequest) Validate() error {
	return dara.Validate(s)
}
