// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoReadOption interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v *VideoReadKeywordOption) *VideoReadOption
	GetKeyword() *VideoReadKeywordOption
	SetPPT(v *VideoReadPPTOption) *VideoReadOption
	GetPPT() *VideoReadPPTOption
	SetQuestion(v *VideoReadQuestionOption) *VideoReadOption
	GetQuestion() *VideoReadQuestionOption
	SetSummary(v *VideoReadSummaryOption) *VideoReadOption
	GetSummary() *VideoReadSummaryOption
}

type VideoReadOption struct {
	Keyword  *VideoReadKeywordOption  `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PPT      *VideoReadPPTOption      `json:"PPT,omitempty" xml:"PPT,omitempty"`
	Question *VideoReadQuestionOption `json:"Question,omitempty" xml:"Question,omitempty"`
	Summary  *VideoReadSummaryOption  `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s VideoReadOption) String() string {
	return dara.Prettify(s)
}

func (s VideoReadOption) GoString() string {
	return s.String()
}

func (s *VideoReadOption) GetKeyword() *VideoReadKeywordOption {
	return s.Keyword
}

func (s *VideoReadOption) GetPPT() *VideoReadPPTOption {
	return s.PPT
}

func (s *VideoReadOption) GetQuestion() *VideoReadQuestionOption {
	return s.Question
}

func (s *VideoReadOption) GetSummary() *VideoReadSummaryOption {
	return s.Summary
}

func (s *VideoReadOption) SetKeyword(v *VideoReadKeywordOption) *VideoReadOption {
	s.Keyword = v
	return s
}

func (s *VideoReadOption) SetPPT(v *VideoReadPPTOption) *VideoReadOption {
	s.PPT = v
	return s
}

func (s *VideoReadOption) SetQuestion(v *VideoReadQuestionOption) *VideoReadOption {
	s.Question = v
	return s
}

func (s *VideoReadOption) SetSummary(v *VideoReadSummaryOption) *VideoReadOption {
	s.Summary = v
	return s
}

func (s *VideoReadOption) Validate() error {
	if s.Keyword != nil {
		if err := s.Keyword.Validate(); err != nil {
			return err
		}
	}
	if s.PPT != nil {
		if err := s.PPT.Validate(); err != nil {
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
