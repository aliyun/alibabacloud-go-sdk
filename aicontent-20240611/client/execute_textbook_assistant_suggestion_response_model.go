// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantSuggestionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantSuggestionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantSuggestionResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTextbookAssistantSuggestionResponseBody) *ExecuteTextbookAssistantSuggestionResponse
  GetBody() *ExecuteTextbookAssistantSuggestionResponseBody 
}

type ExecuteTextbookAssistantSuggestionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTextbookAssistantSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantSuggestionResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantSuggestionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantSuggestionResponse) GetBody() *ExecuteTextbookAssistantSuggestionResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantSuggestionResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantSuggestionResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantSuggestionResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponse) SetBody(v *ExecuteTextbookAssistantSuggestionResponseBody) *ExecuteTextbookAssistantSuggestionResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

