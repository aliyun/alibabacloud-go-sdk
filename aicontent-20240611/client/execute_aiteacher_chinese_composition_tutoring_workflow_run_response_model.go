// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
  GetStatusCode() *int32 
  SetId(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
  GetId() *string 
  SetEvent(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
  GetEvent() *string 
  SetBody(v *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
  GetBody() *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody 
}

type ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Event *string `json:"event,omitempty" xml:"event,omitempty"`
  Body *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) GetId() *string  {
  return s.Id
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) GetEvent() *string  {
  return s.Event
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) GetBody() *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) SetStatusCode(v int32) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) SetId(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse {
  s.Id = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) SetEvent(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse {
  s.Event = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) SetBody(v *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

