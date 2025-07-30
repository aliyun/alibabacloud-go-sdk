// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModifyBEClusterInquiryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetModifyBEClusterInquiryResponseBodyData) *GetModifyBEClusterInquiryResponseBody
	GetData() *GetModifyBEClusterInquiryResponseBodyData
	SetRequestId(v string) *GetModifyBEClusterInquiryResponseBody
	GetRequestId() *string
}

type GetModifyBEClusterInquiryResponseBody struct {
	// The information returned.
	Data *GetModifyBEClusterInquiryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModifyBEClusterInquiryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponseBody) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponseBody) GetData() *GetModifyBEClusterInquiryResponseBodyData {
	return s.Data
}

func (s *GetModifyBEClusterInquiryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModifyBEClusterInquiryResponseBody) SetData(v *GetModifyBEClusterInquiryResponseBodyData) *GetModifyBEClusterInquiryResponseBody {
	s.Data = v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBody) SetRequestId(v string) *GetModifyBEClusterInquiryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetModifyBEClusterInquiryResponseBodyData struct {
	// The currency.
	//
	// example:
	//
	// CNY
	Currency     *string            `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PricingRules map[string]*string `json:"PricingRules,omitempty" xml:"PricingRules,omitempty"`
	// The estimated refund amount when the subscription cluster of a subscription instance is changed to a pay-as-you-go cluster.
	//
	// example:
	//
	// -100
	RefundAmount *string `json:"RefundAmount,omitempty" xml:"RefundAmount,omitempty"`
	// The amount of money.
	//
	// example:
	//
	// 1.76
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetModifyBEClusterInquiryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetPricingRules() map[string]*string {
	return s.PricingRules
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetRefundAmount() *string {
	return s.RefundAmount
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetTradeAmount() *string {
	return s.TradeAmount
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetCurrency(v string) *GetModifyBEClusterInquiryResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetPricingRules(v map[string]*string) *GetModifyBEClusterInquiryResponseBodyData {
	s.PricingRules = v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetRefundAmount(v string) *GetModifyBEClusterInquiryResponseBodyData {
	s.RefundAmount = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetTradeAmount(v string) *GetModifyBEClusterInquiryResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
