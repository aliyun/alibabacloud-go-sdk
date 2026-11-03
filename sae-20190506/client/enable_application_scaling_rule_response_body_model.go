// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationScalingRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableApplicationScalingRuleResponseBody
  GetCode() *string 
  SetErrorCode(v string) *EnableApplicationScalingRuleResponseBody
  GetErrorCode() *string 
  SetMessage(v string) *EnableApplicationScalingRuleResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableApplicationScalingRuleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableApplicationScalingRuleResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *EnableApplicationScalingRuleResponseBody
  GetTraceId() *string 
}

type EnableApplicationScalingRuleResponseBody struct {
  // The API status code or POP error code. Valid values:
  // 
  // - **2xx**: Success.
  // 
  // - **3xx**: Redirect.
  // 
  // - **4xx**: Request error.
  // 
  // - **5xx**: Server error.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The error code. Valid values:
  // 
  // - If the request is successful, the **ErrorCode*	- parameter is not returned.
  // 
  // - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
  // 
  // example:
  // 
  // user.indebt
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The additional information. Valid values:
  // 
  // - **success**: The request is normal.
  // 
  // - A specific error code: The request is abnormal.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the application instance is restarted successfully. Valid values:
  // 
  // - **true**: The application instance is restarted successfully.
  // 
  // - **false**: The application instance fails to be restarted.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  // The trace ID used to accurately query the call information.
  // 
  // example:
  // 
  // 0a98a02315955564772843261e****
  TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s EnableApplicationScalingRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableApplicationScalingRuleResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableApplicationScalingRuleResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableApplicationScalingRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApplicationScalingRuleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableApplicationScalingRuleResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *EnableApplicationScalingRuleResponseBody) SetCode(v string) *EnableApplicationScalingRuleResponseBody {
  s.Code = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetErrorCode(v string) *EnableApplicationScalingRuleResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetMessage(v string) *EnableApplicationScalingRuleResponseBody {
  s.Message = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetRequestId(v string) *EnableApplicationScalingRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetSuccess(v bool) *EnableApplicationScalingRuleResponseBody {
  s.Success = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetTraceId(v string) *EnableApplicationScalingRuleResponseBody {
  s.TraceId = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

