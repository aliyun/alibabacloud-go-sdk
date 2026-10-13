// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListProductCategoriesRequest
	GetName() *string
	SetPageNumber(v int32) *ListProductCategoriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductCategoriesRequest
	GetPageSize() *int32
}

type ListProductCategoriesRequest struct {
	// The keyword for the category name. This parameter is used to search for product categories by name.
	//
	// example:
	//
	// ai-capabilities
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. The value starts from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. The minimum value is 1.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListProductCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListProductCategoriesRequest) GetName() *string {
	return s.Name
}

func (s *ListProductCategoriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductCategoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductCategoriesRequest) SetName(v string) *ListProductCategoriesRequest {
	s.Name = &v
	return s
}

func (s *ListProductCategoriesRequest) SetPageNumber(v int32) *ListProductCategoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductCategoriesRequest) SetPageSize(v int32) *ListProductCategoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
