// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAssetCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *DeleteAssetCategoryRequest
	GetCategoryId() *int64
}

type DeleteAssetCategoryRequest struct {
	// The category ID. You can obtain the value by using one of the following methods:
	//
	// - When you create a category by calling the CreateAssetCategory operation, the category ID is the value of CategoryId in the response.
	//
	// - When you query categories by calling the ListAssetCategories operation, the category ID is the value of CategoryId in the corresponding entry in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DeleteAssetCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAssetCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteAssetCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DeleteAssetCategoryRequest) SetCategoryId(v int64) *DeleteAssetCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *DeleteAssetCategoryRequest) Validate() error {
	return dara.Validate(s)
}
