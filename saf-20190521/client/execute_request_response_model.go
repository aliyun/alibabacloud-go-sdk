// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteRequestResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteRequestResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteRequestResponseBody) *ExecuteRequestResponse
  GetBody() *ExecuteRequestResponseBody 
}

type ExecuteRequestResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteRequestResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestResponse) GoString() string {
  return s.String()
}

func (s *ExecuteRequestResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteRequestResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteRequestResponse) GetBody() *ExecuteRequestResponseBody  {
  return s.Body
}

func (s *ExecuteRequestResponse) SetHeaders(v map[string]*string) *ExecuteRequestResponse {
  s.Headers = v
  return s
}

func (s *ExecuteRequestResponse) SetStatusCode(v int32) *ExecuteRequestResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteRequestResponse) SetBody(v *ExecuteRequestResponseBody) *ExecuteRequestResponse {
  s.Body = v
  return s
}

func (s *ExecuteRequestResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

