// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentAnalysisOption interface {
	dara.Model
	String() string
	GoString() string
	SetChapterSummary(v bool) *GetDocumentAnalysisOption
	GetChapterSummary() *bool
	SetChapterSummaryOption(v *DocumentChapterSummarizeOption) *GetDocumentAnalysisOption
	GetChapterSummaryOption() *DocumentChapterSummarizeOption
	SetImages(v bool) *GetDocumentAnalysisOption
	GetImages() *bool
	SetKeyword(v bool) *GetDocumentAnalysisOption
	GetKeyword() *bool
	SetLayouts(v bool) *GetDocumentAnalysisOption
	GetLayouts() *bool
	SetNarrator(v bool) *GetDocumentAnalysisOption
	GetNarrator() *bool
	SetQuestion(v bool) *GetDocumentAnalysisOption
	GetQuestion() *bool
	SetSummary(v bool) *GetDocumentAnalysisOption
	GetSummary() *bool
}

type GetDocumentAnalysisOption struct {
	// Specifies whether to generate summaries for each chapter. Set to `true` to enable this feature. Use `ChapterSummaryOption` to configure detailed options. Defaults to `false`.
	ChapterSummary       *bool                           `json:"ChapterSummary,omitempty" xml:"ChapterSummary,omitempty"`
	ChapterSummaryOption *DocumentChapterSummarizeOption `json:"ChapterSummaryOption,omitempty" xml:"ChapterSummaryOption,omitempty"`
	// Specifies whether to analyze each image within the document. Set to `true` to enable this feature. Defaults to `false`.
	Images *bool `json:"Images,omitempty" xml:"Images,omitempty"`
	// Specifies whether to extract keywords from the document. Set to `true` to enable this feature. Defaults to `false`.
	Keyword *bool `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Specifies whether to analyze the document layout. Set to `true` to enable this feature. Defaults to `false`.
	Layouts *bool `json:"Layouts,omitempty" xml:"Layouts,omitempty"`
	// Specifies whether to identify the narrator in the document. Set to `true` to enable this feature. Defaults to `false`.
	Narrator *bool `json:"Narrator,omitempty" xml:"Narrator,omitempty"`
	// Specifies whether to generate questions based on the document content. Set to `true` to enable this feature. Defaults to `false`.
	Question *bool `json:"Question,omitempty" xml:"Question,omitempty"`
	// Specifies whether to generate a summary for the document. Set to `true` to enable this feature. Defaults to `false`.
	Summary *bool `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetDocumentAnalysisOption) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalysisOption) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalysisOption) GetChapterSummary() *bool {
	return s.ChapterSummary
}

func (s *GetDocumentAnalysisOption) GetChapterSummaryOption() *DocumentChapterSummarizeOption {
	return s.ChapterSummaryOption
}

func (s *GetDocumentAnalysisOption) GetImages() *bool {
	return s.Images
}

func (s *GetDocumentAnalysisOption) GetKeyword() *bool {
	return s.Keyword
}

func (s *GetDocumentAnalysisOption) GetLayouts() *bool {
	return s.Layouts
}

func (s *GetDocumentAnalysisOption) GetNarrator() *bool {
	return s.Narrator
}

func (s *GetDocumentAnalysisOption) GetQuestion() *bool {
	return s.Question
}

func (s *GetDocumentAnalysisOption) GetSummary() *bool {
	return s.Summary
}

func (s *GetDocumentAnalysisOption) SetChapterSummary(v bool) *GetDocumentAnalysisOption {
	s.ChapterSummary = &v
	return s
}

func (s *GetDocumentAnalysisOption) SetChapterSummaryOption(v *DocumentChapterSummarizeOption) *GetDocumentAnalysisOption {
	s.ChapterSummaryOption = v
	return s
}

func (s *GetDocumentAnalysisOption) SetImages(v bool) *GetDocumentAnalysisOption {
	s.Images = &v
	return s
}

func (s *GetDocumentAnalysisOption) SetKeyword(v bool) *GetDocumentAnalysisOption {
	s.Keyword = &v
	return s
}

func (s *GetDocumentAnalysisOption) SetLayouts(v bool) *GetDocumentAnalysisOption {
	s.Layouts = &v
	return s
}

func (s *GetDocumentAnalysisOption) SetNarrator(v bool) *GetDocumentAnalysisOption {
	s.Narrator = &v
	return s
}

func (s *GetDocumentAnalysisOption) SetQuestion(v bool) *GetDocumentAnalysisOption {
	s.Question = &v
	return s
}

func (s *GetDocumentAnalysisOption) SetSummary(v bool) *GetDocumentAnalysisOption {
	s.Summary = &v
	return s
}

func (s *GetDocumentAnalysisOption) Validate() error {
	if s.ChapterSummaryOption != nil {
		if err := s.ChapterSummaryOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}
