// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectSupplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDirectSupplyResponseBody
	GetCode() *string
	SetData(v *GetDirectSupplyResponseBodyData) *GetDirectSupplyResponseBody
	GetData() *GetDirectSupplyResponseBodyData
	SetMessage(v string) *GetDirectSupplyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDirectSupplyResponseBody
	GetRequestId() *string
}

type GetDirectSupplyResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetDirectSupplyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDirectSupplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDirectSupplyResponseBody) GetData() *GetDirectSupplyResponseBodyData {
	return s.Data
}

func (s *GetDirectSupplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDirectSupplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDirectSupplyResponseBody) SetCode(v string) *GetDirectSupplyResponseBody {
	s.Code = &v
	return s
}

func (s *GetDirectSupplyResponseBody) SetData(v *GetDirectSupplyResponseBodyData) *GetDirectSupplyResponseBody {
	s.Data = v
	return s
}

func (s *GetDirectSupplyResponseBody) SetMessage(v string) *GetDirectSupplyResponseBody {
	s.Message = &v
	return s
}

func (s *GetDirectSupplyResponseBody) SetRequestId(v string) *GetDirectSupplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDirectSupplyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDirectSupplyResponseBodyData struct {
	// The list of candidate models that can be added to direct supply.
	CandidateModels []*GetDirectSupplyResponseBodyDataCandidateModels `json:"candidateModels,omitempty" xml:"candidateModels,omitempty" type:"Repeated"`
	// The error message returned when the model catalog query fails. This parameter is not returned when the query succeeds.
	CatalogError *GetDirectSupplyResponseBodyDataCatalogError `json:"catalogError,omitempty" xml:"catalogError,omitempty" type:"Struct"`
	// Indicates whether direct supply is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The ID of the gateway instance.
	//
	// example:
	//
	// gw-xxxxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The list of models for which direct supply is enabled.
	Models []*GetDirectSupplyResponseBodyDataModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
}

func (s GetDirectSupplyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyData) GetCandidateModels() []*GetDirectSupplyResponseBodyDataCandidateModels {
	return s.CandidateModels
}

func (s *GetDirectSupplyResponseBodyData) GetCatalogError() *GetDirectSupplyResponseBodyDataCatalogError {
	return s.CatalogError
}

func (s *GetDirectSupplyResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetDirectSupplyResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetDirectSupplyResponseBodyData) GetModels() []*GetDirectSupplyResponseBodyDataModels {
	return s.Models
}

func (s *GetDirectSupplyResponseBodyData) SetCandidateModels(v []*GetDirectSupplyResponseBodyDataCandidateModels) *GetDirectSupplyResponseBodyData {
	s.CandidateModels = v
	return s
}

func (s *GetDirectSupplyResponseBodyData) SetCatalogError(v *GetDirectSupplyResponseBodyDataCatalogError) *GetDirectSupplyResponseBodyData {
	s.CatalogError = v
	return s
}

func (s *GetDirectSupplyResponseBodyData) SetEnabled(v bool) *GetDirectSupplyResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetDirectSupplyResponseBodyData) SetGatewayId(v string) *GetDirectSupplyResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetDirectSupplyResponseBodyData) SetModels(v []*GetDirectSupplyResponseBodyDataModels) *GetDirectSupplyResponseBodyData {
	s.Models = v
	return s
}

