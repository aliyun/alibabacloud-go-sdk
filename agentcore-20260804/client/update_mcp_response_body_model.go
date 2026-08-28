// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMcpResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateMcpResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMcpResponseBody
	GetSuccess() *bool
}

type UpdateMcpResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Request processed successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMcpResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMcpResponseBody) SetCode(v string) *UpdateMcpResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMcpResponseBody) SetHttpStatusCode(v int32) *UpdateMcpResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMcpResponseBody) SetMessage(v string) *UpdateMcpResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMcpResponseBody) SetRequestId(v string) *UpdateMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcpResponseBody) SetSuccess(v bool) *UpdateMcpResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMcpResponseBody) Validate() error {
	return dara.Validate(s)
}
