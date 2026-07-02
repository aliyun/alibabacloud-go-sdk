// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentReadSummaryOption interface {
	dara.Model
	String() string
	GoString() string
	SetChapterSummarize(v bool) *DocumentReadSummaryOption
	GetChapterSummarize() *bool
	SetChapterSummarizeOption(v *DocumentChapterSummarizeOption) *DocumentReadSummaryOption
	GetChapterSummarizeOption() *DocumentChapterSummarizeOption
	SetSummarize(v bool) *DocumentReadSummaryOption
	GetSummarize() *bool
}

type DocumentReadSummaryOption struct {
	// Specifies whether to extract the chapter-level summary of the article.
	//
	// example:
	//
	// true
	ChapterSummarize *bool `json:"ChapterSummarize,omitempty" xml:"ChapterSummarize,omitempty"`
	// The chapter-level summary options for the article.
	ChapterSummarizeOption *DocumentChapterSummarizeOption `json:"ChapterSummarizeOption,omitempty" xml:"ChapterSummarizeOption,omitempty"`
	// Specifies whether to extract the article summary.
	//
	// example:
	//
	// true
	Summarize *bool `json:"Summarize,omitempty" xml:"Summarize,omitempty"`
}

func (s DocumentReadSummaryOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentReadSummaryOption) GoString() string {
	return s.String()
}

func (s *DocumentReadSummaryOption) GetChapterSummarize() *bool {
	return s.ChapterSummarize
}

func (s *DocumentReadSummaryOption) GetChapterSummarizeOption() *DocumentChapterSummarizeOption {
	return s.ChapterSummarizeOption
}

func (s *DocumentReadSummaryOption) GetSummarize() *bool {
	return s.Summarize
}

func (s *DocumentReadSummaryOption) SetChapterSummarize(v bool) *DocumentReadSummaryOption {
	s.ChapterSummarize = &v
	return s
}

func (s *DocumentReadSummaryOption) SetChapterSummarizeOption(v *DocumentChapterSummarizeOption) *DocumentReadSummaryOption {
	s.ChapterSummarizeOption = v
	return s
}

func (s *DocumentReadSummaryOption) SetSummarize(v bool) *DocumentReadSummaryOption {
	s.Summarize = &v
	return s
}

func (s *DocumentReadSummaryOption) Validate() error {
	if s.ChapterSummarizeOption != nil {
		if err := s.ChapterSummarizeOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}
