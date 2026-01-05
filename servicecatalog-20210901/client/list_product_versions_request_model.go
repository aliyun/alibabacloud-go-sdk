// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *ListProductVersionsRequest
	GetProductId() *string
}

type ListProductVersionsRequest struct {
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListProductVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductVersionsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListProductVersionsRequest) SetProductId(v string) *ListProductVersionsRequest {
	s.ProductId = &v
	return s
}

func (s *ListProductVersionsRequest) Validate() error {
	return dara.Validate(s)
}
