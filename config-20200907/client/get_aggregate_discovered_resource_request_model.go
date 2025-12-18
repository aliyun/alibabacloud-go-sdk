// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateDiscoveredResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateDiscoveredResourceRequest
	GetAggregatorId() *string
	SetComplianceOption(v int32) *GetAggregateDiscoveredResourceRequest
	GetComplianceOption() *int32
	SetRegion(v string) *GetAggregateDiscoveredResourceRequest
	GetRegion() *string
	SetResourceAccountId(v int64) *GetAggregateDiscoveredResourceRequest
	GetResourceAccountId() *int64
	SetResourceId(v string) *GetAggregateDiscoveredResourceRequest
	GetResourceId() *string
	SetResourceOwnerId(v int64) *GetAggregateDiscoveredResourceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *GetAggregateDiscoveredResourceRequest
	GetResourceType() *string
}

type GetAggregateDiscoveredResourceRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-5885626622af0008****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// Specifies whether to query the compliance results of the resource. Valid values:
	//
	// 	- 0 (default): does not query the compliance results of the resource.
	//
	// 	- 1: queries the compliance results of the resource.
	//
	// example:
	//
	// 0
	ComplianceOption *int32 `json:"ComplianceOption,omitempty" xml:"ComplianceOption,omitempty"`
	// The ID of the region in which the resource resides.
	//
	// For more information about how to query the ID of a region in which the resource resides, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/411691.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Required. The ID of the Alibaba Cloud account to which the specified resource belongs in the account group.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The resource ID.
	//
	// For more information about how to obtain the ID of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/411691.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// new-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Deprecated
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource.
	//
	// For more information about how to obtain the type of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/411691.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateDiscoveredResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateDiscoveredResourceRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateDiscoveredResourceRequest) GetComplianceOption() *int32 {
	return s.ComplianceOption
}

func (s *GetAggregateDiscoveredResourceRequest) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateDiscoveredResourceRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *GetAggregateDiscoveredResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetAggregateDiscoveredResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAggregateDiscoveredResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateDiscoveredResourceRequest) SetAggregatorId(v string) *GetAggregateDiscoveredResourceRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetComplianceOption(v int32) *GetAggregateDiscoveredResourceRequest {
	s.ComplianceOption = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetRegion(v string) *GetAggregateDiscoveredResourceRequest {
	s.Region = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetResourceAccountId(v int64) *GetAggregateDiscoveredResourceRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetResourceId(v string) *GetAggregateDiscoveredResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetResourceOwnerId(v int64) *GetAggregateDiscoveredResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetResourceType(v string) *GetAggregateDiscoveredResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) Validate() error {
	return dara.Validate(s)
}