func (s *GetDirectSupplyResponseBodyData) Validate() error {
	if s.CandidateModels != nil {
		for _, item := range s.CandidateModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CatalogError != nil {
		if err := s.CatalogError.Validate(); err != nil {
			return err
		}
	}
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

type GetDirectSupplyResponseBodyDataCandidateModels struct {
	// The list of available invocation paths.
	AvailablePaths []*GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	// The billing information of the model.
	Credit *GetDirectSupplyResponseBodyDataCandidateModelsCredit `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// The capabilities and features supported by the model.
	//
	// example:
	//
	// {"nativeStreaming":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// The ID of the gateway instance.
	//
	// example:
	//
	// gw-xxxxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The capability information of the model.
	Meta *GetDirectSupplyResponseBodyDataCandidateModelsMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The ID of the model card, which starts with msc-.
	//
	// example:
	//
	// msc-xxxxxxxx
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model provider.
	//
	// example:
	//
	// qwen
	ModelProvider *string `json:"modelProvider,omitempty" xml:"modelProvider,omitempty"`
	// The ID of the model source.
	//
	// example:
	//
	// model-source-xxxxxxxx
	ModelSourceId *string `json:"modelSourceId,omitempty" xml:"modelSourceId,omitempty"`
	// The source of the model.
	//
	// example:
	//
	// source-xxxxxxxx
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The last update time in ISO 8601 (RFC 3339) format in the UTC time zone, such as 2026-05-15T10:00:00Z.
	//
	// example:
	//
	// 2026-05-15T10:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetDirectSupplyResponseBodyDataCandidateModels) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataCandidateModels) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetAvailablePaths() []*GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths {
	return s.AvailablePaths
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetCredit() *GetDirectSupplyResponseBodyDataCandidateModelsCredit {
	return s.Credit
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetMeta() *GetDirectSupplyResponseBodyDataCandidateModelsMeta {
	return s.Meta
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetModelName() *string {
	return s.ModelName
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetModelSourceId() *string {
	return s.ModelSourceId
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetSource() *string {
	return s.Source
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetAvailablePaths(v []*GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.AvailablePaths = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetCredit(v *GetDirectSupplyResponseBodyDataCandidateModelsCredit) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.Credit = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetFeatures(v map[string]interface{}) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.Features = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetGatewayId(v string) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.GatewayId = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetMeta(v *GetDirectSupplyResponseBodyDataCandidateModelsMeta) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.Meta = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetModelCardId(v string) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.ModelCardId = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetModelName(v string) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.ModelName = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetModelProvider(v string) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.ModelProvider = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetModelSourceId(v string) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.ModelSourceId = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetSource(v string) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.Source = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) SetUpdateTime(v string) *GetDirectSupplyResponseBodyDataCandidateModels {
	s.UpdateTime = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModels) Validate() error {
	if s.AvailablePaths != nil {
		for _, item := range s.AvailablePaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Credit != nil {
		if err := s.Credit.Validate(); err != nil {
			return err
		}
	}
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths struct {
	// The invocation path.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The path type.
	//
	// example:
	//
	// type-xxxxxxxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths) GetType() *string {
	return s.Type
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths) SetPath(v string) *GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths {
	s.Path = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths) SetType(v string) *GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths {
	s.Type = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type GetDirectSupplyResponseBodyDataCandidateModelsCredit struct {
	// The unit price of cache-hit tokens.
	//
	// example:
	//
	// 0.001
	CacheCost *float32 `json:"cacheCost,omitempty" xml:"cacheCost,omitempty"`
	// The unit price of input tokens.
	//
	// example:
	//
	// 0.004
	InputCost *float32 `json:"inputCost,omitempty" xml:"inputCost,omitempty"`
	// The unit price of output tokens.
	//
	// example:
	//
	// 0.012
	OutputCost *float32 `json:"outputCost,omitempty" xml:"outputCost,omitempty"`
	// The billing type.
	//
	// example:
	//
	// type-xxxxxxxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDirectSupplyResponseBodyDataCandidateModelsCredit) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataCandidateModelsCredit) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) GetType() *string {
	return s.Type
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) SetCacheCost(v float32) *GetDirectSupplyResponseBodyDataCandidateModelsCredit {
	s.CacheCost = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) SetInputCost(v float32) *GetDirectSupplyResponseBodyDataCandidateModelsCredit {
	s.InputCost = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) SetOutputCost(v float32) *GetDirectSupplyResponseBodyDataCandidateModelsCredit {
	s.OutputCost = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) SetType(v string) *GetDirectSupplyResponseBodyDataCandidateModelsCredit {
	s.Type = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsCredit) Validate() error {
	return dara.Validate(s)
}

type GetDirectSupplyResponseBodyDataCandidateModelsMeta struct {
	// The maximum number of input tokens.
	//
	// example:
	//
	// 128000
	MaxInputTokens *int64 `json:"maxInputTokens,omitempty" xml:"maxInputTokens,omitempty"`
	// The maximum number of output tokens.
	//
	// example:
	//
	// 16384
	MaxOutputTokens *int64 `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	// The maximum total number of tokens.
	//
	// example:
	//
	// 16384
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The list of supported input modalities.
	SupportedInputModalities []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	// The list of supported output modalities.
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s GetDirectSupplyResponseBodyDataCandidateModelsMeta) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataCandidateModelsMeta) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) SetMaxInputTokens(v int64) *GetDirectSupplyResponseBodyDataCandidateModelsMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) SetMaxOutputTokens(v int64) *GetDirectSupplyResponseBodyDataCandidateModelsMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) SetMaxTokens(v int64) *GetDirectSupplyResponseBodyDataCandidateModelsMeta {
	s.MaxTokens = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) SetSupportedInputModalities(v []*string) *GetDirectSupplyResponseBodyDataCandidateModelsMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) SetSupportedOutputModalities(v []*string) *GetDirectSupplyResponseBodyDataCandidateModelsMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCandidateModelsMeta) Validate() error {
	return dara.Validate(s)
}

