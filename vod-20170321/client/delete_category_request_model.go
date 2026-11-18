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
	// The category ID. You can specify only one category ID. You can obtain the category ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration management*	- > **Media asset management configuration*	- > **Category management*	- to view the category ID.
	//
	// - Call the [AddCategory](~~AddCategory~~) operation to create a category. The category ID is returned in the response.
	//
	// > If the specified category ID is the ID of a parent category, the parent category and all its subcategories are deleted. Proceed with caution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3300****
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
