// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductVersionId(v string) *UpdateProductVersionResponseBody
	GetProductVersionId() *string
	SetRequestId(v string) *UpdateProductVersionResponseBody
	GetRequestId() *string
}

type UpdateProductVersionResponseBody struct {
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

func (s UpdateProductVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionResponseBody) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *UpdateProductVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProductVersionResponseBody) SetProductVersionId(v string) *UpdateProductVersionResponseBody {
	s.ProductVersionId = &v
	return s
}

func (s *UpdateProductVersionResponseBody) SetRequestId(v string) *UpdateProductVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProductVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
