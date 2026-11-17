// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicalModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLogicalModelResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteLogicalModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLogicalModelResponseBody
	GetRequestId() *string
}

type DeleteLogicalModelResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
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
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteLogicalModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicalModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogicalModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLogicalModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLogicalModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogicalModelResponseBody) SetCode(v string) *DeleteLogicalModelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLogicalModelResponseBody) SetMessage(v string) *DeleteLogicalModelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLogicalModelResponseBody) SetRequestId(v string) *DeleteLogicalModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogicalModelResponseBody) Validate() error {
	return dara.Validate(s)
}