type GetDirectSupplyResponseBodyDataCatalogError struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Model catalog query failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDirectSupplyResponseBodyDataCatalogError) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataCatalogError) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataCatalogError) GetCode() *string {
	return s.Code
}

func (s *GetDirectSupplyResponseBodyDataCatalogError) GetMessage() *string {
	return s.Message
}

func (s *GetDirectSupplyResponseBodyDataCatalogError) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDirectSupplyResponseBodyDataCatalogError) SetCode(v string) *GetDirectSupplyResponseBodyDataCatalogError {
	s.Code = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCatalogError) SetMessage(v string) *GetDirectSupplyResponseBodyDataCatalogError {
	s.Message = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCatalogError) SetRequestId(v string) *GetDirectSupplyResponseBodyDataCatalogError {
	s.RequestId = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataCatalogError) Validate() error {
	return dara.Validate(s)
}

type GetDirectSupplyResponseBodyDataModels struct {
	// The list of available invocation paths.
	AvailablePaths []*GetDirectSupplyResponseBodyDataModelsAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	// The billing information of the model.
	Credit *GetDirectSupplyResponseBodyDataModelsCredit `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// The capabilities and features supported by the model.
	//
	// example:
	//
	// {"nativeStreaming":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// The ID of the gateway instance.
	//
	// example:
	//
	// gw-xxxxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The capability information of the model.
	Meta *GetDirectSupplyResponseBodyDataModelsMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The ID of the model card, which starts with msc-.
	//
	// example:
	//
	// msc-xxxxxxxx
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model provider.
	//
	// example:
	//
	// qwen
	ModelProvider *string `json:"modelProvider,omitempty" xml:"modelProvider,omitempty"`
	// The ID of the model source.
	//
	// example:
	//
	// model-source-xxxxxxxx
	ModelSourceId *string `json:"modelSourceId,omitempty" xml:"modelSourceId,omitempty"`
	// The source of the model.
	//
	// example:
	//
	// source-xxxxxxxx
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The last update time in ISO 8601 (RFC 3339) format in the UTC time zone, such as 2026-05-15T10:00:00Z.
	//
	// example:
	//
	// 2026-05-15T10:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetDirectSupplyResponseBodyDataModels) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataModels) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataModels) GetAvailablePaths() []*GetDirectSupplyResponseBodyDataModelsAvailablePaths {
	return s.AvailablePaths
}

func (s *GetDirectSupplyResponseBodyDataModels) GetCredit() *GetDirectSupplyResponseBodyDataModelsCredit {
	return s.Credit
}

func (s *GetDirectSupplyResponseBodyDataModels) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *GetDirectSupplyResponseBodyDataModels) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetDirectSupplyResponseBodyDataModels) GetMeta() *GetDirectSupplyResponseBodyDataModelsMeta {
	return s.Meta
}

func (s *GetDirectSupplyResponseBodyDataModels) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *GetDirectSupplyResponseBodyDataModels) GetModelName() *string {
	return s.ModelName
}

func (s *GetDirectSupplyResponseBodyDataModels) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *GetDirectSupplyResponseBodyDataModels) GetModelSourceId() *string {
	return s.ModelSourceId
}

func (s *GetDirectSupplyResponseBodyDataModels) GetSource() *string {
	return s.Source
}

func (s *GetDirectSupplyResponseBodyDataModels) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDirectSupplyResponseBodyDataModels) SetAvailablePaths(v []*GetDirectSupplyResponseBodyDataModelsAvailablePaths) *GetDirectSupplyResponseBodyDataModels {
	s.AvailablePaths = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetCredit(v *GetDirectSupplyResponseBodyDataModelsCredit) *GetDirectSupplyResponseBodyDataModels {
	s.Credit = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetFeatures(v map[string]interface{}) *GetDirectSupplyResponseBodyDataModels {
	s.Features = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetGatewayId(v string) *GetDirectSupplyResponseBodyDataModels {
	s.GatewayId = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetMeta(v *GetDirectSupplyResponseBodyDataModelsMeta) *GetDirectSupplyResponseBodyDataModels {
	s.Meta = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetModelCardId(v string) *GetDirectSupplyResponseBodyDataModels {
	s.ModelCardId = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetModelName(v string) *GetDirectSupplyResponseBodyDataModels {
	s.ModelName = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetModelProvider(v string) *GetDirectSupplyResponseBodyDataModels {
	s.ModelProvider = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetModelSourceId(v string) *GetDirectSupplyResponseBodyDataModels {
	s.ModelSourceId = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetSource(v string) *GetDirectSupplyResponseBodyDataModels {
	s.Source = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) SetUpdateTime(v string) *GetDirectSupplyResponseBodyDataModels {
	s.UpdateTime = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModels) Validate() error {
	if s.AvailablePaths != nil {
		for _, item := range s.AvailablePaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Credit != nil {
		if err := s.Credit.Validate(); err != nil {
			return err
		}
	}
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDirectSupplyResponseBodyDataModelsAvailablePaths struct {
	// The invocation path.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The path type.
	//
	// example:
	//
	// type-xxxxxxxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDirectSupplyResponseBodyDataModelsAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataModelsAvailablePaths) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataModelsAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *GetDirectSupplyResponseBodyDataModelsAvailablePaths) GetType() *string {
	return s.Type
}

func (s *GetDirectSupplyResponseBodyDataModelsAvailablePaths) SetPath(v string) *GetDirectSupplyResponseBodyDataModelsAvailablePaths {
	s.Path = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsAvailablePaths) SetType(v string) *GetDirectSupplyResponseBodyDataModelsAvailablePaths {
	s.Type = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type GetDirectSupplyResponseBodyDataModelsCredit struct {
	// The unit price of cache-hit tokens.
	//
	// example:
	//
	// 0.001
	CacheCost *float32 `json:"cacheCost,omitempty" xml:"cacheCost,omitempty"`
	// The unit price of input tokens.
	//
	// example:
	//
	// 0.004
	InputCost *float32 `json:"inputCost,omitempty" xml:"inputCost,omitempty"`
	// The unit price of output tokens.
	//
	// example:
	//
	// 0.012
	OutputCost *float32 `json:"outputCost,omitempty" xml:"outputCost,omitempty"`
	// The billing type.
	//
	// example:
	//
	// type-xxxxxxxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDirectSupplyResponseBodyDataModelsCredit) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataModelsCredit) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) GetType() *string {
	return s.Type
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) SetCacheCost(v float32) *GetDirectSupplyResponseBodyDataModelsCredit {
	s.CacheCost = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) SetInputCost(v float32) *GetDirectSupplyResponseBodyDataModelsCredit {
	s.InputCost = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) SetOutputCost(v float32) *GetDirectSupplyResponseBodyDataModelsCredit {
	s.OutputCost = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) SetType(v string) *GetDirectSupplyResponseBodyDataModelsCredit {
	s.Type = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsCredit) Validate() error {
	return dara.Validate(s)
}

type GetDirectSupplyResponseBodyDataModelsMeta struct {
	// The maximum number of input tokens.
	//
	// example:
	//
	// 128000
	MaxInputTokens *int64 `json:"maxInputTokens,omitempty" xml:"maxInputTokens,omitempty"`
	// The maximum number of output tokens.
	//
	// example:
	//
	// 16384
	MaxOutputTokens *int64 `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	// The maximum total number of tokens.
	//
	// example:
	//
	// 16384
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The list of supported input modalities.
	SupportedInputModalities []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	// The list of supported output modalities.
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s GetDirectSupplyResponseBodyDataModelsMeta) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponseBodyDataModelsMeta) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) SetMaxInputTokens(v int64) *GetDirectSupplyResponseBodyDataModelsMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) SetMaxOutputTokens(v int64) *GetDirectSupplyResponseBodyDataModelsMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) SetMaxTokens(v int64) *GetDirectSupplyResponseBodyDataModelsMeta {
	s.MaxTokens = &v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) SetSupportedInputModalities(v []*string) *GetDirectSupplyResponseBodyDataModelsMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) SetSupportedOutputModalities(v []*string) *GetDirectSupplyResponseBodyDataModelsMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *GetDirectSupplyResponseBodyDataModelsMeta) Validate() error {
	return dara.Validate(s)
}
