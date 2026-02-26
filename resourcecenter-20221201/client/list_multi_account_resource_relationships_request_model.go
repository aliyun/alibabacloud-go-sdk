// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountResourceRelationshipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMultiAccountResourceRelationshipsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountResourceRelationshipsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListMultiAccountResourceRelationshipsRequest
	GetRegionId() *string
	SetRelatedResourceFilter(v []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) *ListMultiAccountResourceRelationshipsRequest
	GetRelatedResourceFilter() []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter
	SetResourceId(v string) *ListMultiAccountResourceRelationshipsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListMultiAccountResourceRelationshipsRequest
	GetResourceType() *string
	SetScope(v string) *ListMultiAccountResourceRelationshipsRequest
	GetScope() *string
}

type ListMultiAccountResourceRelationshipsRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId              *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelatedResourceFilter []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter `json:"RelatedResourceFilter,omitempty" xml:"RelatedResourceFilter,omitempty" type:"Repeated"`
	// This parameter is required.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetRelatedResourceFilter() []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	return s.RelatedResourceFilter
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetMaxResults(v int32) *ListMultiAccountResourceRelationshipsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetNextToken(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetRegionId(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.RegionId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetRelatedResourceFilter(v []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) *ListMultiAccountResourceRelationshipsRequest {
	s.RelatedResourceFilter = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetResourceId(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetResourceType(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetScope(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) Validate() error {
	if s.RelatedResourceFilter != nil {
		for _, item := range s.RelatedResourceFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter struct {
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	MatchType *string   `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Value     []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GetKey() *string {
	return s.Key
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GetValue() []*string {
	return s.Value
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetKey(v string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.Key = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetMatchType(v string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetValue(v []*string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.Value = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) Validate() error {
	return dara.Validate(s)
}
