// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateBEClusterInquiryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCreateBEClusterInquiryResponseBodyData) *GetCreateBEClusterInquiryResponseBody
	GetData() *GetCreateBEClusterInquiryResponseBodyData
	SetRequestId(v string) *GetCreateBEClusterInquiryResponseBody
	GetRequestId() *string
}

type GetCreateBEClusterInquiryResponseBody struct {
	// The information returned.
	Data *GetCreateBEClusterInquiryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCreateBEClusterInquiryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponseBody) GetData() *GetCreateBEClusterInquiryResponseBodyData {
	return s.Data
}

func (s *GetCreateBEClusterInquiryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCreateBEClusterInquiryResponseBody) SetData(v *GetCreateBEClusterInquiryResponseBodyData) *GetCreateBEClusterInquiryResponseBody {
	s.Data = v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBody) SetRequestId(v string) *GetCreateBEClusterInquiryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCreateBEClusterInquiryResponseBodyData struct {
	// The currency.
	//
	// example:
	//
	// CNY
	Currency     *string            `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PricingRules map[string]*string `json:"PricingRules,omitempty" xml:"PricingRules,omitempty"`
	// The amount of money.
	//
	// example:
	//
	// 1.76
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetCreateBEClusterInquiryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetCreateBEClusterInquiryResponseBodyData) GetPricingRules() map[string]*string {
	return s.PricingRules
}

func (s *GetCreateBEClusterInquiryResponseBodyData) GetTradeAmount() *string {
	return s.TradeAmount
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetCurrency(v string) *GetCreateBEClusterInquiryResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetPricingRules(v map[string]*string) *GetCreateBEClusterInquiryResponseBodyData {
	s.PricingRules = v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetTradeAmount(v string) *GetCreateBEClusterInquiryResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
