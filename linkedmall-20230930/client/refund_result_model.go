// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundResult interface {
	dara.Model
	String() string
	GoString() string
	SetApplyDisputeDesc(v string) *RefundResult
	GetApplyDisputeDesc() *string
	SetApplyReason(v *ApplyReason) *RefundResult
	GetApplyReason() *ApplyReason
	SetBizClaimType(v int32) *RefundResult
	GetBizClaimType() *int32
	SetDisputeCreateTime(v string) *RefundResult
	GetDisputeCreateTime() *string
	SetDisputeDesc(v string) *RefundResult
	GetDisputeDesc() *string
	SetDisputeEndTime(v string) *RefundResult
	GetDisputeEndTime() *string
	SetDisputeId(v string) *RefundResult
	GetDisputeId() *string
	SetDisputeStatus(v int32) *RefundResult
	GetDisputeStatus() *int32
	SetOrderId(v string) *RefundResult
	GetOrderId() *string
	SetOrderLineId(v string) *RefundResult
	GetOrderLineId() *string
	SetOrderLogisticsStatus(v int32) *RefundResult
	GetOrderLogisticsStatus() *int32
	SetRefundFee(v int64) *RefundResult
	GetRefundFee() *int64
	SetRefundFeeData(v *RefundFeeData) *RefundResult
	GetRefundFeeData() *RefundFeeData
	SetRefunderAddress(v string) *RefundResult
	GetRefunderAddress() *string
	SetRefunderName(v string) *RefundResult
	GetRefunderName() *string
	SetRefunderTel(v string) *RefundResult
	GetRefunderTel() *string
	SetRefunderZipCode(v string) *RefundResult
	GetRefunderZipCode() *string
	SetRequestId(v string) *RefundResult
	GetRequestId() *string
	SetReturnGoodLogisticsStatus(v int32) *RefundResult
	GetReturnGoodLogisticsStatus() *int32
	SetSellerAgreeMsg(v string) *RefundResult
	GetSellerAgreeMsg() *string
	SetSellerRefuseAgreementMessage(v string) *RefundResult
	GetSellerRefuseAgreementMessage() *string
	SetSellerRefuseReason(v string) *RefundResult
	GetSellerRefuseReason() *string
}

