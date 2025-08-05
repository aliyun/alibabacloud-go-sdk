// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableRuleResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableRuleResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableRuleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableRuleResponseBody
  GetSuccess() *bool 
}

type EnableRuleResponseBody struct {
  // The error code. The value Success indicates that the request is successful.
  // 
  // example:
  // 
  // Success
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The error message that is returned if the request failed.
  // 
  // example:
  // 
  // EventRuleNotExisted
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 580A938B-6107-586C-8EC7-F22EEBEDA9E6
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the operation is successful. Valid values: true and false.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableRuleResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableRuleResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableRuleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableRuleResponseBody) SetCode(v string) *EnableRuleResponseBody {
  s.Code = &v
  return s
}

func (s *EnableRuleResponseBody) SetMessage(v string) *EnableRuleResponseBody {
  s.Message = &v
  return s
}

func (s *EnableRuleResponseBody) SetRequestId(v string) *EnableRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableRuleResponseBody) SetSuccess(v bool) *EnableRuleResponseBody {
  s.Success = &v
  return s
}

func (s *EnableRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

