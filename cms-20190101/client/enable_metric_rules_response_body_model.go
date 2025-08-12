// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricRulesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableMetricRulesResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableMetricRulesResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableMetricRulesResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableMetricRulesResponseBody
  GetSuccess() *bool 
}

type EnableMetricRulesResponseBody struct {
  // The HTTP status code.
  // 
  // > The status code 200 indicates that the request was successful.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned message.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // FF38D33A-67C1-40EB-AB65-FAEE51EDB644
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

func (s EnableMetricRulesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableMetricRulesResponseBody) GoString() string {
  return s.String()
}

func (s *EnableMetricRulesResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableMetricRulesResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableMetricRulesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableMetricRulesResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableMetricRulesResponseBody) SetCode(v string) *EnableMetricRulesResponseBody {
  s.Code = &v
  return s
}

func (s *EnableMetricRulesResponseBody) SetMessage(v string) *EnableMetricRulesResponseBody {
  s.Message = &v
  return s
}

func (s *EnableMetricRulesResponseBody) SetRequestId(v string) *EnableMetricRulesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableMetricRulesResponseBody) SetSuccess(v bool) *EnableMetricRulesResponseBody {
  s.Success = &v
  return s
}

func (s *EnableMetricRulesResponseBody) Validate() error {
  return dara.Validate(s)
}

