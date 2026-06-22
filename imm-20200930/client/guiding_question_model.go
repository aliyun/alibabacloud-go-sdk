// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGuidingQuestion interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *GuidingQuestion
	GetAnswer() *string
	SetQuestion(v string) *GuidingQuestion
	GetQuestion() *string
}

type GuidingQuestion struct {
	// The answer to the question.
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The guiding question.
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
}

func (s GuidingQuestion) String() string {
	return dara.Prettify(s)
}

func (s GuidingQuestion) GoString() string {
	return s.String()
}

func (s *GuidingQuestion) GetAnswer() *string {
	return s.Answer
}

func (s *GuidingQuestion) GetQuestion() *string {
	return s.Question
}

func (s *GuidingQuestion) SetAnswer(v string) *GuidingQuestion {
	s.Answer = &v
	return s
}

func (s *GuidingQuestion) SetQuestion(v string) *GuidingQuestion {
	s.Question = &v
	return s
}

func (s *GuidingQuestion) Validate() error {
	return dara.Validate(s)
}
