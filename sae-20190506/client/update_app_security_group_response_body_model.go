// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAppSecurityGroupResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateAppSecurityGroupResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateAppSecurityGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAppSecurityGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAppSecurityGroupResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateAppSecurityGroupResponseBody
	GetTraceId() *string
}

type UpdateAppSecurityGroupResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
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
	// - If the request is successful, the **ErrorCode*	- field is not returned.
	//
	// - If the request fails, the **ErrorCode*	- field is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - **success**: The request is successful.
	//
	// - Specific error code: An exception occurred during the request.
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
	// Indicates whether the security group of the application is updated. Valid values:
	//
	// - **true**: The update is successful.
	//
	// - **false**: The update failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to accurately query call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateAppSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppSecurityGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAppSecurityGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAppSecurityGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAppSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppSecurityGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAppSecurityGroupResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateAppSecurityGroupResponseBody) SetCode(v string) *UpdateAppSecurityGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetErrorCode(v string) *UpdateAppSecurityGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetMessage(v string) *UpdateAppSecurityGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetRequestId(v string) *UpdateAppSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetSuccess(v bool) *UpdateAppSecurityGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetTraceId(v string) *UpdateAppSecurityGroupResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
