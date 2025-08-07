// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFirewallRulesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMessage(v string) *EnableFirewallRulesResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableFirewallRulesResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableFirewallRulesResponseBody
  GetSuccess() *bool 
}

type EnableFirewallRulesResponseBody struct {
  // The message that is returned for the request.
  // 
  // > If the request was successful, Successful is returned. If the request failed, an error message that contains information such as an error code is returned.
  // 
  // example:
  // 
  // Message
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 99B355CE-526C-478B-B730-AD9D7C******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false**
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableFirewallRulesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableFirewallRulesResponseBody) GoString() string {
  return s.String()
}

func (s *EnableFirewallRulesResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableFirewallRulesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableFirewallRulesResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableFirewallRulesResponseBody) SetMessage(v string) *EnableFirewallRulesResponseBody {
  s.Message = &v
  return s
}

func (s *EnableFirewallRulesResponseBody) SetRequestId(v string) *EnableFirewallRulesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableFirewallRulesResponseBody) SetSuccess(v bool) *EnableFirewallRulesResponseBody {
  s.Success = &v
  return s
}

func (s *EnableFirewallRulesResponseBody) Validate() error {
  return dara.Validate(s)
}

