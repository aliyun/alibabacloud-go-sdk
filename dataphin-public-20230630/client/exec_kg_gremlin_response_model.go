// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecKgGremlinResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecKgGremlinResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecKgGremlinResponse
  GetStatusCode() *int32 
  SetBody(v *ExecKgGremlinResponseBody) *ExecKgGremlinResponse
  GetBody() *ExecKgGremlinResponseBody 
}

type ExecKgGremlinResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecKgGremlinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecKgGremlinResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponse) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecKgGremlinResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecKgGremlinResponse) GetBody() *ExecKgGremlinResponseBody  {
  return s.Body
}

func (s *ExecKgGremlinResponse) SetHeaders(v map[string]*string) *ExecKgGremlinResponse {
  s.Headers = v
  return s
}

func (s *ExecKgGremlinResponse) SetStatusCode(v int32) *ExecKgGremlinResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecKgGremlinResponse) SetBody(v *ExecKgGremlinResponseBody) *ExecKgGremlinResponse {
  s.Body = v
  return s
}

func (s *ExecKgGremlinResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

