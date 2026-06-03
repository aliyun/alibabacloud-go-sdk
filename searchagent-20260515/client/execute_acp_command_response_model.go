// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAcpCommandResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAcpCommandResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAcpCommandResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAcpCommandResponseBody) *ExecuteAcpCommandResponse
  GetBody() *ExecuteAcpCommandResponseBody 
}

type ExecuteAcpCommandResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAcpCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAcpCommandResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAcpCommandResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAcpCommandResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAcpCommandResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAcpCommandResponse) GetBody() *ExecuteAcpCommandResponseBody  {
  return s.Body
}

func (s *ExecuteAcpCommandResponse) SetHeaders(v map[string]*string) *ExecuteAcpCommandResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAcpCommandResponse) SetStatusCode(v int32) *ExecuteAcpCommandResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAcpCommandResponse) SetBody(v *ExecuteAcpCommandResponseBody) *ExecuteAcpCommandResponse {
  s.Body = v
  return s
}

func (s *ExecuteAcpCommandResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

