// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAnswerSampleByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *QueryAnswerSampleByPageRequest
	GetAnswer() *string
	SetCurrentPage(v int32) *QueryAnswerSampleByPageRequest
	GetCurrentPage() *int32
	SetLibId(v string) *QueryAnswerSampleByPageRequest
	GetLibId() *string
	SetPageSize(v int32) *QueryAnswerSampleByPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *QueryAnswerSampleByPageRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *QueryAnswerSampleByPageRequest
	GetSort() map[string]*string
}

type QueryAnswerSampleByPageRequest struct {
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
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s QueryAnswerSampleByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAnswerSampleByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageRequest) GetAnswer() *string {
	return s.Answer
}

func (s *QueryAnswerSampleByPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryAnswerSampleByPageRequest) GetLibId() *string {
	return s.LibId
}

func (s *QueryAnswerSampleByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAnswerSampleByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryAnswerSampleByPageRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *QueryAnswerSampleByPageRequest) SetAnswer(v string) *QueryAnswerSampleByPageRequest {
	s.Answer = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetCurrentPage(v int32) *QueryAnswerSampleByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetLibId(v string) *QueryAnswerSampleByPageRequest {
	s.LibId = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetPageSize(v int32) *QueryAnswerSampleByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetRegionId(v string) *QueryAnswerSampleByPageRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetSort(v map[string]*string) *QueryAnswerSampleByPageRequest {
	s.Sort = v
	return s
}

func (s *QueryAnswerSampleByPageRequest) Validate() error {
	return dara.Validate(s)
}
