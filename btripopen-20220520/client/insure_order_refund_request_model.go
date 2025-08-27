// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *InsureOrderRefundRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureOrderRefundRequest
	GetBuyerName() *string
	SetIsvName(v string) *InsureOrderRefundRequest
	GetIsvName() *string
	SetOutApplyId(v string) *InsureOrderRefundRequest
	GetOutApplyId() *string
	SetPolicyNoList(v []*string) *InsureOrderRefundRequest
	GetPolicyNoList() []*string
	SetSubInsOrderIds(v []*string) *InsureOrderRefundRequest
	GetSubInsOrderIds() []*string
	SetSupplierCode(v string) *InsureOrderRefundRequest
	GetSupplierCode() *string
}

type InsureOrderRefundRequest struct {
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
	OutApplyId     *string   `json:"out_apply_id,omitempty" xml:"out_apply_id,omitempty"`
	PolicyNoList   []*string `json:"policy_no_list,omitempty" xml:"policy_no_list,omitempty" type:"Repeated"`
	SubInsOrderIds []*string `json:"sub_ins_order_ids,omitempty" xml:"sub_ins_order_ids,omitempty" type:"Repeated"`
	// example:
	//
	// fliggy
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s InsureOrderRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderRefundRequest) GoString() string {
	return s.String()
}

func (s *InsureOrderRefundRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderRefundRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureOrderRefundRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureOrderRefundRequest) GetOutApplyId() *string {
	return s.OutApplyId
}

func (s *InsureOrderRefundRequest) GetPolicyNoList() []*string {
	return s.PolicyNoList
}

func (s *InsureOrderRefundRequest) GetSubInsOrderIds() []*string {
	return s.SubInsOrderIds
}

func (s *InsureOrderRefundRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureOrderRefundRequest) SetBtripUserId(v string) *InsureOrderRefundRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderRefundRequest) SetBuyerName(v string) *InsureOrderRefundRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureOrderRefundRequest) SetIsvName(v string) *InsureOrderRefundRequest {
	s.IsvName = &v
	return s
}

func (s *InsureOrderRefundRequest) SetOutApplyId(v string) *InsureOrderRefundRequest {
	s.OutApplyId = &v
	return s
}

func (s *InsureOrderRefundRequest) SetPolicyNoList(v []*string) *InsureOrderRefundRequest {
	s.PolicyNoList = v
	return s
}

func (s *InsureOrderRefundRequest) SetSubInsOrderIds(v []*string) *InsureOrderRefundRequest {
	s.SubInsOrderIds = v
	return s
}

func (s *InsureOrderRefundRequest) SetSupplierCode(v string) *InsureOrderRefundRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureOrderRefundRequest) Validate() error {
	return dara.Validate(s)
}
