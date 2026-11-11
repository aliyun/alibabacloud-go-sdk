// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelParameters interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCacheControl(v bool) *ModelParameters
	GetEnableCacheControl() *bool
	SetEnableThinking(v bool) *ModelParameters
	GetEnableThinking() *bool
	SetExtraBody(v string) *ModelParameters
	GetExtraBody() *string
	SetFrequencyPenalty(v float32) *ModelParameters
	GetFrequencyPenalty() *float32
	SetMaxCompletionTokens(v int32) *ModelParameters
	GetMaxCompletionTokens() *int32
	SetMaxTokens(v int64) *ModelParameters
	GetMaxTokens() *int64
	SetPresencePenalty(v float32) *ModelParameters
	GetPresencePenalty() *float32
	SetReasoningEffort(v string) *ModelParameters
	GetReasoningEffort() *string
	SetRepetitionPenalty(v float64) *ModelParameters
	GetRepetitionPenalty() *float64
	SetResponseFormat(v *ModelParametersResponseFormat) *ModelParameters
	GetResponseFormat() *ModelParametersResponseFormat
	SetSeed(v int32) *ModelParameters
	GetSeed() *int32
	SetStopSequences(v []*string) *ModelParameters
	GetStopSequences() []*string
	SetTemperature(v float32) *ModelParameters
	GetTemperature() *float32
	SetThinkingBudget(v int32) *ModelParameters
	GetThinkingBudget() *int32
	SetTopK(v int32) *ModelParameters
	GetTopK() *int32
	SetTopP(v float32) *ModelParameters
	GetTopP() *float32
}

type ModelParameters struct {
	// Specifies whether to enable explicit caching for system messages.
	//
	// example:
	//
	// true
	EnableCacheControl *bool `json:"enableCacheControl,omitempty" xml:"enableCacheControl,omitempty"`
	// Specifies whether to enable thinking mode. This takes effect only for hybrid thinking models and defaults to false if not specified.
	//
	// example:
	//
	// true
	EnableThinking *bool `json:"enableThinking,omitempty" xml:"enableThinking,omitempty"`
	// The extended parameters, which are a serialized string of a JSON object. These are merged into the top level of the model request body to pass through other Model Studio parameters.
	//
	// example:
	//
	// {"enable_search":false}
	ExtraBody *string `json:"extraBody,omitempty" xml:"extraBody,omitempty"`
	// The frequency penalty where positive values reduce repetition, with a range of -2.0 to 2.0.
	//
	// example:
	//
	// 0.2
	FrequencyPenalty *float32 `json:"frequencyPenalty,omitempty" xml:"frequencyPenalty,omitempty"`
	// The maximum number of output tokens, including the chain-of-thought and the response. This is recommended for thinking models.
	//
	// example:
	//
	// 4096
	MaxCompletionTokens *int32 `json:"maxCompletionTokens,omitempty" xml:"maxCompletionTokens,omitempty"`
	// The maximum number of output tokens.
	//
	// example:
	//
	// 1024
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The presence penalty where positive values encourage new topics, with a range of -2.0 to 2.0.
	//
	// example:
	//
	// 0.1
	PresencePenalty *float32 `json:"presencePenalty,omitempty" xml:"presencePenalty,omitempty"`
	// The reasoning effort level. This takes effect only when thinking mode is enabled. For the Qwen3.8 series, valid values are low, medium, and xhigh. This cannot be set simultaneously with thinkingBudget.
	//
	// example:
	//
	// low
	ReasoningEffort *string `json:"reasoningEffort,omitempty" xml:"reasoningEffort,omitempty"`
	// The repetition penalty that controls the repetition of consecutive sequences. The value must be greater than 0. A value of 1.0 indicates no penalty, and a larger value results in less repetition.
	//
	// example:
	//
	// 1.05
	RepetitionPenalty *float64 `json:"repetitionPenalty,omitempty" xml:"repetitionPenalty,omitempty"`
	// The response content format, which defaults to text if not specified.
	//
	// example:
	//
	// {"type":"json_schema","jsonSchema":{"name":"answer","strict":true,"schema":"{\\"type\\":\\"object\\",\\"properties\\":{\\"answer\\":{\\"type\\":\\"string\\"}},\\"required\\":[\\"answer\\"]}"}}
	ResponseFormat *ModelParametersResponseFormat `json:"responseFormat,omitempty" xml:"responseFormat,omitempty" type:"Struct"`
	// The random seed, with a range of 0 to 2147483647, which attempts to return the same result for the same input and parameters.
	//
	// example:
	//
	// 1234
	Seed *int32 `json:"seed,omitempty" xml:"seed,omitempty"`
	// The stop sequences that immediately halt generation when encountered.
	//
	// example:
	//
	// ["END"]
	StopSequences []*string `json:"stopSequences,omitempty" xml:"stopSequences,omitempty" type:"Repeated"`
	// The temperature parameter that controls the randomness of the output, with a range of 0 to 2.
	//
	// example:
	//
	// 0.7
	Temperature *float32 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// The maximum number of tokens for the thinking process. This takes effect only when thinking mode is enabled and cannot be set simultaneously with reasoningEffort.
	//
	// example:
	//
	// 2048
	ThinkingBudget *int32 `json:"thinkingBudget,omitempty" xml:"thinkingBudget,omitempty"`
	// The Top K sampling parameter, with a range of 1 to 100.
	//
	// example:
	//
	// 50
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
	// The Top P sampling parameter, with a range of 0 to 1.
	//
	// example:
	//
	// 0.9
	TopP *float32 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s ModelParameters) String() string {
	return dara.Prettify(s)
}

