// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *InsureOrderApplyRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureOrderApplyRequest
	GetBuyerName() *string
	SetInsOrderId(v string) *InsureOrderApplyRequest
	GetInsOrderId() *string
	SetIsvName(v string) *InsureOrderApplyRequest
	GetIsvName() *string
	SetOutOrderId(v string) *InsureOrderApplyRequest
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *InsureOrderApplyRequest
	GetOutSubOrderId() *string
	SetSupplierCode(v string) *InsureOrderApplyRequest
	GetSupplierCode() *string
}

type InsureOrderApplyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100000102
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000003000000490
	InsOrderId *string `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
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
	// example:
	//
	// fliggy
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s InsureOrderApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderApplyRequest) GoString() string {
	return s.String()
}

func (s *InsureOrderApplyRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderApplyRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureOrderApplyRequest) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *InsureOrderApplyRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureOrderApplyRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *InsureOrderApplyRequest) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *InsureOrderApplyRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureOrderApplyRequest) SetBtripUserId(v string) *InsureOrderApplyRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderApplyRequest) SetBuyerName(v string) *InsureOrderApplyRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureOrderApplyRequest) SetInsOrderId(v string) *InsureOrderApplyRequest {
	s.InsOrderId = &v
	return s
}

func (s *InsureOrderApplyRequest) SetIsvName(v string) *InsureOrderApplyRequest {
	s.IsvName = &v
	return s
}

func (s *InsureOrderApplyRequest) SetOutOrderId(v string) *InsureOrderApplyRequest {
	s.OutOrderId = &v
	return s
}

func (s *InsureOrderApplyRequest) SetOutSubOrderId(v string) *InsureOrderApplyRequest {
	s.OutSubOrderId = &v
	return s
}

func (s *InsureOrderApplyRequest) SetSupplierCode(v string) *InsureOrderApplyRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureOrderApplyRequest) Validate() error {
	return dara.Validate(s)
}
