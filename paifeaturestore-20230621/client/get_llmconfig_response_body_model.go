// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLLMConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *GetLLMConfigResponseBody
	GetApiKey() *string
	SetBaseUrl(v string) *GetLLMConfigResponseBody
	GetBaseUrl() *string
	SetBatchSize(v int32) *GetLLMConfigResponseBody
	GetBatchSize() *int32
	SetEmbeddingDimension(v int32) *GetLLMConfigResponseBody
	GetEmbeddingDimension() *int32
	SetGmtCreateTime(v string) *GetLLMConfigResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetLLMConfigResponseBody
	GetGmtModifiedTime() *string
	SetLLMConfigId(v string) *GetLLMConfigResponseBody
	GetLLMConfigId() *string
	SetMaxTokens(v int32) *GetLLMConfigResponseBody
	GetMaxTokens() *int32
	SetModel(v string) *GetLLMConfigResponseBody
	GetModel() *string
	SetName(v string) *GetLLMConfigResponseBody
	GetName() *string
	SetRequestId(v string) *GetLLMConfigResponseBody
	GetRequestId() *string
	SetRps(v int32) *GetLLMConfigResponseBody
	GetRps() *int32
	SetWorkspaceId(v string) *GetLLMConfigResponseBody
	GetWorkspaceId() *string
}

type GetLLMConfigResponseBody struct {
	// example:
	//
	// api-abcdxy
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
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
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// llm_config1
	LLMConfigId *string `json:"LLMConfigId,omitempty" xml:"LLMConfigId,omitempty"`
	// example:
	//
	// 2048
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// example:
	//
	// text-embedding-v1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// llm_config_name1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	Rps *int32 `json:"Rps,omitempty" xml:"Rps,omitempty"`
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetLLMConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLLMConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetLLMConfigResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetLLMConfigResponseBody) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *GetLLMConfigResponseBody) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *GetLLMConfigResponseBody) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *GetLLMConfigResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetLLMConfigResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetLLMConfigResponseBody) GetLLMConfigId() *string {
	return s.LLMConfigId
}

func (s *GetLLMConfigResponseBody) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *GetLLMConfigResponseBody) GetModel() *string {
	return s.Model
}

func (s *GetLLMConfigResponseBody) GetName() *string {
	return s.Name
}

func (s *GetLLMConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLLMConfigResponseBody) GetRps() *int32 {
	return s.Rps
}

func (s *GetLLMConfigResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetLLMConfigResponseBody) SetApiKey(v string) *GetLLMConfigResponseBody {
	s.ApiKey = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetBaseUrl(v string) *GetLLMConfigResponseBody {
	s.BaseUrl = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetBatchSize(v int32) *GetLLMConfigResponseBody {
	s.BatchSize = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetEmbeddingDimension(v int32) *GetLLMConfigResponseBody {
	s.EmbeddingDimension = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetGmtCreateTime(v string) *GetLLMConfigResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetGmtModifiedTime(v string) *GetLLMConfigResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetLLMConfigId(v string) *GetLLMConfigResponseBody {
	s.LLMConfigId = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetMaxTokens(v int32) *GetLLMConfigResponseBody {
	s.MaxTokens = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetModel(v string) *GetLLMConfigResponseBody {
	s.Model = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetName(v string) *GetLLMConfigResponseBody {
	s.Name = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetRequestId(v string) *GetLLMConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetRps(v int32) *GetLLMConfigResponseBody {
	s.Rps = &v
	return s
}

func (s *GetLLMConfigResponseBody) SetWorkspaceId(v string) *GetLLMConfigResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetLLMConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
