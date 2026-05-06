// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEssayCorrectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *RunEssayCorrectionRequest
	GetAnswer() *string
	SetDimensions(v []*RunEssayCorrectionRequestDimensions) *RunEssayCorrectionRequest
	GetDimensions() []*RunEssayCorrectionRequestDimensions
	SetGrade(v string) *RunEssayCorrectionRequest
	GetGrade() *string
	SetModelId(v string) *RunEssayCorrectionRequest
	GetModelId() *string
	SetOtherReviewPoints(v string) *RunEssayCorrectionRequest
	GetOtherReviewPoints() *string
	SetQuestion(v string) *RunEssayCorrectionRequest
	GetQuestion() *string
	SetSubject(v string) *RunEssayCorrectionRequest
	GetSubject() *string
	SetTotalScore(v int32) *RunEssayCorrectionRequest
	GetTotalScore() *int32
}

type RunEssayCorrectionRequest struct {
	// example:
	//
	// xxx
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// example:
	//
	// [{"name": "内容完整度", "rubric": "文章内容是否完整，是否涵盖了题目的核心要求", "maxScore": 30}]
	Dimensions []*RunEssayCorrectionRequestDimensions `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
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

func (s RunEssayCorrectionRequest) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionRequest) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionRequest) GetAnswer() *string {
	return s.Answer
}

func (s *RunEssayCorrectionRequest) GetDimensions() []*RunEssayCorrectionRequestDimensions {
	return s.Dimensions
}

func (s *RunEssayCorrectionRequest) GetGrade() *string {
	return s.Grade
}

func (s *RunEssayCorrectionRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunEssayCorrectionRequest) GetOtherReviewPoints() *string {
	return s.OtherReviewPoints
}

func (s *RunEssayCorrectionRequest) GetQuestion() *string {
	return s.Question
}

func (s *RunEssayCorrectionRequest) GetSubject() *string {
	return s.Subject
}

func (s *RunEssayCorrectionRequest) GetTotalScore() *int32 {
	return s.TotalScore
}

func (s *RunEssayCorrectionRequest) SetAnswer(v string) *RunEssayCorrectionRequest {
	s.Answer = &v
	return s
}

func (s *RunEssayCorrectionRequest) SetDimensions(v []*RunEssayCorrectionRequestDimensions) *RunEssayCorrectionRequest {
	s.Dimensions = v
	return s
}

func (s *RunEssayCorrectionRequest) SetGrade(v string) *RunEssayCorrectionRequest {
	s.Grade = &v
	return s
}

func (s *RunEssayCorrectionRequest) SetModelId(v string) *RunEssayCorrectionRequest {
	s.ModelId = &v
	return s
}

func (s *RunEssayCorrectionRequest) SetOtherReviewPoints(v string) *RunEssayCorrectionRequest {
	s.OtherReviewPoints = &v
	return s
}

func (s *RunEssayCorrectionRequest) SetQuestion(v string) *RunEssayCorrectionRequest {
	s.Question = &v
	return s
}

func (s *RunEssayCorrectionRequest) SetSubject(v string) *RunEssayCorrectionRequest {
	s.Subject = &v
	return s
}

func (s *RunEssayCorrectionRequest) SetTotalScore(v int32) *RunEssayCorrectionRequest {
	s.TotalScore = &v
	return s
}

func (s *RunEssayCorrectionRequest) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunEssayCorrectionRequestDimensions struct {
	// example:
	//
	// 30
	MaxScore *int32 `json:"maxScore,omitempty" xml:"maxScore,omitempty"`
	// example:
	//
	// 内容完整度
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 文章内容是否完整，是否涵盖了题目的核心要求
	Rubric *string `json:"rubric,omitempty" xml:"rubric,omitempty"`
}

func (s RunEssayCorrectionRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionRequestDimensions) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionRequestDimensions) GetMaxScore() *int32 {
	return s.MaxScore
}

func (s *RunEssayCorrectionRequestDimensions) GetName() *string {
	return s.Name
}

func (s *RunEssayCorrectionRequestDimensions) GetRubric() *string {
	return s.Rubric
}

func (s *RunEssayCorrectionRequestDimensions) SetMaxScore(v int32) *RunEssayCorrectionRequestDimensions {
	s.MaxScore = &v
	return s
}

func (s *RunEssayCorrectionRequestDimensions) SetName(v string) *RunEssayCorrectionRequestDimensions {
	s.Name = &v
	return s
}

func (s *RunEssayCorrectionRequestDimensions) SetRubric(v string) *RunEssayCorrectionRequestDimensions {
	s.Rubric = &v
	return s
}

func (s *RunEssayCorrectionRequestDimensions) Validate() error {
	return dara.Validate(s)
}
