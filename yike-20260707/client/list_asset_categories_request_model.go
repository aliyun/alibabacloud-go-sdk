// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListAssetCategoriesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListAssetCategoriesRequest
	GetPageSize() *int32
}

type ListAssetCategoriesRequest struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAssetCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAssetCategoriesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListAssetCategoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAssetCategoriesRequest) SetPageNo(v int32) *ListAssetCategoriesRequest {
	s.PageNo = &v
	return s
}

func (s *ListAssetCategoriesRequest) SetPageSize(v int32) *ListAssetCategoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAssetCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
