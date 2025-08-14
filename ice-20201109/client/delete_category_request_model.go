// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *DeleteCategoryRequest
	GetCateId() *int64
}

type DeleteCategoryRequest struct {
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
	// 46
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
}

func (s DeleteCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteCategoryRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *DeleteCategoryRequest) SetCateId(v int64) *DeleteCategoryRequest {
	s.CateId = &v
	return s
}

func (s *DeleteCategoryRequest) Validate() error {
	return dara.Validate(s)
}
