// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherSyncDialogueTranslateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherSyncDialogueTranslateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherSyncDialogueTranslateResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAITeacherSyncDialogueTranslateResponseBody) *ExecuteAITeacherSyncDialogueTranslateResponse
  GetBody() *ExecuteAITeacherSyncDialogueTranslateResponseBody 
}

type ExecuteAITeacherSyncDialogueTranslateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAITeacherSyncDialogueTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) GetBody() *ExecuteAITeacherSyncDialogueTranslateResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherSyncDialogueTranslateResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetStatusCode(v int32) *ExecuteAITeacherSyncDialogueTranslateResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetBody(v *ExecuteAITeacherSyncDialogueTranslateResponseBody) *ExecuteAITeacherSyncDialogueTranslateResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

