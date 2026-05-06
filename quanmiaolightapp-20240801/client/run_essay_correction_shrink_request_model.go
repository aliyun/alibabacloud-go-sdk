// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEssayCorrectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *RunEssayCorrectionShrinkRequest
	GetAnswer() *string
	SetDimensionsShrink(v string) *RunEssayCorrectionShrinkRequest
	GetDimensionsShrink() *string
	SetGrade(v string) *RunEssayCorrectionShrinkRequest
	GetGrade() *string
	SetModelId(v string) *RunEssayCorrectionShrinkRequest
	GetModelId() *string
	SetOtherReviewPoints(v string) *RunEssayCorrectionShrinkRequest
	GetOtherReviewPoints() *string
	SetQuestion(v string) *RunEssayCorrectionShrinkRequest
	GetQuestion() *string
	SetSubject(v string) *RunEssayCorrectionShrinkRequest
	GetSubject() *string
	SetTotalScore(v int32) *RunEssayCorrectionShrinkRequest
	GetTotalScore() *int32
}

type RunEssayCorrectionShrinkRequest struct {
	// example:
	//
	// xxx
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// example:
	//
	// [{"name": "内容完整度", "rubric": "文章内容是否完整，是否涵盖了题目的核心要求", "maxScore": 30}]
	DimensionsShrink *string `json:"dimensions,omitempty" xml:"dimensions,omitempty"`
	// example:
	//
	// 高一
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// qwen-custom-correction-v1
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// xxx
	OtherReviewPoints *string `json:"otherReviewPoints,omitempty" xml:"otherReviewPoints,omitempty"`
	// example:
	//
	// xx
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// example:
	//
	// 语文
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// example:
	//
	// 60
	TotalScore *int32 `json:"totalScore,omitempty" xml:"totalScore,omitempty"`
}

func (s RunEssayCorrectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionShrinkRequest) GetAnswer() *string {
	return s.Answer
}

func (s *RunEssayCorrectionShrinkRequest) GetDimensionsShrink() *string {
	return s.DimensionsShrink
}

func (s *RunEssayCorrectionShrinkRequest) GetGrade() *string {
	return s.Grade
}

func (s *RunEssayCorrectionShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunEssayCorrectionShrinkRequest) GetOtherReviewPoints() *string {
	return s.OtherReviewPoints
}

func (s *RunEssayCorrectionShrinkRequest) GetQuestion() *string {
	return s.Question
}

func (s *RunEssayCorrectionShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *RunEssayCorrectionShrinkRequest) GetTotalScore() *int32 {
	return s.TotalScore
}

func (s *RunEssayCorrectionShrinkRequest) SetAnswer(v string) *RunEssayCorrectionShrinkRequest {
	s.Answer = &v
	return s
}

func (s *RunEssayCorrectionShrinkRequest) SetDimensionsShrink(v string) *RunEssayCorrectionShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *RunEssayCorrectionShrinkRequest) SetGrade(v string) *RunEssayCorrectionShrinkRequest {
	s.Grade = &v
	return s
}

func (s *RunEssayCorrectionShrinkRequest) SetModelId(v string) *RunEssayCorrectionShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunEssayCorrectionShrinkRequest) SetOtherReviewPoints(v string) *RunEssayCorrectionShrinkRequest {
	s.OtherReviewPoints = &v
	return s
}

func (s *RunEssayCorrectionShrinkRequest) SetQuestion(v string) *RunEssayCorrectionShrinkRequest {
	s.Question = &v
	return s
}

func (s *RunEssayCorrectionShrinkRequest) SetSubject(v string) *RunEssayCorrectionShrinkRequest {
	s.Subject = &v
	return s
}

func (s *RunEssayCorrectionShrinkRequest) SetTotalScore(v int32) *RunEssayCorrectionShrinkRequest {
	s.TotalScore = &v
	return s
}

func (s *RunEssayCorrectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
