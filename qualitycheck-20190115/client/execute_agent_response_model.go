// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAgentResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAgentResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAgentResponse
  GetStatusCode() *int32 
  SetId(v string) *ExecuteAgentResponse
  GetId() *string 
  SetEvent(v string) *ExecuteAgentResponse
  GetEvent() *string 
  SetBody(v *ExecuteAgentResponseBody) *ExecuteAgentResponse
  GetBody() *ExecuteAgentResponseBody 
}

type ExecuteAgentResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Event *string `json:"event,omitempty" xml:"event,omitempty"`
  Body *ExecuteAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAgentResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAgentResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAgentResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAgentResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAgentResponse) GetId() *string  {
  return s.Id
}

func (s *ExecuteAgentResponse) GetEvent() *string  {
  return s.Event
}

func (s *ExecuteAgentResponse) GetBody() *ExecuteAgentResponseBody  {
  return s.Body
}

func (s *ExecuteAgentResponse) SetHeaders(v map[string]*string) *ExecuteAgentResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAgentResponse) SetStatusCode(v int32) *ExecuteAgentResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAgentResponse) SetId(v string) *ExecuteAgentResponse {
  s.Id = &v
  return s
}

func (s *ExecuteAgentResponse) SetEvent(v string) *ExecuteAgentResponse {
  s.Event = &v
  return s
}

func (s *ExecuteAgentResponse) SetBody(v *ExecuteAgentResponseBody) *ExecuteAgentResponse {
  s.Body = v
  return s
}

func (s *ExecuteAgentResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

