// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*SearchResourcesRequestFilter) *SearchResourcesRequest
	GetFilter() []*SearchResourcesRequestFilter
	SetIncludeDeletedResources(v bool) *SearchResourcesRequest
	GetIncludeDeletedResources() *bool
	SetMaxResults(v int32) *SearchResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchResourcesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *SearchResourcesRequest
	GetResourceGroupId() *string
	SetSearchExpression(v string) *SearchResourcesRequest
	GetSearchExpression() *string
	SetSortCriterion(v *SearchResourcesRequestSortCriterion) *SearchResourcesRequest
	GetSortCriterion() *SearchResourcesRequestSortCriterion
}

type SearchResourcesRequest struct {
	// The filter conditions.
	Filter                  []*SearchResourcesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	IncludeDeletedResources *bool                           `json:"IncludeDeletedResources,omitempty" xml:"IncludeDeletedResources,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SearchExpression *string `json:"SearchExpression,omitempty" xml:"SearchExpression,omitempty"`
	// The method that is used to sort the entries returned.
	SortCriterion *SearchResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
}

func (s SearchResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesRequest) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequest) GetFilter() []*SearchResourcesRequestFilter {
	return s.Filter
}

func (s *SearchResourcesRequest) GetIncludeDeletedResources() *bool {
	return s.IncludeDeletedResources
}

func (s *SearchResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchResourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchResourcesRequest) GetSearchExpression() *string {
	return s.SearchExpression
}

func (s *SearchResourcesRequest) GetSortCriterion() *SearchResourcesRequestSortCriterion {
	return s.SortCriterion
}

func (s *SearchResourcesRequest) SetFilter(v []*SearchResourcesRequestFilter) *SearchResourcesRequest {
	s.Filter = v
	return s
}

func (s *SearchResourcesRequest) SetIncludeDeletedResources(v bool) *SearchResourcesRequest {
	s.IncludeDeletedResources = &v
	return s
}

func (s *SearchResourcesRequest) SetMaxResults(v int32) *SearchResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchResourcesRequest) SetNextToken(v string) *SearchResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchResourcesRequest) SetResourceGroupId(v string) *SearchResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchResourcesRequest) SetSearchExpression(v string) *SearchResourcesRequest {
	s.SearchExpression = &v
	return s
}

func (s *SearchResourcesRequest) SetSortCriterion(v *SearchResourcesRequestSortCriterion) *SearchResourcesRequest {
	s.SortCriterion = v
	return s
}

func (s *SearchResourcesRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SortCriterion != nil {
		if err := s.SortCriterion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchResourcesRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// The value Equals indicates an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchResourcesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *SearchResourcesRequestFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *SearchResourcesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *SearchResourcesRequestFilter) SetKey(v string) *SearchResourcesRequestFilter {
	s.Key = &v
	return s
}

func (s *SearchResourcesRequestFilter) SetMatchType(v string) *SearchResourcesRequestFilter {
	s.MatchType = &v
	return s
}

func (s *SearchResourcesRequestFilter) SetValue(v []*string) *SearchResourcesRequestFilter {
	s.Value = v
	return s
}

func (s *SearchResourcesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type SearchResourcesRequestSortCriterion struct {
	// The attribute based on which the entries are sorted.
	//
	// The value CreateTime indicates the creation time of resources.
	//
	// example:
	//
	// CreateTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- ASC: The entries are sorted in ascending order. This value is the default value.
	//
	// 	- DESC: The entries are sorted in descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s SearchResourcesRequestSortCriterion) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequestSortCriterion) GetKey() *string {
	return s.Key
}

func (s *SearchResourcesRequestSortCriterion) GetOrder() *string {
	return s.Order
}

func (s *SearchResourcesRequestSortCriterion) SetKey(v string) *SearchResourcesRequestSortCriterion {
	s.Key = &v
	return s
}

func (s *SearchResourcesRequestSortCriterion) SetOrder(v string) *SearchResourcesRequestSortCriterion {
	s.Order = &v
	return s
}

func (s *SearchResourcesRequestSortCriterion) Validate() error {
	return dara.Validate(s)
}
