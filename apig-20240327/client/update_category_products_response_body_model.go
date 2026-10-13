// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCategoryProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCategoryProductsResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateCategoryProductsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCategoryProductsResponseBody
	GetRequestId() *string
}

type UpdateCategoryProductsResponseBody struct {
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

func (s UpdateCategoryProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCategoryProductsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCategoryProductsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCategoryProductsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCategoryProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCategoryProductsResponseBody) SetCode(v string) *UpdateCategoryProductsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCategoryProductsResponseBody) SetMessage(v string) *UpdateCategoryProductsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCategoryProductsResponseBody) SetRequestId(v string) *UpdateCategoryProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCategoryProductsResponseBody) Validate() error {
	return dara.Validate(s)
}
