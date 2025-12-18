// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceCountsGroupByRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceCountsGroupByRegionRequest
	GetAggregatorId() *string
	SetFolderId(v string) *GetAggregateResourceCountsGroupByRegionRequest
	GetFolderId() *string
	SetResourceAccountId(v int64) *GetAggregateResourceCountsGroupByRegionRequest
	GetResourceAccountId() *int64
	SetResourceOwnerId(v int64) *GetAggregateResourceCountsGroupByRegionRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *GetAggregateResourceCountsGroupByRegionRequest
	GetResourceType() *string
}

type GetAggregateResourceCountsGroupByRegionRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a260626622af0005****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the folder in the resource directory. For more information about how to obtain the ID of a folder, see [View the basic information of a folder](https://help.aliyun.com/document_detail/111223.html).
	//
	// example:
	//
	// r-BU****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resources in the account group belong.
	//
	// > You can use either the ResourceAccountId or ResourceOwnerId parameter. We recommend that you use the ResourceAccountId parameter.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// Deprecated
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type.
	//
	// For more information about how to obtain the type of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateResourceCountsGroupByRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetAggregatorId(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetFolderId(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.FolderId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetResourceAccountId(v int64) *GetAggregateResourceCountsGroupByRegionRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetResourceOwnerId(v int64) *GetAggregateResourceCountsGroupByRegionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetResourceType(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) Validate() error {
	return dara.Validate(s)
}
