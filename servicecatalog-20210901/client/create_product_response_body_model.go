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
	SetProductVersionId(v string) *CreateProductResponseBody
	GetProductVersionId() *string
	SetRequestId(v string) *CreateProductResponseBody
	GetRequestId() *string
}

type CreateProductResponseBody struct {
	// The ID of the product.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
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

func (s *CreateProductResponseBody) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *CreateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductResponseBody) SetProductId(v string) *CreateProductResponseBody {
	s.ProductId = &v
	return s
}

func (s *CreateProductResponseBody) SetProductVersionId(v string) *CreateProductResponseBody {
	s.ProductVersionId = &v
	return s
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductResponseBody) Validate() error {
	return dara.Validate(s)
}
