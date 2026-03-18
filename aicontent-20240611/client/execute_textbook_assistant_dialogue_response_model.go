// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantDialogueResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantDialogueResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantDialogueResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTextbookAssistantDialogueResponseBody) *ExecuteTextbookAssistantDialogueResponse
  GetBody() *ExecuteTextbookAssistantDialogueResponseBody 
}

type ExecuteTextbookAssistantDialogueResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTextbookAssistantDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDialogueResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantDialogueResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantDialogueResponse) GetBody() *ExecuteTextbookAssistantDialogueResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantDialogueResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantDialogueResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantDialogueResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponse) SetBody(v *ExecuteTextbookAssistantDialogueResponseBody) *ExecuteTextbookAssistantDialogueResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantDialogueResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

