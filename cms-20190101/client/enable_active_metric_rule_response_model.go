// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableActiveMetricRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableActiveMetricRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableActiveMetricRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableActiveMetricRuleResponseBody) *EnableActiveMetricRuleResponse
  GetBody() *EnableActiveMetricRuleResponseBody 
}

type EnableActiveMetricRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableActiveMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableActiveMetricRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableActiveMetricRuleResponse) GoString() string {
  return s.String()
}

func (s *EnableActiveMetricRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableActiveMetricRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableActiveMetricRuleResponse) GetBody() *EnableActiveMetricRuleResponseBody  {
  return s.Body
}

func (s *EnableActiveMetricRuleResponse) SetHeaders(v map[string]*string) *EnableActiveMetricRuleResponse {
  s.Headers = v
  return s
}

func (s *EnableActiveMetricRuleResponse) SetStatusCode(v int32) *EnableActiveMetricRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableActiveMetricRuleResponse) SetBody(v *EnableActiveMetricRuleResponseBody) *EnableActiveMetricRuleResponse {
  s.Body = v
  return s
}

func (s *EnableActiveMetricRuleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

