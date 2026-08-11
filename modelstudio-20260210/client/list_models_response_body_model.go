// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelsResponseBody
	GetCode() *string
	SetErrorMessage(v string) *ListModelsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int64) *ListModelsResponseBody
	GetHttpStatusCode() *int64
	SetMaxResults(v int64) *ListModelsResponseBody
	GetMaxResults() *int64
	SetModels(v []*ListModelsResponseBodyModels) *ListModelsResponseBody
	GetModels() []*ListModelsResponseBodyModels
	SetNextToken(v string) *ListModelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListModelsResponseBody
	GetTotalCount() *int64
}

type ListModelsResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// The specified parameter is invalid
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64                          `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Models     []*ListModelsResponseBodyModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	// example:
	//
	// lwytFRtLdNk=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 36045E0A-551D-592D-B1BC-4C56596CE59E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListModelsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListModelsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListModelsResponseBody) GetModels() []*ListModelsResponseBodyModels {
	return s.Models
}

func (s *ListModelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelsResponseBody) SetCode(v string) *ListModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelsResponseBody) SetErrorMessage(v string) *ListModelsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListModelsResponseBody) SetHttpStatusCode(v int64) *ListModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListModelsResponseBody) SetMaxResults(v int64) *ListModelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelsResponseBody) SetModels(v []*ListModelsResponseBodyModels) *ListModelsResponseBody {
	s.Models = v
	return s
}

func (s *ListModelsResponseBody) SetNextToken(v string) *ListModelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetSuccess(v bool) *ListModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelsResponseBody) SetTotalCount(v int64) *ListModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelsResponseBody) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelsResponseBodyModels struct {
	Capabilities []*string `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Repeated"`
	// example:
	//
	// provided by qwen
	Description       *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	Features          []*string                                      `json:"features,omitempty" xml:"features,omitempty" type:"Repeated"`
	InferenceMetadata *ListModelsResponseBodyModelsInferenceMetadata `json:"inferenceMetadata,omitempty" xml:"inferenceMetadata,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Model     *string                                `json:"model,omitempty" xml:"model,omitempty"`
	ModelInfo *ListModelsResponseBodyModelsModelInfo `json:"modelInfo,omitempty" xml:"modelInfo,omitempty" type:"Struct"`
	// example:
	//
	// qwen-plus
	Name   *string                               `json:"name,omitempty" xml:"name,omitempty"`
	Prices []*ListModelsResponseBodyModelsPrices `json:"prices,omitempty" xml:"prices,omitempty" type:"Repeated"`
	// example:
	//
	// qwen
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// 1779268196000
	PublishedTime *int64 `json:"publishedTime,omitempty" xml:"publishedTime,omitempty"`
}

func (s ListModelsResponseBodyModels) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyModels) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyModels) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *ListModelsResponseBodyModels) GetDescription() *string {
	return s.Description
}

func (s *ListModelsResponseBodyModels) GetFeatures() []*string {
	return s.Features
}

func (s *ListModelsResponseBodyModels) GetInferenceMetadata() *ListModelsResponseBodyModelsInferenceMetadata {
	return s.InferenceMetadata
}

func (s *ListModelsResponseBodyModels) GetModel() *string {
	return s.Model
}

func (s *ListModelsResponseBodyModels) GetModelInfo() *ListModelsResponseBodyModelsModelInfo {
	return s.ModelInfo
}

func (s *ListModelsResponseBodyModels) GetName() *string {
	return s.Name
}

func (s *ListModelsResponseBodyModels) GetPrices() []*ListModelsResponseBodyModelsPrices {
	return s.Prices
}

func (s *ListModelsResponseBodyModels) GetProvider() *string {
	return s.Provider
}

func (s *ListModelsResponseBodyModels) GetPublishedTime() *int64 {
	return s.PublishedTime
}

func (s *ListModelsResponseBodyModels) SetCapabilities(v []*string) *ListModelsResponseBodyModels {
	s.Capabilities = v
	return s
}

func (s *ListModelsResponseBodyModels) SetDescription(v string) *ListModelsResponseBodyModels {
	s.Description = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetFeatures(v []*string) *ListModelsResponseBodyModels {
	s.Features = v
	return s
}

func (s *ListModelsResponseBodyModels) SetInferenceMetadata(v *ListModelsResponseBodyModelsInferenceMetadata) *ListModelsResponseBodyModels {
	s.InferenceMetadata = v
	return s
}

func (s *ListModelsResponseBodyModels) SetModel(v string) *ListModelsResponseBodyModels {
	s.Model = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetModelInfo(v *ListModelsResponseBodyModelsModelInfo) *ListModelsResponseBodyModels {
	s.ModelInfo = v
	return s
}

func (s *ListModelsResponseBodyModels) SetName(v string) *ListModelsResponseBodyModels {
	s.Name = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetPrices(v []*ListModelsResponseBodyModelsPrices) *ListModelsResponseBodyModels {
	s.Prices = v
	return s
}

func (s *ListModelsResponseBodyModels) SetProvider(v string) *ListModelsResponseBodyModels {
	s.Provider = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetPublishedTime(v int64) *ListModelsResponseBodyModels {
	s.PublishedTime = &v
	return s
}

func (s *ListModelsResponseBodyModels) Validate() error {
	if s.InferenceMetadata != nil {
		if err := s.InferenceMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.ModelInfo != nil {
		if err := s.ModelInfo.Validate(); err != nil {
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

type ListModelsResponseBodyModelsInferenceMetadata struct {
	RequestModality  []*string `json:"requestModality,omitempty" xml:"requestModality,omitempty" type:"Repeated"`
	ResponseModality []*string `json:"responseModality,omitempty" xml:"responseModality,omitempty" type:"Repeated"`
}

func (s ListModelsResponseBodyModelsInferenceMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyModelsInferenceMetadata) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyModelsInferenceMetadata) GetRequestModality() []*string {
	return s.RequestModality
}

func (s *ListModelsResponseBodyModelsInferenceMetadata) GetResponseModality() []*string {
	return s.ResponseModality
}

func (s *ListModelsResponseBodyModelsInferenceMetadata) SetRequestModality(v []*string) *ListModelsResponseBodyModelsInferenceMetadata {
	s.RequestModality = v
	return s
}

func (s *ListModelsResponseBodyModelsInferenceMetadata) SetResponseModality(v []*string) *ListModelsResponseBodyModelsInferenceMetadata {
	s.ResponseModality = v
	return s
}

func (s *ListModelsResponseBodyModelsInferenceMetadata) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyModelsModelInfo struct {
	// example:
	//
	// 10
	ContextWindow *int64 `json:"contextWindow,omitempty" xml:"contextWindow,omitempty"`
	// example:
	//
	// 10
	MaxInputTokens *int64 `json:"maxInputTokens,omitempty" xml:"maxInputTokens,omitempty"`
	// example:
	//
	// 10
	MaxOutputTokens *int64 `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	// example:
	//
	// 10
	MaxReasoningTokens *int64 `json:"maxReasoningTokens,omitempty" xml:"maxReasoningTokens,omitempty"`
	// example:
	//
	// 10
	ReasoningMaxInputTokens *int64 `json:"reasoningMaxInputTokens,omitempty" xml:"reasoningMaxInputTokens,omitempty"`
	// example:
	//
	// 10
	ReasoningMaxOutputTokens *int64 `json:"reasoningMaxOutputTokens,omitempty" xml:"reasoningMaxOutputTokens,omitempty"`
}

