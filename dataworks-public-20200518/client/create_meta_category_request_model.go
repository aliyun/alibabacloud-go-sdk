// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateMetaCategoryRequest
	GetComment() *string
	SetName(v string) *CreateMetaCategoryRequest
	GetName() *string
	SetParentId(v int64) *CreateMetaCategoryRequest
	GetParentId() *int64
}

type CreateMetaCategoryRequest struct {
	// The remarks of the category.
	//
	// example:
	//
	// category 1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the category.
	//
	// This parameter is required.
	//
	// example:
	//
	// category_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the parent category.
	//
	// example:
	//
	// 0
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s CreateMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaCategoryRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateMetaCategoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetaCategoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateMetaCategoryRequest) SetComment(v string) *CreateMetaCategoryRequest {
	s.Comment = &v
	return s
}

func (s *CreateMetaCategoryRequest) SetName(v string) *CreateMetaCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateMetaCategoryRequest) SetParentId(v int64) *CreateMetaCategoryRequest {
	s.ParentId = &v
	return s
}

func (s *CreateMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
