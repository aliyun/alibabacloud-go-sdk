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
	// - It can be up to 64 bytes in length.
	//
	// - It is encoded in UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// Comedy
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The ID of the parent category.
	//
	// Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Asset Management Configuration*	- > **Category Management*	- > **Audio/Video/Image Categories*	- or **Short Video Material Categories*	- to view the category ID.
	//
	// > - If you specify this parameter, a subcategory is created under the specified parent category. If you do not specify this parameter, a level-1 category is created.
	//
	// > - All level-1 categories for short video materials are built-in and cannot be modified, added, or deleted. You can only create subcategories under level-1 categories. Therefore, this parameter is required when `Type` is set to `material`.
	//
	// example:
	//
	// 100012****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The category type. Valid values:
	//
	// - **default**: audio, video, and image categories. This is the default value.
	//
	// - **material**: short video material categories.
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
