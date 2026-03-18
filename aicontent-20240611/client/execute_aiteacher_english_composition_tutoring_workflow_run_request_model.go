// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEssayOutline(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
  GetEssayOutline() *string 
  SetEssayRequirements(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
  GetEssayRequirements() *string 
  SetEssayTopic(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
  GetEssayTopic() *string 
  SetEssayType(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
  GetEssayType() *string 
  SetEssayWordCount(v int64) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
  GetEssayWordCount() *int64 
  SetGrade(v int64) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
  GetGrade() *int64 
  SetResponseMode(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
  GetResponseMode() *string 
  SetUserId(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
  GetUserId() *string 
}

type ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest struct {
  // example:
  // 
  // Title: The Importance of Reading
  // 
  // I. Introduction
  // 
  // II. Body
  // 
  // III. Conclusion
  EssayOutline *string `json:"essayOutline,omitempty" xml:"essayOutline,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // No less than 100 words
  EssayRequirements *string `json:"essayRequirements,omitempty" xml:"essayRequirements,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // Discuss what to eat
  EssayTopic *string `json:"essayTopic,omitempty" xml:"essayTopic,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // essay
  EssayType *string `json:"essayType,omitempty" xml:"essayType,omitempty"`
  // example:
  // 
  // 100
  EssayWordCount *int64 `json:"essayWordCount,omitempty" xml:"essayWordCount,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 3
  Grade *int64 `json:"grade,omitempty" xml:"grade,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // streaming
  ResponseMode *string `json:"responseMode,omitempty" xml:"responseMode,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // xxxxxxx
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GetEssayOutline() *string  {
  return s.EssayOutline
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GetEssayRequirements() *string  {
  return s.EssayRequirements
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GetEssayTopic() *string  {
  return s.EssayTopic
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GetEssayType() *string  {
  return s.EssayType
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GetEssayWordCount() *int64  {
  return s.EssayWordCount
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GetGrade() *int64  {
  return s.Grade
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GetResponseMode() *string  {
  return s.ResponseMode
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayOutline(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
  s.EssayOutline = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayRequirements(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
  s.EssayRequirements = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayTopic(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
  s.EssayTopic = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayType(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
  s.EssayType = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayWordCount(v int64) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
  s.EssayWordCount = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetGrade(v int64) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
  s.Grade = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetResponseMode(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
  s.ResponseMode = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetUserId(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) Validate() error {
  return dara.Validate(s)
}

