// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelCardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAiModelCardsResponseBody
	GetCode() *string
	SetData(v *ListAiModelCardsResponseBodyData) *ListAiModelCardsResponseBody
	GetData() *ListAiModelCardsResponseBodyData
	SetMessage(v string) *ListAiModelCardsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAiModelCardsResponseBody
	GetRequestId() *string
}

type ListAiModelCardsResponseBody struct {
	// example:
	//
	// Ok
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListAiModelCardsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAiModelCardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelCardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiModelCardsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAiModelCardsResponseBody) GetData() *ListAiModelCardsResponseBodyData {
	return s.Data
}

func (s *ListAiModelCardsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAiModelCardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAiModelCardsResponseBody) SetCode(v string) *ListAiModelCardsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAiModelCardsResponseBody) SetData(v *ListAiModelCardsResponseBodyData) *ListAiModelCardsResponseBody {
	s.Data = v
	return s
}

func (s *ListAiModelCardsResponseBody) SetMessage(v string) *ListAiModelCardsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAiModelCardsResponseBody) SetRequestId(v string) *ListAiModelCardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiModelCardsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAiModelCardsResponseBodyData struct {
	Items []*ListAiModelCardsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s ListAiModelCardsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelCardsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAiModelCardsResponseBodyData) GetItems() []*ListAiModelCardsResponseBodyDataItems {
	return s.Items
}

func (s *ListAiModelCardsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAiModelCardsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiModelCardsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListAiModelCardsResponseBodyData) SetItems(v []*ListAiModelCardsResponseBodyDataItems) *ListAiModelCardsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListAiModelCardsResponseBodyData) SetPageNumber(v int32) *ListAiModelCardsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAiModelCardsResponseBodyData) SetPageSize(v int32) *ListAiModelCardsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAiModelCardsResponseBodyData) SetTotalSize(v int32) *ListAiModelCardsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListAiModelCardsResponseBodyData) Validate() error {
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

type ListAiModelCardsResponseBodyDataItems struct {
	AvailablePaths []*ListAiModelCardsResponseBodyDataItemsAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	Credit         *ListAiModelCardsResponseBodyDataItemsCredit           `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// example:
	//
	// {"functionCalling":true,"toolChoice":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string                                    `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Meta      *ListAiModelCardsResponseBodyDataItemsMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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

func (s ListAiModelCardsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelCardsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListAiModelCardsResponseBodyDataItems) GetAvailablePaths() []*ListAiModelCardsResponseBodyDataItemsAvailablePaths {
	return s.AvailablePaths
}

func (s *ListAiModelCardsResponseBodyDataItems) GetCredit() *ListAiModelCardsResponseBodyDataItemsCredit {
	return s.Credit
}

func (s *ListAiModelCardsResponseBodyDataItems) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *ListAiModelCardsResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListAiModelCardsResponseBodyDataItems) GetMeta() *ListAiModelCardsResponseBodyDataItemsMeta {
	return s.Meta
}

func (s *ListAiModelCardsResponseBodyDataItems) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ListAiModelCardsResponseBodyDataItems) GetModelName() *string {
	return s.ModelName
}

func (s *ListAiModelCardsResponseBodyDataItems) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *ListAiModelCardsResponseBodyDataItems) GetSource() *string {
	return s.Source
}

func (s *ListAiModelCardsResponseBodyDataItems) GetSourceURL() *string {
	return s.SourceURL
}

func (s *ListAiModelCardsResponseBodyDataItems) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAiModelCardsResponseBodyDataItems) SetAvailablePaths(v []*ListAiModelCardsResponseBodyDataItemsAvailablePaths) *ListAiModelCardsResponseBodyDataItems {
	s.AvailablePaths = v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetCredit(v *ListAiModelCardsResponseBodyDataItemsCredit) *ListAiModelCardsResponseBodyDataItems {
	s.Credit = v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetFeatures(v map[string]interface{}) *ListAiModelCardsResponseBodyDataItems {
	s.Features = v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetGatewayId(v string) *ListAiModelCardsResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetMeta(v *ListAiModelCardsResponseBodyDataItemsMeta) *ListAiModelCardsResponseBodyDataItems {
	s.Meta = v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetModelCardId(v string) *ListAiModelCardsResponseBodyDataItems {
	s.ModelCardId = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetModelName(v string) *ListAiModelCardsResponseBodyDataItems {
	s.ModelName = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetModelProvider(v string) *ListAiModelCardsResponseBodyDataItems {
	s.ModelProvider = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetSource(v string) *ListAiModelCardsResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetSourceURL(v string) *ListAiModelCardsResponseBodyDataItems {
	s.SourceURL = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) SetUpdateTime(v string) *ListAiModelCardsResponseBodyDataItems {
	s.UpdateTime = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItems) Validate() error {
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

type ListAiModelCardsResponseBodyDataItemsAvailablePaths struct {
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAiModelCardsResponseBodyDataItemsAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelCardsResponseBodyDataItemsAvailablePaths) GoString() string {
	return s.String()
}

func (s *ListAiModelCardsResponseBodyDataItemsAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *ListAiModelCardsResponseBodyDataItemsAvailablePaths) GetType() *string {
	return s.Type
}

func (s *ListAiModelCardsResponseBodyDataItemsAvailablePaths) SetPath(v string) *ListAiModelCardsResponseBodyDataItemsAvailablePaths {
	s.Path = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsAvailablePaths) SetType(v string) *ListAiModelCardsResponseBodyDataItemsAvailablePaths {
	s.Type = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type ListAiModelCardsResponseBodyDataItemsCredit struct {
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

func (s ListAiModelCardsResponseBodyDataItemsCredit) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelCardsResponseBodyDataItemsCredit) GoString() string {
	return s.String()
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) GetType() *string {
	return s.Type
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) SetCacheCost(v float32) *ListAiModelCardsResponseBodyDataItemsCredit {
	s.CacheCost = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) SetInputCost(v float32) *ListAiModelCardsResponseBodyDataItemsCredit {
	s.InputCost = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) SetOutputCost(v float32) *ListAiModelCardsResponseBodyDataItemsCredit {
	s.OutputCost = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) SetType(v string) *ListAiModelCardsResponseBodyDataItemsCredit {
	s.Type = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsCredit) Validate() error {
	return dara.Validate(s)
}

type ListAiModelCardsResponseBodyDataItemsMeta struct {
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

func (s ListAiModelCardsResponseBodyDataItemsMeta) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelCardsResponseBodyDataItemsMeta) GoString() string {
	return s.String()
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) SetMaxInputTokens(v int64) *ListAiModelCardsResponseBodyDataItemsMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) SetMaxOutputTokens(v int64) *ListAiModelCardsResponseBodyDataItemsMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) SetMaxTokens(v int64) *ListAiModelCardsResponseBodyDataItemsMeta {
	s.MaxTokens = &v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) SetSupportedInputModalities(v []*string) *ListAiModelCardsResponseBodyDataItemsMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) SetSupportedOutputModalities(v []*string) *ListAiModelCardsResponseBodyDataItemsMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *ListAiModelCardsResponseBodyDataItemsMeta) Validate() error {
	return dara.Validate(s)
}
