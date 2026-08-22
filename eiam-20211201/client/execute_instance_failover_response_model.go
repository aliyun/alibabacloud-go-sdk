// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteInstanceFailoverResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteInstanceFailoverResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteInstanceFailoverResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteInstanceFailoverResponseBody) *ExecuteInstanceFailoverResponse
  GetBody() *ExecuteInstanceFailoverResponseBody 
}

type ExecuteInstanceFailoverResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteInstanceFailoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteInstanceFailoverResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteInstanceFailoverResponse) GoString() string {
  return s.String()
}

func (s *ExecuteInstanceFailoverResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteInstanceFailoverResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteInstanceFailoverResponse) GetBody() *ExecuteInstanceFailoverResponseBody  {
  return s.Body
}

func (s *ExecuteInstanceFailoverResponse) SetHeaders(v map[string]*string) *ExecuteInstanceFailoverResponse {
  s.Headers = v
  return s
}

func (s *ExecuteInstanceFailoverResponse) SetStatusCode(v int32) *ExecuteInstanceFailoverResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteInstanceFailoverResponse) SetBody(v *ExecuteInstanceFailoverResponseBody) *ExecuteInstanceFailoverResponse {
  s.Body = v
  return s
}

func (s *ExecuteInstanceFailoverResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

