// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiModelCardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailablePaths(v []*CreateAiModelCardRequestAvailablePaths) *CreateAiModelCardRequest
	GetAvailablePaths() []*CreateAiModelCardRequestAvailablePaths
	SetCredit(v *CreateAiModelCardRequestCredit) *CreateAiModelCardRequest
	GetCredit() *CreateAiModelCardRequestCredit
	SetFeatures(v map[string]interface{}) *CreateAiModelCardRequest
	GetFeatures() map[string]interface{}
	SetGatewayId(v string) *CreateAiModelCardRequest
	GetGatewayId() *string
	SetMeta(v *CreateAiModelCardRequestMeta) *CreateAiModelCardRequest
	GetMeta() *CreateAiModelCardRequestMeta
	SetModelName(v string) *CreateAiModelCardRequest
	GetModelName() *string
	SetModelProvider(v string) *CreateAiModelCardRequest
	GetModelProvider() *string
}

type CreateAiModelCardRequest struct {
	// The list of API paths supported by the model. Each item must include both path and type.
	//
	// example:
	//
	// [{"path":"/v1/chat/completions","type":"OpenAICompatible"}]
	AvailablePaths []*CreateAiModelCardRequestAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	// The credit billing information of the model. Currently, only the fixed type is supported. The cost unit is Credits per million tokens. If not specified, all cost values default to 0.
	//
	// example:
	//
	// {"type":"fixed","inputCost":1.5,"outputCost":3,"cacheCost":0.5}
	Credit *CreateAiModelCardRequestCredit `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// The model capability switches. Keys must be model capability names supported by API Gateway, and values must be Boolean.
	//
	// example:
	//
	// {"functionCalling":true,"toolChoice":true,"promptCaching":false}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// The AI gateway instance ID. The target instance must exist, belong to the current account, and be of the AI gateway type.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The token limits and input/output modality information of the model.
	//
	// example:
	//
	// {"maxInputTokens":131072,"maxOutputTokens":8192,"maxTokens":131072,"supportedInputModalities":["text","image"],"supportedOutputModalities":["text"]}
	Meta *CreateAiModelCardRequestMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The model name. The model name must be unique within the same AI gateway instance and the same model provider. The value can be up to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model provider identifier. The value must reference an existing model provider in the target AI gateway instance. The value can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen
	ModelProvider *string `json:"modelProvider,omitempty" xml:"modelProvider,omitempty"`
}

func (s CreateAiModelCardRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardRequest) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardRequest) GetAvailablePaths() []*CreateAiModelCardRequestAvailablePaths {
	return s.AvailablePaths
}

func (s *CreateAiModelCardRequest) GetCredit() *CreateAiModelCardRequestCredit {
	return s.Credit
}

func (s *CreateAiModelCardRequest) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *CreateAiModelCardRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateAiModelCardRequest) GetMeta() *CreateAiModelCardRequestMeta {
	return s.Meta
}

func (s *CreateAiModelCardRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateAiModelCardRequest) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *CreateAiModelCardRequest) SetAvailablePaths(v []*CreateAiModelCardRequestAvailablePaths) *CreateAiModelCardRequest {
	s.AvailablePaths = v
	return s
}

func (s *CreateAiModelCardRequest) SetCredit(v *CreateAiModelCardRequestCredit) *CreateAiModelCardRequest {
	s.Credit = v
	return s
}

func (s *CreateAiModelCardRequest) SetFeatures(v map[string]interface{}) *CreateAiModelCardRequest {
	s.Features = v
	return s
}

func (s *CreateAiModelCardRequest) SetGatewayId(v string) *CreateAiModelCardRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateAiModelCardRequest) SetMeta(v *CreateAiModelCardRequestMeta) *CreateAiModelCardRequest {
	s.Meta = v
	return s
}

func (s *CreateAiModelCardRequest) SetModelName(v string) *CreateAiModelCardRequest {
	s.ModelName = &v
	return s
}

func (s *CreateAiModelCardRequest) SetModelProvider(v string) *CreateAiModelCardRequest {
	s.ModelProvider = &v
	return s
}

func (s *CreateAiModelCardRequest) Validate() error {
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

type CreateAiModelCardRequestAvailablePaths struct {
	// The model invocation path. The value can be up to 2048 characters in length.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The protocol type corresponding to the path. The value can be up to 64 characters in length.
	//
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAiModelCardRequestAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardRequestAvailablePaths) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardRequestAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *CreateAiModelCardRequestAvailablePaths) GetType() *string {
	return s.Type
}

func (s *CreateAiModelCardRequestAvailablePaths) SetPath(v string) *CreateAiModelCardRequestAvailablePaths {
	s.Path = &v
	return s
}

func (s *CreateAiModelCardRequestAvailablePaths) SetType(v string) *CreateAiModelCardRequestAvailablePaths {
	s.Type = &v
	return s
}

func (s *CreateAiModelCardRequestAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type CreateAiModelCardRequestCredit struct {
	// The cache hit token cost, in Credits per million tokens. The value must be greater than or equal to 0. If not specified, the value defaults to 0.
	//
	// example:
	//
	// 0.5
	CacheCost *float32 `json:"cacheCost,omitempty" xml:"cacheCost,omitempty"`
	// The input token cost, in Credits per million tokens. The value must be greater than or equal to 0. If not specified, the value defaults to 0.
	//
	// example:
	//
	// 1.5
	InputCost *float32 `json:"inputCost,omitempty" xml:"inputCost,omitempty"`
	// The output token cost, in Credits per million tokens. The value must be greater than or equal to 0. If not specified, the value defaults to 0.
	//
	// example:
	//
	// 3
	OutputCost *float32 `json:"outputCost,omitempty" xml:"outputCost,omitempty"`
	// The billing type. Currently, only fixed is supported. If not specified, the value defaults to fixed.
	//
	// example:
	//
	// fixed
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAiModelCardRequestCredit) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardRequestCredit) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardRequestCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *CreateAiModelCardRequestCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *CreateAiModelCardRequestCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *CreateAiModelCardRequestCredit) GetType() *string {
	return s.Type
}

func (s *CreateAiModelCardRequestCredit) SetCacheCost(v float32) *CreateAiModelCardRequestCredit {
	s.CacheCost = &v
	return s
}

func (s *CreateAiModelCardRequestCredit) SetInputCost(v float32) *CreateAiModelCardRequestCredit {
	s.InputCost = &v
	return s
}

func (s *CreateAiModelCardRequestCredit) SetOutputCost(v float32) *CreateAiModelCardRequestCredit {
	s.OutputCost = &v
	return s
}

func (s *CreateAiModelCardRequestCredit) SetType(v string) *CreateAiModelCardRequestCredit {
	s.Type = &v
	return s
}

func (s *CreateAiModelCardRequestCredit) Validate() error {
	return dara.Validate(s)
}

type CreateAiModelCardRequestMeta struct {
	// The maximum number of input tokens supported by the model. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 131072
	MaxInputTokens *int64 `json:"maxInputTokens,omitempty" xml:"maxInputTokens,omitempty"`
	// The maximum number of output tokens supported by the model. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 8192
	MaxOutputTokens *int64 `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	// The maximum total number of context tokens supported by the model. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 131072
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The list of input modalities supported by the model. The list can contain up to 16 items, and each item must not be empty.
	//
	// example:
	//
	// ["text","image"]
	SupportedInputModalities []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	// The list of output modalities supported by the model. The list can contain up to 16 items, and each item must not be empty.
	//
	// example:
	//
	// ["text"]
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s CreateAiModelCardRequestMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelCardRequestMeta) GoString() string {
	return s.String()
}

func (s *CreateAiModelCardRequestMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *CreateAiModelCardRequestMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *CreateAiModelCardRequestMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *CreateAiModelCardRequestMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *CreateAiModelCardRequestMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *CreateAiModelCardRequestMeta) SetMaxInputTokens(v int64) *CreateAiModelCardRequestMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *CreateAiModelCardRequestMeta) SetMaxOutputTokens(v int64) *CreateAiModelCardRequestMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *CreateAiModelCardRequestMeta) SetMaxTokens(v int64) *CreateAiModelCardRequestMeta {
	s.MaxTokens = &v
	return s
}

func (s *CreateAiModelCardRequestMeta) SetSupportedInputModalities(v []*string) *CreateAiModelCardRequestMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *CreateAiModelCardRequestMeta) SetSupportedOutputModalities(v []*string) *CreateAiModelCardRequestMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *CreateAiModelCardRequestMeta) Validate() error {
	return dara.Validate(s)
}
