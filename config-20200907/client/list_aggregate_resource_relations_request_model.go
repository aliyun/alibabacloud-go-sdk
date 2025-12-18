// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateResourceRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateResourceRelationsRequest
	GetAggregatorId() *string
	SetMaxResults(v int32) *ListAggregateResourceRelationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggregateResourceRelationsRequest
	GetNextToken() *string
	SetRegion(v string) *ListAggregateResourceRelationsRequest
	GetRegion() *string
	SetRelationType(v string) *ListAggregateResourceRelationsRequest
	GetRelationType() *string
	SetResourceAccountId(v int64) *ListAggregateResourceRelationsRequest
	GetResourceAccountId() *int64
	SetResourceId(v string) *ListAggregateResourceRelationsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListAggregateResourceRelationsRequest
	GetResourceType() *string
	SetTargetResourceId(v string) *ListAggregateResourceRelationsRequest
	GetTargetResourceId() *string
	SetTargetResourceType(v string) *ListAggregateResourceRelationsRequest
	GetTargetResourceType() *string
}

type ListAggregateResourceRelationsRequest struct {
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-856a626622af0033****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The maximum number of entries to return for a single request. Valid values: 1 to 1000.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AcBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the relationship between the resource and the object.
	//
	// Valid values:
	//
	// 	- IsContained: The object is included as part of the resource.
	//
	// 	- IsAttachedTo: The object is added to the resource.
	//
	// 	- IsAssociatedIn: The object is associated with the resource.
	//
	// 	- Contains: The object contains the resource.
	//
	// example:
	//
	// IsAttachedTo
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The ID of the Alibaba Cloud account to which the resources in the account group belong.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-j6cajg9yrfoh4sas****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the resource that is associated with the object.
	//
	// example:
	//
	// d-j6c8k731qbrc7fxi****
	TargetResourceId *string `json:"TargetResourceId,omitempty" xml:"TargetResourceId,omitempty"`
	// The type of the resource that is associated with the object.
	//
	// example:
	//
	// ACS::ECS::Disk
	TargetResourceType *string `json:"TargetResourceType,omitempty" xml:"TargetResourceType,omitempty"`
}

func (s ListAggregateResourceRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourceRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceRelationsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateResourceRelationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateResourceRelationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateResourceRelationsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListAggregateResourceRelationsRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *ListAggregateResourceRelationsRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *ListAggregateResourceRelationsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAggregateResourceRelationsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAggregateResourceRelationsRequest) GetTargetResourceId() *string {
	return s.TargetResourceId
}

func (s *ListAggregateResourceRelationsRequest) GetTargetResourceType() *string {
	return s.TargetResourceType
}

func (s *ListAggregateResourceRelationsRequest) SetAggregatorId(v string) *ListAggregateResourceRelationsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetMaxResults(v int32) *ListAggregateResourceRelationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetNextToken(v string) *ListAggregateResourceRelationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetRegion(v string) *ListAggregateResourceRelationsRequest {
	s.Region = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetRelationType(v string) *ListAggregateResourceRelationsRequest {
	s.RelationType = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetResourceAccountId(v int64) *ListAggregateResourceRelationsRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetResourceId(v string) *ListAggregateResourceRelationsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetResourceType(v string) *ListAggregateResourceRelationsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetTargetResourceId(v string) *ListAggregateResourceRelationsRequest {
	s.TargetResourceId = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) SetTargetResourceType(v string) *ListAggregateResourceRelationsRequest {
	s.TargetResourceType = &v
	return s
}

func (s *ListAggregateResourceRelationsRequest) Validate() error {
	return dara.Validate(s)
}
