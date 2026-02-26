// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentParseOption interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v *DocumentParseKeywordOption) *DocumentParseOption
	GetKeyword() *DocumentParseKeywordOption
	SetNarrator(v *DocumentParseNarratorOption) *DocumentParseOption
	GetNarrator() *DocumentParseNarratorOption
	SetQuestion(v *DocumentParseQuestionOption) *DocumentParseOption
	GetQuestion() *DocumentParseQuestionOption
	SetSummary(v *DocumentParseSummaryOption) *DocumentParseOption
	GetSummary() *DocumentParseSummaryOption
}

type DocumentParseOption struct {
	Keyword  *DocumentParseKeywordOption  `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Narrator *DocumentParseNarratorOption `json:"Narrator,omitempty" xml:"Narrator,omitempty"`
	Question *DocumentParseQuestionOption `json:"Question,omitempty" xml:"Question,omitempty"`
	Summary  *DocumentParseSummaryOption  `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s DocumentParseOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseOption) GoString() string {
	return s.String()
}

func (s *DocumentParseOption) GetKeyword() *DocumentParseKeywordOption {
	return s.Keyword
}

func (s *DocumentParseOption) GetNarrator() *DocumentParseNarratorOption {
	return s.Narrator
}

func (s *DocumentParseOption) GetQuestion() *DocumentParseQuestionOption {
	return s.Question
}

func (s *DocumentParseOption) GetSummary() *DocumentParseSummaryOption {
	return s.Summary
}

func (s *DocumentParseOption) SetKeyword(v *DocumentParseKeywordOption) *DocumentParseOption {
	s.Keyword = v
	return s
}

func (s *DocumentParseOption) SetNarrator(v *DocumentParseNarratorOption) *DocumentParseOption {
	s.Narrator = v
	return s
}

func (s *DocumentParseOption) SetQuestion(v *DocumentParseQuestionOption) *DocumentParseOption {
	s.Question = v
	return s
}

func (s *DocumentParseOption) SetSummary(v *DocumentParseSummaryOption) *DocumentParseOption {
	s.Summary = v
	return s
}

func (s *DocumentParseOption) Validate() error {
	if s.Keyword != nil {
		if err := s.Keyword.Validate(); err != nil {
			return err
		}
	}
	if s.Narrator != nil {
		if err := s.Narrator.Validate(); err != nil {
			return err
		}
	}
	if s.Question != nil {
		if err := s.Question.Validate(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}