func (s ModelParameters) GoString() string {
	return s.String()
}

func (s *ModelParameters) GetEnableCacheControl() *bool {
	return s.EnableCacheControl
}

func (s *ModelParameters) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *ModelParameters) GetExtraBody() *string {
	return s.ExtraBody
}

func (s *ModelParameters) GetFrequencyPenalty() *float32 {
	return s.FrequencyPenalty
}

func (s *ModelParameters) GetMaxCompletionTokens() *int32 {
	return s.MaxCompletionTokens
}

func (s *ModelParameters) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ModelParameters) GetPresencePenalty() *float32 {
	return s.PresencePenalty
}

func (s *ModelParameters) GetReasoningEffort() *string {
	return s.ReasoningEffort
}

func (s *ModelParameters) GetRepetitionPenalty() *float64 {
	return s.RepetitionPenalty
}

func (s *ModelParameters) GetResponseFormat() *ModelParametersResponseFormat {
	return s.ResponseFormat
}

func (s *ModelParameters) GetSeed() *int32 {
	return s.Seed
}

func (s *ModelParameters) GetStopSequences() []*string {
	return s.StopSequences
}

func (s *ModelParameters) GetTemperature() *float32 {
	return s.Temperature
}

func (s *ModelParameters) GetThinkingBudget() *int32 {
	return s.ThinkingBudget
}

func (s *ModelParameters) GetTopK() *int32 {
	return s.TopK
}

func (s *ModelParameters) GetTopP() *float32 {
	return s.TopP
}

func (s *ModelParameters) SetEnableCacheControl(v bool) *ModelParameters {
	s.EnableCacheControl = &v
	return s
}

func (s *ModelParameters) SetEnableThinking(v bool) *ModelParameters {
	s.EnableThinking = &v
	return s
}

func (s *ModelParameters) SetExtraBody(v string) *ModelParameters {
	s.ExtraBody = &v
	return s
}

func (s *ModelParameters) SetFrequencyPenalty(v float32) *ModelParameters {
	s.FrequencyPenalty = &v
	return s
}

func (s *ModelParameters) SetMaxCompletionTokens(v int32) *ModelParameters {
	s.MaxCompletionTokens = &v
	return s
}

func (s *ModelParameters) SetMaxTokens(v int64) *ModelParameters {
	s.MaxTokens = &v
	return s
}

func (s *ModelParameters) SetPresencePenalty(v float32) *ModelParameters {
	s.PresencePenalty = &v
	return s
}

func (s *ModelParameters) SetReasoningEffort(v string) *ModelParameters {
	s.ReasoningEffort = &v
	return s
}

func (s *ModelParameters) SetRepetitionPenalty(v float64) *ModelParameters {
	s.RepetitionPenalty = &v
	return s
}

