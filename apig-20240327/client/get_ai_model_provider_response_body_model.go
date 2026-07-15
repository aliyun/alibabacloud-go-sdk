// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAiModelProviderResponseBody
	GetCode() *string
	SetData(v *GetAiModelProviderResponseBodyData) *GetAiModelProviderResponseBody
	GetData() *GetAiModelProviderResponseBodyData
	SetMessage(v string) *GetAiModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiModelProviderResponseBody
	GetRequestId() *string
}

type GetAiModelProviderResponseBody struct {
	// example:
	//
	// Ok
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetAiModelProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAiModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiModelProviderResponseBody) GetData() *GetAiModelProviderResponseBodyData {
	return s.Data
}

func (s *GetAiModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiModelProviderResponseBody) SetCode(v string) *GetAiModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiModelProviderResponseBody) SetData(v *GetAiModelProviderResponseBodyData) *GetAiModelProviderResponseBody {
	s.Data = v
	return s
}

func (s *GetAiModelProviderResponseBody) SetMessage(v string) *GetAiModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiModelProviderResponseBody) SetRequestId(v string) *GetAiModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiModelProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiModelProviderResponseBodyData struct {
	BoundServices []*ServiceInfo `json:"boundServices,omitempty" xml:"boundServices,omitempty" type:"Repeated"`
	// example:
	//
	// 千问云 / 阿里云百炼
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId  *string                                         `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	ModelCards []*GetAiModelProviderResponseBodyDataModelCards `json:"modelCards,omitempty" xml:"modelCards,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	ModelCount *int32 `json:"modelCount,omitempty" xml:"modelCount,omitempty"`
	// example:
	//
	// mp-8c13d2b4f8a1
	ModelProviderId *string `json:"modelProviderId,omitempty" xml:"modelProviderId,omitempty"`
	// example:
	//
	// qwen
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// user
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// 2026-07-14 18:30:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetAiModelProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiModelProviderResponseBodyData) GetBoundServices() []*ServiceInfo {
	return s.BoundServices
}

func (s *GetAiModelProviderResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAiModelProviderResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetAiModelProviderResponseBodyData) GetModelCards() []*GetAiModelProviderResponseBodyDataModelCards {
	return s.ModelCards
}

func (s *GetAiModelProviderResponseBodyData) GetModelCount() *int32 {
	return s.ModelCount
}

func (s *GetAiModelProviderResponseBodyData) GetModelProviderId() *string {
	return s.ModelProviderId
}

func (s *GetAiModelProviderResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *GetAiModelProviderResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetAiModelProviderResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAiModelProviderResponseBodyData) SetBoundServices(v []*ServiceInfo) *GetAiModelProviderResponseBodyData {
	s.BoundServices = v
	return s
}

func (s *GetAiModelProviderResponseBodyData) SetDisplayName(v string) *GetAiModelProviderResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetAiModelProviderResponseBodyData) SetGatewayId(v string) *GetAiModelProviderResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetAiModelProviderResponseBodyData) SetModelCards(v []*GetAiModelProviderResponseBodyDataModelCards) *GetAiModelProviderResponseBodyData {
	s.ModelCards = v
	return s
}

func (s *GetAiModelProviderResponseBodyData) SetModelCount(v int32) *GetAiModelProviderResponseBodyData {
	s.ModelCount = &v
	return s
}

func (s *GetAiModelProviderResponseBodyData) SetModelProviderId(v string) *GetAiModelProviderResponseBodyData {
	s.ModelProviderId = &v
	return s
}

