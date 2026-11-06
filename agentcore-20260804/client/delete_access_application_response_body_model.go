// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAccessApplicationResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteAccessApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAccessApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAccessApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAccessApplicationResponseBody
	GetSuccess() *bool
}

type DeleteAccessApplicationResponseBody struct {
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
	// The message about the request processing result.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAccessApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAccessApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAccessApplicationResponseBody) SetCode(v string) *DeleteAccessApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAccessApplicationResponseBody) SetHttpStatusCode(v int32) *DeleteAccessApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAccessApplicationResponseBody) SetMessage(v string) *DeleteAccessApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAccessApplicationResponseBody) SetRequestId(v string) *DeleteAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessApplicationResponseBody) SetSuccess(v bool) *DeleteAccessApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAccessApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
