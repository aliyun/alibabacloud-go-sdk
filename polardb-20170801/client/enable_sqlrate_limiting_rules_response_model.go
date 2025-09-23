// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSQLRateLimitingRulesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSQLRateLimitingRulesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSQLRateLimitingRulesResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSQLRateLimitingRulesResponseBody) *EnableSQLRateLimitingRulesResponse
  GetBody() *EnableSQLRateLimitingRulesResponseBody 
}

type EnableSQLRateLimitingRulesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSQLRateLimitingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSQLRateLimitingRulesResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSQLRateLimitingRulesResponse) GoString() string {
  return s.String()
}

func (s *EnableSQLRateLimitingRulesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSQLRateLimitingRulesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSQLRateLimitingRulesResponse) GetBody() *EnableSQLRateLimitingRulesResponseBody  {
  return s.Body
}

func (s *EnableSQLRateLimitingRulesResponse) SetHeaders(v map[string]*string) *EnableSQLRateLimitingRulesResponse {
  s.Headers = v
  return s
}

func (s *EnableSQLRateLimitingRulesResponse) SetStatusCode(v int32) *EnableSQLRateLimitingRulesResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSQLRateLimitingRulesResponse) SetBody(v *EnableSQLRateLimitingRulesResponseBody) *EnableSQLRateLimitingRulesResponse {
  s.Body = v
  return s
}

func (s *EnableSQLRateLimitingRulesResponse) Validate() error {
  return dara.Validate(s)
}

