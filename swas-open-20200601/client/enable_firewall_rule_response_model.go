// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFirewallRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableFirewallRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableFirewallRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableFirewallRuleResponseBody) *EnableFirewallRuleResponse
  GetBody() *EnableFirewallRuleResponseBody 
}

type EnableFirewallRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableFirewallRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableFirewallRuleResponse) GoString() string {
  return s.String()
}

func (s *EnableFirewallRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableFirewallRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableFirewallRuleResponse) GetBody() *EnableFirewallRuleResponseBody  {
  return s.Body
}

func (s *EnableFirewallRuleResponse) SetHeaders(v map[string]*string) *EnableFirewallRuleResponse {
  s.Headers = v
  return s
}

func (s *EnableFirewallRuleResponse) SetStatusCode(v int32) *EnableFirewallRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableFirewallRuleResponse) SetBody(v *EnableFirewallRuleResponseBody) *EnableFirewallRuleResponse {
  s.Body = v
  return s
}

func (s *EnableFirewallRuleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

