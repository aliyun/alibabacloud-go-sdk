// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateProductResponseBody
	GetCode() *string
	SetData(v *CreateProductResponseBodyData) *CreateProductResponseBody
	GetData() *CreateProductResponseBodyData
	SetMessage(v string) *CreateProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateProductResponseBody
	GetRequestId() *string
}

type CreateProductResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *CreateProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateProductResponseBody) GetData() *CreateProductResponseBodyData {
	return s.Data
}

func (s *CreateProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductResponseBody) SetCode(v string) *CreateProductResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductResponseBody) SetData(v *CreateProductResponseBodyData) *CreateProductResponseBody {
	s.Data = v
	return s
}

func (s *CreateProductResponseBody) SetMessage(v string) *CreateProductResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProductResponseBodyData struct {
	// The ID of the newly created API product, which can be used to locate the product in API operations such as GetProduct, UpdateProduct, PublishProduct, and GetProductRef.
	//
	// example:
	//
	// p-xxxxxxxx
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
}

func (s CreateProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBodyData) GetProductId() *string {
	return s.ProductId
}

func (s *CreateProductResponseBodyData) SetProductId(v string) *CreateProductResponseBodyData {
	s.ProductId = &v
	return s
}

func (s *CreateProductResponseBodyData) Validate() error {
	return dara.Validate(s)
}