func (s ListModelsResponseBodyModelsModelInfo) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyModelsModelInfo) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyModelsModelInfo) GetContextWindow() *int64 {
	return s.ContextWindow
}

func (s *ListModelsResponseBodyModelsModelInfo) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *ListModelsResponseBodyModelsModelInfo) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *ListModelsResponseBodyModelsModelInfo) GetMaxReasoningTokens() *int64 {
	return s.MaxReasoningTokens
}

func (s *ListModelsResponseBodyModelsModelInfo) GetReasoningMaxInputTokens() *int64 {
	return s.ReasoningMaxInputTokens
}

func (s *ListModelsResponseBodyModelsModelInfo) GetReasoningMaxOutputTokens() *int64 {
	return s.ReasoningMaxOutputTokens
}

func (s *ListModelsResponseBodyModelsModelInfo) SetContextWindow(v int64) *ListModelsResponseBodyModelsModelInfo {
	s.ContextWindow = &v
	return s
}

func (s *ListModelsResponseBodyModelsModelInfo) SetMaxInputTokens(v int64) *ListModelsResponseBodyModelsModelInfo {
	s.MaxInputTokens = &v
	return s
}

func (s *ListModelsResponseBodyModelsModelInfo) SetMaxOutputTokens(v int64) *ListModelsResponseBodyModelsModelInfo {
	s.MaxOutputTokens = &v
	return s
}

