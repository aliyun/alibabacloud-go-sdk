// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceRelationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceRelationsRequest
	GetNextToken() *string
	SetRegion(v string) *ListResourceRelationsRequest
	GetRegion() *string
	SetRelationType(v string) *ListResourceRelationsRequest
	GetRelationType() *string
	SetResourceId(v string) *ListResourceRelationsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListResourceRelationsRequest
	GetResourceType() *string
	SetTargetResourceId(v string) *ListResourceRelationsRequest
	GetTargetResourceId() *string
	SetTargetResourceType(v string) *ListResourceRelationsRequest
	GetTargetResourceType() *string
}

type ListResourceRelationsRequest struct {
	// The maximum number of entries to return in a request. Valid values: 1 to 1000.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// C2DjqMYSy0is7zSMGf21****
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

func (s ListResourceRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceRelationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceRelationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceRelationsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListResourceRelationsRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *ListResourceRelationsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourceRelationsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceRelationsRequest) GetTargetResourceId() *string {
	return s.TargetResourceId
}

func (s *ListResourceRelationsRequest) GetTargetResourceType() *string {
	return s.TargetResourceType
}

func (s *ListResourceRelationsRequest) SetMaxResults(v int32) *ListResourceRelationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceRelationsRequest) SetNextToken(v string) *ListResourceRelationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceRelationsRequest) SetRegion(v string) *ListResourceRelationsRequest {
	s.Region = &v
	return s
}

func (s *ListResourceRelationsRequest) SetRelationType(v string) *ListResourceRelationsRequest {
	s.RelationType = &v
	return s
}

func (s *ListResourceRelationsRequest) SetResourceId(v string) *ListResourceRelationsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourceRelationsRequest) SetResourceType(v string) *ListResourceRelationsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceRelationsRequest) SetTargetResourceId(v string) *ListResourceRelationsRequest {
	s.TargetResourceId = &v
	return s
}

func (s *ListResourceRelationsRequest) SetTargetResourceType(v string) *ListResourceRelationsRequest {
	s.TargetResourceType = &v
	return s
}

func (s *ListResourceRelationsRequest) Validate() error {
	return dara.Validate(s)
}
