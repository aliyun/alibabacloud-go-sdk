// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelCardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAiModelCardResponseBody
	GetCode() *string
	SetData(v *GetAiModelCardResponseBodyData) *GetAiModelCardResponseBody
	GetData() *GetAiModelCardResponseBodyData
	SetMessage(v string) *GetAiModelCardResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiModelCardResponseBody
	GetRequestId() *string
}

type GetAiModelCardResponseBody struct {
	// example:
	//
	// Ok
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetAiModelCardResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAiModelCardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelCardResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiModelCardResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiModelCardResponseBody) GetData() *GetAiModelCardResponseBodyData {
	return s.Data
}

func (s *GetAiModelCardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiModelCardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiModelCardResponseBody) SetCode(v string) *GetAiModelCardResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiModelCardResponseBody) SetData(v *GetAiModelCardResponseBodyData) *GetAiModelCardResponseBody {
	s.Data = v
	return s
}

func (s *GetAiModelCardResponseBody) SetMessage(v string) *GetAiModelCardResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiModelCardResponseBody) SetRequestId(v string) *GetAiModelCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiModelCardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiModelCardResponseBodyData struct {
	AvailablePaths []*GetAiModelCardResponseBodyDataAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	Credit         *GetAiModelCardResponseBodyDataCredit           `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// example:
	//
	// {"functionCalling":true,"toolChoice":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string                             `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Meta      *GetAiModelCardResponseBodyDataMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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
	// 2026-07-14 18:30:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetAiModelCardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiModelCardResponseBodyData) GetAvailablePaths() []*GetAiModelCardResponseBodyDataAvailablePaths {
	return s.AvailablePaths
}

func (s *GetAiModelCardResponseBodyData) GetCredit() *GetAiModelCardResponseBodyDataCredit {
	return s.Credit
}

func (s *GetAiModelCardResponseBodyData) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *GetAiModelCardResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetAiModelCardResponseBodyData) GetMeta() *GetAiModelCardResponseBodyDataMeta {
	return s.Meta
}

func (s *GetAiModelCardResponseBodyData) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *GetAiModelCardResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *GetAiModelCardResponseBodyData) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *GetAiModelCardResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetAiModelCardResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAiModelCardResponseBodyData) SetAvailablePaths(v []*GetAiModelCardResponseBodyDataAvailablePaths) *GetAiModelCardResponseBodyData {
	s.AvailablePaths = v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetCredit(v *GetAiModelCardResponseBodyDataCredit) *GetAiModelCardResponseBodyData {
	s.Credit = v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetFeatures(v map[string]interface{}) *GetAiModelCardResponseBodyData {
	s.Features = v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetGatewayId(v string) *GetAiModelCardResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetMeta(v *GetAiModelCardResponseBodyDataMeta) *GetAiModelCardResponseBodyData {
	s.Meta = v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetModelCardId(v string) *GetAiModelCardResponseBodyData {
	s.ModelCardId = &v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetModelName(v string) *GetAiModelCardResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetModelProvider(v string) *GetAiModelCardResponseBodyData {
	s.ModelProvider = &v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetSource(v string) *GetAiModelCardResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetAiModelCardResponseBodyData) SetUpdateTime(v string) *GetAiModelCardResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAiModelCardResponseBodyData) Validate() error {
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

type GetAiModelCardResponseBodyDataAvailablePaths struct {
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAiModelCardResponseBodyDataAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelCardResponseBodyDataAvailablePaths) GoString() string {
	return s.String()
}

func (s *GetAiModelCardResponseBodyDataAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *GetAiModelCardResponseBodyDataAvailablePaths) GetType() *string {
	return s.Type
}

func (s *GetAiModelCardResponseBodyDataAvailablePaths) SetPath(v string) *GetAiModelCardResponseBodyDataAvailablePaths {
	s.Path = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataAvailablePaths) SetType(v string) *GetAiModelCardResponseBodyDataAvailablePaths {
	s.Type = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type GetAiModelCardResponseBodyDataCredit struct {
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

func (s GetAiModelCardResponseBodyDataCredit) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelCardResponseBodyDataCredit) GoString() string {
	return s.String()
}

func (s *GetAiModelCardResponseBodyDataCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *GetAiModelCardResponseBodyDataCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *GetAiModelCardResponseBodyDataCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *GetAiModelCardResponseBodyDataCredit) GetType() *string {
	return s.Type
}

func (s *GetAiModelCardResponseBodyDataCredit) SetCacheCost(v float32) *GetAiModelCardResponseBodyDataCredit {
	s.CacheCost = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataCredit) SetInputCost(v float32) *GetAiModelCardResponseBodyDataCredit {
	s.InputCost = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataCredit) SetOutputCost(v float32) *GetAiModelCardResponseBodyDataCredit {
	s.OutputCost = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataCredit) SetType(v string) *GetAiModelCardResponseBodyDataCredit {
	s.Type = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataCredit) Validate() error {
	return dara.Validate(s)
}

type GetAiModelCardResponseBodyDataMeta struct {
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

func (s GetAiModelCardResponseBodyDataMeta) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelCardResponseBodyDataMeta) GoString() string {
	return s.String()
}

func (s *GetAiModelCardResponseBodyDataMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *GetAiModelCardResponseBodyDataMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *GetAiModelCardResponseBodyDataMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *GetAiModelCardResponseBodyDataMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *GetAiModelCardResponseBodyDataMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *GetAiModelCardResponseBodyDataMeta) SetMaxInputTokens(v int64) *GetAiModelCardResponseBodyDataMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataMeta) SetMaxOutputTokens(v int64) *GetAiModelCardResponseBodyDataMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataMeta) SetMaxTokens(v int64) *GetAiModelCardResponseBodyDataMeta {
	s.MaxTokens = &v
	return s
}

func (s *GetAiModelCardResponseBodyDataMeta) SetSupportedInputModalities(v []*string) *GetAiModelCardResponseBodyDataMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *GetAiModelCardResponseBodyDataMeta) SetSupportedOutputModalities(v []*string) *GetAiModelCardResponseBodyDataMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *GetAiModelCardResponseBodyDataMeta) Validate() error {
	return dara.Validate(s)
}
