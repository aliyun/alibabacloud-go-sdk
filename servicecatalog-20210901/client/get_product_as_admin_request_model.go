// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductAsAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *GetProductAsAdminRequest
	GetProductId() *string
}

type GetProductAsAdminRequest struct {
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s GetProductAsAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsAdminRequest) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminRequest) GetProductId() *string {
	return s.ProductId
}

func (s *GetProductAsAdminRequest) SetProductId(v string) *GetProductAsAdminRequest {
	s.ProductId = &v
	return s
}

func (s *GetProductAsAdminRequest) Validate() error {
	return dara.Validate(s)
}
