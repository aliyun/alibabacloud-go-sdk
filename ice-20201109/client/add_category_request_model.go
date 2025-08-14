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
	// 	- The value can be up to 64 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// This parameter is required.
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The ID of the parent category.
	//
	// example:
	//
	// 5
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The type of the category. Valid values:
	//
	// 	- default: audio, video, and image files. This is the default value.
	//
	// 	- material: short video materials.
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
