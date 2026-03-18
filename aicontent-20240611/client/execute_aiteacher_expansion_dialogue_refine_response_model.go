// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueRefineResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueRefineResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueRefineResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAITeacherExpansionDialogueRefineResponseBody) *ExecuteAITeacherExpansionDialogueRefineResponse
  GetBody() *ExecuteAITeacherExpansionDialogueRefineResponseBody 
}

type ExecuteAITeacherExpansionDialogueRefineResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAITeacherExpansionDialogueRefineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) GetBody() *ExecuteAITeacherExpansionDialogueRefineResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueRefineResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueRefineResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetBody(v *ExecuteAITeacherExpansionDialogueRefineResponseBody) *ExecuteAITeacherExpansionDialogueRefineResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

