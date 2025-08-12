// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricRuleBlackListResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableMetricRuleBlackListResponseBody
  GetCode() *string 
  SetCount(v int32) *EnableMetricRuleBlackListResponseBody
  GetCount() *int32 
  SetMessage(v string) *EnableMetricRuleBlackListResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableMetricRuleBlackListResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableMetricRuleBlackListResponseBody
  GetSuccess() *bool 
}

type EnableMetricRuleBlackListResponseBody struct {
  // The status code.
  // 
  // > The status code 200 indicates that the request was successful.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The number of blacklist policies that are enabled or disabled.
  // 
  // example:
  // 
  // 1
  Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // The specified resource is not found.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 94869866-1621-3652-BBC9-72A47B2AC2F5
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

func (s EnableMetricRuleBlackListResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableMetricRuleBlackListResponseBody) GoString() string {
  return s.String()
}

func (s *EnableMetricRuleBlackListResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableMetricRuleBlackListResponseBody) GetCount() *int32  {
  return s.Count
}

func (s *EnableMetricRuleBlackListResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableMetricRuleBlackListResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableMetricRuleBlackListResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableMetricRuleBlackListResponseBody) SetCode(v string) *EnableMetricRuleBlackListResponseBody {
  s.Code = &v
  return s
}

func (s *EnableMetricRuleBlackListResponseBody) SetCount(v int32) *EnableMetricRuleBlackListResponseBody {
  s.Count = &v
  return s
}

func (s *EnableMetricRuleBlackListResponseBody) SetMessage(v string) *EnableMetricRuleBlackListResponseBody {
  s.Message = &v
  return s
}

func (s *EnableMetricRuleBlackListResponseBody) SetRequestId(v string) *EnableMetricRuleBlackListResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableMetricRuleBlackListResponseBody) SetSuccess(v bool) *EnableMetricRuleBlackListResponseBody {
  s.Success = &v
  return s
}

func (s *EnableMetricRuleBlackListResponseBody) Validate() error {
  return dara.Validate(s)
}

