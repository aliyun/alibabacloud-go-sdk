// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateProductResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateProductResponseBody
	GetRequestId() *string
}

type UpdateProductResponseBody struct {
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

func (s UpdateProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProductResponseBody) SetCode(v string) *UpdateProductResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProductResponseBody) SetMessage(v string) *UpdateProductResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProductResponseBody) SetRequestId(v string) *UpdateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProductResponseBody) Validate() error {
	return dara.Validate(s)
}
