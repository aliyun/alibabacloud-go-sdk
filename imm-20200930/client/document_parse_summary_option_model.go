// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentParseSummaryOption interface {
	dara.Model
	String() string
	GoString() string
	SetChapterSummarize(v bool) *DocumentParseSummaryOption
	GetChapterSummarize() *bool
	SetSummarize(v bool) *DocumentParseSummaryOption
	GetSummarize() *bool
}

type DocumentParseSummaryOption struct {
	ChapterSummarize *bool `json:"ChapterSummarize,omitempty" xml:"ChapterSummarize,omitempty"`
	Summarize        *bool `json:"Summarize,omitempty" xml:"Summarize,omitempty"`
}

func (s DocumentParseSummaryOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseSummaryOption) GoString() string {
	return s.String()
}

func (s *DocumentParseSummaryOption) GetChapterSummarize() *bool {
	return s.ChapterSummarize
}

func (s *DocumentParseSummaryOption) GetSummarize() *bool {
	return s.Summarize
}

func (s *DocumentParseSummaryOption) SetChapterSummarize(v bool) *DocumentParseSummaryOption {
	s.ChapterSummarize = &v
	return s
}

func (s *DocumentParseSummaryOption) SetSummarize(v bool) *DocumentParseSummaryOption {
	s.Summarize = &v
	return s
}

func (s *DocumentParseSummaryOption) Validate() error {
	return dara.Validate(s)
}
