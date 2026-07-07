// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmDisburseCmd interface {
	dara.Model
	String() string
	GoString() string
	SetDisputeId(v string) *ConfirmDisburseCmd
	GetDisputeId() *string
	SetOrderId(v string) *ConfirmDisburseCmd
	GetOrderId() *string
	SetPurchaseOrderId(v string) *ConfirmDisburseCmd
	GetPurchaseOrderId() *string
}

type ConfirmDisburseCmd struct {
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
	// The primary distribution order ID.
	//
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The distribution transaction ID.
	//
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

func (s *ConfirmDisburseCmd) GetDisputeId() *string {
	return s.DisputeId
}

func (s *ConfirmDisburseCmd) GetOrderId() *string {
	return s.OrderId
}

func (s *ConfirmDisburseCmd) GetPurchaseOrderId() *string {
	return s.PurchaseOrderId
}

func (s *ConfirmDisburseCmd) SetDisputeId(v string) *ConfirmDisburseCmd {
	s.DisputeId = &v
	return s
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
