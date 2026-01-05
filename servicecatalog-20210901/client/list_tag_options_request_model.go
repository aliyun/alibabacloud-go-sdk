// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v *ListTagOptionsRequestFilters) *ListTagOptionsRequest
	GetFilters() *ListTagOptionsRequestFilters
	SetPageNumber(v int32) *ListTagOptionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagOptionsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListTagOptionsRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListTagOptionsRequest
	GetSortOrder() *string
}

type ListTagOptionsRequest struct {
	// The filter condition.
	Filters *ListTagOptionsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information based on which you want to sort the query results.
	//
	// Set the value to CreateTime, which specifies the creation time of tag options.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- Asc: the ascending order
	//
	// 	- Desc (default): the descending order
	//
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListTagOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagOptionsRequest) GoString() string {
	return s.String()
}

func (s *ListTagOptionsRequest) GetFilters() *ListTagOptionsRequestFilters {
	return s.Filters
}

func (s *ListTagOptionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagOptionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagOptionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTagOptionsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTagOptionsRequest) SetFilters(v *ListTagOptionsRequestFilters) *ListTagOptionsRequest {
	s.Filters = v
	return s
}

func (s *ListTagOptionsRequest) SetPageNumber(v int32) *ListTagOptionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagOptionsRequest) SetPageSize(v int32) *ListTagOptionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagOptionsRequest) SetSortBy(v string) *ListTagOptionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTagOptionsRequest) SetSortOrder(v string) *ListTagOptionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTagOptionsRequest) Validate() error {
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagOptionsRequestFilters struct {
	// Specifies whether to enable the tag option. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The full-text search method.
	//
	// example:
	//
	// k1
	FullTextSearch *string `json:"FullTextSearch,omitempty" xml:"FullTextSearch,omitempty"`
	// The key of the tag option.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag option.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagOptionsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListTagOptionsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListTagOptionsRequestFilters) GetActive() *bool {
	return s.Active
}

func (s *ListTagOptionsRequestFilters) GetFullTextSearch() *string {
	return s.FullTextSearch
}

func (s *ListTagOptionsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListTagOptionsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListTagOptionsRequestFilters) SetActive(v bool) *ListTagOptionsRequestFilters {
	s.Active = &v
	return s
}

func (s *ListTagOptionsRequestFilters) SetFullTextSearch(v string) *ListTagOptionsRequestFilters {
	s.FullTextSearch = &v
	return s
}

func (s *ListTagOptionsRequestFilters) SetKey(v string) *ListTagOptionsRequestFilters {
	s.Key = &v
	return s
}

func (s *ListTagOptionsRequestFilters) SetValue(v string) *ListTagOptionsRequestFilters {
	s.Value = &v
	return s
}

func (s *ListTagOptionsRequestFilters) Validate() error {
	return dara.Validate(s)
}
