// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricRulesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableMetricRulesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableMetricRulesResponse
  GetStatusCode() *int32 
  SetBody(v *EnableMetricRulesResponseBody) *EnableMetricRulesResponse
  GetBody() *EnableMetricRulesResponseBody 
}

type EnableMetricRulesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableMetricRulesResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableMetricRulesResponse) GoString() string {
  return s.String()
}

func (s *EnableMetricRulesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableMetricRulesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableMetricRulesResponse) GetBody() *EnableMetricRulesResponseBody  {
  return s.Body
}

func (s *EnableMetricRulesResponse) SetHeaders(v map[string]*string) *EnableMetricRulesResponse {
  s.Headers = v
  return s
}

func (s *EnableMetricRulesResponse) SetStatusCode(v int32) *EnableMetricRulesResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableMetricRulesResponse) SetBody(v *EnableMetricRulesResponseBody) *EnableMetricRulesResponse {
  s.Body = v
  return s
}

func (s *EnableMetricRulesResponse) Validate() error {
  return dara.Validate(s)
}

