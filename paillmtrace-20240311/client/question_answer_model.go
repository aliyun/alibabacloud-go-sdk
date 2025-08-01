// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuestionAnswer interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v *QuestionAnswerAnswer) *QuestionAnswer
	GetAnswer() *QuestionAnswerAnswer
	SetGroundTruth(v *QuestionAnswerGroundTruth) *QuestionAnswer
	GetGroundTruth() *QuestionAnswerGroundTruth
	SetQuestion(v string) *QuestionAnswer
	GetQuestion() *string
}

type QuestionAnswer struct {
	Answer      *QuestionAnswerAnswer      `json:"Answer,omitempty" xml:"Answer,omitempty" type:"Struct"`
	GroundTruth *QuestionAnswerGroundTruth `json:"GroundTruth,omitempty" xml:"GroundTruth,omitempty" type:"Struct"`
	Question    *string                    `json:"Question,omitempty" xml:"Question,omitempty"`
}

func (s QuestionAnswer) String() string {
	return dara.Prettify(s)
}

func (s QuestionAnswer) GoString() string {
	return s.String()
}

func (s *QuestionAnswer) GetAnswer() *QuestionAnswerAnswer {
	return s.Answer
}

func (s *QuestionAnswer) GetGroundTruth() *QuestionAnswerGroundTruth {
	return s.GroundTruth
}

func (s *QuestionAnswer) GetQuestion() *string {
	return s.Question
}

func (s *QuestionAnswer) SetAnswer(v *QuestionAnswerAnswer) *QuestionAnswer {
	s.Answer = v
	return s
}

func (s *QuestionAnswer) SetGroundTruth(v *QuestionAnswerGroundTruth) *QuestionAnswer {
	s.GroundTruth = v
	return s
}

func (s *QuestionAnswer) SetQuestion(v string) *QuestionAnswer {
	s.Question = &v
	return s
}

func (s *QuestionAnswer) Validate() error {
	return dara.Validate(s)
}

type QuestionAnswerAnswer struct {
	Contexts []*string `json:"Contexts,omitempty" xml:"Contexts,omitempty" type:"Repeated"`
	Text     *string   `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QuestionAnswerAnswer) String() string {
	return dara.Prettify(s)
}

func (s QuestionAnswerAnswer) GoString() string {
	return s.String()
}

func (s *QuestionAnswerAnswer) GetContexts() []*string {
	return s.Contexts
}

func (s *QuestionAnswerAnswer) GetText() *string {
	return s.Text
}

func (s *QuestionAnswerAnswer) SetContexts(v []*string) *QuestionAnswerAnswer {
	s.Contexts = v
	return s
}

func (s *QuestionAnswerAnswer) SetText(v string) *QuestionAnswerAnswer {
	s.Text = &v
	return s
}

func (s *QuestionAnswerAnswer) Validate() error {
	return dara.Validate(s)
}

type QuestionAnswerGroundTruth struct {
	Contexts []*string `json:"Contexts,omitempty" xml:"Contexts,omitempty" type:"Repeated"`
	Text     *string   `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QuestionAnswerGroundTruth) String() string {
	return dara.Prettify(s)
}

func (s QuestionAnswerGroundTruth) GoString() string {
	return s.String()
}

func (s *QuestionAnswerGroundTruth) GetContexts() []*string {
	return s.Contexts
}

func (s *QuestionAnswerGroundTruth) GetText() *string {
	return s.Text
}

func (s *QuestionAnswerGroundTruth) SetContexts(v []*string) *QuestionAnswerGroundTruth {
	s.Contexts = v
	return s
}

func (s *QuestionAnswerGroundTruth) SetText(v string) *QuestionAnswerGroundTruth {
	s.Text = &v
	return s
}

func (s *QuestionAnswerGroundTruth) Validate() error {
	return dara.Validate(s)
}
