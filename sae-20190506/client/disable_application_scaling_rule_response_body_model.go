// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableApplicationScalingRuleResponseBody
	GetCode() *string
	SetErrorCode(v string) *DisableApplicationScalingRuleResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DisableApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableApplicationScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableApplicationScalingRuleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DisableApplicationScalingRuleResponseBody
	GetTraceId() *string
}

type DisableApplicationScalingRuleResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error codes. Valid values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, **success*	- is returned.
	//
	// 	- If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the auto scaling policy was disabled. Valid values:
	//
	// 	- **true**: The auto scaling policy was disabled.
	//
	// 	- **false**: The auto scaling policy failed to be disabled.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DisableApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableApplicationScalingRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DisableApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableApplicationScalingRuleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DisableApplicationScalingRuleResponseBody) SetCode(v string) *DisableApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetErrorCode(v string) *DisableApplicationScalingRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetMessage(v string) *DisableApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetRequestId(v string) *DisableApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetSuccess(v bool) *DisableApplicationScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetTraceId(v string) *DisableApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
