// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *ListProjectFeaturesRequest
	GetAliasName() *string
	SetFilter(v string) *ListProjectFeaturesRequest
	GetFilter() *string
	SetName(v string) *ListProjectFeaturesRequest
	GetName() *string
	SetOrder(v string) *ListProjectFeaturesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListProjectFeaturesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectFeaturesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListProjectFeaturesRequest
	GetSortBy() *string
}

type ListProjectFeaturesRequest struct {
	// example:
	//
	// ff1
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// {"feature_view_name":"fv1"}
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
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
	// ModelFeatureCount
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListProjectFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectFeaturesRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *ListProjectFeaturesRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListProjectFeaturesRequest) GetName() *string {
	return s.Name
}

func (s *ListProjectFeaturesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListProjectFeaturesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectFeaturesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectFeaturesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProjectFeaturesRequest) SetAliasName(v string) *ListProjectFeaturesRequest {
	s.AliasName = &v
	return s
}

func (s *ListProjectFeaturesRequest) SetFilter(v string) *ListProjectFeaturesRequest {
	s.Filter = &v
	return s
}

func (s *ListProjectFeaturesRequest) SetName(v string) *ListProjectFeaturesRequest {
	s.Name = &v
	return s
}

func (s *ListProjectFeaturesRequest) SetOrder(v string) *ListProjectFeaturesRequest {
	s.Order = &v
	return s
}

func (s *ListProjectFeaturesRequest) SetPageNumber(v int32) *ListProjectFeaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectFeaturesRequest) SetPageSize(v int32) *ListProjectFeaturesRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectFeaturesRequest) SetSortBy(v string) *ListProjectFeaturesRequest {
	s.SortBy = &v
	return s
}

func (s *ListProjectFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
