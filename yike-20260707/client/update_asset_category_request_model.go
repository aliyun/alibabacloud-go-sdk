// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *UpdateAssetCategoryRequest
	GetCategoryId() *int64
	SetCategoryName(v string) *UpdateAssetCategoryRequest
	GetCategoryName() *string
}

type UpdateAssetCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 50
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scenery
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s UpdateAssetCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssetCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateAssetCategoryRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *UpdateAssetCategoryRequest) SetCategoryId(v int64) *UpdateAssetCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateAssetCategoryRequest) SetCategoryName(v string) *UpdateAssetCategoryRequest {
	s.CategoryName = &v
	return s
}

func (s *UpdateAssetCategoryRequest) Validate() error {
	return dara.Validate(s)
}
