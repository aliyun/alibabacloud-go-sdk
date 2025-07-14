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
  // The HTTP status code. Take note of the following rules:
  // 
  // 	- **2xx**: The call was successful.
  // 
  // 	- **3xx**: The call was redirected.
  // 
  // 	- **4xx**: The call failed.
  // 
  // 	- **5xx**: A server error occurred.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The error code returned if the request failed. Take note of the following rules:
  // 
  // 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
  // 
  // 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
  // 
  // example:
  // 
  // Null
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The additional information that is returned. Take note of the following rules:
  // 
  // 	- success: If the call is successful, **success*	- is returned.
  // 
  // 	- An error code: If the call fails, an error code is returned.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Specifies whether the instances are successfully restarted. Take note of the following rules:
  // 
  // 	- **true**
  // 
  // 	- **false**: The restart failed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  // The trace ID that is used to query the details of the request.
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

