// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteApplicationScalingRuleResponseBody
	GetCode() *string
	SetErrorCode(v string) *DeleteApplicationScalingRuleResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApplicationScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteApplicationScalingRuleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteApplicationScalingRuleResponseBody
	GetTraceId() *string
}

type DeleteApplicationScalingRuleResponseBody struct {
	// The status of the API call or a POP error code. Valid values:
	//
	// - **2xx**: success.
	//
	// - **3xx**: redirection.
	//
	// - **4xx**: request error.
	//
	// - **5xx**: server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code. The following rules apply:
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the list of **error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. The following values may be returned:
	//
	// - **success*	- is returned if the request is normal.
	//
	// - A specific error code is returned if the request is abnormal.
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
	// Indicates whether the Auto Scaling policy was successfully deleted. Valid values:
	//
	// - **true**: The deletion was successful.
	//
	// - **false**: The deletion failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. You can use this ID to query the details of a call.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApplicationScalingRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteApplicationScalingRuleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteApplicationScalingRuleResponseBody) SetCode(v string) *DeleteApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) SetErrorCode(v string) *DeleteApplicationScalingRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) SetMessage(v string) *DeleteApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) SetRequestId(v string) *DeleteApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) SetSuccess(v bool) *DeleteApplicationScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) SetTraceId(v string) *DeleteApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
