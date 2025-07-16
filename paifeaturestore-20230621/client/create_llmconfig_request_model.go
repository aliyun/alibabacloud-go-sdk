// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLLMConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateLLMConfigRequest
	GetApiKey() *string
	SetBaseUrl(v string) *CreateLLMConfigRequest
	GetBaseUrl() *string
	SetBatchSize(v int32) *CreateLLMConfigRequest
	GetBatchSize() *int32
	SetMaxTokens(v int32) *CreateLLMConfigRequest
	GetMaxTokens() *int32
	SetModel(v string) *CreateLLMConfigRequest
	GetModel() *string
	SetName(v string) *CreateLLMConfigRequest
	GetName() *string
	SetRps(v int32) *CreateLLMConfigRequest
	GetRps() *int32
	SetWorkspaceId(v string) *CreateLLMConfigRequest
	GetWorkspaceId() *string
}

type CreateLLMConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// api-xyz
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// example:
	//
	// 8
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
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
	// llm_config1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30
	Rps *int32 `json:"Rps,omitempty" xml:"Rps,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateLLMConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLLMConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLLMConfigRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateLLMConfigRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *CreateLLMConfigRequest) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *CreateLLMConfigRequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *CreateLLMConfigRequest) GetModel() *string {
	return s.Model
}

func (s *CreateLLMConfigRequest) GetName() *string {
	return s.Name
}

func (s *CreateLLMConfigRequest) GetRps() *int32 {
	return s.Rps
}

func (s *CreateLLMConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateLLMConfigRequest) SetApiKey(v string) *CreateLLMConfigRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateLLMConfigRequest) SetBaseUrl(v string) *CreateLLMConfigRequest {
	s.BaseUrl = &v
	return s
}

func (s *CreateLLMConfigRequest) SetBatchSize(v int32) *CreateLLMConfigRequest {
	s.BatchSize = &v
	return s
}

func (s *CreateLLMConfigRequest) SetMaxTokens(v int32) *CreateLLMConfigRequest {
	s.MaxTokens = &v
	return s
}

func (s *CreateLLMConfigRequest) SetModel(v string) *CreateLLMConfigRequest {
	s.Model = &v
	return s
}

func (s *CreateLLMConfigRequest) SetName(v string) *CreateLLMConfigRequest {
	s.Name = &v
	return s
}

func (s *CreateLLMConfigRequest) SetRps(v int32) *CreateLLMConfigRequest {
	s.Rps = &v
	return s
}

func (s *CreateLLMConfigRequest) SetWorkspaceId(v string) *CreateLLMConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateLLMConfigRequest) Validate() error {
	return dara.Validate(s)
}
