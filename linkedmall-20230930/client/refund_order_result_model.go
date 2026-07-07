// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundOrderResult interface {
	dara.Model
	String() string
	GoString() string
	SetDisputeId(v string) *RefundOrderResult
	GetDisputeId() *string
	SetDisputeStatus(v int32) *RefundOrderResult
	GetDisputeStatus() *int32
	SetOrderLineId(v string) *RefundOrderResult
	GetOrderLineId() *string
	SetRequestId(v string) *RefundOrderResult
	GetRequestId() *string
}

type RefundOrderResult struct {
	// Dispute ID
	//
	// example:
	//
	// 6693****4352
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
	// Status of the reverse operation
	//
	// 1 – Return pending
	//
	// 2 – Awaiting buyer return
	//
	// 3 – Awaiting merchant receipt
	//
	// 4 – Refund closed
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
	// Sub-distribution order ID for which the reverse operation was initiated
	//
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	// API request ID
	//
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RefundOrderResult) String() string {
	return dara.Prettify(s)
}

func (s RefundOrderResult) GoString() string {
	return s.String()
}

func (s *RefundOrderResult) GetDisputeId() *string {
	return s.DisputeId
}

func (s *RefundOrderResult) GetDisputeStatus() *int32 {
	return s.DisputeStatus
}

func (s *RefundOrderResult) GetOrderLineId() *string {
	return s.OrderLineId
}

func (s *RefundOrderResult) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundOrderResult) SetDisputeId(v string) *RefundOrderResult {
	s.DisputeId = &v
	return s
}

func (s *RefundOrderResult) SetDisputeStatus(v int32) *RefundOrderResult {
	s.DisputeStatus = &v
	return s
}

func (s *RefundOrderResult) SetOrderLineId(v string) *RefundOrderResult {
	s.OrderLineId = &v
	return s
}

func (s *RefundOrderResult) SetRequestId(v string) *RefundOrderResult {
	s.RequestId = &v
	return s
}

func (s *RefundOrderResult) Validate() error {
	return dara.Validate(s)
}
