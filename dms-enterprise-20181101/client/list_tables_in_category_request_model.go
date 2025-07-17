// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesInCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *ListTablesInCategoryRequest
	GetCategoryId() *int64
	SetPageNumber(v int32) *ListTablesInCategoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTablesInCategoryRequest
	GetPageSize() *int32
	SetTid(v int64) *ListTablesInCategoryRequest
	GetTid() *int64
}

type ListTablesInCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30000235593
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTablesInCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesInCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListTablesInCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ListTablesInCategoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTablesInCategoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesInCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTablesInCategoryRequest) SetCategoryId(v int64) *ListTablesInCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *ListTablesInCategoryRequest) SetPageNumber(v int32) *ListTablesInCategoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTablesInCategoryRequest) SetPageSize(v int32) *ListTablesInCategoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablesInCategoryRequest) SetTid(v int64) *ListTablesInCategoryRequest {
	s.Tid = &v
	return s
}

func (s *ListTablesInCategoryRequest) Validate() error {
	return dara.Validate(s)
}
