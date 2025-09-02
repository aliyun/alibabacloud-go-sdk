// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *UpdateMetaCategoryRequest
	GetCategoryId() *int64
	SetComment(v string) *UpdateMetaCategoryRequest
	GetComment() *string
	SetName(v string) *UpdateMetaCategoryRequest
	GetName() *string
}

type UpdateMetaCategoryRequest struct {
	// The ID of the category.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The remarks of the category.
	//
	// example:
	//
	// category name
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// category name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateMetaCategoryRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateMetaCategoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMetaCategoryRequest) SetCategoryId(v int64) *UpdateMetaCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateMetaCategoryRequest) SetComment(v string) *UpdateMetaCategoryRequest {
	s.Comment = &v
	return s
}

func (s *UpdateMetaCategoryRequest) SetName(v string) *UpdateMetaCategoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