func (s *ListModelsResponseBodyModelsModelInfo) SetMaxReasoningTokens(v int64) *ListModelsResponseBodyModelsModelInfo {
	s.MaxReasoningTokens = &v
	return s
}

func (s *ListModelsResponseBodyModelsModelInfo) SetReasoningMaxInputTokens(v int64) *ListModelsResponseBodyModelsModelInfo {
	s.ReasoningMaxInputTokens = &v
	return s
}

func (s *ListModelsResponseBodyModelsModelInfo) SetReasoningMaxOutputTokens(v int64) *ListModelsResponseBodyModelsModelInfo {
	s.ReasoningMaxOutputTokens = &v
	return s
}

func (s *ListModelsResponseBodyModelsModelInfo) Validate() error {
	return dara.Validate(s)
}

type ListModelsResponseBodyModelsPrices struct {
	Prices []*ListModelsResponseBodyModelsPricesPrices `json:"prices,omitempty" xml:"prices,omitempty" type:"Repeated"`
	// example:
	//
	// Default
	RangeName *string `json:"rangeName,omitempty" xml:"rangeName,omitempty"`
}

func (s ListModelsResponseBodyModelsPrices) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyModelsPrices) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyModelsPrices) GetPrices() []*ListModelsResponseBodyModelsPricesPrices {
	return s.Prices
}

func (s *ListModelsResponseBodyModelsPrices) GetRangeName() *string {
	return s.RangeName
}

func (s *ListModelsResponseBodyModelsPrices) SetPrices(v []*ListModelsResponseBodyModelsPricesPrices) *ListModelsResponseBodyModelsPrices {
	s.Prices = v
	return s
}

func (s *ListModelsResponseBodyModelsPrices) SetRangeName(v string) *ListModelsResponseBodyModelsPrices {
	s.RangeName = &v
	return s
}

func (s *ListModelsResponseBodyModelsPrices) Validate() error {
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

type ListModelsResponseBodyModelsPricesPrices struct {
	// example:
	//
	// 12
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// input
	PriceName *string `json:"priceName,omitempty" xml:"priceName,omitempty"`
	// example:
	//
	// Per 1M tokens
	PriceUnit *string `json:"priceUnit,omitempty" xml:"priceUnit,omitempty"`
}

func (s ListModelsResponseBodyModelsPricesPrices) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyModelsPricesPrices) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyModelsPricesPrices) GetPrice() *string {
	return s.Price
}

func (s *ListModelsResponseBodyModelsPricesPrices) GetPriceName() *string {
	return s.PriceName
}

func (s *ListModelsResponseBodyModelsPricesPrices) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *ListModelsResponseBodyModelsPricesPrices) SetPrice(v string) *ListModelsResponseBodyModelsPricesPrices {
	s.Price = &v
	return s
}

func (s *ListModelsResponseBodyModelsPricesPrices) SetPriceName(v string) *ListModelsResponseBodyModelsPricesPrices {
	s.PriceName = &v
	return s
}

func (s *ListModelsResponseBodyModelsPricesPrices) SetPriceUnit(v string) *ListModelsResponseBodyModelsPricesPrices {
	s.PriceUnit = &v
	return s
}

func (s *ListModelsResponseBodyModelsPricesPrices) Validate() error {
	return dara.Validate(s)
}
