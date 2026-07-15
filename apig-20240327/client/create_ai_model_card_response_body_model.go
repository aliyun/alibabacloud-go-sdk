// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiModelCardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAiModelCardResponseBody
	GetCode() *string
	SetData(v *CreateAiModelCardResponseBodyData) *CreateAiModelCardResponseBody
	GetData() *CreateAiModelCardResponseBodyData
	SetMessage(v string) *CreateAiModelCardResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAiModelCardResponseBody
	GetRequestId() *string
}

type CreateAiModelCardResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {"modelCardId":"mc-8c13d2b4f8a1","gatewayId":"gw-8c13d2b4f8a1","modelProvider":"qwen","modelName":"qwen-plus","source":"user","credit":{"type":"fixed","inputCost":1.5,"outputCost":3,"cacheCost":0.5},"features":{"functionCalling":true,"toolChoice":true},"meta":{"maxInputTokens":131072,"maxOutputTokens":8192,"maxTokens":131072,"supportedInputModalities":["text","image"],"supportedOutputModalities":["text"]},"availablePaths":[{"path":"/v1/chat/completions","type":"OpenAICompatible"}],"updateTime":"2026-07-14 18:30:00"}
	Data *CreateAiModelCardResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAiModelCardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAiModelCardResponseBody) GetData() *CreateAiModelCardResponseBodyData {
	return s.Data
}

func (s *CreateAiModelCardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAiModelCardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAiModelCardResponseBody) SetCode(v string) *CreateAiModelCardResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAiModelCardResponseBody) SetData(v *CreateAiModelCardResponseBodyData) *CreateAiModelCardResponseBody {
	s.Data = v
	return s
}

func (s *CreateAiModelCardResponseBody) SetMessage(v string) *CreateAiModelCardResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAiModelCardResponseBody) SetRequestId(v string) *CreateAiModelCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAiModelCardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAiModelCardResponseBodyData struct {
	// example:
	//
	// [{"path":"/v1/chat/completions","type":"OpenAICompatible"}]
	AvailablePaths []*CreateAiModelCardResponseBodyDataAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	// example:
	//
	// {"type":"fixed","inputCost":1.5,"outputCost":3,"cacheCost":0.5}
	Credit *CreateAiModelCardResponseBodyDataCredit `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// example:
	//
	// {"functionCalling":true,"toolChoice":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// {"maxInputTokens":131072,"maxOutputTokens":8192,"maxTokens":131072,"supportedInputModalities":["text","image"],"supportedOutputModalities":["text"]}
	Meta *CreateAiModelCardResponseBodyDataMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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

func (s CreateAiModelCardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardResponseBodyData) GetAvailablePaths() []*CreateAiModelCardResponseBodyDataAvailablePaths {
	return s.AvailablePaths
}

func (s *CreateAiModelCardResponseBodyData) GetCredit() *CreateAiModelCardResponseBodyDataCredit {
	return s.Credit
}

func (s *CreateAiModelCardResponseBodyData) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *CreateAiModelCardResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateAiModelCardResponseBodyData) GetMeta() *CreateAiModelCardResponseBodyDataMeta {
	return s.Meta
}

func (s *CreateAiModelCardResponseBodyData) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *CreateAiModelCardResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *CreateAiModelCardResponseBodyData) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *CreateAiModelCardResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *CreateAiModelCardResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateAiModelCardResponseBodyData) SetAvailablePaths(v []*CreateAiModelCardResponseBodyDataAvailablePaths) *CreateAiModelCardResponseBodyData {
	s.AvailablePaths = v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetCredit(v *CreateAiModelCardResponseBodyDataCredit) *CreateAiModelCardResponseBodyData {
	s.Credit = v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetFeatures(v map[string]interface{}) *CreateAiModelCardResponseBodyData {
	s.Features = v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetGatewayId(v string) *CreateAiModelCardResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetMeta(v *CreateAiModelCardResponseBodyDataMeta) *CreateAiModelCardResponseBodyData {
	s.Meta = v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetModelCardId(v string) *CreateAiModelCardResponseBodyData {
	s.ModelCardId = &v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetModelName(v string) *CreateAiModelCardResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetModelProvider(v string) *CreateAiModelCardResponseBodyData {
	s.ModelProvider = &v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetSource(v string) *CreateAiModelCardResponseBodyData {
	s.Source = &v
	return s
}

func (s *CreateAiModelCardResponseBodyData) SetUpdateTime(v string) *CreateAiModelCardResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateAiModelCardResponseBodyData) Validate() error {
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

type CreateAiModelCardResponseBodyDataAvailablePaths struct {
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAiModelCardResponseBodyDataAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardResponseBodyDataAvailablePaths) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardResponseBodyDataAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *CreateAiModelCardResponseBodyDataAvailablePaths) GetType() *string {
	return s.Type
}

func (s *CreateAiModelCardResponseBodyDataAvailablePaths) SetPath(v string) *CreateAiModelCardResponseBodyDataAvailablePaths {
	s.Path = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataAvailablePaths) SetType(v string) *CreateAiModelCardResponseBodyDataAvailablePaths {
	s.Type = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type CreateAiModelCardResponseBodyDataCredit struct {
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

func (s CreateAiModelCardResponseBodyDataCredit) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardResponseBodyDataCredit) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardResponseBodyDataCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *CreateAiModelCardResponseBodyDataCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *CreateAiModelCardResponseBodyDataCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *CreateAiModelCardResponseBodyDataCredit) GetType() *string {
	return s.Type
}

func (s *CreateAiModelCardResponseBodyDataCredit) SetCacheCost(v float32) *CreateAiModelCardResponseBodyDataCredit {
	s.CacheCost = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataCredit) SetInputCost(v float32) *CreateAiModelCardResponseBodyDataCredit {
	s.InputCost = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataCredit) SetOutputCost(v float32) *CreateAiModelCardResponseBodyDataCredit {
	s.OutputCost = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataCredit) SetType(v string) *CreateAiModelCardResponseBodyDataCredit {
	s.Type = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataCredit) Validate() error {
	return dara.Validate(s)
}

type CreateAiModelCardResponseBodyDataMeta struct {
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
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// example:
	//
	// ["text","image"]
	SupportedInputModalities []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	// example:
	//
	// ["text"]
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s CreateAiModelCardResponseBodyDataMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardResponseBodyDataMeta) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardResponseBodyDataMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *CreateAiModelCardResponseBodyDataMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *CreateAiModelCardResponseBodyDataMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *CreateAiModelCardResponseBodyDataMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *CreateAiModelCardResponseBodyDataMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *CreateAiModelCardResponseBodyDataMeta) SetMaxInputTokens(v int64) *CreateAiModelCardResponseBodyDataMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataMeta) SetMaxOutputTokens(v int64) *CreateAiModelCardResponseBodyDataMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataMeta) SetMaxTokens(v int64) *CreateAiModelCardResponseBodyDataMeta {
	s.MaxTokens = &v
	return s
}

func (s *CreateAiModelCardResponseBodyDataMeta) SetSupportedInputModalities(v []*string) *CreateAiModelCardResponseBodyDataMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *CreateAiModelCardResponseBodyDataMeta) SetSupportedOutputModalities(v []*string) *CreateAiModelCardResponseBodyDataMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *CreateAiModelCardResponseBodyDataMeta) Validate() error {
	return dara.Validate(s)
}
