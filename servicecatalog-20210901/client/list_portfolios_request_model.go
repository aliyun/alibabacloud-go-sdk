// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortfoliosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*ListPortfoliosRequestFilters) *ListPortfoliosRequest
	GetFilters() []*ListPortfoliosRequestFilters
	SetPageNumber(v int32) *ListPortfoliosRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPortfoliosRequest
	GetPageSize() *int32
	SetProductId(v string) *ListPortfoliosRequest
	GetProductId() *string
	SetScope(v string) *ListPortfoliosRequest
	GetScope() *string
	SetSortBy(v string) *ListPortfoliosRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListPortfoliosRequest
	GetSortOrder() *string
}

type ListPortfoliosRequest struct {
	Filters []*ListPortfoliosRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// Local
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListPortfoliosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPortfoliosRequest) GoString() string {
	return s.String()
}

func (s *ListPortfoliosRequest) GetFilters() []*ListPortfoliosRequestFilters {
	return s.Filters
}

func (s *ListPortfoliosRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPortfoliosRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPortfoliosRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListPortfoliosRequest) GetScope() *string {
	return s.Scope
}

func (s *ListPortfoliosRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListPortfoliosRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListPortfoliosRequest) SetFilters(v []*ListPortfoliosRequestFilters) *ListPortfoliosRequest {
	s.Filters = v
	return s
}

func (s *ListPortfoliosRequest) SetPageNumber(v int32) *ListPortfoliosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPortfoliosRequest) SetPageSize(v int32) *ListPortfoliosRequest {
	s.PageSize = &v
	return s
}

func (s *ListPortfoliosRequest) SetProductId(v string) *ListPortfoliosRequest {
	s.ProductId = &v
	return s
}

func (s *ListPortfoliosRequest) SetScope(v string) *ListPortfoliosRequest {
	s.Scope = &v
	return s
}

func (s *ListPortfoliosRequest) SetSortBy(v string) *ListPortfoliosRequest {
	s.SortBy = &v
	return s
}

func (s *ListPortfoliosRequest) SetSortOrder(v string) *ListPortfoliosRequest {
	s.SortOrder = &v
	return s
}

func (s *ListPortfoliosRequest) Validate() error {
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

type ListPortfoliosRequestFilters struct {
	// example:
	//
	// PortfolioName
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPortfoliosRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListPortfoliosRequestFilters) GoString() string {
	return s.String()
}

func (s *ListPortfoliosRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListPortfoliosRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListPortfoliosRequestFilters) SetKey(v string) *ListPortfoliosRequestFilters {
	s.Key = &v
	return s
}

func (s *ListPortfoliosRequestFilters) SetValue(v string) *ListPortfoliosRequestFilters {
	s.Value = &v
	return s
}

func (s *ListPortfoliosRequestFilters) Validate() error {
	return dara.Validate(s)
}
