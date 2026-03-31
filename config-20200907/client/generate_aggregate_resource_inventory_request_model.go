// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateResourceInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v string) *GenerateAggregateResourceInventoryRequest
	GetAccountIds() *string
	SetAggregatorId(v string) *GenerateAggregateResourceInventoryRequest
	GetAggregatorId() *string
	SetRegions(v string) *GenerateAggregateResourceInventoryRequest
	GetRegions() *string
	SetResourceDeleted(v int32) *GenerateAggregateResourceInventoryRequest
	GetResourceDeleted() *int32
	SetResourceTypes(v string) *GenerateAggregateResourceInventoryRequest
	GetResourceTypes() *string
}

type GenerateAggregateResourceInventoryRequest struct {
	// The IDs of member accounts in the account group. Separate multiple member account IDs with commas (,).
	//
	// example:
	//
	// 126672004088****
	AccountIds *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a91d626622af0035****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The IDs of the regions to which the resources belong. Separate multiple region IDs with commas (,).
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

func (s GenerateAggregateResourceInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateResourceInventoryRequest) GoString() string {
	return s.String()
}

func (s *GenerateAggregateResourceInventoryRequest) GetAccountIds() *string {
	return s.AccountIds
}

func (s *GenerateAggregateResourceInventoryRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GenerateAggregateResourceInventoryRequest) GetRegions() *string {
	return s.Regions
}

func (s *GenerateAggregateResourceInventoryRequest) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *GenerateAggregateResourceInventoryRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *GenerateAggregateResourceInventoryRequest) SetAccountIds(v string) *GenerateAggregateResourceInventoryRequest {
	s.AccountIds = &v
	return s
}

func (s *GenerateAggregateResourceInventoryRequest) SetAggregatorId(v string) *GenerateAggregateResourceInventoryRequest {
	s.AggregatorId = &v
	return s
}

func (s *GenerateAggregateResourceInventoryRequest) SetRegions(v string) *GenerateAggregateResourceInventoryRequest {
	s.Regions = &v
	return s
}

func (s *GenerateAggregateResourceInventoryRequest) SetResourceDeleted(v int32) *GenerateAggregateResourceInventoryRequest {
	s.ResourceDeleted = &v
	return s
}

func (s *GenerateAggregateResourceInventoryRequest) SetResourceTypes(v string) *GenerateAggregateResourceInventoryRequest {
	s.ResourceTypes = &v
	return s
}

func (s *GenerateAggregateResourceInventoryRequest) Validate() error {
	return dara.Validate(s)
}
