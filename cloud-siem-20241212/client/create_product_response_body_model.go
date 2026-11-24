// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *CreateProductResponseBody
	GetProductId() *string
	SetRequestId(v string) *CreateProductResponseBody
	GetRequestId() *string
}

type CreateProductResponseBody struct {
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) GetProductId() *string {
	return s.ProductId
}

func (s *CreateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductResponseBody) SetProductId(v string) *CreateProductResponseBody {
	s.ProductId = &v
	return s
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductResponseBody) Validate() error {
	return dara.Validate(s)
}
