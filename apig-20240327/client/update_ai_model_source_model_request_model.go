// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelSourceModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailablePaths(v []*UpdateAiModelSourceModelRequestAvailablePaths) *UpdateAiModelSourceModelRequest
	GetAvailablePaths() []*UpdateAiModelSourceModelRequestAvailablePaths
	SetCredit(v *UpdateAiModelSourceModelRequestCredit) *UpdateAiModelSourceModelRequest
	GetCredit() *UpdateAiModelSourceModelRequestCredit
	SetExpectedVersion(v int64) *UpdateAiModelSourceModelRequest
	GetExpectedVersion() *int64
	SetFeatures(v map[string]interface{}) *UpdateAiModelSourceModelRequest
	GetFeatures() map[string]interface{}
	SetMeta(v *UpdateAiModelSourceModelRequestMeta) *UpdateAiModelSourceModelRequest
	GetMeta() *UpdateAiModelSourceModelRequestMeta
	SetModelName(v string) *UpdateAiModelSourceModelRequest
	GetModelName() *string
}

type UpdateAiModelSourceModelRequest struct {
	// The protocols and invocation paths supported by the model.
	AvailablePaths []*UpdateAiModelSourceModelRequestAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	// The credit pricing configuration of the model.
	Credit *UpdateAiModelSourceModelRequestCredit `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	// The expected object version. Use the version returned by the most recent query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
	// The set of capabilities declared by the model. The key is the capability name, and the value indicates whether the capability is supported.
	//
	// example:
	//
	// {"nativeStreaming":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// The token limits and input and output modality information of the model.
	Meta *UpdateAiModelSourceModelRequestMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The actual model name used by the upstream service.
	//
	// example:
	//
	// qwen-plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
}

func (s UpdateAiModelSourceModelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelSourceModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiModelSourceModelRequest) GetAvailablePaths() []*UpdateAiModelSourceModelRequestAvailablePaths {
	return s.AvailablePaths
}

func (s *UpdateAiModelSourceModelRequest) GetCredit() *UpdateAiModelSourceModelRequestCredit {
	return s.Credit
}

func (s *UpdateAiModelSourceModelRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateAiModelSourceModelRequest) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *UpdateAiModelSourceModelRequest) GetMeta() *UpdateAiModelSourceModelRequestMeta {
	return s.Meta
}

func (s *UpdateAiModelSourceModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateAiModelSourceModelRequest) SetAvailablePaths(v []*UpdateAiModelSourceModelRequestAvailablePaths) *UpdateAiModelSourceModelRequest {
	s.AvailablePaths = v
	return s
}

func (s *UpdateAiModelSourceModelRequest) SetCredit(v *UpdateAiModelSourceModelRequestCredit) *UpdateAiModelSourceModelRequest {
	s.Credit = v
	return s
}

func (s *UpdateAiModelSourceModelRequest) SetExpectedVersion(v int64) *UpdateAiModelSourceModelRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateAiModelSourceModelRequest) SetFeatures(v map[string]interface{}) *UpdateAiModelSourceModelRequest {
	s.Features = v
	return s
}

func (s *UpdateAiModelSourceModelRequest) SetMeta(v *UpdateAiModelSourceModelRequestMeta) *UpdateAiModelSourceModelRequest {
	s.Meta = v
	return s
}

func (s *UpdateAiModelSourceModelRequest) SetModelName(v string) *UpdateAiModelSourceModelRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateAiModelSourceModelRequest) Validate() error {
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

type UpdateAiModelSourceModelRequestAvailablePaths struct {
	// The invocation path of the model.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The protocol type used by the invocation path.
	//
	// example:
	//
	// OpenAI/v1
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAiModelSourceModelRequestAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelSourceModelRequestAvailablePaths) GoString() string {
	return s.String()
}

func (s *UpdateAiModelSourceModelRequestAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *UpdateAiModelSourceModelRequestAvailablePaths) GetType() *string {
	return s.Type
}

func (s *UpdateAiModelSourceModelRequestAvailablePaths) SetPath(v string) *UpdateAiModelSourceModelRequestAvailablePaths {
	s.Path = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestAvailablePaths) SetType(v string) *UpdateAiModelSourceModelRequestAvailablePaths {
	s.Type = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type UpdateAiModelSourceModelRequestCredit struct {
	// The number of credits per million cached input tokens. The value must be a non-negative number. A value of 0 indicates that cached input is free of charge.
	//
	// example:
	//
	// 0.1
	CacheCost *float32 `json:"cacheCost,omitempty" xml:"cacheCost,omitempty"`
	// The number of credits per million input tokens. The value must be a non-negative number. A value of 0 indicates that the input is free of charge.
	//
	// example:
	//
	// 1
	InputCost *float32 `json:"inputCost,omitempty" xml:"inputCost,omitempty"`
	// The number of credits per million output tokens. The value must be a non-negative number. A value of 0 indicates that the output is free of charge.
	//
	// example:
	//
	// 2
	OutputCost *float32 `json:"outputCost,omitempty" xml:"outputCost,omitempty"`
	// The pricing method. Valid value: fixed (fixed credit pricing).
	//
	// example:
	//
	// fixed
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAiModelSourceModelRequestCredit) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelSourceModelRequestCredit) GoString() string {
	return s.String()
}

func (s *UpdateAiModelSourceModelRequestCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *UpdateAiModelSourceModelRequestCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *UpdateAiModelSourceModelRequestCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *UpdateAiModelSourceModelRequestCredit) GetType() *string {
	return s.Type
}

func (s *UpdateAiModelSourceModelRequestCredit) SetCacheCost(v float32) *UpdateAiModelSourceModelRequestCredit {
	s.CacheCost = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestCredit) SetInputCost(v float32) *UpdateAiModelSourceModelRequestCredit {
	s.InputCost = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestCredit) SetOutputCost(v float32) *UpdateAiModelSourceModelRequestCredit {
	s.OutputCost = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestCredit) SetType(v string) *UpdateAiModelSourceModelRequestCredit {
	s.Type = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestCredit) Validate() error {
	return dara.Validate(s)
}

type UpdateAiModelSourceModelRequestMeta struct {
	// The maximum number of input tokens supported by the model.
	//
	// example:
	//
	// 32768
	MaxInputTokens *int64 `json:"maxInputTokens,omitempty" xml:"maxInputTokens,omitempty"`
	// The maximum number of output tokens supported by the model.
	//
	// example:
	//
	// 8192
	MaxOutputTokens *int64 `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	// The maximum number of context tokens supported by the model.
	//
	// example:
	//
	// 40960
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The input modalities supported by the model.
	SupportedInputModalities []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	// The output modalities supported by the model.
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s UpdateAiModelSourceModelRequestMeta) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelSourceModelRequestMeta) GoString() string {
	return s.String()
}

func (s *UpdateAiModelSourceModelRequestMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *UpdateAiModelSourceModelRequestMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *UpdateAiModelSourceModelRequestMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *UpdateAiModelSourceModelRequestMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *UpdateAiModelSourceModelRequestMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *UpdateAiModelSourceModelRequestMeta) SetMaxInputTokens(v int64) *UpdateAiModelSourceModelRequestMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestMeta) SetMaxOutputTokens(v int64) *UpdateAiModelSourceModelRequestMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestMeta) SetMaxTokens(v int64) *UpdateAiModelSourceModelRequestMeta {
	s.MaxTokens = &v
	return s
}

func (s *UpdateAiModelSourceModelRequestMeta) SetSupportedInputModalities(v []*string) *UpdateAiModelSourceModelRequestMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *UpdateAiModelSourceModelRequestMeta) SetSupportedOutputModalities(v []*string) *UpdateAiModelSourceModelRequestMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *UpdateAiModelSourceModelRequestMeta) Validate() error {
	return dara.Validate(s)
}
