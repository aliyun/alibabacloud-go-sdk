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
	// The category ID. You can obtain the category ID by using the following methods:
	//
	// - Log on to the [IMS console](https://ims.console.aliyun.com), and choose **Media Asset Management*	- > **Category Management*	- to view the category ID.
	//
	// - When you create a category by calling the create category operation, the category ID is the value of CateId in the response.
	//
	// - When you query a category by calling the get category operation, the category ID is the value of CateId in the response.
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
