// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLLMConfigs(v []*ListLLMConfigsResponseBodyLLMConfigs) *ListLLMConfigsResponseBody
	GetLLMConfigs() []*ListLLMConfigsResponseBodyLLMConfigs
	SetMaxResults(v int32) *ListLLMConfigsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListLLMConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListLLMConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListLLMConfigsResponseBody
	GetTotalCount() *int64
}

type ListLLMConfigsResponseBody struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	LLMConfigs []*ListLLMConfigsResponseBodyLLMConfigs `json:"LLMConfigs,omitempty" xml:"LLMConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6mbU5D/SFHCHMApYkMcWlp5
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 898DB17C-09E9-5C41-909D-269BA183EB92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLLMConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLLMConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLLMConfigsResponseBody) GetLLMConfigs() []*ListLLMConfigsResponseBodyLLMConfigs {
	return s.LLMConfigs
}

func (s *ListLLMConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLLMConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLLMConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLLMConfigsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLLMConfigsResponseBody) SetLLMConfigs(v []*ListLLMConfigsResponseBodyLLMConfigs) *ListLLMConfigsResponseBody {
	s.LLMConfigs = v
	return s
}

func (s *ListLLMConfigsResponseBody) SetMaxResults(v int32) *ListLLMConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLLMConfigsResponseBody) SetNextToken(v string) *ListLLMConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLLMConfigsResponseBody) SetRequestId(v string) *ListLLMConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLLMConfigsResponseBody) SetTotalCount(v int64) *ListLLMConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLLMConfigsResponseBody) Validate() error {
	if s.LLMConfigs != nil {
		for _, item := range s.LLMConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLLMConfigsResponseBodyLLMConfigs struct {
	// example:
	//
	// apikey-abcdxy
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
	// example:
	//
	// rg-aek2vtzqjaohzqi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 30
	Rps *int32 `json:"Rps,omitempty" xml:"Rps,omitempty"`
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListLLMConfigsResponseBodyLLMConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListLLMConfigsResponseBodyLLMConfigs) GoString() string {
	return s.String()
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetLLMConfigId() *string {
	return s.LLMConfigId
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetModel() *string {
	return s.Model
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetName() *string {
	return s.Name
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetRps() *int32 {
	return s.Rps
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetApiKey(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.ApiKey = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetBaseUrl(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.BaseUrl = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetBatchSize(v int32) *ListLLMConfigsResponseBodyLLMConfigs {
	s.BatchSize = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetEmbeddingDimension(v int32) *ListLLMConfigsResponseBodyLLMConfigs {
	s.EmbeddingDimension = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetGmtCreateTime(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetGmtModifiedTime(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetLLMConfigId(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.LLMConfigId = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetMaxTokens(v int32) *ListLLMConfigsResponseBodyLLMConfigs {
	s.MaxTokens = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetModel(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.Model = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetName(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.Name = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetResourceGroupId(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetRps(v int32) *ListLLMConfigsResponseBodyLLMConfigs {
	s.Rps = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) SetWorkspaceId(v string) *ListLLMConfigsResponseBodyLLMConfigs {
	s.WorkspaceId = &v
	return s
}

func (s *ListLLMConfigsResponseBodyLLMConfigs) Validate() error {
	return dara.Validate(s)
}
