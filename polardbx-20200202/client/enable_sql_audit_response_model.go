// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlAuditResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSqlAuditResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSqlAuditResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSqlAuditResponseBody) *EnableSqlAuditResponse
  GetBody() *EnableSqlAuditResponseBody 
}

type EnableSqlAuditResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSqlAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSqlAuditResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlAuditResponse) GoString() string {
  return s.String()
}

func (s *EnableSqlAuditResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSqlAuditResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSqlAuditResponse) GetBody() *EnableSqlAuditResponseBody  {
  return s.Body
}

func (s *EnableSqlAuditResponse) SetHeaders(v map[string]*string) *EnableSqlAuditResponse {
  s.Headers = v
  return s
}

func (s *EnableSqlAuditResponse) SetStatusCode(v int32) *EnableSqlAuditResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSqlAuditResponse) SetBody(v *EnableSqlAuditResponseBody) *EnableSqlAuditResponse {
  s.Body = v
  return s
}

func (s *EnableSqlAuditResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

