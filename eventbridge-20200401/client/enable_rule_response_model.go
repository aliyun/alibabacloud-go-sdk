// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableRuleResponseBody) *EnableRuleResponse
  GetBody() *EnableRuleResponseBody 
}

type EnableRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableRuleResponse) GoString() string {
  return s.String()
}

func (s *EnableRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableRuleResponse) GetBody() *EnableRuleResponseBody  {
  return s.Body
}

func (s *EnableRuleResponse) SetHeaders(v map[string]*string) *EnableRuleResponse {
  s.Headers = v
  return s
}

func (s *EnableRuleResponse) SetStatusCode(v int32) *EnableRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableRuleResponse) SetBody(v *EnableRuleResponseBody) *EnableRuleResponse {
  s.Body = v
  return s
}

func (s *EnableRuleResponse) Validate() error {
  return dara.Validate(s)
}

