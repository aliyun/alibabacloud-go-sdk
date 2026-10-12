// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductRefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateProductRefResponseBody
	GetCode() *string
	SetMessage(v string) *CreateProductRefResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateProductRefResponseBody
	GetRequestId() *string
}

type CreateProductRefResponseBody struct {
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

func (s CreateProductRefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRefResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductRefResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateProductRefResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateProductRefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductRefResponseBody) SetCode(v string) *CreateProductRefResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductRefResponseBody) SetMessage(v string) *CreateProductRefResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProductRefResponseBody) SetRequestId(v string) *CreateProductRefResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductRefResponseBody) Validate() error {
	return dara.Validate(s)
}
