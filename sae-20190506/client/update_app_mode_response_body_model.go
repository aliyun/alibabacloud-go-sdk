// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAppModeResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateAppModeResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateAppModeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAppModeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateAppModeResponseBody
	GetSuccess() *string
	SetTraceId(v string) *UpdateAppModeResponseBody
	GetTraceId() *string
}

type UpdateAppModeResponseBody struct {
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
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If an exception occurs, a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the application was created successfully. Valid values:
	//
	// - true: The application was created successfully.
	//
	// - false: The application failed to be created.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query invocation details.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateAppModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppModeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAppModeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAppModeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAppModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppModeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateAppModeResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateAppModeResponseBody) SetCode(v string) *UpdateAppModeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppModeResponseBody) SetErrorCode(v string) *UpdateAppModeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAppModeResponseBody) SetMessage(v string) *UpdateAppModeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppModeResponseBody) SetRequestId(v string) *UpdateAppModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppModeResponseBody) SetSuccess(v string) *UpdateAppModeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAppModeResponseBody) SetTraceId(v string) *UpdateAppModeResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateAppModeResponseBody) Validate() error {
	return dara.Validate(s)
}
