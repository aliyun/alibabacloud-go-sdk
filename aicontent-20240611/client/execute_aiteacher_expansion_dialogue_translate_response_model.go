// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueTranslateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueTranslateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueTranslateResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAITeacherExpansionDialogueTranslateResponseBody) *ExecuteAITeacherExpansionDialogueTranslateResponse
  GetBody() *ExecuteAITeacherExpansionDialogueTranslateResponseBody 
}

type ExecuteAITeacherExpansionDialogueTranslateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAITeacherExpansionDialogueTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) GetBody() *ExecuteAITeacherExpansionDialogueTranslateResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueTranslateResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueTranslateResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetBody(v *ExecuteAITeacherExpansionDialogueTranslateResponseBody) *ExecuteAITeacherExpansionDialogueTranslateResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

