// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateName(v string) *AddCategoryRequest
	GetCateName() *string
	SetParentId(v int64) *AddCategoryRequest
	GetParentId() *int64
	SetType(v string) *AddCategoryRequest
	GetType() *string
}

type AddCategoryRequest struct {
	// The category name.
	//
	// - The maximum length is 64 bytes.
	//
	// - UTF-8 encoding.
	//
	// This parameter is required.
	//
	// example:
	//
	// Third-level subcategory
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The parent category ID.
	//
	// example:
	//
	// 5
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The category type. Valid values:
	//
	// - default (default): audio, video, and image category.
	//
	// - material: short video material category.
	//
	// example:
	//
	// default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddCategoryRequest) GetCateName() *string {
	return s.CateName
}

func (s *AddCategoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *AddCategoryRequest) GetType() *string {
	return s.Type
}

func (s *AddCategoryRequest) SetCateName(v string) *AddCategoryRequest {
	s.CateName = &v
	return s
}

func (s *AddCategoryRequest) SetParentId(v int64) *AddCategoryRequest {
	s.ParentId = &v
	return s
}

func (s *AddCategoryRequest) SetType(v string) *AddCategoryRequest {
	s.Type = &v
	return s
}

func (s *AddCategoryRequest) Validate() error {
	return dara.Validate(s)
}
