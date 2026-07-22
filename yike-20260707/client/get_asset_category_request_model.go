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
