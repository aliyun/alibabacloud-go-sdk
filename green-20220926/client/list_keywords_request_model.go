// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeywordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListKeywordsRequest
	GetCurrentPage() *int32
	SetLibId(v string) *ListKeywordsRequest
	GetLibId() *string
	SetPageSize(v int32) *ListKeywordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListKeywordsRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *ListKeywordsRequest
	GetSort() map[string]*string
	SetWord(v string) *ListKeywordsRequest
	GetWord() *string
}

type ListKeywordsRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Library ID.
	//
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Sort field.
	Sort map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Keyword.
	//
	// example:
	//
	// 测试词
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ListKeywordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordsRequest) GoString() string {
	return s.String()
}

func (s *ListKeywordsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListKeywordsRequest) GetLibId() *string {
	return s.LibId
}

func (s *ListKeywordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKeywordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKeywordsRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *ListKeywordsRequest) GetWord() *string {
	return s.Word
}

func (s *ListKeywordsRequest) SetCurrentPage(v int32) *ListKeywordsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListKeywordsRequest) SetLibId(v string) *ListKeywordsRequest {
	s.LibId = &v
	return s
}

func (s *ListKeywordsRequest) SetPageSize(v int32) *ListKeywordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeywordsRequest) SetRegionId(v string) *ListKeywordsRequest {
	s.RegionId = &v
	return s
}

func (s *ListKeywordsRequest) SetSort(v map[string]*string) *ListKeywordsRequest {
	s.Sort = v
	return s
}

func (s *ListKeywordsRequest) SetWord(v string) *ListKeywordsRequest {
	s.Word = &v
	return s
}

func (s *ListKeywordsRequest) Validate() error {
	return dara.Validate(s)
}
