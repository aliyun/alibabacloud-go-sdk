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
	SetMaxTokens(v int32) *UpdateLLMConfigRequest
	GetMaxTokens() *int32
	SetModel(v string) *UpdateLLMConfigRequest
	GetModel() *string
	SetName(v string) *UpdateLLMConfigRequest
	GetName() *string
	SetRps(v int32) *UpdateLLMConfigRequest
	GetRps() *int32
}

type UpdateLLMConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// apikey-***
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// example:
	//
	// 8
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// example:
	//
	// 1024
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2048
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text-embedding-v1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-config1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *UpdateLLMConfigRequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *UpdateLLMConfigRequest) GetModel() *string {
	return s.Model
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

func (s *UpdateLLMConfigRequest) SetMaxTokens(v int32) *UpdateLLMConfigRequest {
	s.MaxTokens = &v
	return s
}

func (s *UpdateLLMConfigRequest) SetModel(v string) *UpdateLLMConfigRequest {
	s.Model = &v
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
