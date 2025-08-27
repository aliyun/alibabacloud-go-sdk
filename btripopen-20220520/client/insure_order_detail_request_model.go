// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *InsureOrderDetailRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureOrderDetailRequest
	GetBuyerName() *string
	SetInsOrderId(v string) *InsureOrderDetailRequest
	GetInsOrderId() *string
	SetIsvName(v string) *InsureOrderDetailRequest
	GetIsvName() *string
	SetSupplierCode(v string) *InsureOrderDetailRequest
	GetSupplierCode() *string
}

type InsureOrderDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1021000196440356901
	InsOrderId *string `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	// example:
	//
	// PostalSavingsBank
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// fliggy
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s InsureOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderDetailRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureOrderDetailRequest) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *InsureOrderDetailRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureOrderDetailRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureOrderDetailRequest) SetBtripUserId(v string) *InsureOrderDetailRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderDetailRequest) SetBuyerName(v string) *InsureOrderDetailRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureOrderDetailRequest) SetInsOrderId(v string) *InsureOrderDetailRequest {
	s.InsOrderId = &v
	return s
}

func (s *InsureOrderDetailRequest) SetIsvName(v string) *InsureOrderDetailRequest {
	s.IsvName = &v
	return s
}

func (s *InsureOrderDetailRequest) SetSupplierCode(v string) *InsureOrderDetailRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
