// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherGrammarCheckResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAITeacherGrammarCheckResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAITeacherGrammarCheckResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAITeacherGrammarCheckResponseBody) *ExecuteAITeacherGrammarCheckResponse
  GetBody() *ExecuteAITeacherGrammarCheckResponseBody 
}

type ExecuteAITeacherGrammarCheckResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAITeacherGrammarCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAITeacherGrammarCheckResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAITeacherGrammarCheckResponse) GetBody() *ExecuteAITeacherGrammarCheckResponseBody  {
  return s.Body
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherGrammarCheckResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetStatusCode(v int32) *ExecuteAITeacherGrammarCheckResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetBody(v *ExecuteAITeacherGrammarCheckResponseBody) *ExecuteAITeacherGrammarCheckResponse {
  s.Body = v
  return s
}

func (s *ExecuteAITeacherGrammarCheckResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

