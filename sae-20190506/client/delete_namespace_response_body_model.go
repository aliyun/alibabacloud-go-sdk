// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteNamespaceResponseBody
	GetCode() *string
	SetErrorCode(v string) *DeleteNamespaceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNamespaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteNamespaceResponseBody
	GetTraceId() *string
}

type DeleteNamespaceResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A request error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code.
	//
	// - If the request is successful, this parameter is not returned.
	//
	// - If the request fails, this parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned.
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
	// Indicates whether the namespace was deleted. Valid values:
	//
	// - **true**: The namespace was deleted.
	//
	// - **false**: The namespace failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. You can use this ID to query the details of a request.
	//
	// example:
	//
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNamespaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteNamespaceResponseBody) SetCode(v string) *DeleteNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetErrorCode(v string) *DeleteNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetMessage(v string) *DeleteNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetSuccess(v bool) *DeleteNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetTraceId(v string) *DeleteNamespaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
