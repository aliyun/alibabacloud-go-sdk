// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLLMConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *UpdateLLMConfigRequest
	GetApiKey() *string
	SetBaseUrl(v string) *UpdateLLMConfigRequest
	GetBaseUrl() *string
	SetBatchSize(v int32) *UpdateLLMConfigRequest
	GetBatchSize() *int32
	SetEmbeddingDimension(v int32) *UpdateLLMConfigRequest
	GetEmbeddingDimension() *int32
	SetEnableFusion(v bool) *UpdateLLMConfigRequest
	GetEnableFusion() *bool
	SetMaxTokens(v int32) *UpdateLLMConfigRequest
	GetMaxTokens() *int32
	SetModel(v string) *UpdateLLMConfigRequest
	GetModel() *string
	SetModelType(v string) *UpdateLLMConfigRequest
	GetModelType() *string
	SetName(v string) *UpdateLLMConfigRequest
	GetName() *string
	SetRps(v int32) *UpdateLLMConfigRequest
	GetRps() *int32
}

type UpdateLLMConfigRequest struct {
	// The API key used to call the large language model (LLM).
	//
	// This parameter is required.
	//
	// example:
	//
	// apikey-***
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The base URL for calling the large language model (LLM).
	//
	// This parameter is required.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The batch size.
	//
	// example:
	//
	// 8
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// The embedding dimension. If you omit this parameter or set it to 0, the model uses its default dimension.
	//
	// example:
	//
	// 1024
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// Whether to enable data fusion.
	EnableFusion *bool `json:"EnableFusion,omitempty" xml:"EnableFusion,omitempty"`
	// The maximum number of input tokens per row.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2048
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// The model name.
	//
	// This parameter is required.
	//
	// example:
	//
	// text-embedding-v1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The model type.
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The name of the large language model (LLM) call configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-config1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of requests per second (RPS).
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Rps *int32 `json:"Rps,omitempty" xml:"Rps,omitempty"`
}

func (s UpdateLLMConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLLMConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLLMConfigRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateLLMConfigRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *UpdateLLMConfigRequest) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *UpdateLLMConfigRequest) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *UpdateLLMConfigRequest) GetEnableFusion() *bool {
	return s.EnableFusion
}

func (s *UpdateLLMConfigRequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *UpdateLLMConfigRequest) GetModel() *string {
	return s.Model
}

func (s *UpdateLLMConfigRequest) GetModelType() *string {
	return s.ModelType
}

func (s *UpdateLLMConfigRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLLMConfigRequest) GetRps() *int32 {
	return s.Rps
}

func (s *UpdateLLMConfigRequest) SetApiKey(v string) *UpdateLLMConfigRequest {
	s.ApiKey = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetBaseUrl(v string) *UpdateLLMConfigRequest {
	s.BaseUrl = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetBatchSize(v int32) *UpdateLLMConfigRequest {
	s.BatchSize = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetEmbeddingDimension(v int32) *UpdateLLMConfigRequest {
	s.EmbeddingDimension = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetEnableFusion(v bool) *UpdateLLMConfigRequest {
	s.EnableFusion = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetMaxTokens(v int32) *UpdateLLMConfigRequest {
	s.MaxTokens = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetModel(v string) *UpdateLLMConfigRequest {
	s.Model = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetModelType(v string) *UpdateLLMConfigRequest {
	s.ModelType = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetName(v string) *UpdateLLMConfigRequest {
	s.Name = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetRps(v int32) *UpdateLLMConfigRequest {
	s.Rps = &v
	return s
}

func (s *UpdateLLMConfigRequest) Validate() error {
	return dara.Validate(s)
}
