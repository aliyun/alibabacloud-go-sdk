// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListLlmTemplatesResponseBodyData) *ListLlmTemplatesResponseBody
	GetData() []*ListLlmTemplatesResponseBodyData
	SetPageNumber(v int32) *ListLlmTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLlmTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLlmTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLlmTemplatesResponseBody
	GetTotalCount() *int32
}

type ListLlmTemplatesResponseBody struct {
	// The list of returned data objects.
	Data []*ListLlmTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The current page number of the query results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of query results per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of query results.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLlmTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponseBody) GetData() []*ListLlmTemplatesResponseBodyData {
	return s.Data
}

func (s *ListLlmTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLlmTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLlmTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLlmTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLlmTemplatesResponseBody) SetData(v []*ListLlmTemplatesResponseBodyData) *ListLlmTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListLlmTemplatesResponseBody) SetPageNumber(v int32) *ListLlmTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLlmTemplatesResponseBody) SetPageSize(v int32) *ListLlmTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLlmTemplatesResponseBody) SetRequestId(v string) *ListLlmTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLlmTemplatesResponseBody) SetTotalCount(v int32) *ListLlmTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLlmTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLlmTemplatesResponseBodyData struct {
	// The model configuration JSON object.
	//
	// example:
	//
	// {
	//
	// 	"id": "qwen3.6-plus",
	//
	// 	"cost": {
	//
	// 		"input": 0,
	//
	// 		"output": 0,
	//
	// 		"cacheRead": 0,
	//
	// 		"cacheWrite": 0
	//
	// 	},
	//
	// 	"name": "Qwen3.6-Plus",
	//
	// 	"input": ["image", "text"],
	//
	// 	"compat": {
	//
	// 		"supportsUsageInStreaming": true
	//
	// 	},
	//
	// 	"maxTokens": 65536,
	//
	// 	"reasoning": false,
	//
	// 	"contextWindow": 1000000
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The credit consumption multiplier (rate). A null value indicates that the model does not participate in credit-based billing.
	CreditMultiplier *ListLlmTemplatesResponseBodyDataCreditMultiplier `json:"CreditMultiplier,omitempty" xml:"CreditMultiplier,omitempty" type:"Struct"`
	// The template description.
	//
	// example:
	//
	// Qwen Plus series models
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of model features, such as function-calling, web-search, and structured-outputs.
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// The inference metadata, including request and response modalities.
	InferenceMetadata *ListLlmTemplatesResponseBodyDataInferenceMetadata `json:"InferenceMetadata,omitempty" xml:"InferenceMetadata,omitempty" type:"Struct"`
	// Indicates whether this is the default model under the associated model group.
	//
	// example:
	//
	// true
	IsDefaultModel *bool `json:"IsDefaultModel,omitempty" xml:"IsDefaultModel,omitempty"`
	// The model code.
	//
	// example:
	//
	// qwen3.6-plus
	LlmCode *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	// The model template ID.
	//
	// example:
	//
	// llmt-xxxx
	LlmTemplateId *string `json:"LlmTemplateId,omitempty" xml:"LlmTemplateId,omitempty"`
	// The model information, including context window size and maximum input/output tokens.
	ModelInfo map[string]interface{} `json:"ModelInfo,omitempty" xml:"ModelInfo,omitempty"`
	// The template name.
	//
	// example:
	//
	// Qwen3.6-Plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of price information.
	Prices []*ListLlmTemplatesResponseBodyDataPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Repeated"`
	// The ID of the model provider template.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
	// The publish time in ISO 8601 format, such as 2026-03-04T06:25:17.000+00:00.
	PublishedTime *string `json:"PublishedTime,omitempty" xml:"PublishedTime,omitempty"`
	// The authorization scope of the associated model group. Valid values: ALL_USER (all users), USER_MIXED (specified users and user groups), RESOURCE_MIXED (specified resources). Returned only when SmartModel is set to true.
	RefScope *string `json:"RefScope,omitempty" xml:"RefScope,omitempty"`
	// The number of route policies configured under this model tier. Returned only when SmartModel is set to true. Returns 0 for tiers without configured policies.
	RoutePolicyCount *int32 `json:"RoutePolicyCount,omitempty" xml:"RoutePolicyCount,omitempty"`
}

