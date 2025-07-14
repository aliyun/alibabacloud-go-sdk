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
	// Indicates whether the VPC information was updated. Valid values:
	//
	// 	- **true**: indicates that the information was updated.
	//
	// 	- **false**: indicates that the information could not be updated.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the trace. It can be used to query the details of a request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned message.
	//
	// 	- **success*	- is returned when the request succeeds.
	//
	// 	- An error code is returned when the request fails.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The error code.
	//
	// 	- The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
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
