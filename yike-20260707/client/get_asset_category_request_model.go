// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *GetAssetCategoryRequest
	GetCategoryId() *int64
}

type GetAssetCategoryRequest struct {
	// The category ID. You can obtain the category ID by using the following methods:
	//
	// - When you create a category by calling the CreateAssetCategory operation, the category ID is the value of CategoryId in the response.
	//
	// - When you query the category list by calling the ListAssetCategories operation, the category ID is the value of CategoryId in the corresponding entry in the response.
	//
	// example:
	//
	// scenery
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s GetAssetCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetAssetCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetAssetCategoryRequest) SetCategoryId(v int64) *GetAssetCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *GetAssetCategoryRequest) Validate() error {
	return dara.Validate(s)
}
