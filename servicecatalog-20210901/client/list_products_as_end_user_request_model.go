// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsAsEndUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*ListProductsAsEndUserRequestFilters) *ListProductsAsEndUserRequest
	GetFilters() []*ListProductsAsEndUserRequestFilters
	SetPageNumber(v int32) *ListProductsAsEndUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductsAsEndUserRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListProductsAsEndUserRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListProductsAsEndUserRequest
	GetSortOrder() *string
}

type ListProductsAsEndUserRequest struct {
	// The filter conditions.
	Filters []*ListProductsAsEndUserRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
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
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field that is used to sort data.
	//
	// The default value is CreateTime, which specifies the time when the product was created.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The method that is used to sort the returned entries. Valid values:
	//
	// 	- Asc: sorts the entries in ascending order.
	//
	// 	- Desc (default): sorts the entries in descending order.
	//
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProductsAsEndUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsEndUserRequest) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserRequest) GetFilters() []*ListProductsAsEndUserRequestFilters {
	return s.Filters
}

func (s *ListProductsAsEndUserRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductsAsEndUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductsAsEndUserRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProductsAsEndUserRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListProductsAsEndUserRequest) SetFilters(v []*ListProductsAsEndUserRequestFilters) *ListProductsAsEndUserRequest {
	s.Filters = v
	return s
}

func (s *ListProductsAsEndUserRequest) SetPageNumber(v int32) *ListProductsAsEndUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductsAsEndUserRequest) SetPageSize(v int32) *ListProductsAsEndUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductsAsEndUserRequest) SetSortBy(v string) *ListProductsAsEndUserRequest {
	s.SortBy = &v
	return s
}

func (s *ListProductsAsEndUserRequest) SetSortOrder(v string) *ListProductsAsEndUserRequest {
	s.SortOrder = &v
	return s
}

func (s *ListProductsAsEndUserRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsAsEndUserRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// 	- ProductName: performs exact matches by product name. Product names are not case-sensitive.
	//
	// 	- FullTextSearch: performs full-text searches by product name, product provider, or product description. Fuzzy match is supported.
	//
	// example:
	//
	// ProductName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// DEMO-Create an ECS instance
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductsAsEndUserRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsEndUserRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListProductsAsEndUserRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListProductsAsEndUserRequestFilters) SetKey(v string) *ListProductsAsEndUserRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProductsAsEndUserRequestFilters) SetValue(v string) *ListProductsAsEndUserRequestFilters {
	s.Value = &v
	return s
}

func (s *ListProductsAsEndUserRequestFilters) Validate() error {
	return dara.Validate(s)
}
