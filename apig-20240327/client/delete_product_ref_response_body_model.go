// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductRefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteProductRefResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteProductRefResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteProductRefResponseBody
	GetRequestId() *string
}

type DeleteProductRefResponseBody struct {
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

func (s DeleteProductRefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductRefResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductRefResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteProductRefResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteProductRefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProductRefResponseBody) SetCode(v string) *DeleteProductRefResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProductRefResponseBody) SetMessage(v string) *DeleteProductRefResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProductRefResponseBody) SetRequestId(v string) *DeleteProductRefResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProductRefResponseBody) Validate() error {
	return dara.Validate(s)
}
