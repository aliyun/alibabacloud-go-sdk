// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*GetResourceCountsRequestFilter) *GetResourceCountsRequest
	GetFilter() []*GetResourceCountsRequestFilter
	SetGroupByKey(v string) *GetResourceCountsRequest
	GetGroupByKey() *string
	SetIncludeDeletedResources(v bool) *GetResourceCountsRequest
	GetIncludeDeletedResources() *bool
	SetSearchExpression(v string) *GetResourceCountsRequest
	GetSearchExpression() *string
}

type GetResourceCountsRequest struct {
	// The filter conditions.
	Filter []*GetResourceCountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The dimension by which the queried resources are grouped. Valid values:
	//
	// - ResourceType: The resource type.
	//
	// - RegionId: The region.
	//
	// - ResourceGroupId: The resource group ID.
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// Specifies whether to include deleted resources. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	IncludeDeletedResources *bool `json:"IncludeDeletedResources,omitempty" xml:"IncludeDeletedResources,omitempty"`
	// The search keyword. Resource Center filters the search results based on relevance.
	//
	// example:
	//
	// keywords
	SearchExpression *string `json:"SearchExpression,omitempty" xml:"SearchExpression,omitempty"`
}

func (s GetResourceCountsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceCountsRequest) GetFilter() []*GetResourceCountsRequestFilter {
	return s.Filter
}

func (s *GetResourceCountsRequest) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetResourceCountsRequest) GetIncludeDeletedResources() *bool {
	return s.IncludeDeletedResources
}

func (s *GetResourceCountsRequest) GetSearchExpression() *string {
	return s.SearchExpression
}

func (s *GetResourceCountsRequest) SetFilter(v []*GetResourceCountsRequestFilter) *GetResourceCountsRequest {
	s.Filter = v
	return s
}

func (s *GetResourceCountsRequest) SetGroupByKey(v string) *GetResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

func (s *GetResourceCountsRequest) SetIncludeDeletedResources(v bool) *GetResourceCountsRequest {
	s.IncludeDeletedResources = &v
	return s
}

func (s *GetResourceCountsRequest) SetSearchExpression(v string) *GetResourceCountsRequest {
	s.SearchExpression = &v
	return s
}

func (s *GetResourceCountsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceCountsRequestFilter struct {
	// The key of the filter condition. For information about valid values, see the "`Supported filter parameters`" section below.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method.
	//
	// Set this parameter to `Equals`, which means an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s GetResourceCountsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCountsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetResourceCountsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetResourceCountsRequestFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *GetResourceCountsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *GetResourceCountsRequestFilter) SetKey(v string) *GetResourceCountsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetResourceCountsRequestFilter) SetMatchType(v string) *GetResourceCountsRequestFilter {
	s.MatchType = &v
	return s
}

func (s *GetResourceCountsRequestFilter) SetValue(v []*string) *GetResourceCountsRequestFilter {
	s.Value = v
	return s
}

func (s *GetResourceCountsRequestFilter) Validate() error {
	return dara.Validate(s)
}
