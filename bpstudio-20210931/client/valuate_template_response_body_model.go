// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValuateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ValuateTemplateResponseBody
	GetCode() *string
	SetData(v *ValuateTemplateResponseBodyData) *ValuateTemplateResponseBody
	GetData() *ValuateTemplateResponseBodyData
	SetMessage(v string) *ValuateTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ValuateTemplateResponseBody
	GetRequestId() *string
}

type ValuateTemplateResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the inquiry.
	Data *ValuateTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 847C9D0A-BABD-589C-8A9C-6464409EDED9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValuateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValuateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ValuateTemplateResponseBody) GetData() *ValuateTemplateResponseBodyData {
	return s.Data
}

func (s *ValuateTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ValuateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValuateTemplateResponseBody) SetCode(v string) *ValuateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ValuateTemplateResponseBody) SetData(v *ValuateTemplateResponseBodyData) *ValuateTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ValuateTemplateResponseBody) SetMessage(v string) *ValuateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ValuateTemplateResponseBody) SetRequestId(v string) *ValuateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValuateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ValuateTemplateResponseBodyData struct {
	// The result set of the inquiry.
	ResourceList []*ValuateTemplateResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
}

func (s ValuateTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ValuateTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponseBodyData) GetResourceList() []*ValuateTemplateResponseBodyDataResourceList {
	return s.ResourceList
}

func (s *ValuateTemplateResponseBodyData) SetResourceList(v []*ValuateTemplateResponseBodyDataResourceList) *ValuateTemplateResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *ValuateTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ValuateTemplateResponseBodyDataResourceList struct {
	// The discount amount.
	//
	// example:
	//
	// 73
	DiscountAmount *float64 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// InvalidSaleComponentFault : The request not refer to the correct order sale component.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ecs
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The original price.
	//
	// example:
	//
	// 83.0
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The information about the price.
	PriceList []*ValuateTemplateResponseBodyDataResourceListPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// The pricing unit.
	//
	// example:
	//
	// USD
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The discount information.
	//
	// example:
	//
	// The discount information.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The price at which the transaction is made.
	//
	// example:
	//
	// 10.0
	TradePrice *float64 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s ValuateTemplateResponseBodyDataResourceList) String() string {
	return dara.Prettify(s)
}

func (s ValuateTemplateResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponseBodyDataResourceList) GetDiscountAmount() *float64 {
	return s.DiscountAmount
}

func (s *ValuateTemplateResponseBodyDataResourceList) GetError() *string {
	return s.Error
}

func (s *ValuateTemplateResponseBodyDataResourceList) GetNodeType() *string {
	return s.NodeType
}

func (s *ValuateTemplateResponseBodyDataResourceList) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *ValuateTemplateResponseBodyDataResourceList) GetPriceList() []*ValuateTemplateResponseBodyDataResourceListPriceList {
	return s.PriceList
}

func (s *ValuateTemplateResponseBodyDataResourceList) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *ValuateTemplateResponseBodyDataResourceList) GetPromotionName() *string {
	return s.PromotionName
}

func (s *ValuateTemplateResponseBodyDataResourceList) GetTradePrice() *float64 {
	return s.TradePrice
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetDiscountAmount(v float64) *ValuateTemplateResponseBodyDataResourceList {
	s.DiscountAmount = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetError(v string) *ValuateTemplateResponseBodyDataResourceList {
	s.Error = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetNodeType(v string) *ValuateTemplateResponseBodyDataResourceList {
	s.NodeType = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetOriginalPrice(v float64) *ValuateTemplateResponseBodyDataResourceList {
	s.OriginalPrice = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetPriceList(v []*ValuateTemplateResponseBodyDataResourceListPriceList) *ValuateTemplateResponseBodyDataResourceList {
	s.PriceList = v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetPriceUnit(v string) *ValuateTemplateResponseBodyDataResourceList {
	s.PriceUnit = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetPromotionName(v string) *ValuateTemplateResponseBodyDataResourceList {
	s.PromotionName = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetTradePrice(v float64) *ValuateTemplateResponseBodyDataResourceList {
	s.TradePrice = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) Validate() error {
	return dara.Validate(s)
}

type ValuateTemplateResponseBodyDataResourceListPriceList struct {
	// The discount amount.
	//
	// example:
	//
	// 82.99
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// ServiceUnavailable : The request has failed due to a temporary failure of the server.\\r\\nRequestId : 4AA302DB-3286-5589-8637-FF6D8507B7A9.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The resource type.
	//
	// example:
	//
	// eip
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The original price.
	//
	// example:
	//
	// 83.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The pricing unit.
	//
	// example:
	//
	// USD
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The discount information.
	//
	// example:
	//
	// The discount information.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 1687225092
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The price at which the transaction is made.
	//
	// example:
	//
	// 0.01
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	// Indicates whether the instance is newly created. Valid values:\\
	//
	// 1: The instance is newly created.\\
	//
	// 2: The instance already exists.\\
	//
	// 0: The price of the instance is not included.
	//
	// example:
	//
	// "1"
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ValuateTemplateResponseBodyDataResourceListPriceList) String() string {
	return dara.Prettify(s)
}

func (s ValuateTemplateResponseBodyDataResourceListPriceList) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetError() *string {
	return s.Error
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetNodeType() *string {
	return s.NodeType
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetPromotionName() *string {
	return s.PromotionName
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) GetType() *string {
	return s.Type
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetDiscountAmount(v float32) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.DiscountAmount = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetError(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.Error = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetNodeType(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.NodeType = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetOriginalPrice(v float32) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.OriginalPrice = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetPriceUnit(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.PriceUnit = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetPromotionName(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.PromotionName = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetResourceId(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.ResourceId = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetTradePrice(v float32) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.TradePrice = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetType(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.Type = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) Validate() error {
	return dara.Validate(s)
}
