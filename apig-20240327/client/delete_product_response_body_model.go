// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteProductResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteProductResponseBody
	GetRequestId() *string
}

type DeleteProductResponseBody struct {
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

func (s DeleteProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProductResponseBody) SetCode(v string) *DeleteProductResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProductResponseBody) SetMessage(v string) *DeleteProductResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProductResponseBody) SetRequestId(v string) *DeleteProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProductResponseBody) Validate() error {
	return dara.Validate(s)
}