func (s ListLlmTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *ListLlmTemplatesResponseBodyData) GetCreditMultiplier() *ListLlmTemplatesResponseBodyDataCreditMultiplier {
	return s.CreditMultiplier
}

func (s *ListLlmTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListLlmTemplatesResponseBodyData) GetFeatures() []*string {
	return s.Features
}

func (s *ListLlmTemplatesResponseBodyData) GetInferenceMetadata() *ListLlmTemplatesResponseBodyDataInferenceMetadata {
	return s.InferenceMetadata
}

func (s *ListLlmTemplatesResponseBodyData) GetIsDefaultModel() *bool {
	return s.IsDefaultModel
}

func (s *ListLlmTemplatesResponseBodyData) GetLlmCode() *string {
	return s.LlmCode
}

func (s *ListLlmTemplatesResponseBodyData) GetLlmTemplateId() *string {
	return s.LlmTemplateId
}

func (s *ListLlmTemplatesResponseBodyData) GetModelInfo() map[string]interface{} {
	return s.ModelInfo
}

func (s *ListLlmTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListLlmTemplatesResponseBodyData) GetPrices() []*ListLlmTemplatesResponseBodyDataPrices {
	return s.Prices
}

func (s *ListLlmTemplatesResponseBodyData) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *ListLlmTemplatesResponseBodyData) GetPublishedTime() *string {
	return s.PublishedTime
}

func (s *ListLlmTemplatesResponseBodyData) GetRefScope() *string {
	return s.RefScope
}

func (s *ListLlmTemplatesResponseBodyData) GetRoutePolicyCount() *int32 {
	return s.RoutePolicyCount
}

