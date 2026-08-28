// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMcpResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteMcpResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMcpResponseBody
	GetSuccess() *bool
}

type DeleteMcpResponseBody struct {
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

func (s DeleteMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMcpResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcpResponseBody) SetCode(v string) *DeleteMcpResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMcpResponseBody) SetHttpStatusCode(v int32) *DeleteMcpResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMcpResponseBody) SetMessage(v string) *DeleteMcpResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMcpResponseBody) SetRequestId(v string) *DeleteMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcpResponseBody) SetSuccess(v bool) *DeleteMcpResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMcpResponseBody) Validate() error {
	return dara.Validate(s)
}
