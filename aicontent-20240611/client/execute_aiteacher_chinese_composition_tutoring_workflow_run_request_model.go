// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEssayOutline(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
  GetEssayOutline() *string 
  SetEssayRequirements(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
  GetEssayRequirements() *string 
  SetEssayTopic(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
  GetEssayTopic() *string 
  SetEssayType(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
  GetEssayType() *string 
  SetEssayWordCount(v int64) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
  GetEssayWordCount() *int64 
  SetGrade(v int64) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
  GetGrade() *int64 
  SetResponseMode(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
  GetResponseMode() *string 
  SetUserId(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
  GetUserId() *string 
}

type ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest struct {
  EssayOutline *string `json:"essayOutline,omitempty" xml:"essayOutline,omitempty"`
  // This parameter is required.
  EssayRequirements *string `json:"essayRequirements,omitempty" xml:"essayRequirements,omitempty"`
  // This parameter is required.
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
  // xxxxxxxxx
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GetEssayOutline() *string  {
  return s.EssayOutline
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GetEssayRequirements() *string  {
  return s.EssayRequirements
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GetEssayTopic() *string  {
  return s.EssayTopic
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GetEssayType() *string  {
  return s.EssayType
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GetEssayWordCount() *int64  {
  return s.EssayWordCount
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GetGrade() *int64  {
  return s.Grade
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GetResponseMode() *string  {
  return s.ResponseMode
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayOutline(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
  s.EssayOutline = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayRequirements(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
  s.EssayRequirements = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayTopic(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
  s.EssayTopic = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayType(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
  s.EssayType = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayWordCount(v int64) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
  s.EssayWordCount = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetGrade(v int64) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
  s.Grade = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetResponseMode(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
  s.ResponseMode = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetUserId(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) Validate() error {
  return dara.Validate(s)
}

