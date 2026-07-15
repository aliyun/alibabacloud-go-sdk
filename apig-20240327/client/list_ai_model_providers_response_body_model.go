// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAiModelProvidersResponseBody
	GetCode() *string
	SetData(v *ListAiModelProvidersResponseBodyData) *ListAiModelProvidersResponseBody
	GetData() *ListAiModelProvidersResponseBodyData
	SetMessage(v string) *ListAiModelProvidersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAiModelProvidersResponseBody
	GetRequestId() *string
}

type ListAiModelProvidersResponseBody struct {
	// example:
	//
	// Ok
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListAiModelProvidersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAiModelProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAiModelProvidersResponseBody) GetData() *ListAiModelProvidersResponseBodyData {
	return s.Data
}

func (s *ListAiModelProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAiModelProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAiModelProvidersResponseBody) SetCode(v string) *ListAiModelProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *ListAiModelProvidersResponseBody) SetData(v *ListAiModelProvidersResponseBodyData) *ListAiModelProvidersResponseBody {
	s.Data = v
	return s
}

func (s *ListAiModelProvidersResponseBody) SetMessage(v string) *ListAiModelProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListAiModelProvidersResponseBody) SetRequestId(v string) *ListAiModelProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiModelProvidersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAiModelProvidersResponseBodyData struct {
	Items []*ListAiModelProvidersResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListAiModelProvidersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersResponseBodyData) GetItems() []*ListAiModelProvidersResponseBodyDataItems {
	return s.Items
}

func (s *ListAiModelProvidersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAiModelProvidersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiModelProvidersResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListAiModelProvidersResponseBodyData) SetItems(v []*ListAiModelProvidersResponseBodyDataItems) *ListAiModelProvidersResponseBodyData {
	s.Items = v
	return s
}

func (s *ListAiModelProvidersResponseBodyData) SetPageNumber(v int32) *ListAiModelProvidersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyData) SetPageSize(v int32) *ListAiModelProvidersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyData) SetTotalSize(v int32) *ListAiModelProvidersResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAiModelProvidersResponseBodyDataItems struct {
	BoundServices []*ServiceInfo `json:"boundServices,omitempty" xml:"boundServices,omitempty" type:"Repeated"`
	// example:
	//
	// 千问云 / 阿里云百炼
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId  *string                                                `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	ModelCards []*ListAiModelProvidersResponseBodyDataItemsModelCards `json:"modelCards,omitempty" xml:"modelCards,omitempty" type:"Repeated"`
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

func (s ListAiModelProvidersResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetBoundServices() []*ServiceInfo {
	return s.BoundServices
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetModelCards() []*ListAiModelProvidersResponseBodyDataItemsModelCards {
	return s.ModelCards
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetModelCount() *int32 {
	return s.ModelCount
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetModelProviderId() *string {
	return s.ModelProviderId
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetProvider() *string {
	return s.Provider
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetSource() *string {
	return s.Source
}

func (s *ListAiModelProvidersResponseBodyDataItems) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetBoundServices(v []*ServiceInfo) *ListAiModelProvidersResponseBodyDataItems {
	s.BoundServices = v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetDisplayName(v string) *ListAiModelProvidersResponseBodyDataItems {
	s.DisplayName = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetGatewayId(v string) *ListAiModelProvidersResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetModelCards(v []*ListAiModelProvidersResponseBodyDataItemsModelCards) *ListAiModelProvidersResponseBodyDataItems {
	s.ModelCards = v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetModelCount(v int32) *ListAiModelProvidersResponseBodyDataItems {
	s.ModelCount = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetModelProviderId(v string) *ListAiModelProvidersResponseBodyDataItems {
	s.ModelProviderId = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetProvider(v string) *ListAiModelProvidersResponseBodyDataItems {
	s.Provider = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetSource(v string) *ListAiModelProvidersResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) SetUpdateTime(v string) *ListAiModelProvidersResponseBodyDataItems {
	s.UpdateTime = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItems) Validate() error {
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

type ListAiModelProvidersResponseBodyDataItemsModelCards struct {
	AvailablePaths []*ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	Credit         *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit           `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// example:
	//
	// {"functionCalling":true,"toolChoice":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string                                                  `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Meta      *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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

func (s ListAiModelProvidersResponseBodyDataItemsModelCards) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersResponseBodyDataItemsModelCards) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetAvailablePaths() []*ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths {
	return s.AvailablePaths
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetCredit() *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit {
	return s.Credit
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetMeta() *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta {
	return s.Meta
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetModelName() *string {
	return s.ModelName
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetSource() *string {
	return s.Source
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetSourceURL() *string {
	return s.SourceURL
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetAvailablePaths(v []*ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.AvailablePaths = v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetCredit(v *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.Credit = v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetFeatures(v map[string]interface{}) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.Features = v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetGatewayId(v string) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.GatewayId = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetMeta(v *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.Meta = v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetModelCardId(v string) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.ModelCardId = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetModelName(v string) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.ModelName = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetModelProvider(v string) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.ModelProvider = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetSource(v string) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.Source = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetSourceURL(v string) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.SourceURL = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) SetUpdateTime(v string) *ListAiModelProvidersResponseBodyDataItemsModelCards {
	s.UpdateTime = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCards) Validate() error {
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

type ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths struct {
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths) GetType() *string {
	return s.Type
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths) SetPath(v string) *ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths {
	s.Path = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths) SetType(v string) *ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths {
	s.Type = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type ListAiModelProvidersResponseBodyDataItemsModelCardsCredit struct {
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

func (s ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) GetType() *string {
	return s.Type
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) SetCacheCost(v float32) *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit {
	s.CacheCost = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) SetInputCost(v float32) *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit {
	s.InputCost = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) SetOutputCost(v float32) *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit {
	s.OutputCost = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) SetType(v string) *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit {
	s.Type = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsCredit) Validate() error {
	return dara.Validate(s)
}

type ListAiModelProvidersResponseBodyDataItemsModelCardsMeta struct {
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

func (s ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) SetMaxInputTokens(v int64) *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) SetMaxOutputTokens(v int64) *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) SetMaxTokens(v int64) *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta {
	s.MaxTokens = &v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) SetSupportedInputModalities(v []*string) *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) SetSupportedOutputModalities(v []*string) *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *ListAiModelProvidersResponseBodyDataItemsModelCardsMeta) Validate() error {
	return dara.Validate(s)
}
