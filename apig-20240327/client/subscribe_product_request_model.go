// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *SubscribeProductRequest
	GetProductId() *string
}

type SubscribeProductRequest struct {
	// The ID of the API product to subscribe to. This is the primary ID of the API product resource. You can obtain this ID from the productId field in the response of the ListProducts operation.
	//
	// example:
	//
	// product-xxxxxxxx
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
}

func (s SubscribeProductRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscribeProductRequest) GoString() string {
	return s.String()
}

func (s *SubscribeProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *SubscribeProductRequest) SetProductId(v string) *SubscribeProductRequest {
	s.ProductId = &v
	return s
}

func (s *SubscribeProductRequest) Validate() error {
	return dara.Validate(s)
}
