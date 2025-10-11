// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRelationshipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceRelationshipsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceRelationshipsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListResourceRelationshipsRequest
	GetRegionId() *string
	SetRelatedResourceFilter(v []*ListResourceRelationshipsRequestRelatedResourceFilter) *ListResourceRelationshipsRequest
	GetRelatedResourceFilter() []*ListResourceRelationshipsRequestRelatedResourceFilter
	SetResourceId(v string) *ListResourceRelationshipsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListResourceRelationshipsRequest
	GetResourceType() *string
}

type ListResourceRelationshipsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The filter conditions for resources associated with the resource.
	RelatedResourceFilter []*ListResourceRelationshipsRequestRelatedResourceFilter `json:"RelatedResourceFilter,omitempty" xml:"RelatedResourceFilter,omitempty" type:"Repeated"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceRelationshipsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationshipsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceRelationshipsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceRelationshipsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceRelationshipsRequest) GetRelatedResourceFilter() []*ListResourceRelationshipsRequestRelatedResourceFilter {
	return s.RelatedResourceFilter
}

func (s *ListResourceRelationshipsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourceRelationshipsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceRelationshipsRequest) SetMaxResults(v int32) *ListResourceRelationshipsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetNextToken(v string) *ListResourceRelationshipsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetRegionId(v string) *ListResourceRelationshipsRequest {
	s.RegionId = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetRelatedResourceFilter(v []*ListResourceRelationshipsRequestRelatedResourceFilter) *ListResourceRelationshipsRequest {
	s.RelatedResourceFilter = v
	return s
}

func (s *ListResourceRelationshipsRequest) SetResourceId(v string) *ListResourceRelationshipsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetResourceType(v string) *ListResourceRelationshipsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceRelationshipsRequest) Validate() error {
	return dara.Validate(s)
}

type ListResourceRelationshipsRequestRelatedResourceFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RelatedResourceRegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListResourceRelationshipsRequestRelatedResourceFilter) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationshipsRequestRelatedResourceFilter) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) GetKey() *string {
	return s.Key
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) GetValue() []*string {
	return s.Value
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) SetKey(v string) *ListResourceRelationshipsRequestRelatedResourceFilter {
	s.Key = &v
	return s
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) SetMatchType(v string) *ListResourceRelationshipsRequestRelatedResourceFilter {
	s.MatchType = &v
	return s
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) SetValue(v []*string) *ListResourceRelationshipsRequestRelatedResourceFilter {
	s.Value = v
	return s
}

func (s *ListResourceRelationshipsRequestRelatedResourceFilter) Validate() error {
	return dara.Validate(s)
}
