// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLaunchOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *ListLaunchOptionsRequest
	GetProductId() *string
}

type ListLaunchOptionsRequest struct {
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListLaunchOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLaunchOptionsRequest) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListLaunchOptionsRequest) SetProductId(v string) *ListLaunchOptionsRequest {
	s.ProductId = &v
	return s
}

func (s *ListLaunchOptionsRequest) Validate() error {
	return dara.Validate(s)
}
