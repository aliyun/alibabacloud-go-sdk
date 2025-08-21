// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageGatewayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderDetailsShrink(v string) *CreateStorageGatewayShrinkRequest
	GetOrderDetailsShrink() *string
}

type CreateStorageGatewayShrinkRequest struct {
	// The array of orders.
	//
	// This parameter is required.
	OrderDetailsShrink *string `json:"OrderDetails,omitempty" xml:"OrderDetails,omitempty"`
}

func (s CreateStorageGatewayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageGatewayShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStorageGatewayShrinkRequest) GetOrderDetailsShrink() *string {
	return s.OrderDetailsShrink
}

func (s *CreateStorageGatewayShrinkRequest) SetOrderDetailsShrink(v string) *CreateStorageGatewayShrinkRequest {
	s.OrderDetailsShrink = &v
	return s
}

func (s *CreateStorageGatewayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
