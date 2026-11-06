// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderLlmCodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListModelProviderLlmCodesResponseBodyData) *ListModelProviderLlmCodesResponseBody
	GetData() *ListModelProviderLlmCodesResponseBodyData
	SetPageNumber(v int32) *ListModelProviderLlmCodesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelProviderLlmCodesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListModelProviderLlmCodesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListModelProviderLlmCodesResponseBody
	GetTotalCount() *int32
}

type ListModelProviderLlmCodesResponseBody struct {
	// The paginated data object that contains the list of LLM information.
	Data *ListModelProviderLlmCodesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records after filtering.
	//
	// example:
	//
	// 128
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelProviderLlmCodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesResponseBody) GetData() *ListModelProviderLlmCodesResponseBodyData {
	return s.Data
}

func (s *ListModelProviderLlmCodesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelProviderLlmCodesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelProviderLlmCodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelProviderLlmCodesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelProviderLlmCodesResponseBody) SetData(v *ListModelProviderLlmCodesResponseBodyData) *ListModelProviderLlmCodesResponseBody {
	s.Data = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBody) SetPageNumber(v int32) *ListModelProviderLlmCodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBody) SetPageSize(v int32) *ListModelProviderLlmCodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBody) SetRequestId(v string) *ListModelProviderLlmCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBody) SetTotalCount(v int32) *ListModelProviderLlmCodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListModelProviderLlmCodesResponseBodyData struct {
	// The list of LLM information.
	LlmInfoList []*ListModelProviderLlmCodesResponseBodyDataLlmInfoList `json:"LlmInfoList,omitempty" xml:"LlmInfoList,omitempty" type:"Repeated"`
}

func (s ListModelProviderLlmCodesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesResponseBodyData) GetLlmInfoList() []*ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	return s.LlmInfoList
}

