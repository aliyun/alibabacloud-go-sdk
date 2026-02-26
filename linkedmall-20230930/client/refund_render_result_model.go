// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundRenderResult interface {
	dara.Model
	String() string
	GoString() string
	SetBizClaimType(v int32) *RefundRenderResult
	GetBizClaimType() *int32
	SetMaxRefundFeeData(v *DistributionMaxRefundFee) *RefundRenderResult
	GetMaxRefundFeeData() *DistributionMaxRefundFee
	SetOrderLineId(v string) *RefundRenderResult
	GetOrderLineId() *string
	SetRefundReasonList(v []*RefundReason) *RefundRenderResult
	GetRefundReasonList() []*RefundReason
	SetRequestId(v string) *RefundRenderResult
	GetRequestId() *string
}

type RefundRenderResult struct {
	// example:
	//
	// 1
	BizClaimType     *int32                    `json:"bizClaimType,omitempty" xml:"bizClaimType,omitempty"`
	MaxRefundFeeData *DistributionMaxRefundFee `json:"maxRefundFeeData,omitempty" xml:"maxRefundFeeData,omitempty"`
	// example:
	//
	// 6692****5458
	OrderLineId      *string         `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	RefundReasonList []*RefundReason `json:"refundReasonList,omitempty" xml:"refundReasonList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RefundRenderResult) String() string {
	return dara.Prettify(s)
}

func (s RefundRenderResult) GoString() string {
	return s.String()
}

func (s *RefundRenderResult) GetBizClaimType() *int32 {
	return s.BizClaimType
}

func (s *RefundRenderResult) GetMaxRefundFeeData() *DistributionMaxRefundFee {
	return s.MaxRefundFeeData
}

func (s *RefundRenderResult) GetOrderLineId() *string {
	return s.OrderLineId
}

func (s *RefundRenderResult) GetRefundReasonList() []*RefundReason {
	return s.RefundReasonList
}

func (s *RefundRenderResult) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundRenderResult) SetBizClaimType(v int32) *RefundRenderResult {
	s.BizClaimType = &v
	return s
}

func (s *RefundRenderResult) SetMaxRefundFeeData(v *DistributionMaxRefundFee) *RefundRenderResult {
	s.MaxRefundFeeData = v
	return s
}

func (s *RefundRenderResult) SetOrderLineId(v string) *RefundRenderResult {
	s.OrderLineId = &v
	return s
}

func (s *RefundRenderResult) SetRefundReasonList(v []*RefundReason) *RefundRenderResult {
	s.RefundReasonList = v
	return s
}

func (s *RefundRenderResult) SetRequestId(v string) *RefundRenderResult {
	s.RequestId = &v
	return s
}

func (s *RefundRenderResult) Validate() error {
	if s.MaxRefundFeeData != nil {
		if err := s.MaxRefundFeeData.Validate(); err != nil {
			return err
		}
	}
	if s.RefundReasonList != nil {
		for _, item := range s.RefundReasonList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
