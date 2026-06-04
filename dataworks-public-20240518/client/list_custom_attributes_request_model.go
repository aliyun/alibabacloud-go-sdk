// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListCustomAttributesRequest
	GetComment() *string
	SetDisplayName(v string) *ListCustomAttributesRequest
	GetDisplayName() *string
	SetEntityTypes(v string) *ListCustomAttributesRequest
	GetEntityTypes() *string
	SetOrder(v string) *ListCustomAttributesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListCustomAttributesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomAttributesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListCustomAttributesRequest
	GetSortBy() *string
}

type ListCustomAttributesRequest struct {
	// example:
	//
	// owner
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// Owner
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// maxcompute-table
	EntityTypes *string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty"`
	// example:
	//
	// Asc
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
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListCustomAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAttributesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAttributesRequest) GetComment() *string {
	return s.Comment
}

func (s *ListCustomAttributesRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListCustomAttributesRequest) GetEntityTypes() *string {
	return s.EntityTypes
}

func (s *ListCustomAttributesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListCustomAttributesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomAttributesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomAttributesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListCustomAttributesRequest) SetComment(v string) *ListCustomAttributesRequest {
	s.Comment = &v
	return s
}

func (s *ListCustomAttributesRequest) SetDisplayName(v string) *ListCustomAttributesRequest {
	s.DisplayName = &v
	return s
}

func (s *ListCustomAttributesRequest) SetEntityTypes(v string) *ListCustomAttributesRequest {
	s.EntityTypes = &v
	return s
}

func (s *ListCustomAttributesRequest) SetOrder(v string) *ListCustomAttributesRequest {
	s.Order = &v
	return s
}

func (s *ListCustomAttributesRequest) SetPageNumber(v int32) *ListCustomAttributesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomAttributesRequest) SetPageSize(v int32) *ListCustomAttributesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomAttributesRequest) SetSortBy(v string) *ListCustomAttributesRequest {
	s.SortBy = &v
	return s
}

func (s *ListCustomAttributesRequest) Validate() error {
	return dara.Validate(s)
}
