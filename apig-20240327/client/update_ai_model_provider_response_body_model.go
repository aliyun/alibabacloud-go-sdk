// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAiModelProviderResponseBody
	GetCode() *string
	SetData(v *UpdateAiModelProviderResponseBodyData) *UpdateAiModelProviderResponseBody
	GetData() *UpdateAiModelProviderResponseBodyData
	SetMessage(v string) *UpdateAiModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAiModelProviderResponseBody
	GetRequestId() *string
}

type UpdateAiModelProviderResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateAiModelProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAiModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAiModelProviderResponseBody) GetData() *UpdateAiModelProviderResponseBodyData {
	return s.Data
}

func (s *UpdateAiModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAiModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAiModelProviderResponseBody) SetCode(v string) *UpdateAiModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAiModelProviderResponseBody) SetData(v *UpdateAiModelProviderResponseBodyData) *UpdateAiModelProviderResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAiModelProviderResponseBody) SetMessage(v string) *UpdateAiModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAiModelProviderResponseBody) SetRequestId(v string) *UpdateAiModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAiModelProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAiModelProviderResponseBodyData struct {
	BoundServices []*ServiceInfo `json:"boundServices,omitempty" xml:"boundServices,omitempty" type:"Repeated"`
	// example:
	//
	// 千问云 / 阿里云百炼
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId  *string                                            `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	ModelCards []*UpdateAiModelProviderResponseBodyDataModelCards `json:"modelCards,omitempty" xml:"modelCards,omitempty" type:"Repeated"`
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

func (s UpdateAiModelProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderResponseBodyData) GetBoundServices() []*ServiceInfo {
	return s.BoundServices
}

func (s *UpdateAiModelProviderResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAiModelProviderResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateAiModelProviderResponseBodyData) GetModelCards() []*UpdateAiModelProviderResponseBodyDataModelCards {
	return s.ModelCards
}

func (s *UpdateAiModelProviderResponseBodyData) GetModelCount() *int32 {
	return s.ModelCount
}

func (s *UpdateAiModelProviderResponseBodyData) GetModelProviderId() *string {
	return s.ModelProviderId
}

func (s *UpdateAiModelProviderResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *UpdateAiModelProviderResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *UpdateAiModelProviderResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateAiModelProviderResponseBodyData) SetBoundServices(v []*ServiceInfo) *UpdateAiModelProviderResponseBodyData {
	s.BoundServices = v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) SetDisplayName(v string) *UpdateAiModelProviderResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) SetGatewayId(v string) *UpdateAiModelProviderResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) SetModelCards(v []*UpdateAiModelProviderResponseBodyDataModelCards) *UpdateAiModelProviderResponseBodyData {
	s.ModelCards = v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) SetModelCount(v int32) *UpdateAiModelProviderResponseBodyData {
	s.ModelCount = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) SetModelProviderId(v string) *UpdateAiModelProviderResponseBodyData {
	s.ModelProviderId = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) SetProvider(v string) *UpdateAiModelProviderResponseBodyData {
	s.Provider = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) SetSource(v string) *UpdateAiModelProviderResponseBodyData {
	s.Source = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) SetUpdateTime(v string) *UpdateAiModelProviderResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyData) Validate() error {
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

type UpdateAiModelProviderResponseBodyDataModelCards struct {
	AvailablePaths []*UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	Credit         *UpdateAiModelProviderResponseBodyDataModelCardsCredit           `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// example:
	//
	// {"functionCalling":true,"toolChoice":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string                                              `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Meta      *UpdateAiModelProviderResponseBodyDataModelCardsMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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

func (s UpdateAiModelProviderResponseBodyDataModelCards) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderResponseBodyDataModelCards) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetAvailablePaths() []*UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	return s.AvailablePaths
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetCredit() *UpdateAiModelProviderResponseBodyDataModelCardsCredit {
	return s.Credit
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetMeta() *UpdateAiModelProviderResponseBodyDataModelCardsMeta {
	return s.Meta
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetSource() *string {
	return s.Source
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetSourceURL() *string {
	return s.SourceURL
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetAvailablePaths(v []*UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.AvailablePaths = v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetCredit(v *UpdateAiModelProviderResponseBodyDataModelCardsCredit) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.Credit = v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetFeatures(v map[string]interface{}) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.Features = v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetGatewayId(v string) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.GatewayId = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetMeta(v *UpdateAiModelProviderResponseBodyDataModelCardsMeta) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.Meta = v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetModelCardId(v string) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.ModelCardId = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetModelName(v string) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.ModelName = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetModelProvider(v string) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.ModelProvider = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetSource(v string) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.Source = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetSourceURL(v string) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.SourceURL = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) SetUpdateTime(v string) *UpdateAiModelProviderResponseBodyDataModelCards {
	s.UpdateTime = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCards) Validate() error {
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

type UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths struct {
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths) GetType() *string {
	return s.Type
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths) SetPath(v string) *UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	s.Path = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths) SetType(v string) *UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	s.Type = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type UpdateAiModelProviderResponseBodyDataModelCardsCredit struct {
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

func (s UpdateAiModelProviderResponseBodyDataModelCardsCredit) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderResponseBodyDataModelCardsCredit) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) GetType() *string {
	return s.Type
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) SetCacheCost(v float32) *UpdateAiModelProviderResponseBodyDataModelCardsCredit {
	s.CacheCost = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) SetInputCost(v float32) *UpdateAiModelProviderResponseBodyDataModelCardsCredit {
	s.InputCost = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) SetOutputCost(v float32) *UpdateAiModelProviderResponseBodyDataModelCardsCredit {
	s.OutputCost = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) SetType(v string) *UpdateAiModelProviderResponseBodyDataModelCardsCredit {
	s.Type = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsCredit) Validate() error {
	return dara.Validate(s)
}

type UpdateAiModelProviderResponseBodyDataModelCardsMeta struct {
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

func (s UpdateAiModelProviderResponseBodyDataModelCardsMeta) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderResponseBodyDataModelCardsMeta) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) SetMaxInputTokens(v int64) *UpdateAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) SetMaxOutputTokens(v int64) *UpdateAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) SetMaxTokens(v int64) *UpdateAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxTokens = &v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) SetSupportedInputModalities(v []*string) *UpdateAiModelProviderResponseBodyDataModelCardsMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) SetSupportedOutputModalities(v []*string) *UpdateAiModelProviderResponseBodyDataModelCardsMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *UpdateAiModelProviderResponseBodyDataModelCardsMeta) Validate() error {
	return dara.Validate(s)
}
