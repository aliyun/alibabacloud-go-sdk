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
	// The category ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [Intelligent Media Services (IMS) console](https://ims.console.aliyun.com) and choose **Media Asset Management*	- > **Category Management*	- to view the category ID.
	//
	// 	- View the value of CateId returned by the AddCategory operation that you called to create a category.
	//
	// 	- View the value of CateId returned by the GetCategories operation that you called to query a category.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// This parameter is required.
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
