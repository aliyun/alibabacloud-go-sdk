// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFirewallRulesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableFirewallRulesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableFirewallRulesResponse
  GetStatusCode() *int32 
  SetBody(v *EnableFirewallRulesResponseBody) *EnableFirewallRulesResponse
  GetBody() *EnableFirewallRulesResponseBody 
}

type EnableFirewallRulesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableFirewallRulesResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableFirewallRulesResponse) GoString() string {
  return s.String()
}

func (s *EnableFirewallRulesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableFirewallRulesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableFirewallRulesResponse) GetBody() *EnableFirewallRulesResponseBody  {
  return s.Body
}

func (s *EnableFirewallRulesResponse) SetHeaders(v map[string]*string) *EnableFirewallRulesResponse {
  s.Headers = v
  return s
}

func (s *EnableFirewallRulesResponse) SetStatusCode(v int32) *EnableFirewallRulesResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableFirewallRulesResponse) SetBody(v *EnableFirewallRulesResponseBody) *EnableFirewallRulesResponse {
  s.Body = v
  return s
}

func (s *EnableFirewallRulesResponse) Validate() error {
  return dara.Validate(s)
}

