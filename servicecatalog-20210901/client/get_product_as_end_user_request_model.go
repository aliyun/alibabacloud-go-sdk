// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductAsEndUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *GetProductAsEndUserRequest
	GetProductId() *string
}

type GetProductAsEndUserRequest struct {
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s GetProductAsEndUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsEndUserRequest) GoString() string {
	return s.String()
}

func (s *GetProductAsEndUserRequest) GetProductId() *string {
	return s.ProductId
}

func (s *GetProductAsEndUserRequest) SetProductId(v string) *GetProductAsEndUserRequest {
	s.ProductId = &v
	return s
}

func (s *GetProductAsEndUserRequest) Validate() error {
	return dara.Validate(s)
}