func (s *ListLlmTemplatesResponseBodyData) SetConfig(v string) *ListLlmTemplatesResponseBodyData {
	s.Config = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetCreditMultiplier(v *ListLlmTemplatesResponseBodyDataCreditMultiplier) *ListLlmTemplatesResponseBodyData {
	s.CreditMultiplier = v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetDescription(v string) *ListLlmTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetFeatures(v []*string) *ListLlmTemplatesResponseBodyData {
	s.Features = v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetInferenceMetadata(v *ListLlmTemplatesResponseBodyDataInferenceMetadata) *ListLlmTemplatesResponseBodyData {
	s.InferenceMetadata = v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetIsDefaultModel(v bool) *ListLlmTemplatesResponseBodyData {
	s.IsDefaultModel = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetLlmCode(v string) *ListLlmTemplatesResponseBodyData {
	s.LlmCode = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetLlmTemplateId(v string) *ListLlmTemplatesResponseBodyData {
	s.LlmTemplateId = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetModelInfo(v map[string]interface{}) *ListLlmTemplatesResponseBodyData {
	s.ModelInfo = v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetName(v string) *ListLlmTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetPrices(v []*ListLlmTemplatesResponseBodyDataPrices) *ListLlmTemplatesResponseBodyData {
	s.Prices = v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetProviderTemplateId(v string) *ListLlmTemplatesResponseBodyData {
	s.ProviderTemplateId = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetPublishedTime(v string) *ListLlmTemplatesResponseBodyData {
	s.PublishedTime = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetRefScope(v string) *ListLlmTemplatesResponseBodyData {
	s.RefScope = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetRoutePolicyCount(v int32) *ListLlmTemplatesResponseBodyData {
	s.RoutePolicyCount = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) Validate() error {
	if s.CreditMultiplier != nil {
		if err := s.CreditMultiplier.Validate(); err != nil {
			return err
		}
	}
	if s.InferenceMetadata != nil {
		if err := s.InferenceMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.Prices != nil {
		for _, item := range s.Prices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLlmTemplatesResponseBodyDataCreditMultiplier struct {
	// The maximum multiplier. A null value indicates no upper limit. For example, Min=1 with Max as null is displayed as 1x and above.
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum multiplier. When equal to Max, it represents a fixed multiplier. For example, Min=Max=2 is displayed as 2x.
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s ListLlmTemplatesResponseBodyDataCreditMultiplier) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponseBodyDataCreditMultiplier) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponseBodyDataCreditMultiplier) GetMax() *float32 {
	return s.Max
}

func (s *ListLlmTemplatesResponseBodyDataCreditMultiplier) GetMin() *float32 {
	return s.Min
}

func (s *ListLlmTemplatesResponseBodyDataCreditMultiplier) SetMax(v float32) *ListLlmTemplatesResponseBodyDataCreditMultiplier {
	s.Max = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataCreditMultiplier) SetMin(v float32) *ListLlmTemplatesResponseBodyDataCreditMultiplier {
	s.Min = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataCreditMultiplier) Validate() error {
	return dara.Validate(s)
}

type ListLlmTemplatesResponseBodyDataInferenceMetadata struct {
	// The list of request modalities, such as Text, Image, and Audio.
	RequestModality []*string `json:"RequestModality,omitempty" xml:"RequestModality,omitempty" type:"Repeated"`
	// The list of response modalities, such as Text, Image, and Audio.
	ResponseModality []*string `json:"ResponseModality,omitempty" xml:"ResponseModality,omitempty" type:"Repeated"`
}

func (s ListLlmTemplatesResponseBodyDataInferenceMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponseBodyDataInferenceMetadata) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponseBodyDataInferenceMetadata) GetRequestModality() []*string {
	return s.RequestModality
}

func (s *ListLlmTemplatesResponseBodyDataInferenceMetadata) GetResponseModality() []*string {
	return s.ResponseModality
}

func (s *ListLlmTemplatesResponseBodyDataInferenceMetadata) SetRequestModality(v []*string) *ListLlmTemplatesResponseBodyDataInferenceMetadata {
	s.RequestModality = v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataInferenceMetadata) SetResponseModality(v []*string) *ListLlmTemplatesResponseBodyDataInferenceMetadata {
	s.ResponseModality = v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataInferenceMetadata) Validate() error {
	return dara.Validate(s)
}

type ListLlmTemplatesResponseBodyDataPrices struct {
	// The list of prices within the range.
	Prices []*ListLlmTemplatesResponseBodyDataPricesPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Repeated"`
	// The range name, such as Default or 0-1M tokens.
	RangeName *string `json:"RangeName,omitempty" xml:"RangeName,omitempty"`
}

func (s ListLlmTemplatesResponseBodyDataPrices) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponseBodyDataPrices) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponseBodyDataPrices) GetPrices() []*ListLlmTemplatesResponseBodyDataPricesPrices {
	return s.Prices
}

func (s *ListLlmTemplatesResponseBodyDataPrices) GetRangeName() *string {
	return s.RangeName
}

func (s *ListLlmTemplatesResponseBodyDataPrices) SetPrices(v []*ListLlmTemplatesResponseBodyDataPricesPrices) *ListLlmTemplatesResponseBodyDataPrices {
	s.Prices = v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataPrices) SetRangeName(v string) *ListLlmTemplatesResponseBodyDataPrices {
	s.RangeName = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataPrices) Validate() error {
	if s.Prices != nil {
		for _, item := range s.Prices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLlmTemplatesResponseBodyDataPricesPrices struct {
	// The price in string format, such as 0.2.
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
	// The price name, such as Input, Output, or Image Generation.
	PriceName *string `json:"PriceName,omitempty" xml:"PriceName,omitempty"`
	// The price unit, such as per image or per thousand tokens.
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
}

func (s ListLlmTemplatesResponseBodyDataPricesPrices) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponseBodyDataPricesPrices) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponseBodyDataPricesPrices) GetPrice() *string {
	return s.Price
}

func (s *ListLlmTemplatesResponseBodyDataPricesPrices) GetPriceName() *string {
	return s.PriceName
}

func (s *ListLlmTemplatesResponseBodyDataPricesPrices) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *ListLlmTemplatesResponseBodyDataPricesPrices) SetPrice(v string) *ListLlmTemplatesResponseBodyDataPricesPrices {
	s.Price = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataPricesPrices) SetPriceName(v string) *ListLlmTemplatesResponseBodyDataPricesPrices {
	s.PriceName = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataPricesPrices) SetPriceUnit(v string) *ListLlmTemplatesResponseBodyDataPricesPrices {
	s.PriceUnit = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyDataPricesPrices) Validate() error {
	return dara.Validate(s)
}
