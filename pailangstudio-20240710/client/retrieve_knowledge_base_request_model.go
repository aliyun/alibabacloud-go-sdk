// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHybridStrategyConfig(v string) *RetrieveKnowledgeBaseRequest
	GetHybridStrategyConfig() *string
	SetMetaDataFilterConditions(v string) *RetrieveKnowledgeBaseRequest
	GetMetaDataFilterConditions() *string
	SetQuery(v string) *RetrieveKnowledgeBaseRequest
	GetQuery() *string
	SetQueryMode(v string) *RetrieveKnowledgeBaseRequest
	GetQueryMode() *string
	SetRerankConfig(v string) *RetrieveKnowledgeBaseRequest
	GetRerankConfig() *string
	SetRewriteConfig(v string) *RetrieveKnowledgeBaseRequest
	GetRewriteConfig() *string
	SetScoreThreshold(v float32) *RetrieveKnowledgeBaseRequest
	GetScoreThreshold() *float32
	SetTopK(v int32) *RetrieveKnowledgeBaseRequest
	GetTopK() *int32
	SetVersionName(v string) *RetrieveKnowledgeBaseRequest
	GetVersionName() *string
	SetWorkspaceId(v string) *RetrieveKnowledgeBaseRequest
	GetWorkspaceId() *string
}

type RetrieveKnowledgeBaseRequest struct {
	// example:
	//
	// {
	//
	//   "Strategy": "rrf",
	//
	//   "RRFK":60,
	//
	//   "Weight": 0.5
	//
	// }
	HybridStrategyConfig *string `json:"HybridStrategyConfig,omitempty" xml:"HybridStrategyConfig,omitempty"`
	// example:
	//
	// {
	//
	//     "FilterCondition": "and",
	//
	//     "MetaDataFilters": [
	//
	//         {
	//
	//             "Key": "key1",
	//
	//             "Value": "value1",
	//
	//             "Operator": "=="
	//
	//         },
	//
	//         {
	//
	//             "Key": "key2",
	//
	//             "Value": "value2",
	//
	//             "Operator": "!="
	//
	//         },
	//
	//         {
	//
	//             "Key": "file_name",
	//
	//             "Value": "prefix",
	//
	//             "Operator": "contains"
	//
	//         }
	//
	//     ]
	//
	// }
	MetaDataFilterConditions *string `json:"MetaDataFilterConditions,omitempty" xml:"MetaDataFilterConditions,omitempty"`
	// This parameter is required.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// dense
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// example:
	//
	// {
	//
	//    "ConnectionId":"conn-xxx",
	//
	//     "Model": "qwen-max",
	//
	//     "TopK": 5
	//
	// }
	RerankConfig *string `json:"RerankConfig,omitempty" xml:"RerankConfig,omitempty"`
	// example:
	//
	// {
	//
	//     "ConnectionId":"conn-xxx",
	//
	//     "Model": "qwen-max",
	//
	//     "Temperature": 0.7,
	//
	//     "TopP": 0.9,
	//
	//     "PresencePenalty": 0.5,
	//
	//     "FrequencyPenalty": 0.5,
	//
	//     "Seed": 0,
	//
	//     "MaxTokens": 1024,
	//
	//     "Stop": [],
	//
	//     "EnableThinking": true
	//
	// }
	RewriteConfig  *string  `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty"`
	ScoreThreshold *float32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// example:
	//
	// 5
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 174***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RetrieveKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseRequest) GetHybridStrategyConfig() *string {
	return s.HybridStrategyConfig
}

func (s *RetrieveKnowledgeBaseRequest) GetMetaDataFilterConditions() *string {
	return s.MetaDataFilterConditions
}

func (s *RetrieveKnowledgeBaseRequest) GetQuery() *string {
	return s.Query
}

func (s *RetrieveKnowledgeBaseRequest) GetQueryMode() *string {
	return s.QueryMode
}

func (s *RetrieveKnowledgeBaseRequest) GetRerankConfig() *string {
	return s.RerankConfig
}

func (s *RetrieveKnowledgeBaseRequest) GetRewriteConfig() *string {
	return s.RewriteConfig
}

func (s *RetrieveKnowledgeBaseRequest) GetScoreThreshold() *float32 {
	return s.ScoreThreshold
}

func (s *RetrieveKnowledgeBaseRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RetrieveKnowledgeBaseRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *RetrieveKnowledgeBaseRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RetrieveKnowledgeBaseRequest) SetHybridStrategyConfig(v string) *RetrieveKnowledgeBaseRequest {
	s.HybridStrategyConfig = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetMetaDataFilterConditions(v string) *RetrieveKnowledgeBaseRequest {
	s.MetaDataFilterConditions = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetQuery(v string) *RetrieveKnowledgeBaseRequest {
	s.Query = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetQueryMode(v string) *RetrieveKnowledgeBaseRequest {
	s.QueryMode = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetRerankConfig(v string) *RetrieveKnowledgeBaseRequest {
	s.RerankConfig = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetRewriteConfig(v string) *RetrieveKnowledgeBaseRequest {
	s.RewriteConfig = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetScoreThreshold(v float32) *RetrieveKnowledgeBaseRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetTopK(v int32) *RetrieveKnowledgeBaseRequest {
	s.TopK = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetVersionName(v string) *RetrieveKnowledgeBaseRequest {
	s.VersionName = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetWorkspaceId(v string) *RetrieveKnowledgeBaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
