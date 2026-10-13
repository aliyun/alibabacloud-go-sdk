// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMemoryResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteMemoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMemoryResponseBody
	GetSuccess() *bool
}

type DeleteMemoryResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
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
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMemoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMemoryResponseBody) SetCode(v string) *DeleteMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetHttpStatusCode(v int32) *DeleteMemoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetMessage(v string) *DeleteMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetRequestId(v string) *DeleteMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetSuccess(v bool) *DeleteMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
