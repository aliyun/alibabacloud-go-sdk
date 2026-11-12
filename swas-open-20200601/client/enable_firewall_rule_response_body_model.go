// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFirewallRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableFirewallRuleResponseBody
  GetRequestId() *string 
}

type EnableFirewallRuleResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 20758A-585D-4A41-A9B2-28DA8F4F****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableFirewallRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableFirewallRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableFirewallRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableFirewallRuleResponseBody) SetRequestId(v string) *EnableFirewallRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableFirewallRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

