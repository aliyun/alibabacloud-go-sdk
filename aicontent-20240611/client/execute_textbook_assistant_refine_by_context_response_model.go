// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantRefineByContextResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTextbookAssistantRefineByContextResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTextbookAssistantRefineByContextResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTextbookAssistantRefineByContextResponseBody) *ExecuteTextbookAssistantRefineByContextResponse
  GetBody() *ExecuteTextbookAssistantRefineByContextResponseBody 
}

type ExecuteTextbookAssistantRefineByContextResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTextbookAssistantRefineByContextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantRefineByContextResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) GetBody() *ExecuteTextbookAssistantRefineByContextResponseBody  {
  return s.Body
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantRefineByContextResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantRefineByContextResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) SetBody(v *ExecuteTextbookAssistantRefineByContextResponseBody) *ExecuteTextbookAssistantRefineByContextResponse {
  s.Body = v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

