// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantRetryConversationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantRetryConversationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantRetryConversationResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTextbookAssistantRetryConversationResponseBody) *ExecuteTextbookAssistantRetryConversationResponse
  GetBody() *ExecuteTextbookAssistantRetryConversationResponseBody 
}

type ExecuteTextbookAssistantRetryConversationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTextbookAssistantRetryConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) GetBody() *ExecuteTextbookAssistantRetryConversationResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantRetryConversationResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantRetryConversationResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) SetBody(v *ExecuteTextbookAssistantRetryConversationResponseBody) *ExecuteTextbookAssistantRetryConversationResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

