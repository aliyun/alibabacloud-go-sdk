// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsAsAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*ListProductsAsAdminRequestFilters) *ListProductsAsAdminRequest
	GetFilters() []*ListProductsAsAdminRequestFilters
	SetPageNumber(v int32) *ListProductsAsAdminRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductsAsAdminRequest
	GetPageSize() *int32
	SetPortfolioId(v string) *ListProductsAsAdminRequest
	GetPortfolioId() *string
	SetScope(v string) *ListProductsAsAdminRequest
	GetScope() *string
	SetSortBy(v string) *ListProductsAsAdminRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListProductsAsAdminRequest
	GetSortOrder() *string
}

type ListProductsAsAdminRequest struct {
	// The filter conditions.
	Filters []*ListProductsAsAdminRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
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
	// The ID of the product portfolio.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The query scope. Valid values:
	//
	// 	- Local: the products that are created by using the current account. This is the default value.
	//
	// 	- Import: the products that are imported from other accounts.
	//
	// 	- All: all available products.
	//
	// example:
	//
	// Local
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The field that is used to sort the queried data.
	//
	// Set the value to CreateTime, which specifies the time when the product was created.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the queried data. Valid values:
	//
	// 	- Asc: the ascending order
	//
	// 	- Desc: the descending order
	//
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProductsAsAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsAdminRequest) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminRequest) GetFilters() []*ListProductsAsAdminRequestFilters {
	return s.Filters
}

func (s *ListProductsAsAdminRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductsAsAdminRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductsAsAdminRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *ListProductsAsAdminRequest) GetScope() *string {
	return s.Scope
}

func (s *ListProductsAsAdminRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProductsAsAdminRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListProductsAsAdminRequest) SetFilters(v []*ListProductsAsAdminRequestFilters) *ListProductsAsAdminRequest {
	s.Filters = v
	return s
}

func (s *ListProductsAsAdminRequest) SetPageNumber(v int32) *ListProductsAsAdminRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetPageSize(v int32) *ListProductsAsAdminRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetPortfolioId(v string) *ListProductsAsAdminRequest {
	s.PortfolioId = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetScope(v string) *ListProductsAsAdminRequest {
	s.Scope = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetSortBy(v string) *ListProductsAsAdminRequest {
	s.SortBy = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetSortOrder(v string) *ListProductsAsAdminRequest {
	s.SortOrder = &v
	return s
}

func (s *ListProductsAsAdminRequest) Validate() error {
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

type ListProductsAsAdminRequestFilters struct {
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

func (s ListProductsAsAdminRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsAdminRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListProductsAsAdminRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListProductsAsAdminRequestFilters) SetKey(v string) *ListProductsAsAdminRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProductsAsAdminRequestFilters) SetValue(v string) *ListProductsAsAdminRequestFilters {
	s.Value = &v
	return s
}

func (s *ListProductsAsAdminRequestFilters) Validate() error {
	return dara.Validate(s)
}
