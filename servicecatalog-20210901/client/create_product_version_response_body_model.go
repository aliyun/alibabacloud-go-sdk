// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductVersionId(v string) *CreateProductVersionResponseBody
	GetProductVersionId() *string
	SetRequestId(v string) *CreateProductVersionResponseBody
	GetRequestId() *string
}

type CreateProductVersionResponseBody struct {
	// The ID of the product version.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProductVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductVersionResponseBody) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *CreateProductVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductVersionResponseBody) SetProductVersionId(v string) *CreateProductVersionResponseBody {
	s.ProductVersionId = &v
	return s
}

func (s *CreateProductVersionResponseBody) SetRequestId(v string) *CreateProductVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
