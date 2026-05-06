// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEssayCorrectionTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensionsShrink(v string) *SubmitEssayCorrectionTaskShrinkRequest
	GetDimensionsShrink() *string
	SetGrade(v string) *SubmitEssayCorrectionTaskShrinkRequest
	GetGrade() *string
	SetModelId(v string) *SubmitEssayCorrectionTaskShrinkRequest
	GetModelId() *string
	SetOtherReviewPoints(v string) *SubmitEssayCorrectionTaskShrinkRequest
	GetOtherReviewPoints() *string
	SetQuestion(v string) *SubmitEssayCorrectionTaskShrinkRequest
	GetQuestion() *string
	SetSubject(v string) *SubmitEssayCorrectionTaskShrinkRequest
	GetSubject() *string
	SetTasksShrink(v string) *SubmitEssayCorrectionTaskShrinkRequest
	GetTasksShrink() *string
	SetTotalScore(v int32) *SubmitEssayCorrectionTaskShrinkRequest
	GetTotalScore() *int32
}

type SubmitEssayCorrectionTaskShrinkRequest struct {
	// example:
	//
	// [{"name": "内容完整度", "rubric": "文章内容是否完整，是否涵盖了题目的核心要求", "maxScore": 30}]
	DimensionsShrink *string `json:"dimensions,omitempty" xml:"dimensions,omitempty"`
	// example:
	//
	// 高中二年级
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// xxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// xxx
	OtherReviewPoints *string `json:"otherReviewPoints,omitempty" xml:"otherReviewPoints,omitempty"`
	// example:
	//
	// xxx
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// example:
	//
	// 语文
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// example:
	//
	// [{"grade": "高中二年级", "subject": "语文", "totalScore": 60, "question": "请以我的梦想为主题写一篇作文", "answer": "我的梦想是成为一名科学家...", "customId": "task-001"}]
	TasksShrink *string `json:"tasks,omitempty" xml:"tasks,omitempty"`
	// example:
	//
	// 60
	TotalScore *int32 `json:"totalScore,omitempty" xml:"totalScore,omitempty"`
}

func (s SubmitEssayCorrectionTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitEssayCorrectionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) GetDimensionsShrink() *string {
	return s.DimensionsShrink
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) GetGrade() *string {
	return s.Grade
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) GetOtherReviewPoints() *string {
	return s.OtherReviewPoints
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) GetQuestion() *string {
	return s.Question
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) GetTasksShrink() *string {
	return s.TasksShrink
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) GetTotalScore() *int32 {
	return s.TotalScore
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) SetDimensionsShrink(v string) *SubmitEssayCorrectionTaskShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) SetGrade(v string) *SubmitEssayCorrectionTaskShrinkRequest {
	s.Grade = &v
	return s
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) SetModelId(v string) *SubmitEssayCorrectionTaskShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) SetOtherReviewPoints(v string) *SubmitEssayCorrectionTaskShrinkRequest {
	s.OtherReviewPoints = &v
	return s
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) SetQuestion(v string) *SubmitEssayCorrectionTaskShrinkRequest {
	s.Question = &v
	return s
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) SetSubject(v string) *SubmitEssayCorrectionTaskShrinkRequest {
	s.Subject = &v
	return s
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) SetTasksShrink(v string) *SubmitEssayCorrectionTaskShrinkRequest {
	s.TasksShrink = &v
	return s
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) SetTotalScore(v int32) *SubmitEssayCorrectionTaskShrinkRequest {
	s.TotalScore = &v
	return s
}

func (s *SubmitEssayCorrectionTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
