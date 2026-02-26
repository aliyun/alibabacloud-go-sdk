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
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId              *string                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelatedResourceFilter []*ListResourceRelationshipsRequestRelatedResourceFilter `json:"RelatedResourceFilter,omitempty" xml:"RelatedResourceFilter,omitempty" type:"Repeated"`
	// This parameter is required.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
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

type ListResourceRelationshipsRequestRelatedResourceFilter struct {
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	MatchType *string   `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Value     []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
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
