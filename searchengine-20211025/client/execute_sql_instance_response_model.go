// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSqlInstanceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteSqlInstanceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteSqlInstanceResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteSqlInstanceResponseBody) *ExecuteSqlInstanceResponse
  GetBody() *ExecuteSqlInstanceResponseBody 
}

type ExecuteSqlInstanceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteSqlInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSqlInstanceResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSqlInstanceResponse) GoString() string {
  return s.String()
}

func (s *ExecuteSqlInstanceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteSqlInstanceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteSqlInstanceResponse) GetBody() *ExecuteSqlInstanceResponseBody  {
  return s.Body
}

func (s *ExecuteSqlInstanceResponse) SetHeaders(v map[string]*string) *ExecuteSqlInstanceResponse {
  s.Headers = v
  return s
}

func (s *ExecuteSqlInstanceResponse) SetStatusCode(v int32) *ExecuteSqlInstanceResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteSqlInstanceResponse) SetBody(v *ExecuteSqlInstanceResponseBody) *ExecuteSqlInstanceResponse {
  s.Body = v
  return s
}

func (s *ExecuteSqlInstanceResponse) Validate() error {
  return dara.Validate(s)
}

