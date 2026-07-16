// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentReadOption interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v *DocumentReadKeywordOption) *DocumentReadOption
	GetKeyword() *DocumentReadKeywordOption
	SetNarrator(v *DocumentReadNarratorOption) *DocumentReadOption
	GetNarrator() *DocumentReadNarratorOption
	SetQuestion(v *DocumentReadQuestionOption) *DocumentReadOption
	GetQuestion() *DocumentReadQuestionOption
	SetSummary(v *DocumentReadSummaryOption) *DocumentReadOption
	GetSummary() *DocumentReadSummaryOption
}

type DocumentReadOption struct {
	// The document intensive reading keyword extraction options.
	Keyword *DocumentReadKeywordOption `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The document intensive reading guide options.
	Narrator *DocumentReadNarratorOption `json:"Narrator,omitempty" xml:"Narrator,omitempty"`
	// The document intensive reading question guide options.
	Question *DocumentReadQuestionOption `json:"Question,omitempty" xml:"Question,omitempty"`
	// The document intensive reading summary options.
	Summary *DocumentReadSummaryOption `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s DocumentReadOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentReadOption) GoString() string {
	return s.String()
}

func (s *DocumentReadOption) GetKeyword() *DocumentReadKeywordOption {
	return s.Keyword
}

func (s *DocumentReadOption) GetNarrator() *DocumentReadNarratorOption {
	return s.Narrator
}

func (s *DocumentReadOption) GetQuestion() *DocumentReadQuestionOption {
	return s.Question
}

func (s *DocumentReadOption) GetSummary() *DocumentReadSummaryOption {
	return s.Summary
}

func (s *DocumentReadOption) SetKeyword(v *DocumentReadKeywordOption) *DocumentReadOption {
	s.Keyword = v
	return s
}

func (s *DocumentReadOption) SetNarrator(v *DocumentReadNarratorOption) *DocumentReadOption {
	s.Narrator = v
	return s
}

func (s *DocumentReadOption) SetQuestion(v *DocumentReadQuestionOption) *DocumentReadOption {
	s.Question = v
	return s
}

func (s *DocumentReadOption) SetSummary(v *DocumentReadSummaryOption) *DocumentReadOption {
	s.Summary = v
	return s
}

func (s *DocumentReadOption) Validate() error {
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
