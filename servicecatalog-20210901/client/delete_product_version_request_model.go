// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductVersionId(v string) *DeleteProductVersionRequest
	GetProductVersionId() *string
}

type DeleteProductVersionRequest struct {
	// The ID of the product version.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
}

func (s DeleteProductVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionRequest) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *DeleteProductVersionRequest) SetProductVersionId(v string) *DeleteProductVersionRequest {
	s.ProductVersionId = &v
	return s
}

func (s *DeleteProductVersionRequest) Validate() error {
	return dara.Validate(s)
}