type RefundResult struct {
	ApplyDisputeDesc *string      `json:"applyDisputeDesc,omitempty" xml:"applyDisputeDesc,omitempty"`
	ApplyReason      *ApplyReason `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	// example:
	//
	// 1
	BizClaimType *int32 `json:"bizClaimType,omitempty" xml:"bizClaimType,omitempty"`
	// example:
	//
	// 2023-09-02T00:00:00.000Z
	DisputeCreateTime *string `json:"disputeCreateTime,omitempty" xml:"disputeCreateTime,omitempty"`
	DisputeDesc       *string `json:"disputeDesc,omitempty" xml:"disputeDesc,omitempty"`
	// example:
	//
	// 2023-09-02T12:00:00.000Z
	DisputeEndTime *string `json:"disputeEndTime,omitempty" xml:"disputeEndTime,omitempty"`
	// example:
	//
	// 6693****4352
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
	// example:
	//
	// 1
	DisputeStatus *int32 `json:"disputeStatus,omitempty" xml:"disputeStatus,omitempty"`
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	// example:
	//
	// 1
	OrderLogisticsStatus *int32 `json:"orderLogisticsStatus,omitempty" xml:"orderLogisticsStatus,omitempty"`
	// example:
	//
	// 1
	RefundFee       *int64         `json:"refundFee,omitempty" xml:"refundFee,omitempty"`
	RefundFeeData   *RefundFeeData `json:"refundFeeData,omitempty" xml:"refundFeeData,omitempty"`
	RefunderAddress *string        `json:"refunderAddress,omitempty" xml:"refunderAddress,omitempty"`
	RefunderName    *string        `json:"refunderName,omitempty" xml:"refunderName,omitempty"`
	// example:
	//
	// 182****1334
	RefunderTel *string `json:"refunderTel,omitempty" xml:"refunderTel,omitempty"`
	// example:
	//
	// 331001
	RefunderZipCode *string `json:"refunderZipCode,omitempty" xml:"refunderZipCode,omitempty"`
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0
	ReturnGoodLogisticsStatus    *int32  `json:"returnGoodLogisticsStatus,omitempty" xml:"returnGoodLogisticsStatus,omitempty"`
	SellerAgreeMsg               *string `json:"sellerAgreeMsg,omitempty" xml:"sellerAgreeMsg,omitempty"`
	SellerRefuseAgreementMessage *string `json:"sellerRefuseAgreementMessage,omitempty" xml:"sellerRefuseAgreementMessage,omitempty"`
	SellerRefuseReason           *string `json:"sellerRefuseReason,omitempty" xml:"sellerRefuseReason,omitempty"`
}

func (s RefundResult) String() string {
	return dara.Prettify(s)
}

func (s RefundResult) GoString() string {
	return s.String()
}

func (s *RefundResult) GetApplyDisputeDesc() *string {
	return s.ApplyDisputeDesc
}

func (s *RefundResult) GetApplyReason() *ApplyReason {
	return s.ApplyReason
}

func (s *RefundResult) GetBizClaimType() *int32 {
	return s.BizClaimType
}

func (s *RefundResult) GetDisputeCreateTime() *string {
	return s.DisputeCreateTime
}

func (s *RefundResult) GetDisputeDesc() *string {
	return s.DisputeDesc
}

func (s *RefundResult) GetDisputeEndTime() *string {
	return s.DisputeEndTime
}

func (s *RefundResult) GetDisputeId() *string {
	return s.DisputeId
}

func (s *RefundResult) GetDisputeStatus() *int32 {
	return s.DisputeStatus
}

func (s *RefundResult) GetOrderId() *string {
	return s.OrderId
}

func (s *RefundResult) GetOrderLineId() *string {
	return s.OrderLineId
}

func (s *RefundResult) GetOrderLogisticsStatus() *int32 {
	return s.OrderLogisticsStatus
}

func (s *RefundResult) GetRefundFee() *int64 {
	return s.RefundFee
}

func (s *RefundResult) GetRefundFeeData() *RefundFeeData {
	return s.RefundFeeData
}

func (s *RefundResult) GetRefunderAddress() *string {
	return s.RefunderAddress
}

func (s *RefundResult) GetRefunderName() *string {
	return s.RefunderName
}

func (s *RefundResult) GetRefunderTel() *string {
	return s.RefunderTel
}

func (s *RefundResult) GetRefunderZipCode() *string {
	return s.RefunderZipCode
}

func (s *RefundResult) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundResult) GetReturnGoodLogisticsStatus() *int32 {
	return s.ReturnGoodLogisticsStatus
}

func (s *RefundResult) GetSellerAgreeMsg() *string {
	return s.SellerAgreeMsg
}

func (s *RefundResult) GetSellerRefuseAgreementMessage() *string {
	return s.SellerRefuseAgreementMessage
}

func (s *RefundResult) GetSellerRefuseReason() *string {
	return s.SellerRefuseReason
}

func (s *RefundResult) SetApplyDisputeDesc(v string) *RefundResult {
	s.ApplyDisputeDesc = &v
	return s
}

func (s *RefundResult) SetApplyReason(v *ApplyReason) *RefundResult {
	s.ApplyReason = v
	return s
}

func (s *RefundResult) SetBizClaimType(v int32) *RefundResult {
	s.BizClaimType = &v
	return s
}

func (s *RefundResult) SetDisputeCreateTime(v string) *RefundResult {
	s.DisputeCreateTime = &v
	return s
}

func (s *RefundResult) SetDisputeDesc(v string) *RefundResult {
	s.DisputeDesc = &v
	return s
}

func (s *RefundResult) SetDisputeEndTime(v string) *RefundResult {
	s.DisputeEndTime = &v
	return s
}

func (s *RefundResult) SetDisputeId(v string) *RefundResult {
	s.DisputeId = &v
	return s
}

func (s *RefundResult) SetDisputeStatus(v int32) *RefundResult {
	s.DisputeStatus = &v
	return s
}

func (s *RefundResult) SetOrderId(v string) *RefundResult {
	s.OrderId = &v
	return s
}

func (s *RefundResult) SetOrderLineId(v string) *RefundResult {
	s.OrderLineId = &v
	return s
}

func (s *RefundResult) SetOrderLogisticsStatus(v int32) *RefundResult {
	s.OrderLogisticsStatus = &v
	return s
}

func (s *RefundResult) SetRefundFee(v int64) *RefundResult {
	s.RefundFee = &v
	return s
}

func (s *RefundResult) SetRefundFeeData(v *RefundFeeData) *RefundResult {
	s.RefundFeeData = v
	return s
}

func (s *RefundResult) SetRefunderAddress(v string) *RefundResult {
	s.RefunderAddress = &v
	return s
}

func (s *RefundResult) SetRefunderName(v string) *RefundResult {
	s.RefunderName = &v
	return s
}

func (s *RefundResult) SetRefunderTel(v string) *RefundResult {
	s.RefunderTel = &v
	return s
}

func (s *RefundResult) SetRefunderZipCode(v string) *RefundResult {
	s.RefunderZipCode = &v
	return s
}

func (s *RefundResult) SetRequestId(v string) *RefundResult {
	s.RequestId = &v
	return s
}

func (s *RefundResult) SetReturnGoodLogisticsStatus(v int32) *RefundResult {
	s.ReturnGoodLogisticsStatus = &v
	return s
}

func (s *RefundResult) SetSellerAgreeMsg(v string) *RefundResult {
	s.SellerAgreeMsg = &v
	return s
}

func (s *RefundResult) SetSellerRefuseAgreementMessage(v string) *RefundResult {
	s.SellerRefuseAgreementMessage = &v
	return s
}

func (s *RefundResult) SetSellerRefuseReason(v string) *RefundResult {
	s.SellerRefuseReason = &v
	return s
}

func (s *RefundResult) Validate() error {
	if s.ApplyReason != nil {
		if err := s.ApplyReason.Validate(); err != nil {
			return err
		}
	}
	if s.RefundFeeData != nil {
		if err := s.RefundFeeData.Validate(); err != nil {
			return err
		}
	}
	return nil
}
