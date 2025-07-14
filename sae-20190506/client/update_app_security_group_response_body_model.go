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
	// The HTTP status code. Valid values:
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
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
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
	// Indicates whether the security group of the application was updated. Valid values:
	//
	// 	- **true**: The security group was updated.
	//
	// 	- **false**: The security group failed to be updated.
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
