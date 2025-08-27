// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *InsureOrderPayRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureOrderPayRequest
	GetBuyerName() *string
	SetIsvName(v string) *InsureOrderPayRequest
	GetIsvName() *string
	SetOutOrderId(v string) *InsureOrderPayRequest
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *InsureOrderPayRequest
	GetOutSubOrderId() *string
	SetPaymentAmount(v int64) *InsureOrderPayRequest
	GetPaymentAmount() *int64
	SetSupplierCode(v string) *InsureOrderPayRequest
	GetSupplierCode() *string
}

type InsureOrderPayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100000001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// example:
	//
	// PostalSavingsBank
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 202310101026030
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1020030003332000
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3000
	PaymentAmount *int64 `json:"payment_amount,omitempty" xml:"payment_amount,omitempty"`
	// example:
	//
	// fliggy
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s InsureOrderPayRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderPayRequest) GoString() string {
	return s.String()
}

func (s *InsureOrderPayRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderPayRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureOrderPayRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureOrderPayRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *InsureOrderPayRequest) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *InsureOrderPayRequest) GetPaymentAmount() *int64 {
	return s.PaymentAmount
}

func (s *InsureOrderPayRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureOrderPayRequest) SetBtripUserId(v string) *InsureOrderPayRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderPayRequest) SetBuyerName(v string) *InsureOrderPayRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureOrderPayRequest) SetIsvName(v string) *InsureOrderPayRequest {
	s.IsvName = &v
	return s
}

func (s *InsureOrderPayRequest) SetOutOrderId(v string) *InsureOrderPayRequest {
	s.OutOrderId = &v
	return s
}

func (s *InsureOrderPayRequest) SetOutSubOrderId(v string) *InsureOrderPayRequest {
	s.OutSubOrderId = &v
	return s
}

func (s *InsureOrderPayRequest) SetPaymentAmount(v int64) *InsureOrderPayRequest {
	s.PaymentAmount = &v
	return s
}

func (s *InsureOrderPayRequest) SetSupplierCode(v string) *InsureOrderPayRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureOrderPayRequest) Validate() error {
	return dara.Validate(s)
}
