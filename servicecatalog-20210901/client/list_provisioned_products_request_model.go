// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLevelFilter(v string) *ListProvisionedProductsRequest
	GetAccessLevelFilter() *string
	SetFilters(v []*ListProvisionedProductsRequestFilters) *ListProvisionedProductsRequest
	GetFilters() []*ListProvisionedProductsRequestFilters
	SetPageNumber(v int32) *ListProvisionedProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProvisionedProductsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListProvisionedProductsRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListProvisionedProductsRequest
	GetSortOrder() *string
}

type ListProvisionedProductsRequest struct {
	// The access filter. Valid values:
	//
	// 	- User: queries the product instances that are created by the current requester. This is the default value.
	//
	// 	- Account: queries the product instances that belong to the current Alibaba Cloud account.
	//
	// example:
	//
	// User
	AccessLevelFilter *string `json:"AccessLevelFilter,omitempty" xml:"AccessLevelFilter,omitempty"`
	// The filter conditions.
	Filters []*ListProvisionedProductsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
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
	// The field that is used to sort the queried data.
	//
	// Set the value to CreateTime, which specifies the time when the product instance was created.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting method. Valid values:
	//
	// 	- Asc: the ascending order.
	//
	// 	- Desc (default): the descending order.
	//
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProvisionedProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductsRequest) GetAccessLevelFilter() *string {
	return s.AccessLevelFilter
}

func (s *ListProvisionedProductsRequest) GetFilters() []*ListProvisionedProductsRequestFilters {
	return s.Filters
}

func (s *ListProvisionedProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProvisionedProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProvisionedProductsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProvisionedProductsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListProvisionedProductsRequest) SetAccessLevelFilter(v string) *ListProvisionedProductsRequest {
	s.AccessLevelFilter = &v
	return s
}

func (s *ListProvisionedProductsRequest) SetFilters(v []*ListProvisionedProductsRequestFilters) *ListProvisionedProductsRequest {
	s.Filters = v
	return s
}

func (s *ListProvisionedProductsRequest) SetPageNumber(v int32) *ListProvisionedProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProvisionedProductsRequest) SetPageSize(v int32) *ListProvisionedProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProvisionedProductsRequest) SetSortBy(v string) *ListProvisionedProductsRequest {
	s.SortBy = &v
	return s
}

func (s *ListProvisionedProductsRequest) SetSortOrder(v string) *ListProvisionedProductsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListProvisionedProductsRequest) Validate() error {
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

type ListProvisionedProductsRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// 	- ProvisionedProductName: performs exact matches by product instance name. Product instance names are not case-sensitive.
	//
	// 	- FullTextSearch: performs full-text searches by product instance name. Fuzzy match is supported.
	//
	// example:
	//
	// ProvisionedProductName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// DEMO-ECS instance
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProvisionedProductsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListProvisionedProductsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListProvisionedProductsRequestFilters) SetKey(v string) *ListProvisionedProductsRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProvisionedProductsRequestFilters) SetValue(v string) *ListProvisionedProductsRequestFilters {
	s.Value = &v
	return s
}

func (s *ListProvisionedProductsRequestFilters) Validate() error {
	return dara.Validate(s)
}
