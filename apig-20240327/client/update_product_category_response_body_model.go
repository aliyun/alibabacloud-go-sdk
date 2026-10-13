// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateProductCategoryResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateProductCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateProductCategoryResponseBody
	GetRequestId() *string
}

type UpdateProductCategoryResponseBody struct {
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

func (s UpdateProductCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateProductCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateProductCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProductCategoryResponseBody) SetCode(v string) *UpdateProductCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProductCategoryResponseBody) SetMessage(v string) *UpdateProductCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProductCategoryResponseBody) SetRequestId(v string) *UpdateProductCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProductCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
