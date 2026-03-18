// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
  GetStatusCode() *int32 
  SetId(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
  GetId() *string 
  SetEvent(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
  GetEvent() *string 
  SetBody(v *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
  GetBody() *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody 
}

type ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Event *string `json:"event,omitempty" xml:"event,omitempty"`
  Body *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) GetId() *string  {
  return s.Id
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) GetEvent() *string  {
  return s.Event
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) GetBody() *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) SetStatusCode(v int32) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) SetId(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse {
  s.Id = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) SetEvent(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse {
  s.Event = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) SetBody(v *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