func (s *ListModelProviderLlmCodesResponseBodyData) SetLlmInfoList(v []*ListModelProviderLlmCodesResponseBodyDataLlmInfoList) *ListModelProviderLlmCodesResponseBodyData {
	s.LlmInfoList = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyData) Validate() error {
	if s.LlmInfoList != nil {
		for _, item := range s.LlmInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelProviderLlmCodesResponseBodyDataLlmInfoList struct {
	// The credit consumption multiplier. An empty value indicates that the model does not participate in credit billing.
	CreditMultiplier *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier `json:"CreditMultiplier,omitempty" xml:"CreditMultiplier,omitempty" type:"Struct"`
	// The model description.
	//
	// example:
	//
	// Qwen image generation model that supports text-to-image and image editing capabilities
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of model features, such as function-calling, web-search, and structured-outputs.
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// The inference metadata, including request and response modalities.
	InferenceMetadata *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata `json:"InferenceMetadata,omitempty" xml:"InferenceMetadata,omitempty" type:"Struct"`
	// The LLM code, such as qwen-image-2.0.
	//
	// example:
	//
	// qwen-image-2.0
	LlmCode *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	// The model brand, such as Qwen, Wanx, or DeepSeek. This parameter is returned only for models annotated in the platform-side model pool.
	//
	// example:
	//
	// qwen
	ModelBrand *string `json:"ModelBrand,omitempty" xml:"ModelBrand,omitempty"`
	// The model information, such as the context window size and the maximum number of input/output tokens.
	//
	// example:
	//
	// {"contextWindow":1000000,"maxInputTokens":991808,"maxReasoningTokens":262144,"reasoningMaxInputTokens":983616,"maxOutputTokens":131072,"reasoningMaxOutputTokens":131072}
	ModelInfo map[string]interface{} `json:"ModelInfo,omitempty" xml:"ModelInfo,omitempty"`
	// The model series, such as Qwen-Max or DeepSeek. This parameter is returned only for the wuying-credit provider.
	//
	// example:
	//
	// Qwen-Max
	ModelSeries *string `json:"ModelSeries,omitempty" xml:"ModelSeries,omitempty"`
	// The model name, such as Qwen-Image-2.0.
	//
	// example:
	//
	// Qwen-Image-2.0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of price information.
	Prices []*ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Repeated"`
	// The time when the model was published, in ISO 8601 format, such as 2026-03-04T06:25:17.000+00:00.
	//
	// example:
	//
	// 2026-03-04T06:25:17.000+00:00
	PublishedTime *string `json:"PublishedTime,omitempty" xml:"PublishedTime,omitempty"`
	// Indicates whether the model is recommended. Valid values: true (recommended), false (explicitly not recommended), and null (not applicable, such as non-Bailian or unannotated models).
	Recommended *bool `json:"Recommended,omitempty" xml:"Recommended,omitempty"`
	// The model risk type. This parameter is returned only when includeRiskInfo is set to true.
	//
	// example:
	//
	// Normal
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetCreditMultiplier() *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier {
	return s.CreditMultiplier
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetDescription() *string {
	return s.Description
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetFeatures() []*string {
	return s.Features
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetInferenceMetadata() *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata {
	return s.InferenceMetadata
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetLlmCode() *string {
	return s.LlmCode
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetModelBrand() *string {
	return s.ModelBrand
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetModelInfo() map[string]interface{} {
	return s.ModelInfo
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetModelSeries() *string {
	return s.ModelSeries
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetName() *string {
	return s.Name
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetPrices() []*ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices {
	return s.Prices
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetPublishedTime() *string {
	return s.PublishedTime
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetRecommended() *bool {
	return s.Recommended
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) GetRiskType() *string {
	return s.RiskType
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetCreditMultiplier(v *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.CreditMultiplier = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetDescription(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.Description = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetFeatures(v []*string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.Features = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetInferenceMetadata(v *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.InferenceMetadata = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetLlmCode(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.LlmCode = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetModelBrand(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.ModelBrand = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetModelInfo(v map[string]interface{}) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.ModelInfo = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetModelSeries(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.ModelSeries = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetName(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.Name = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetPrices(v []*ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.Prices = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetPublishedTime(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.PublishedTime = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetRecommended(v bool) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.Recommended = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) SetRiskType(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoList {
	s.RiskType = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoList) Validate() error {
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

type ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier struct {
	// The upper limit of the multiplier. An empty value indicates no upper limit. For example, if the minimum multiplier is 1 and the maximum multiplier is empty, it is displayed as starting from 1x.
	//
	// example:
	//
	// 2
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The lower limit of the multiplier. If this value is equal to the maximum multiplier, it is a fixed multiplier. For example, if both the minimum and maximum multipliers are 2, it is displayed as 2x.
	//
	// example:
	//
	// 2
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier) GetMax() *float32 {
	return s.Max
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier) GetMin() *float32 {
	return s.Min
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier) SetMax(v float32) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier {
	s.Max = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier) SetMin(v float32) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier {
	s.Min = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListCreditMultiplier) Validate() error {
	return dara.Validate(s)
}

type ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata struct {
	// The list of request modalities, such as Text, Image, and Audio.
	RequestModality []*string `json:"RequestModality,omitempty" xml:"RequestModality,omitempty" type:"Repeated"`
	// The list of response modalities, such as Text, Image, and Audio.
	ResponseModality []*string `json:"ResponseModality,omitempty" xml:"ResponseModality,omitempty" type:"Repeated"`
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata) GetRequestModality() []*string {
	return s.RequestModality
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata) GetResponseModality() []*string {
	return s.ResponseModality
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata) SetRequestModality(v []*string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata {
	s.RequestModality = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata) SetResponseModality(v []*string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata {
	s.ResponseModality = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListInferenceMetadata) Validate() error {
	return dara.Validate(s)
}

type ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices struct {
	// The list of prices within the range.
	Prices []*ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Repeated"`
	// The name of the range, such as Default or 0-1M tokens.
	//
	// example:
	//
	// Default
	RangeName *string `json:"RangeName,omitempty" xml:"RangeName,omitempty"`
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices) GetPrices() []*ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices {
	return s.Prices
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices) GetRangeName() *string {
	return s.RangeName
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices) SetPrices(v []*ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices {
	s.Prices = v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices) SetRangeName(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices {
	s.RangeName = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPrices) Validate() error {
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

type ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices struct {
	// The price in string format, such as 0.2.
	//
	// example:
	//
	// 0.2
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
	// The name of the price item, such as input, output, or image generation.
	//
	// example:
	//
	// Input
	PriceName *string `json:"PriceName,omitempty" xml:"PriceName,omitempty"`
	// The price unit, such as per image or per 1,000 tokens.
	//
	// example:
	//
	// Per image
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) GetPrice() *string {
	return s.Price
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) GetPriceName() *string {
	return s.PriceName
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) SetPrice(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices {
	s.Price = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) SetPriceName(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices {
	s.PriceName = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) SetPriceUnit(v string) *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices {
	s.PriceUnit = &v
	return s
}

func (s *ListModelProviderLlmCodesResponseBodyDataLlmInfoListPricesPrices) Validate() error {
	return dara.Validate(s)
}
