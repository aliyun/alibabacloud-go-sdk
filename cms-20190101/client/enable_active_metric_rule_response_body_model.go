// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableActiveMetricRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableActiveMetricRuleResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableActiveMetricRuleResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableActiveMetricRuleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableActiveMetricRuleResponseBody
  GetSuccess() *bool 
}

type EnableActiveMetricRuleResponseBody struct {
  // The responses code.
  // 
  // >  The status code 200 indicates that the request was successful.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // The Request is not authorization.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // F82E6667-7811-4BA0-842F-5B2DC42BBAAD
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableActiveMetricRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableActiveMetricRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableActiveMetricRuleResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableActiveMetricRuleResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableActiveMetricRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableActiveMetricRuleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableActiveMetricRuleResponseBody) SetCode(v string) *EnableActiveMetricRuleResponseBody {
  s.Code = &v
  return s
}

func (s *EnableActiveMetricRuleResponseBody) SetMessage(v string) *EnableActiveMetricRuleResponseBody {
  s.Message = &v
  return s
}

func (s *EnableActiveMetricRuleResponseBody) SetRequestId(v string) *EnableActiveMetricRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableActiveMetricRuleResponseBody) SetSuccess(v bool) *EnableActiveMetricRuleResponseBody {
  s.Success = &v
  return s
}

func (s *EnableActiveMetricRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

