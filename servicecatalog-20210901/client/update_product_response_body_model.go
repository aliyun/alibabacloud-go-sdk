// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *UpdateProductResponseBody
	GetProductId() *string
	SetRequestId(v string) *UpdateProductResponseBody
	GetRequestId() *string
}

type UpdateProductResponseBody struct {
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductResponseBody) GetProductId() *string {
	return s.ProductId
}

func (s *UpdateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProductResponseBody) SetProductId(v string) *UpdateProductResponseBody {
	s.ProductId = &v
	return s
}

func (s *UpdateProductResponseBody) SetRequestId(v string) *UpdateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProductResponseBody) Validate() error {
	return dara.Validate(s)
}