func (s *ModelParameters) SetResponseFormat(v *ModelParametersResponseFormat) *ModelParameters {
	s.ResponseFormat = v
	return s
}

func (s *ModelParameters) SetSeed(v int32) *ModelParameters {
	s.Seed = &v
	return s
}

func (s *ModelParameters) SetStopSequences(v []*string) *ModelParameters {
	s.StopSequences = v
	return s
}

func (s *ModelParameters) SetTemperature(v float32) *ModelParameters {
	s.Temperature = &v
	return s
}

func (s *ModelParameters) SetThinkingBudget(v int32) *ModelParameters {
	s.ThinkingBudget = &v
	return s
}

func (s *ModelParameters) SetTopK(v int32) *ModelParameters {
	s.TopK = &v
	return s
}

func (s *ModelParameters) SetTopP(v float32) *ModelParameters {
	s.TopP = &v
	return s
}

func (s *ModelParameters) Validate() error {
	if s.ResponseFormat != nil {
		if err := s.ResponseFormat.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelParametersResponseFormat struct {
	// The JSON Schema constraint, which is required when the type is json_schema.
	//
	// example:
	//
	// {"name":"answer","strict":true,"schema":"{\\"type\\":\\"object\\",\\"properties\\":{\\"answer\\":{\\"type\\":\\"string\\"}},\\"required\\":[\\"answer\\"]}"}
	JsonSchema *ModelParametersResponseFormatJsonSchema `json:"jsonSchema,omitempty" xml:"jsonSchema,omitempty" type:"Struct"`
	// The format type. Valid values: text, json_object, and json_schema. When using json_object, the prompt must explicitly request JSON output.
	//
	// example:
	//
	// json_object
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelParametersResponseFormat) String() string {
	return dara.Prettify(s)
}

func (s ModelParametersResponseFormat) GoString() string {
	return s.String()
}

func (s *ModelParametersResponseFormat) GetJsonSchema() *ModelParametersResponseFormatJsonSchema {
	return s.JsonSchema
}

func (s *ModelParametersResponseFormat) GetType() *string {
	return s.Type
}

func (s *ModelParametersResponseFormat) SetJsonSchema(v *ModelParametersResponseFormatJsonSchema) *ModelParametersResponseFormat {
	s.JsonSchema = v
	return s
}

func (s *ModelParametersResponseFormat) SetType(v string) *ModelParametersResponseFormat {
	s.Type = &v
	return s
}

func (s *ModelParametersResponseFormat) Validate() error {
	if s.JsonSchema != nil {
		if err := s.JsonSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelParametersResponseFormatJsonSchema struct {
	// The schema name.
	//
	// example:
	//
	// answer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON Schema definition, which is a serialized string of a JSON object.
	//
	// example:
	//
	// {"type":"object","properties":{"answer":{"type":"string"}},"required":["answer"]}
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// Specifies whether to strictly adhere to the schema structure. Set this to true.
	//
	// example:
	//
	// true
	Strict *bool `json:"strict,omitempty" xml:"strict,omitempty"`
}

func (s ModelParametersResponseFormatJsonSchema) String() string {
	return dara.Prettify(s)
}

func (s ModelParametersResponseFormatJsonSchema) GoString() string {
	return s.String()
}

func (s *ModelParametersResponseFormatJsonSchema) GetName() *string {
	return s.Name
}

func (s *ModelParametersResponseFormatJsonSchema) GetSchema() *string {
	return s.Schema
}

func (s *ModelParametersResponseFormatJsonSchema) GetStrict() *bool {
	return s.Strict
}

func (s *ModelParametersResponseFormatJsonSchema) SetName(v string) *ModelParametersResponseFormatJsonSchema {
	s.Name = &v
	return s
}

func (s *ModelParametersResponseFormatJsonSchema) SetSchema(v string) *ModelParametersResponseFormatJsonSchema {
	s.Schema = &v
	return s
}

func (s *ModelParametersResponseFormatJsonSchema) SetStrict(v bool) *ModelParametersResponseFormatJsonSchema {
	s.Strict = &v
	return s
}

func (s *ModelParametersResponseFormatJsonSchema) Validate() error {
	return dara.Validate(s)
}
