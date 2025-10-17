// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEssayCorrectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrade(v string) *SubmitEssayCorrectionTaskRequest
	GetGrade() *string
	SetModelId(v string) *SubmitEssayCorrectionTaskRequest
	GetModelId() *string
	SetOtherReviewPoints(v string) *SubmitEssayCorrectionTaskRequest
	GetOtherReviewPoints() *string
	SetQuestion(v string) *SubmitEssayCorrectionTaskRequest
	GetQuestion() *string
	SetSubject(v string) *SubmitEssayCorrectionTaskRequest
	GetSubject() *string
	SetTasks(v []*SubmitEssayCorrectionTaskRequestTasks) *SubmitEssayCorrectionTaskRequest
	GetTasks() []*SubmitEssayCorrectionTaskRequestTasks
	SetTotalScore(v int32) *SubmitEssayCorrectionTaskRequest
	GetTotalScore() *int32
}

type SubmitEssayCorrectionTaskRequest struct {
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
	Subject *string                                  `json:"subject,omitempty" xml:"subject,omitempty"`
	Tasks   []*SubmitEssayCorrectionTaskRequestTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	TotalScore *int32 `json:"totalScore,omitempty" xml:"totalScore,omitempty"`
}

func (s SubmitEssayCorrectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitEssayCorrectionTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitEssayCorrectionTaskRequest) GetGrade() *string {
	return s.Grade
}

func (s *SubmitEssayCorrectionTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitEssayCorrectionTaskRequest) GetOtherReviewPoints() *string {
	return s.OtherReviewPoints
}

func (s *SubmitEssayCorrectionTaskRequest) GetQuestion() *string {
	return s.Question
}

func (s *SubmitEssayCorrectionTaskRequest) GetSubject() *string {
	return s.Subject
}

func (s *SubmitEssayCorrectionTaskRequest) GetTasks() []*SubmitEssayCorrectionTaskRequestTasks {
	return s.Tasks
}

func (s *SubmitEssayCorrectionTaskRequest) GetTotalScore() *int32 {
	return s.TotalScore
}

func (s *SubmitEssayCorrectionTaskRequest) SetGrade(v string) *SubmitEssayCorrectionTaskRequest {
	s.Grade = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequest) SetModelId(v string) *SubmitEssayCorrectionTaskRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequest) SetOtherReviewPoints(v string) *SubmitEssayCorrectionTaskRequest {
	s.OtherReviewPoints = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequest) SetQuestion(v string) *SubmitEssayCorrectionTaskRequest {
	s.Question = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequest) SetSubject(v string) *SubmitEssayCorrectionTaskRequest {
	s.Subject = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequest) SetTasks(v []*SubmitEssayCorrectionTaskRequestTasks) *SubmitEssayCorrectionTaskRequest {
	s.Tasks = v
	return s
}

func (s *SubmitEssayCorrectionTaskRequest) SetTotalScore(v int32) *SubmitEssayCorrectionTaskRequest {
	s.TotalScore = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequest) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitEssayCorrectionTaskRequestTasks struct {
	// example:
	//
	// xxx
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// example:
	//
	// 高中二年级
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
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
	// 60
	TotalScore *int32 `json:"totalScore,omitempty" xml:"totalScore,omitempty"`
}

func (s SubmitEssayCorrectionTaskRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s SubmitEssayCorrectionTaskRequestTasks) GoString() string {
	return s.String()
}

func (s *SubmitEssayCorrectionTaskRequestTasks) GetAnswer() *string {
	return s.Answer
}

func (s *SubmitEssayCorrectionTaskRequestTasks) GetGrade() *string {
	return s.Grade
}

func (s *SubmitEssayCorrectionTaskRequestTasks) GetOtherReviewPoints() *string {
	return s.OtherReviewPoints
}

func (s *SubmitEssayCorrectionTaskRequestTasks) GetQuestion() *string {
	return s.Question
}

func (s *SubmitEssayCorrectionTaskRequestTasks) GetSubject() *string {
	return s.Subject
}

func (s *SubmitEssayCorrectionTaskRequestTasks) GetTotalScore() *int32 {
	return s.TotalScore
}

func (s *SubmitEssayCorrectionTaskRequestTasks) SetAnswer(v string) *SubmitEssayCorrectionTaskRequestTasks {
	s.Answer = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequestTasks) SetGrade(v string) *SubmitEssayCorrectionTaskRequestTasks {
	s.Grade = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequestTasks) SetOtherReviewPoints(v string) *SubmitEssayCorrectionTaskRequestTasks {
	s.OtherReviewPoints = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequestTasks) SetQuestion(v string) *SubmitEssayCorrectionTaskRequestTasks {
	s.Question = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequestTasks) SetSubject(v string) *SubmitEssayCorrectionTaskRequestTasks {
	s.Subject = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequestTasks) SetTotalScore(v int32) *SubmitEssayCorrectionTaskRequestTasks {
	s.TotalScore = &v
	return s
}

func (s *SubmitEssayCorrectionTaskRequestTasks) Validate() error {
	return dara.Validate(s)
}
