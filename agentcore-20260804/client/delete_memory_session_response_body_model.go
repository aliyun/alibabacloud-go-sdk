// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemorySessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMemorySessionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteMemorySessionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMemorySessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMemorySessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMemorySessionResponseBody
	GetSuccess() *bool
}

type DeleteMemorySessionResponseBody struct {
	// code
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// requestId
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteMemorySessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemorySessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemorySessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMemorySessionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMemorySessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMemorySessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemorySessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMemorySessionResponseBody) SetCode(v string) *DeleteMemorySessionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMemorySessionResponseBody) SetHttpStatusCode(v int32) *DeleteMemorySessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMemorySessionResponseBody) SetMessage(v string) *DeleteMemorySessionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMemorySessionResponseBody) SetRequestId(v string) *DeleteMemorySessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemorySessionResponseBody) SetSuccess(v bool) *DeleteMemorySessionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMemorySessionResponseBody) Validate() error {
	return dara.Validate(s)
}
