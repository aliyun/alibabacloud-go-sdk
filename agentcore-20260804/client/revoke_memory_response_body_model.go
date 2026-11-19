// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevokeMemoryResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RevokeMemoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RevokeMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokeMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokeMemoryResponseBody
	GetSuccess() *bool
}

type RevokeMemoryResponseBody struct {
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

func (s RevokeMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevokeMemoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RevokeMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeMemoryResponseBody) SetCode(v string) *RevokeMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeMemoryResponseBody) SetHttpStatusCode(v int32) *RevokeMemoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevokeMemoryResponseBody) SetMessage(v string) *RevokeMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeMemoryResponseBody) SetRequestId(v string) *RevokeMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeMemoryResponseBody) SetSuccess(v bool) *RevokeMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
