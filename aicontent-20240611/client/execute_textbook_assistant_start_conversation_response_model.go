// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantStartConversationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantStartConversationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantStartConversationResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTextbookAssistantStartConversationResponseBody) *ExecuteTextbookAssistantStartConversationResponse
  GetBody() *ExecuteTextbookAssistantStartConversationResponseBody 
}

type ExecuteTextbookAssistantStartConversationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTextbookAssistantStartConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantStartConversationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantStartConversationResponse) GetBody() *ExecuteTextbookAssistantStartConversationResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantStartConversationResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantStartConversationResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantStartConversationResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponse) SetBody(v *ExecuteTextbookAssistantStartConversationResponseBody) *ExecuteTextbookAssistantStartConversationResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

