// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApplicationDescriptionResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateApplicationDescriptionResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateApplicationDescriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationDescriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApplicationDescriptionResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateApplicationDescriptionResponseBody
	GetTraceId() *string
}

type UpdateApplicationDescriptionResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Client error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code. Valid values:
	//
	// - Successful request: The **ErrorCode*	- parameter is not returned.
	//
	// - Failed request: The **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - **success**: The request is successful.
	//
	// - Specific error code: The request is abnormal.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application is deployed successfully. Valid values:
	//
	// - **true**: The deployment is successful.
	//
	// - **false**: The deployment failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query invocation details.
	//
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateApplicationDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApplicationDescriptionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateApplicationDescriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationDescriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApplicationDescriptionResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateApplicationDescriptionResponseBody) SetCode(v string) *UpdateApplicationDescriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetErrorCode(v string) *UpdateApplicationDescriptionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetMessage(v string) *UpdateApplicationDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetRequestId(v string) *UpdateApplicationDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetSuccess(v bool) *UpdateApplicationDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetTraceId(v string) *UpdateApplicationDescriptionResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
