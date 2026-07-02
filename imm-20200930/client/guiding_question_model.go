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
	// The answer.
	//
	// example:
	//
	// "将场景文本检测和布局分析统一起来是重要的，因为这两个任务虽然在文献中通常被独立研究，但实际上是紧密相关的。"
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The question.
	//
	// example:
	//
	// "为什么将场景文本检测和布局分析统一起来是重要的？"
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
