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
	// Current buyer\\"s refund request description
	//
	// example:
	//
	// 多拍不想要
	ApplyDisputeDesc *string `json:"applyDisputeDesc,omitempty" xml:"applyDisputeDesc,omitempty"`
	// Request reason
	ApplyReason *ApplyReason `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	// Order return method
	//
	// 1 – identity indicates refund only
	//
	// 3 – identity indicates return and refund
	//
	// example:
	//
	// 1
	BizClaimType *int32 `json:"bizClaimType,omitempty" xml:"bizClaimType,omitempty"`
	// Dispute creation time
	//
	// example:
	//
	// 2023-09-15T19:23:59.000+08:00
	DisputeCreateTime *string `json:"disputeCreateTime,omitempty" xml:"disputeCreateTime,omitempty"`
	// Reverse request description
	//
	// example:
	//
	// 多拍不想要
	DisputeDesc *string `json:"disputeDesc,omitempty" xml:"disputeDesc,omitempty"`
	// Reverse process end time
	//
	// example:
	//
	// 2023-09-15T19:23:59.000+08:00
	DisputeEndTime *string `json:"disputeEndTime,omitempty" xml:"disputeEndTime,omitempty"`
	// Reverse order ID
	//
	// example:
	//
	// 6693****4352
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
	// Reverse order status
	//
	// 1 – Return pending
	//
	// 2 – Awaiting buyer return
	//
	// 3 – Awaiting merchant receipt
	//
	// 4 – Refund shutdown
	//
	// 5 – Refund succeeded
	//
	// 6 – Refund denied
	//
	// 17 – Canceling refund
	//
	// example:
	//
	// 1
	DisputeStatus *int32 `json:"disputeStatus,omitempty" xml:"disputeStatus,omitempty"`
	// Main order ID
	//
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// Sub-order ID
	//
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	// Order logistics status
	//
	// 1 – Not shipped → Awaiting seller shipment
	//
	// 2 – Shipped → Awaiting buyer confirmation of receipt
	//
	// 3 – Received → Transaction succeeded
	//
	// 6 – Partially shipping
	//
	// 8 – Logistics order not yet created
	//
	// example:
	//
	// 1
	OrderLogisticsStatus *int32 `json:"orderLogisticsStatus,omitempty" xml:"orderLogisticsStatus,omitempty"`
	// Refund amount
	//
	// example:
	//
	// 1
	RefundFee *int64 `json:"refundFee,omitempty" xml:"refundFee,omitempty"`
	// Refund period
	RefundFeeData *RefundFeeData `json:"refundFeeData,omitempty" xml:"refundFeeData,omitempty"`
	// Merchant return address (available when disputeStatus=2, indicating the status is pending buyer return; save the return address during this status if needed)
	//
	// example:
	//
	// 阿里云云谷
	RefunderAddress *string `json:"refunderAddress,omitempty" xml:"refunderAddress,omitempty"`
	// Return recipient name
	//
	// example:
	//
	// 赵先生
	RefunderName *string `json:"refunderName,omitempty" xml:"refunderName,omitempty"`
	// Return contact information
	//
	// example:
	//
	// 182****1334
	RefunderTel *string `json:"refunderTel,omitempty" xml:"refunderTel,omitempty"`
	// Return address ZIP code
	//
	// example:
	//
	// 331001
	RefunderZipCode *string `json:"refunderZipCode,omitempty" xml:"refunderZipCode,omitempty"`
	// Request ID
	//
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return logistics status
	//
	// 0 – Return not initiated
	//
	// 1 – Awaiting pickup
	//
	// 2 – Package picked up
	//
	// 3 – In transit
	//
	// 4 – Out for delivery
	//
	// 5 – Delivered
	//
	// 6 – Delivery failed
	//
	// example:
	//
	// 0
	ReturnGoodLogisticsStatus *int32 `json:"returnGoodLogisticsStatus,omitempty" xml:"returnGoodLogisticsStatus,omitempty"`
	// Seller’s return approval message
	//
	// example:
	//
	// 同意退款
	SellerAgreeMsg *string `json:"sellerAgreeMsg,omitempty" xml:"sellerAgreeMsg,omitempty"`
	// Merchant\\"s message explaining the denial
	//
	// example:
	//
	// 不同意退款
	SellerRefuseAgreementMessage *string `json:"sellerRefuseAgreementMessage,omitempty" xml:"sellerRefuseAgreementMessage,omitempty"`
	// Merchant denial reason
	//
	// example:
	//
	// 商品没问题，买家举证无效
	SellerRefuseReason *string `json:"sellerRefuseReason,omitempty" xml:"sellerRefuseReason,omitempty"`
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
