// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelCardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailablePaths(v []*UpdateAiModelCardRequestAvailablePaths) *UpdateAiModelCardRequest
	GetAvailablePaths() []*UpdateAiModelCardRequestAvailablePaths
	SetCredit(v *UpdateAiModelCardRequestCredit) *UpdateAiModelCardRequest
	GetCredit() *UpdateAiModelCardRequestCredit
	SetFeatures(v map[string]interface{}) *UpdateAiModelCardRequest
	GetFeatures() map[string]interface{}
	SetMeta(v *UpdateAiModelCardRequestMeta) *UpdateAiModelCardRequest
	GetMeta() *UpdateAiModelCardRequestMeta
	SetModelName(v string) *UpdateAiModelCardRequest
	GetModelName() *string
	SetModelProvider(v string) *UpdateAiModelCardRequest
	GetModelProvider() *string
}

type UpdateAiModelCardRequest struct {
	// The list of invocation paths supported by the model. Each item must include both path and type. The list is overwritten as a whole during updates.
	AvailablePaths []*UpdateAiModelCardRequestAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	// The credit billing information of the model. Only the fixed type is supported. The unit is Credits per million tokens. If not specified, all cost values default to 0.
	Credit *UpdateAiModelCardRequestCredit `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// The model capability switches. Keys must be model capability names supported by the API gateway. Values are Boolean.
	//
	// example:
	//
	// {"functionCalling":true,"toolChoice":true,"promptCaching":false}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// The token limits and input/output modality information of the model.
	Meta *UpdateAiModelCardRequestMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The model name. The name must be unique within the same AI gateway instance and model provider. Maximum length: 256 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model provider identifier. The value must reference an existing model provider in the target AI gateway instance. Maximum length: 128 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen
	ModelProvider *string `json:"modelProvider,omitempty" xml:"modelProvider,omitempty"`
}

func (s UpdateAiModelCardRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardRequest) GetAvailablePaths() []*UpdateAiModelCardRequestAvailablePaths {
	return s.AvailablePaths
}

func (s *UpdateAiModelCardRequest) GetCredit() *UpdateAiModelCardRequestCredit {
	return s.Credit
}

func (s *UpdateAiModelCardRequest) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *UpdateAiModelCardRequest) GetMeta() *UpdateAiModelCardRequestMeta {
	return s.Meta
}

func (s *UpdateAiModelCardRequest) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateAiModelCardRequest) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *UpdateAiModelCardRequest) SetAvailablePaths(v []*UpdateAiModelCardRequestAvailablePaths) *UpdateAiModelCardRequest {
	s.AvailablePaths = v
	return s
}

func (s *UpdateAiModelCardRequest) SetCredit(v *UpdateAiModelCardRequestCredit) *UpdateAiModelCardRequest {
	s.Credit = v
	return s
}

func (s *UpdateAiModelCardRequest) SetFeatures(v map[string]interface{}) *UpdateAiModelCardRequest {
	s.Features = v
	return s
}

func (s *UpdateAiModelCardRequest) SetMeta(v *UpdateAiModelCardRequestMeta) *UpdateAiModelCardRequest {
	s.Meta = v
	return s
}

func (s *UpdateAiModelCardRequest) SetModelName(v string) *UpdateAiModelCardRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateAiModelCardRequest) SetModelProvider(v string) *UpdateAiModelCardRequest {
	s.ModelProvider = &v
	return s
}

func (s *UpdateAiModelCardRequest) Validate() error {
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

type UpdateAiModelCardRequestAvailablePaths struct {
	// The model invocation path. Maximum length: 2048 characters.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The protocol type of the path. Maximum length: 64 characters.
	//
	// example:
	//
	// OpenAICompatible
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAiModelCardRequestAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardRequestAvailablePaths) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardRequestAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *UpdateAiModelCardRequestAvailablePaths) GetType() *string {
	return s.Type
}

func (s *UpdateAiModelCardRequestAvailablePaths) SetPath(v string) *UpdateAiModelCardRequestAvailablePaths {
	s.Path = &v
	return s
}

func (s *UpdateAiModelCardRequestAvailablePaths) SetType(v string) *UpdateAiModelCardRequestAvailablePaths {
	s.Type = &v
	return s
}

func (s *UpdateAiModelCardRequestAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type UpdateAiModelCardRequestCredit struct {
	// The cache hit token cost in Credits per million tokens. The value must be greater than or equal to 0. Default value: 0.
	//
	// example:
	//
	// 0.5
	CacheCost *float32 `json:"cacheCost,omitempty" xml:"cacheCost,omitempty"`
	// The input token cost in Credits per million tokens. The value must be greater than or equal to 0. Default value: 0.
	//
	// example:
	//
	// 1.5
	InputCost *float32 `json:"inputCost,omitempty" xml:"inputCost,omitempty"`
	// The output token cost in Credits per million tokens. The value must be greater than or equal to 0. Default value: 0.
	//
	// example:
	//
	// 3
	OutputCost *float32 `json:"outputCost,omitempty" xml:"outputCost,omitempty"`
	// The billing type. Only fixed is supported. Default value: fixed.
	//
	// example:
	//
	// fixed
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAiModelCardRequestCredit) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardRequestCredit) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardRequestCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *UpdateAiModelCardRequestCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *UpdateAiModelCardRequestCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *UpdateAiModelCardRequestCredit) GetType() *string {
	return s.Type
}

func (s *UpdateAiModelCardRequestCredit) SetCacheCost(v float32) *UpdateAiModelCardRequestCredit {
	s.CacheCost = &v
	return s
}

func (s *UpdateAiModelCardRequestCredit) SetInputCost(v float32) *UpdateAiModelCardRequestCredit {
	s.InputCost = &v
	return s
}

func (s *UpdateAiModelCardRequestCredit) SetOutputCost(v float32) *UpdateAiModelCardRequestCredit {
	s.OutputCost = &v
	return s
}

func (s *UpdateAiModelCardRequestCredit) SetType(v string) *UpdateAiModelCardRequestCredit {
	s.Type = &v
	return s
}

func (s *UpdateAiModelCardRequestCredit) Validate() error {
	return dara.Validate(s)
}

type UpdateAiModelCardRequestMeta struct {
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
	// The list of input modalities supported by the model. The list contains up to 16 items, and each item must not be empty.
	SupportedInputModalities []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	// The list of output modalities supported by the model. The list contains up to 16 items, and each item must not be empty.
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s UpdateAiModelCardRequestMeta) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelCardRequestMeta) GoString() string {
	return s.String()
}

func (s *UpdateAiModelCardRequestMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *UpdateAiModelCardRequestMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *UpdateAiModelCardRequestMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *UpdateAiModelCardRequestMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *UpdateAiModelCardRequestMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *UpdateAiModelCardRequestMeta) SetMaxInputTokens(v int64) *UpdateAiModelCardRequestMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *UpdateAiModelCardRequestMeta) SetMaxOutputTokens(v int64) *UpdateAiModelCardRequestMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *UpdateAiModelCardRequestMeta) SetMaxTokens(v int64) *UpdateAiModelCardRequestMeta {
	s.MaxTokens = &v
	return s
}

func (s *UpdateAiModelCardRequestMeta) SetSupportedInputModalities(v []*string) *UpdateAiModelCardRequestMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *UpdateAiModelCardRequestMeta) SetSupportedOutputModalities(v []*string) *UpdateAiModelCardRequestMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *UpdateAiModelCardRequestMeta) Validate() error {
	return dara.Validate(s)
}
