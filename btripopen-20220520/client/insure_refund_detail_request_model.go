// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureRefundDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyId(v string) *InsureRefundDetailRequest
	GetApplyId() *string
	SetBtripUserId(v string) *InsureRefundDetailRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureRefundDetailRequest
	GetBuyerName() *string
	SetInsOrderId(v string) *InsureRefundDetailRequest
	GetInsOrderId() *string
	SetIsvName(v string) *InsureRefundDetailRequest
	GetIsvName() *string
	SetOutApplyId(v string) *InsureRefundDetailRequest
	GetOutApplyId() *string
	SetSupplierCode(v string) *InsureRefundDetailRequest
	GetSupplierCode() *string
}

type InsureRefundDetailRequest struct {
	// example:
	//
	// 1423052318072952023
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 1000000001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
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
	// 1021000196440356908
	OutApplyId *string `json:"out_apply_id,omitempty" xml:"out_apply_id,omitempty"`
	// example:
	//
	// fliggy
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s InsureRefundDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailRequest) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailRequest) GetApplyId() *string {
	return s.ApplyId
}

func (s *InsureRefundDetailRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureRefundDetailRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureRefundDetailRequest) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *InsureRefundDetailRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureRefundDetailRequest) GetOutApplyId() *string {
	return s.OutApplyId
}

func (s *InsureRefundDetailRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureRefundDetailRequest) SetApplyId(v string) *InsureRefundDetailRequest {
	s.ApplyId = &v
	return s
}

func (s *InsureRefundDetailRequest) SetBtripUserId(v string) *InsureRefundDetailRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureRefundDetailRequest) SetBuyerName(v string) *InsureRefundDetailRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureRefundDetailRequest) SetInsOrderId(v string) *InsureRefundDetailRequest {
	s.InsOrderId = &v
	return s
}

func (s *InsureRefundDetailRequest) SetIsvName(v string) *InsureRefundDetailRequest {
	s.IsvName = &v
	return s
}

func (s *InsureRefundDetailRequest) SetOutApplyId(v string) *InsureRefundDetailRequest {
	s.OutApplyId = &v
	return s
}

func (s *InsureRefundDetailRequest) SetSupplierCode(v string) *InsureRefundDetailRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureRefundDetailRequest) Validate() error {
	return dara.Validate(s)
}
