// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int32) *ListProductsRequest
	GetOffset() *int32
	SetPageNum(v int32) *ListProductsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListProductsRequest
	GetPageSize() *int32
	SetProductName(v string) *ListProductsRequest
	GetProductName() *string
	SetSearchKeyWord(v string) *ListProductsRequest
	GetSearchKeyWord() *string
	SetSimple(v bool) *ListProductsRequest
	GetSimple() *bool
	SetSize(v int32) *ListProductsRequest
	GetSize() *int32
}

type ListProductsRequest struct {
	// example:
	//
	// 1
	Offset        *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageNum       *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductName   *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// example:
	//
	// false
	Simple *bool `json:"Simple,omitempty" xml:"Simple,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *ListProductsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductsRequest) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsRequest) GetSearchKeyWord() *string {
	return s.SearchKeyWord
}

func (s *ListProductsRequest) GetSimple() *bool {
	return s.Simple
}

func (s *ListProductsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListProductsRequest) SetOffset(v int32) *ListProductsRequest {
	s.Offset = &v
	return s
}

func (s *ListProductsRequest) SetPageNum(v int32) *ListProductsRequest {
	s.PageNum = &v
	return s
}

func (s *ListProductsRequest) SetPageSize(v int32) *ListProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductsRequest) SetProductName(v string) *ListProductsRequest {
	s.ProductName = &v
	return s
}

func (s *ListProductsRequest) SetSearchKeyWord(v string) *ListProductsRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *ListProductsRequest) SetSimple(v bool) *ListProductsRequest {
	s.Simple = &v
	return s
}

func (s *ListProductsRequest) SetSize(v int32) *ListProductsRequest {
	s.Size = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
