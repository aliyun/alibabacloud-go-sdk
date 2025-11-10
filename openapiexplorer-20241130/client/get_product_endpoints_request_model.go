// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProduct(v string) *GetProductEndpointsRequest
	GetProduct() *string
}

type GetProductEndpointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetProductEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductEndpointsRequest) GoString() string {
	return s.String()
}

func (s *GetProductEndpointsRequest) GetProduct() *string {
	return s.Product
}

func (s *GetProductEndpointsRequest) SetProduct(v string) *GetProductEndpointsRequest {
	s.Product = &v
	return s
}

func (s *GetProductEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
