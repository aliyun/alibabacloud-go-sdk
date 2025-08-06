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
	// The ID of the category. You can specify only one ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). Choose **Configuration Management*	- > **Media Management*	- > **Categories**. On the **Audio and Video / Image Category*	- or **Short Video Material Category*	- tab, view the category ID.
	//
	// 	- Obtain the category ID from the response to the [AddCategory](~~AddCategory~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10020****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// 	- The value can be up to 64 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// beauty
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
