// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeywordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListKeywordsShrinkRequest
	GetCurrentPage() *int32
	SetLibId(v string) *ListKeywordsShrinkRequest
	GetLibId() *string
	SetPageSize(v int32) *ListKeywordsShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListKeywordsShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *ListKeywordsShrinkRequest
	GetSortShrink() *string
	SetWord(v string) *ListKeywordsShrinkRequest
	GetWord() *string
}

type ListKeywordsShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Word       *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ListKeywordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListKeywordsShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListKeywordsShrinkRequest) GetLibId() *string {
	return s.LibId
}

func (s *ListKeywordsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKeywordsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKeywordsShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *ListKeywordsShrinkRequest) GetWord() *string {
	return s.Word
}

func (s *ListKeywordsShrinkRequest) SetCurrentPage(v int32) *ListKeywordsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetLibId(v string) *ListKeywordsShrinkRequest {
	s.LibId = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetPageSize(v int32) *ListKeywordsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetRegionId(v string) *ListKeywordsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetSortShrink(v string) *ListKeywordsShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetWord(v string) *ListKeywordsShrinkRequest {
	s.Word = &v
	return s
}

func (s *ListKeywordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
