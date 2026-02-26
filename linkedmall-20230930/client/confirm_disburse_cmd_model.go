// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmDisburseCmd interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ConfirmDisburseCmd
	GetOrderId() *string
	SetPurchaseOrderId(v string) *ConfirmDisburseCmd
	GetPurchaseOrderId() *string
}

type ConfirmDisburseCmd struct {
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 6692****5696
	PurchaseOrderId *string `json:"purchaseOrderId,omitempty" xml:"purchaseOrderId,omitempty"`
}

func (s ConfirmDisburseCmd) String() string {
	return dara.Prettify(s)
}

func (s ConfirmDisburseCmd) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseCmd) GetOrderId() *string {
	return s.OrderId
}

func (s *ConfirmDisburseCmd) GetPurchaseOrderId() *string {
	return s.PurchaseOrderId
}

func (s *ConfirmDisburseCmd) SetOrderId(v string) *ConfirmDisburseCmd {
	s.OrderId = &v
	return s
}

func (s *ConfirmDisburseCmd) SetPurchaseOrderId(v string) *ConfirmDisburseCmd {
	s.PurchaseOrderId = &v
	return s
}

func (s *ConfirmDisburseCmd) Validate() error {
	return dara.Validate(s)
}
