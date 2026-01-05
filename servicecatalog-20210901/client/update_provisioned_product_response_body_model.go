// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProvisionedProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProvisionedProductId(v string) *UpdateProvisionedProductResponseBody
	GetProvisionedProductId() *string
	SetRequestId(v string) *UpdateProvisionedProductResponseBody
	GetRequestId() *string
}

type UpdateProvisionedProductResponseBody struct {
	// The ID of the product instance.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProvisionedProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductResponseBody) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *UpdateProvisionedProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProvisionedProductResponseBody) SetProvisionedProductId(v string) *UpdateProvisionedProductResponseBody {
	s.ProvisionedProductId = &v
	return s
}

func (s *UpdateProvisionedProductResponseBody) SetRequestId(v string) *UpdateProvisionedProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProvisionedProductResponseBody) Validate() error {
	return dara.Validate(s)
}