func (s *GetAiModelProviderResponseBodyData) SetProvider(v string) *GetAiModelProviderResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetAiModelProviderResponseBodyData) SetSource(v string) *GetAiModelProviderResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetAiModelProviderResponseBodyData) SetUpdateTime(v string) *GetAiModelProviderResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAiModelProviderResponseBodyData) Validate() error {
	if s.BoundServices != nil {
		for _, item := range s.BoundServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModelCards != nil {
		for _, item := range s.ModelCards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiModelProviderResponseBodyDataModelCards struct {
	AvailablePaths []*GetAiModelProviderResponseBodyDataModelCardsAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	Credit         *GetAiModelProviderResponseBodyDataModelCardsCredit           `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// example:
	//
	// {"functionCalling":true,"toolChoice":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string                                           `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Meta      *GetAiModelProviderResponseBodyDataModelCardsMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// example:
	//
	// mc-8c13d2b4f8a1
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	// example:
	//
	// qwen-plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// qwen
	ModelProvider *string `json:"modelProvider,omitempty" xml:"modelProvider,omitempty"`
	// example:
	//
	// user
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// http://https://dashscope-intl.aliyuncs.com
	SourceURL *string `json:"sourceURL,omitempty" xml:"sourceURL,omitempty"`
	// example:
	//
	// 2026-07-14 18:30:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetAiModelProviderResponseBodyDataModelCards) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelProviderResponseBodyDataModelCards) GoString() string {
	return s.String()
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetAvailablePaths() []*GetAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	return s.AvailablePaths
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetCredit() *GetAiModelProviderResponseBodyDataModelCardsCredit {
	return s.Credit
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetMeta() *GetAiModelProviderResponseBodyDataModelCardsMeta {
	return s.Meta
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetModelName() *string {
	return s.ModelName
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetSource() *string {
	return s.Source
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetSourceURL() *string {
	return s.SourceURL
}

func (s *GetAiModelProviderResponseBodyDataModelCards) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetAvailablePaths(v []*GetAiModelProviderResponseBodyDataModelCardsAvailablePaths) *GetAiModelProviderResponseBodyDataModelCards {
	s.AvailablePaths = v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetCredit(v *GetAiModelProviderResponseBodyDataModelCardsCredit) *GetAiModelProviderResponseBodyDataModelCards {
	s.Credit = v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetFeatures(v map[string]interface{}) *GetAiModelProviderResponseBodyDataModelCards {
	s.Features = v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetGatewayId(v string) *GetAiModelProviderResponseBodyDataModelCards {
	s.GatewayId = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetMeta(v *GetAiModelProviderResponseBodyDataModelCardsMeta) *GetAiModelProviderResponseBodyDataModelCards {
	s.Meta = v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetModelCardId(v string) *GetAiModelProviderResponseBodyDataModelCards {
	s.ModelCardId = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetModelName(v string) *GetAiModelProviderResponseBodyDataModelCards {
	s.ModelName = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetModelProvider(v string) *GetAiModelProviderResponseBodyDataModelCards {
	s.ModelProvider = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetSource(v string) *GetAiModelProviderResponseBodyDataModelCards {
	s.Source = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetSourceURL(v string) *GetAiModelProviderResponseBodyDataModelCards {
	s.SourceURL = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) SetUpdateTime(v string) *GetAiModelProviderResponseBodyDataModelCards {
	s.UpdateTime = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCards) Validate() error {
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

type GetAiModelProviderResponseBodyDataModelCardsAvailablePaths struct {
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAiModelProviderResponseBodyDataModelCardsAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelProviderResponseBodyDataModelCardsAvailablePaths) GoString() string {
	return s.String()
}

func (s *GetAiModelProviderResponseBodyDataModelCardsAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *GetAiModelProviderResponseBodyDataModelCardsAvailablePaths) GetType() *string {
	return s.Type
}

func (s *GetAiModelProviderResponseBodyDataModelCardsAvailablePaths) SetPath(v string) *GetAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	s.Path = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsAvailablePaths) SetType(v string) *GetAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	s.Type = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type GetAiModelProviderResponseBodyDataModelCardsCredit struct {
	// example:
	//
	// 0.5
	CacheCost *float32 `json:"cacheCost,omitempty" xml:"cacheCost,omitempty"`
	// example:
	//
	// 1.5
	InputCost *float32 `json:"inputCost,omitempty" xml:"inputCost,omitempty"`
	// example:
	//
	// 3
	OutputCost *float32 `json:"outputCost,omitempty" xml:"outputCost,omitempty"`
	// example:
	//
	// fixed
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAiModelProviderResponseBodyDataModelCardsCredit) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelProviderResponseBodyDataModelCardsCredit) GoString() string {
	return s.String()
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) GetType() *string {
	return s.Type
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) SetCacheCost(v float32) *GetAiModelProviderResponseBodyDataModelCardsCredit {
	s.CacheCost = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) SetInputCost(v float32) *GetAiModelProviderResponseBodyDataModelCardsCredit {
	s.InputCost = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) SetOutputCost(v float32) *GetAiModelProviderResponseBodyDataModelCardsCredit {
	s.OutputCost = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) SetType(v string) *GetAiModelProviderResponseBodyDataModelCardsCredit {
	s.Type = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsCredit) Validate() error {
	return dara.Validate(s)
}

type GetAiModelProviderResponseBodyDataModelCardsMeta struct {
	// example:
	//
	// 131072
	MaxInputTokens *int64 `json:"maxInputTokens,omitempty" xml:"maxInputTokens,omitempty"`
	// example:
	//
	// 8192
	MaxOutputTokens *int64 `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	// example:
	//
	// 131072
	MaxTokens                 *int64    `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	SupportedInputModalities  []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s GetAiModelProviderResponseBodyDataModelCardsMeta) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelProviderResponseBodyDataModelCardsMeta) GoString() string {
	return s.String()
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) SetMaxInputTokens(v int64) *GetAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) SetMaxOutputTokens(v int64) *GetAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) SetMaxTokens(v int64) *GetAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxTokens = &v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) SetSupportedInputModalities(v []*string) *GetAiModelProviderResponseBodyDataModelCardsMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) SetSupportedOutputModalities(v []*string) *GetAiModelProviderResponseBodyDataModelCardsMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *GetAiModelProviderResponseBodyDataModelCardsMeta) Validate() error {
	return dara.Validate(s)
}
