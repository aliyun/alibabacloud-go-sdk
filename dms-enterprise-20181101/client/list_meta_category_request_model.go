// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListMetaCategoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMetaCategoryRequest
	GetPageSize() *int32
	SetParentCategoryId(v int64) *ListMetaCategoryRequest
	GetParentCategoryId() *int64
	SetTid(v int64) *ListMetaCategoryRequest
	GetTid() *int64
}

type ListMetaCategoryRequest struct {
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
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 30000322682
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListMetaCategoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMetaCategoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaCategoryRequest) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *ListMetaCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListMetaCategoryRequest) SetPageNumber(v int32) *ListMetaCategoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMetaCategoryRequest) SetPageSize(v int32) *ListMetaCategoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaCategoryRequest) SetParentCategoryId(v int64) *ListMetaCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *ListMetaCategoryRequest) SetTid(v int64) *ListMetaCategoryRequest {
	s.Tid = &v
	return s
}

func (s *ListMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
