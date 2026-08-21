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
	// - Maximum length: 64 bytes.
	//
	// - UTF-8 encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// Comedy
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The parent category ID.
	//
	// Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Management Configuration*	- > **Category Management*	- > **Audio/Video/Image Categories*	- or **Short Video Material Categories*	- to view category IDs.
	//
	// > - If you specify this parameter, a subcategory is created under the specified parent category. If you do not specify this parameter, a level-0 category is created.
	//
	// > - Because all level-0 categories for short video materials are built-in and cannot be modified, added, or deleted, only subcategories can be created under level-0 categories. Therefore, this parameter is required when `Type` is set to `material`.
	//
	// example:
	//
	// 100012****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The category type. Valid values:
	//
	// - **default*	- (default): audio/video/image category.
	//
	// - **material**: short video material category.
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
