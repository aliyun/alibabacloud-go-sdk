// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherSyncDialogueResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherSyncDialogueResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherSyncDialogueResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAITeacherSyncDialogueResponseBody) *ExecuteAITeacherSyncDialogueResponse
  GetBody() *ExecuteAITeacherSyncDialogueResponseBody 
}

type ExecuteAITeacherSyncDialogueResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAITeacherSyncDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherSyncDialogueResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherSyncDialogueResponse) GetBody() *ExecuteAITeacherSyncDialogueResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherSyncDialogueResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetStatusCode(v int32) *ExecuteAITeacherSyncDialogueResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetBody(v *ExecuteAITeacherSyncDialogueResponseBody) *ExecuteAITeacherSyncDialogueResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

