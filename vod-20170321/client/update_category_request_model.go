// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *UpdateCategoryRequest
	GetCateId() *int64
	SetCateName(v string) *UpdateCategoryRequest
	GetCateName() *string
}

type UpdateCategoryRequest struct {
	// The category ID. Only a single category ID can be specified. You can obtain the category ID by using the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Asset Management Configuration*	- > **Category Management*	- > **Audio/Video/Image Category*	- or **Short Video Material Category*	- to view the category ID.
	//
	// - Obtain the category ID from the response of the [AddCategory](~~AddCategory~~) operation when you create a category.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10020****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// - The name can be up to 64 bytes in length.
	//
	// - The name must be encoded in UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// Landscape
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
}

func (s UpdateCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategoryRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *UpdateCategoryRequest) GetCateName() *string {
	return s.CateName
}

func (s *UpdateCategoryRequest) SetCateId(v int64) *UpdateCategoryRequest {
	s.CateId = &v
	return s
}

func (s *UpdateCategoryRequest) SetCateName(v string) *UpdateCategoryRequest {
	s.CateName = &v
	return s
}

func (s *UpdateCategoryRequest) Validate() error {
	return dara.Validate(s)
}
