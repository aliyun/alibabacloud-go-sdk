// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecKgCypherResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecKgCypherResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecKgCypherResponse
  GetStatusCode() *int32 
  SetBody(v *ExecKgCypherResponseBody) *ExecKgCypherResponse
  GetBody() *ExecKgCypherResponseBody 
}

type ExecKgCypherResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecKgCypherResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecKgCypherResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponse) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecKgCypherResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecKgCypherResponse) GetBody() *ExecKgCypherResponseBody  {
  return s.Body
}

func (s *ExecKgCypherResponse) SetHeaders(v map[string]*string) *ExecKgCypherResponse {
  s.Headers = v
  return s
}

func (s *ExecKgCypherResponse) SetStatusCode(v int32) *ExecKgCypherResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecKgCypherResponse) SetBody(v *ExecKgCypherResponseBody) *ExecKgCypherResponse {
  s.Body = v
  return s
}

func (s *ExecKgCypherResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

