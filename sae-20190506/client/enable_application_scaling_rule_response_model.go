// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationScalingRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationScalingRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationScalingRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationScalingRuleResponseBody) *EnableApplicationScalingRuleResponse
  GetBody() *EnableApplicationScalingRuleResponseBody 
}

type EnableApplicationScalingRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationScalingRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationScalingRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationScalingRuleResponse) GetBody() *EnableApplicationScalingRuleResponseBody  {
  return s.Body
}

func (s *EnableApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *EnableApplicationScalingRuleResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationScalingRuleResponse) SetStatusCode(v int32) *EnableApplicationScalingRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationScalingRuleResponse) SetBody(v *EnableApplicationScalingRuleResponseBody) *EnableApplicationScalingRuleResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationScalingRuleResponse) Validate() error {
  return dara.Validate(s)
}

