// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantGrammarCheckResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantGrammarCheckResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantGrammarCheckResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTextbookAssistantGrammarCheckResponseBody) *ExecuteTextbookAssistantGrammarCheckResponse
  GetBody() *ExecuteTextbookAssistantGrammarCheckResponseBody 
}

type ExecuteTextbookAssistantGrammarCheckResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTextbookAssistantGrammarCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantGrammarCheckResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) GetBody() *ExecuteTextbookAssistantGrammarCheckResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantGrammarCheckResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantGrammarCheckResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) SetBody(v *ExecuteTextbookAssistantGrammarCheckResponseBody) *ExecuteTextbookAssistantGrammarCheckResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

