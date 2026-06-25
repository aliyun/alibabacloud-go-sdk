// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateNamespaceVpcResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateNamespaceVpcResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateNamespaceVpcResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNamespaceVpcResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNamespaceVpcResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateNamespaceVpcResponseBody
	GetTraceId() *string
}

type UpdateNamespaceVpcResponseBody struct {
	// The HTTP status code. The value can be a POP error code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: The request was invalid.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code.
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, an error code is returned.
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
	// Indicates whether the VPC was updated. Valid values:
	//
	// - **true**: The VPC was updated.
	//
	// - **false**: The VPC failed to be updated.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. You can use it to query the details of a call.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateNamespaceVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceVpcResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceVpcResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateNamespaceVpcResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateNamespaceVpcResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNamespaceVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNamespaceVpcResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNamespaceVpcResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateNamespaceVpcResponseBody) SetCode(v string) *UpdateNamespaceVpcResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetErrorCode(v string) *UpdateNamespaceVpcResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetMessage(v string) *UpdateNamespaceVpcResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetRequestId(v string) *UpdateNamespaceVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetSuccess(v bool) *UpdateNamespaceVpcResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetTraceId(v string) *UpdateNamespaceVpcResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
