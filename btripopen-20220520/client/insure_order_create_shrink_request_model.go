// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCreateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicantShrink(v string) *InsureOrderCreateShrinkRequest
	GetApplicantShrink() *string
	SetBtripUserId(v string) *InsureOrderCreateShrinkRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureOrderCreateShrinkRequest
	GetBuyerName() *string
	SetInsPersonAndSegmentListShrink(v string) *InsureOrderCreateShrinkRequest
	GetInsPersonAndSegmentListShrink() *string
	SetIsvName(v string) *InsureOrderCreateShrinkRequest
	GetIsvName() *string
	SetOutInsOrderId(v string) *InsureOrderCreateShrinkRequest
	GetOutInsOrderId() *string
	SetOutOrderId(v string) *InsureOrderCreateShrinkRequest
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *InsureOrderCreateShrinkRequest
	GetOutSubOrderId() *string
	SetSupplierCode(v string) *InsureOrderCreateShrinkRequest
	GetSupplierCode() *string
}

type InsureOrderCreateShrinkRequest struct {
	// This parameter is required.
	ApplicantShrink *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	// example:
	//
	// 20202109390122
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	InsPersonAndSegmentListShrink *string `json:"ins_person_and_segment_list,omitempty" xml:"ins_person_and_segment_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PostalSavingsBank
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 100000320302020
	OutInsOrderId *string `json:"out_ins_order_id,omitempty" xml:"out_ins_order_id,omitempty"`
	// This parameter is required.
	//
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

func (s InsureOrderCreateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateShrinkRequest) GetApplicantShrink() *string {
	return s.ApplicantShrink
}

func (s *InsureOrderCreateShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderCreateShrinkRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureOrderCreateShrinkRequest) GetInsPersonAndSegmentListShrink() *string {
	return s.InsPersonAndSegmentListShrink
}

func (s *InsureOrderCreateShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureOrderCreateShrinkRequest) GetOutInsOrderId() *string {
	return s.OutInsOrderId
}

func (s *InsureOrderCreateShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *InsureOrderCreateShrinkRequest) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *InsureOrderCreateShrinkRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureOrderCreateShrinkRequest) SetApplicantShrink(v string) *InsureOrderCreateShrinkRequest {
	s.ApplicantShrink = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) SetBtripUserId(v string) *InsureOrderCreateShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) SetBuyerName(v string) *InsureOrderCreateShrinkRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) SetInsPersonAndSegmentListShrink(v string) *InsureOrderCreateShrinkRequest {
	s.InsPersonAndSegmentListShrink = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) SetIsvName(v string) *InsureOrderCreateShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) SetOutInsOrderId(v string) *InsureOrderCreateShrinkRequest {
	s.OutInsOrderId = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) SetOutOrderId(v string) *InsureOrderCreateShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) SetOutSubOrderId(v string) *InsureOrderCreateShrinkRequest {
	s.OutSubOrderId = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) SetSupplierCode(v string) *InsureOrderCreateShrinkRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureOrderCreateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
