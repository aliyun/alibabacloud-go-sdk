// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAITeacherExpansionDialogueResponseBody) *ExecuteAITeacherExpansionDialogueResponse
  GetBody() *ExecuteAITeacherExpansionDialogueResponseBody 
}

type ExecuteAITeacherExpansionDialogueResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAITeacherExpansionDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherExpansionDialogueResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherExpansionDialogueResponse) GetBody() *ExecuteAITeacherExpansionDialogueResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetBody(v *ExecuteAITeacherExpansionDialogueResponseBody) *ExecuteAITeacherExpansionDialogueResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

