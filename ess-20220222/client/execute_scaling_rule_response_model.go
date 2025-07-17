// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScalingRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteScalingRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteScalingRuleResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteScalingRuleResponseBody) *ExecuteScalingRuleResponse
  GetBody() *ExecuteScalingRuleResponseBody 
}

type ExecuteScalingRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteScalingRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScalingRuleResponse) GoString() string {
  return s.String()
}

func (s *ExecuteScalingRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteScalingRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteScalingRuleResponse) GetBody() *ExecuteScalingRuleResponseBody  {
  return s.Body
}

func (s *ExecuteScalingRuleResponse) SetHeaders(v map[string]*string) *ExecuteScalingRuleResponse {
  s.Headers = v
  return s
}

func (s *ExecuteScalingRuleResponse) SetStatusCode(v int32) *ExecuteScalingRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteScalingRuleResponse) SetBody(v *ExecuteScalingRuleResponseBody) *ExecuteScalingRuleResponse {
  s.Body = v
  return s
}

func (s *ExecuteScalingRuleResponse) Validate() error {
  return dara.Validate(s)
}

