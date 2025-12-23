// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMultiAccountResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*SearchMultiAccountResourcesRequestFilter) *SearchMultiAccountResourcesRequest
	GetFilter() []*SearchMultiAccountResourcesRequestFilter
	SetMaxResults(v int32) *SearchMultiAccountResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchMultiAccountResourcesRequest
	GetNextToken() *string
	SetScope(v string) *SearchMultiAccountResourcesRequest
	GetScope() *string
	SetSortCriterion(v *SearchMultiAccountResourcesRequestSortCriterion) *SearchMultiAccountResourcesRequest
	GetSortCriterion() *SearchMultiAccountResourcesRequestSortCriterion
}

type SearchMultiAccountResourcesRequest struct {
	// The filter conditions.
	Filter []*SearchMultiAccountResourcesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
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
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the token to initiate another request and obtain the remaining entries.``
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. You can set the value to one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The method that is used to sort the entries returned.
	SortCriterion *SearchMultiAccountResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
}

func (s SearchMultiAccountResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesRequest) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequest) GetFilter() []*SearchMultiAccountResourcesRequestFilter {
	return s.Filter
}

func (s *SearchMultiAccountResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchMultiAccountResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchMultiAccountResourcesRequest) GetScope() *string {
	return s.Scope
}

func (s *SearchMultiAccountResourcesRequest) GetSortCriterion() *SearchMultiAccountResourcesRequestSortCriterion {
	return s.SortCriterion
}

func (s *SearchMultiAccountResourcesRequest) SetFilter(v []*SearchMultiAccountResourcesRequestFilter) *SearchMultiAccountResourcesRequest {
	s.Filter = v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetMaxResults(v int32) *SearchMultiAccountResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetNextToken(v string) *SearchMultiAccountResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetScope(v string) *SearchMultiAccountResourcesRequest {
	s.Scope = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetSortCriterion(v *SearchMultiAccountResourcesRequestSortCriterion) *SearchMultiAccountResourcesRequest {
	s.SortCriterion = v
	return s
}

func (s *SearchMultiAccountResourcesRequest) Validate() error {
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

type SearchMultiAccountResourcesRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// ResourceGroupId
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

func (s SearchMultiAccountResourcesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *SearchMultiAccountResourcesRequestFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *SearchMultiAccountResourcesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *SearchMultiAccountResourcesRequestFilter) SetKey(v string) *SearchMultiAccountResourcesRequestFilter {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) SetMatchType(v string) *SearchMultiAccountResourcesRequestFilter {
	s.MatchType = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) SetValue(v []*string) *SearchMultiAccountResourcesRequestFilter {
	s.Value = v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type SearchMultiAccountResourcesRequestSortCriterion struct {
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

func (s SearchMultiAccountResourcesRequestSortCriterion) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) GetKey() *string {
	return s.Key
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) GetOrder() *string {
	return s.Order
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) SetKey(v string) *SearchMultiAccountResourcesRequestSortCriterion {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) SetOrder(v string) *SearchMultiAccountResourcesRequestSortCriterion {
	s.Order = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) Validate() error {
	return dara.Validate(s)
}
