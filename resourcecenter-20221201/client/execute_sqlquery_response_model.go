// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSQLQueryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteSQLQueryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteSQLQueryResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteSQLQueryResponseBody) *ExecuteSQLQueryResponse
  GetBody() *ExecuteSQLQueryResponseBody 
}

type ExecuteSQLQueryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteSQLQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSQLQueryResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSQLQueryResponse) GoString() string {
  return s.String()
}

func (s *ExecuteSQLQueryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteSQLQueryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteSQLQueryResponse) GetBody() *ExecuteSQLQueryResponseBody  {
  return s.Body
}

func (s *ExecuteSQLQueryResponse) SetHeaders(v map[string]*string) *ExecuteSQLQueryResponse {
  s.Headers = v
  return s
}

func (s *ExecuteSQLQueryResponse) SetStatusCode(v int32) *ExecuteSQLQueryResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteSQLQueryResponse) SetBody(v *ExecuteSQLQueryResponseBody) *ExecuteSQLQueryResponse {
  s.Body = v
  return s
}

func (s *ExecuteSQLQueryResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

