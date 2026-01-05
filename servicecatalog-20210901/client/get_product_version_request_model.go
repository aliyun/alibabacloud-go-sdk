// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductVersionId(v string) *GetProductVersionRequest
	GetProductVersionId() *string
}

type GetProductVersionRequest struct {
	// The ID of the product version.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
}

func (s GetProductVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductVersionRequest) GoString() string {
	return s.String()
}

func (s *GetProductVersionRequest) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *GetProductVersionRequest) SetProductVersionId(v string) *GetProductVersionRequest {
	s.ProductVersionId = &v
	return s
}

func (s *GetProductVersionRequest) Validate() error {
	return dara.Validate(s)
}
