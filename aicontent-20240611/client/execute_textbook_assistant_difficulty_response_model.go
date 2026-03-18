// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantDifficultyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantDifficultyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantDifficultyResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTextbookAssistantDifficultyResponseBody) *ExecuteTextbookAssistantDifficultyResponse
  GetBody() *ExecuteTextbookAssistantDifficultyResponseBody 
}

type ExecuteTextbookAssistantDifficultyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTextbookAssistantDifficultyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantDifficultyResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantDifficultyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantDifficultyResponse) GetBody() *ExecuteTextbookAssistantDifficultyResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantDifficultyResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantDifficultyResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantDifficultyResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponse) SetBody(v *ExecuteTextbookAssistantDifficultyResponseBody) *ExecuteTextbookAssistantDifficultyResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

