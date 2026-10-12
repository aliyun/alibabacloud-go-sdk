// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteProductCategoryResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteProductCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteProductCategoryResponseBody
	GetRequestId() *string
}

type DeleteProductCategoryResponseBody struct {
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

func (s DeleteProductCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteProductCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteProductCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProductCategoryResponseBody) SetCode(v string) *DeleteProductCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProductCategoryResponseBody) SetMessage(v string) *DeleteProductCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProductCategoryResponseBody) SetRequestId(v string) *DeleteProductCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProductCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
