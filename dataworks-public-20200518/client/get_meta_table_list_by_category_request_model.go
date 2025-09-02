// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableListByCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *GetMetaTableListByCategoryRequest
	GetCategoryId() *int64
	SetPageNumber(v int32) *GetMetaTableListByCategoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetMetaTableListByCategoryRequest
	GetPageSize() *int32
}

type GetMetaTableListByCategoryRequest struct {
	// The category ID. You can call the [GetMetaCategory](https://help.aliyun.com/document_detail/2780099.html) operation to obtain the ID of the category. Categories allow you to efficiently organize and manage tables by category. You can search for the desired table by category.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetMetaTableListByCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableListByCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableListByCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetMetaTableListByCategoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTableListByCategoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableListByCategoryRequest) SetCategoryId(v int64) *GetMetaTableListByCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *GetMetaTableListByCategoryRequest) SetPageNumber(v int32) *GetMetaTableListByCategoryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTableListByCategoryRequest) SetPageSize(v int32) *GetMetaTableListByCategoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableListByCategoryRequest) Validate() error {
	return dara.Validate(s)
}
