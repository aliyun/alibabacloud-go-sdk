// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaEntityDefsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ListMetaEntityDefsRequest
	GetDescription() *string
	SetDisplayName(v string) *ListMetaEntityDefsRequest
	GetDisplayName() *string
	SetExtend(v string) *ListMetaEntityDefsRequest
	GetExtend() *string
	SetOrder(v string) *ListMetaEntityDefsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListMetaEntityDefsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMetaEntityDefsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListMetaEntityDefsRequest
	GetSortBy() *string
}

type ListMetaEntityDefsRequest struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// CustomReport
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// NONE
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
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

func (s ListMetaEntityDefsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntityDefsRequest) GoString() string {
	return s.String()
}

func (s *ListMetaEntityDefsRequest) GetDescription() *string {
	return s.Description
}

func (s *ListMetaEntityDefsRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMetaEntityDefsRequest) GetExtend() *string {
	return s.Extend
}

func (s *ListMetaEntityDefsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListMetaEntityDefsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMetaEntityDefsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaEntityDefsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListMetaEntityDefsRequest) SetDescription(v string) *ListMetaEntityDefsRequest {
	s.Description = &v
	return s
}

func (s *ListMetaEntityDefsRequest) SetDisplayName(v string) *ListMetaEntityDefsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListMetaEntityDefsRequest) SetExtend(v string) *ListMetaEntityDefsRequest {
	s.Extend = &v
	return s
}

func (s *ListMetaEntityDefsRequest) SetOrder(v string) *ListMetaEntityDefsRequest {
	s.Order = &v
	return s
}

func (s *ListMetaEntityDefsRequest) SetPageNumber(v int32) *ListMetaEntityDefsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMetaEntityDefsRequest) SetPageSize(v int32) *ListMetaEntityDefsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaEntityDefsRequest) SetSortBy(v string) *ListMetaEntityDefsRequest {
	s.SortBy = &v
	return s
}

func (s *ListMetaEntityDefsRequest) Validate() error {
	return dara.Validate(s)
}
