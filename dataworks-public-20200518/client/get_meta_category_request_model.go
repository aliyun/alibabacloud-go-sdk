// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *GetMetaCategoryRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetMetaCategoryRequest
	GetPageSize() *int32
	SetParentCategoryId(v int64) *GetMetaCategoryRequest
	GetParentCategoryId() *int64
}

type GetMetaCategoryRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The category tree ID.
	//
	// example:
	//
	// 333
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s GetMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetMetaCategoryRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMetaCategoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaCategoryRequest) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *GetMetaCategoryRequest) SetPageNum(v int32) *GetMetaCategoryRequest {
	s.PageNum = &v
	return s
}

func (s *GetMetaCategoryRequest) SetPageSize(v int32) *GetMetaCategoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaCategoryRequest) SetParentCategoryId(v int64) *GetMetaCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *GetMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
