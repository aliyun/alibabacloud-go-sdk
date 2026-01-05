// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *DeleteProductRequest
	GetProductId() *string
}

type DeleteProductRequest struct {
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DeleteProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *DeleteProductRequest) SetProductId(v string) *DeleteProductRequest {
	s.ProductId = &v
	return s
}

func (s *DeleteProductRequest) Validate() error {
	return dara.Validate(s)
}
