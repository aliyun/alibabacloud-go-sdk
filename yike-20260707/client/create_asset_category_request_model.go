// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *CreateAssetCategoryRequest
	GetCategoryName() *string
	SetParentId(v int64) *CreateAssetCategoryRequest
	GetParentId() *int64
}

type CreateAssetCategoryRequest struct {
	// The category name.
	//
	// Maximum length: 64 bytes.
	//
	// UTF-8 encoding.
	//
	// This parameter is required.
	//
	// example:
	//
	// Third-level subcategory
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The parent category ID.
	//
	// example:
	//
	// 5
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s CreateAssetCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateAssetCategoryRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *CreateAssetCategoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateAssetCategoryRequest) SetCategoryName(v string) *CreateAssetCategoryRequest {
	s.CategoryName = &v
	return s
}

func (s *CreateAssetCategoryRequest) SetParentId(v int64) *CreateAssetCategoryRequest {
	s.ParentId = &v
	return s
}

func (s *CreateAssetCategoryRequest) Validate() error {
	return dara.Validate(s)
}
