// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlConcurrencyControlResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSqlConcurrencyControlResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSqlConcurrencyControlResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSqlConcurrencyControlResponseBody) *EnableSqlConcurrencyControlResponse
  GetBody() *EnableSqlConcurrencyControlResponseBody 
}

type EnableSqlConcurrencyControlResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSqlConcurrencyControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSqlConcurrencyControlResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlConcurrencyControlResponse) GoString() string {
  return s.String()
}

func (s *EnableSqlConcurrencyControlResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSqlConcurrencyControlResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSqlConcurrencyControlResponse) GetBody() *EnableSqlConcurrencyControlResponseBody  {
  return s.Body
}

func (s *EnableSqlConcurrencyControlResponse) SetHeaders(v map[string]*string) *EnableSqlConcurrencyControlResponse {
  s.Headers = v
  return s
}

func (s *EnableSqlConcurrencyControlResponse) SetStatusCode(v int32) *EnableSqlConcurrencyControlResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSqlConcurrencyControlResponse) SetBody(v *EnableSqlConcurrencyControlResponseBody) *EnableSqlConcurrencyControlResponse {
  s.Body = v
  return s
}

func (s *EnableSqlConcurrencyControlResponse) Validate() error {
  return dara.Validate(s)
}

