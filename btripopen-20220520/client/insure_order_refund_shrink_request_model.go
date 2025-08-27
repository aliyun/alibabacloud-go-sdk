// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderRefundShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *InsureOrderRefundShrinkRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureOrderRefundShrinkRequest
	GetBuyerName() *string
	SetIsvName(v string) *InsureOrderRefundShrinkRequest
	GetIsvName() *string
	SetOutApplyId(v string) *InsureOrderRefundShrinkRequest
	GetOutApplyId() *string
	SetPolicyNoListShrink(v string) *InsureOrderRefundShrinkRequest
	GetPolicyNoListShrink() *string
	SetSubInsOrderIdsShrink(v string) *InsureOrderRefundShrinkRequest
	GetSubInsOrderIdsShrink() *string
	SetSupplierCode(v string) *InsureOrderRefundShrinkRequest
	GetSupplierCode() *string
}

type InsureOrderRefundShrinkRequest struct {
	// example:
	//
	// 1000001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// example:
	//
	// PostalSavingsBank
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1021000196440356908
	OutApplyId           *string `json:"out_apply_id,omitempty" xml:"out_apply_id,omitempty"`
	PolicyNoListShrink   *string `json:"policy_no_list,omitempty" xml:"policy_no_list,omitempty"`
	SubInsOrderIdsShrink *string `json:"sub_ins_order_ids,omitempty" xml:"sub_ins_order_ids,omitempty"`
	// example:
	//
	// fliggy
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s InsureOrderRefundShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderRefundShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsureOrderRefundShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderRefundShrinkRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureOrderRefundShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureOrderRefundShrinkRequest) GetOutApplyId() *string {
	return s.OutApplyId
}

func (s *InsureOrderRefundShrinkRequest) GetPolicyNoListShrink() *string {
	return s.PolicyNoListShrink
}

func (s *InsureOrderRefundShrinkRequest) GetSubInsOrderIdsShrink() *string {
	return s.SubInsOrderIdsShrink
}

func (s *InsureOrderRefundShrinkRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureOrderRefundShrinkRequest) SetBtripUserId(v string) *InsureOrderRefundShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderRefundShrinkRequest) SetBuyerName(v string) *InsureOrderRefundShrinkRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureOrderRefundShrinkRequest) SetIsvName(v string) *InsureOrderRefundShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *InsureOrderRefundShrinkRequest) SetOutApplyId(v string) *InsureOrderRefundShrinkRequest {
	s.OutApplyId = &v
	return s
}

func (s *InsureOrderRefundShrinkRequest) SetPolicyNoListShrink(v string) *InsureOrderRefundShrinkRequest {
	s.PolicyNoListShrink = &v
	return s
}

func (s *InsureOrderRefundShrinkRequest) SetSubInsOrderIdsShrink(v string) *InsureOrderRefundShrinkRequest {
	s.SubInsOrderIdsShrink = &v
	return s
}

func (s *InsureOrderRefundShrinkRequest) SetSupplierCode(v string) *InsureOrderRefundShrinkRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureOrderRefundShrinkRequest) Validate() error {
	return dara.Validate(s)
}
