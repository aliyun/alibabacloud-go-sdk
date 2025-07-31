// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAnswerSampleByPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *QueryAnswerSampleByPageShrinkRequest
	GetAnswer() *string
	SetCurrentPage(v int32) *QueryAnswerSampleByPageShrinkRequest
	GetCurrentPage() *int32
	SetLibId(v string) *QueryAnswerSampleByPageShrinkRequest
	GetLibId() *string
	SetPageSize(v int32) *QueryAnswerSampleByPageShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *QueryAnswerSampleByPageShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *QueryAnswerSampleByPageShrinkRequest
	GetSortShrink() *string
}

type QueryAnswerSampleByPageShrinkRequest struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
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
}

func (s QueryAnswerSampleByPageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAnswerSampleByPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageShrinkRequest) GetAnswer() *string {
	return s.Answer
}

func (s *QueryAnswerSampleByPageShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryAnswerSampleByPageShrinkRequest) GetLibId() *string {
	return s.LibId
}

func (s *QueryAnswerSampleByPageShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAnswerSampleByPageShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryAnswerSampleByPageShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetAnswer(v string) *QueryAnswerSampleByPageShrinkRequest {
	s.Answer = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetCurrentPage(v int32) *QueryAnswerSampleByPageShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetLibId(v string) *QueryAnswerSampleByPageShrinkRequest {
	s.LibId = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetPageSize(v int32) *QueryAnswerSampleByPageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetRegionId(v string) *QueryAnswerSampleByPageShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetSortShrink(v string) *QueryAnswerSampleByPageShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
