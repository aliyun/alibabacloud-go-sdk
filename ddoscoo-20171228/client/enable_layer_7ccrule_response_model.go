// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLayer7CCRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableLayer7CCRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableLayer7CCRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableLayer7CCRuleResponseBody) *EnableLayer7CCRuleResponse
  GetBody() *EnableLayer7CCRuleResponseBody 
}

type EnableLayer7CCRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableLayer7CCRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableLayer7CCRuleResponse) GoString() string {
  return s.String()
}

func (s *EnableLayer7CCRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableLayer7CCRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableLayer7CCRuleResponse) GetBody() *EnableLayer7CCRuleResponseBody  {
  return s.Body
}

func (s *EnableLayer7CCRuleResponse) SetHeaders(v map[string]*string) *EnableLayer7CCRuleResponse {
  s.Headers = v
  return s
}

func (s *EnableLayer7CCRuleResponse) SetStatusCode(v int32) *EnableLayer7CCRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableLayer7CCRuleResponse) SetBody(v *EnableLayer7CCRuleResponseBody) *EnableLayer7CCRuleResponse {
  s.Body = v
  return s
}

func (s *EnableLayer7CCRuleResponse) Validate() error {
  return dara.Validate(s)
}

