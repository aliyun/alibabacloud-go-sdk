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
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// This parameter is required.
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// This parameter is required.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TargetResourceId   *string `json:"TargetResourceId,omitempty" xml:"TargetResourceId,omitempty"`
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
