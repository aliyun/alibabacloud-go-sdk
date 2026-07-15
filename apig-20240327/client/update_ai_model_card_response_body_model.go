// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelCardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAiModelCardResponseBody
	GetCode() *string
	SetData(v *UpdateAiModelCardResponseBodyData) *UpdateAiModelCardResponseBody
	GetData() *UpdateAiModelCardResponseBodyData
	SetMessage(v string) *UpdateAiModelCardResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAiModelCardResponseBody
	GetRequestId() *string
}

type UpdateAiModelCardResponseBody struct {
	// example:
	//
	// Ok
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateAiModelCardResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAiModelCardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAiModelCardResponseBody) GetData() *UpdateAiModelCardResponseBodyData {
	return s.Data
}

func (s *UpdateAiModelCardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAiModelCardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAiModelCardResponseBody) SetCode(v string) *UpdateAiModelCardResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAiModelCardResponseBody) SetData(v *UpdateAiModelCardResponseBodyData) *UpdateAiModelCardResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAiModelCardResponseBody) SetMessage(v string) *UpdateAiModelCardResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAiModelCardResponseBody) SetRequestId(v string) *UpdateAiModelCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAiModelCardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAiModelCardResponseBodyData struct {
	AvailablePaths []*UpdateAiModelCardResponseBodyDataAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	Credit         *UpdateAiModelCardResponseBodyDataCredit           `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// example:
	//
	// {"functionCalling":true,"toolChoice":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string                                `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Meta      *UpdateAiModelCardResponseBodyDataMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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

func (s UpdateAiModelCardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardResponseBodyData) GetAvailablePaths() []*UpdateAiModelCardResponseBodyDataAvailablePaths {
	return s.AvailablePaths
}

func (s *UpdateAiModelCardResponseBodyData) GetCredit() *UpdateAiModelCardResponseBodyDataCredit {
	return s.Credit
}

func (s *UpdateAiModelCardResponseBodyData) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *UpdateAiModelCardResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateAiModelCardResponseBodyData) GetMeta() *UpdateAiModelCardResponseBodyDataMeta {
	return s.Meta
}

func (s *UpdateAiModelCardResponseBodyData) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *UpdateAiModelCardResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateAiModelCardResponseBodyData) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *UpdateAiModelCardResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *UpdateAiModelCardResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateAiModelCardResponseBodyData) SetAvailablePaths(v []*UpdateAiModelCardResponseBodyDataAvailablePaths) *UpdateAiModelCardResponseBodyData {
	s.AvailablePaths = v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetCredit(v *UpdateAiModelCardResponseBodyDataCredit) *UpdateAiModelCardResponseBodyData {
	s.Credit = v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetFeatures(v map[string]interface{}) *UpdateAiModelCardResponseBodyData {
	s.Features = v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetGatewayId(v string) *UpdateAiModelCardResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetMeta(v *UpdateAiModelCardResponseBodyDataMeta) *UpdateAiModelCardResponseBodyData {
	s.Meta = v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetModelCardId(v string) *UpdateAiModelCardResponseBodyData {
	s.ModelCardId = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetModelName(v string) *UpdateAiModelCardResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetModelProvider(v string) *UpdateAiModelCardResponseBodyData {
	s.ModelProvider = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetSource(v string) *UpdateAiModelCardResponseBodyData {
	s.Source = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) SetUpdateTime(v string) *UpdateAiModelCardResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyData) Validate() error {
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

type UpdateAiModelCardResponseBodyDataAvailablePaths struct {
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAiModelCardResponseBodyDataAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardResponseBodyDataAvailablePaths) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardResponseBodyDataAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *UpdateAiModelCardResponseBodyDataAvailablePaths) GetType() *string {
	return s.Type
}

func (s *UpdateAiModelCardResponseBodyDataAvailablePaths) SetPath(v string) *UpdateAiModelCardResponseBodyDataAvailablePaths {
	s.Path = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataAvailablePaths) SetType(v string) *UpdateAiModelCardResponseBodyDataAvailablePaths {
	s.Type = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type UpdateAiModelCardResponseBodyDataCredit struct {
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

func (s UpdateAiModelCardResponseBodyDataCredit) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardResponseBodyDataCredit) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardResponseBodyDataCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *UpdateAiModelCardResponseBodyDataCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *UpdateAiModelCardResponseBodyDataCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *UpdateAiModelCardResponseBodyDataCredit) GetType() *string {
	return s.Type
}

func (s *UpdateAiModelCardResponseBodyDataCredit) SetCacheCost(v float32) *UpdateAiModelCardResponseBodyDataCredit {
	s.CacheCost = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataCredit) SetInputCost(v float32) *UpdateAiModelCardResponseBodyDataCredit {
	s.InputCost = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataCredit) SetOutputCost(v float32) *UpdateAiModelCardResponseBodyDataCredit {
	s.OutputCost = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataCredit) SetType(v string) *UpdateAiModelCardResponseBodyDataCredit {
	s.Type = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataCredit) Validate() error {
	return dara.Validate(s)
}

type UpdateAiModelCardResponseBodyDataMeta struct {
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

func (s UpdateAiModelCardResponseBodyDataMeta) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardResponseBodyDataMeta) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardResponseBodyDataMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *UpdateAiModelCardResponseBodyDataMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *UpdateAiModelCardResponseBodyDataMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *UpdateAiModelCardResponseBodyDataMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *UpdateAiModelCardResponseBodyDataMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *UpdateAiModelCardResponseBodyDataMeta) SetMaxInputTokens(v int64) *UpdateAiModelCardResponseBodyDataMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataMeta) SetMaxOutputTokens(v int64) *UpdateAiModelCardResponseBodyDataMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataMeta) SetMaxTokens(v int64) *UpdateAiModelCardResponseBodyDataMeta {
	s.MaxTokens = &v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataMeta) SetSupportedInputModalities(v []*string) *UpdateAiModelCardResponseBodyDataMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataMeta) SetSupportedOutputModalities(v []*string) *UpdateAiModelCardResponseBodyDataMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *UpdateAiModelCardResponseBodyDataMeta) Validate() error {
	return dara.Validate(s)
}
