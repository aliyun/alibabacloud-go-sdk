// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSqlStatementResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteSqlStatementResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteSqlStatementResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteSqlStatementResponseBody) *ExecuteSqlStatementResponse
  GetBody() *ExecuteSqlStatementResponseBody 
}

type ExecuteSqlStatementResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSqlStatementResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSqlStatementResponse) GoString() string {
  return s.String()
}

func (s *ExecuteSqlStatementResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteSqlStatementResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteSqlStatementResponse) GetBody() *ExecuteSqlStatementResponseBody  {
  return s.Body
}

func (s *ExecuteSqlStatementResponse) SetHeaders(v map[string]*string) *ExecuteSqlStatementResponse {
  s.Headers = v
  return s
}

func (s *ExecuteSqlStatementResponse) SetStatusCode(v int32) *ExecuteSqlStatementResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteSqlStatementResponse) SetBody(v *ExecuteSqlStatementResponseBody) *ExecuteSqlStatementResponse {
  s.Body = v
  return s
}

func (s *ExecuteSqlStatementResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

