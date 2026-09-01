// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeSpaceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetACLMode(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetACLMode() *string
	SetCreationTime(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetCreationTime() *string
	SetDBClusterId(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetDBClusterId() *string
	SetDBName(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetDBName() *string
	SetDBType(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetDBType() *string
	SetDescription(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetDescription() *string
	SetEmbeddingDimension(v int32) *DescribeKnowledgeSpaceAttributeResponseBody
	GetEmbeddingDimension() *int32
	SetEmbeddingModel(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetEmbeddingModel() *string
	SetKnowledgeBaseCount(v int32) *DescribeKnowledgeSpaceAttributeResponseBody
	GetKnowledgeBaseCount() *int32
	SetKnowledgeSpaceId(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetKnowledgeSpaceId() *string
	SetLLMModel(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetLLMModel() *string
	SetName(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetName() *string
	SetOSSBucket(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetOSSBucket() *string
	SetRequestId(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetRequestId() *string
	SetRerankModel(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetRerankModel() *string
	SetShardSize(v int32) *DescribeKnowledgeSpaceAttributeResponseBody
	GetShardSize() *int32
	SetShardingStrategyConfig(v *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) *DescribeKnowledgeSpaceAttributeResponseBody
	GetShardingStrategyConfig() *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig
	SetStatus(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetStatus() *string
	SetStrategy(v string) *DescribeKnowledgeSpaceAttributeResponseBody
	GetStrategy() *string
	SetTotalDocs(v int32) *DescribeKnowledgeSpaceAttributeResponseBody
	GetTotalDocs() *int32
	SetTotalSizeBytes(v int64) *DescribeKnowledgeSpaceAttributeResponseBody
	GetTotalSizeBytes() *int64
}

type DescribeKnowledgeSpaceAttributeResponseBody struct {
	// The access control list (ACL) mode of the knowledge space. Valid values:
	//
	// - DISABLED
	//
	// - ENFORCED
	//
	// example:
	//
	// ENFORCED
	ACLMode *string `json:"ACLMode,omitempty" xml:"ACLMode,omitempty"`
	// The time when the knowledge space was created.
	//
	// example:
	//
	// 2026-06-25T09:53:44Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the PolarDB instance.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// polar_rag_meta
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The type of the database engine. Valid values:
	//
	// 	- MySQL
	//
	// 	- PostgreSQL
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The description of the knowledge space.
	//
	// example:
	//
	// testDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The vector dimensions.
	//
	// example:
	//
	// 1536
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// The embedding model.
	//
	// example:
	//
	// text-embedding-v4
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// The total number of knowledge bases.
	//
	// example:
	//
	// 1
	KnowledgeBaseCount *int32 `json:"KnowledgeBaseCount,omitempty" xml:"KnowledgeBaseCount,omitempty"`
	// The unique identifier of the knowledge space.
	//
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// The large language model.
	//
	// example:
	//
	// qwen3.6-plus
	LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
	// The name of the knowledge space.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// OSS Bucket
	//
	// example:
	//
	// test-bucket
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The reranking model.
	//
	// example:
	//
	// qwen3-rerank
	RerankModel *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// The chunk size in tokens.
	//
	// example:
	//
	// 512
	ShardSize *int32 `json:"ShardSize,omitempty" xml:"ShardSize,omitempty"`
	// The default chunking strategy configuration of the knowledge space. This parameter may be empty if existing instances do not have the complete configuration saved.
	ShardingStrategyConfig *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig `json:"ShardingStrategyConfig,omitempty" xml:"ShardingStrategyConfig,omitempty" type:"Struct"`
	// The instance status.
	//
	// example:
	//
	// Activation
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The chunking strategy.
	//
	// example:
	//
	// hybrid
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The total number of documents.
	//
	// example:
	//
	// 1
	TotalDocs *int32 `json:"TotalDocs,omitempty" xml:"TotalDocs,omitempty"`
	// The total size in bytes.
	//
	// example:
	//
	// 318881
	TotalSizeBytes *int64 `json:"TotalSizeBytes,omitempty" xml:"TotalSizeBytes,omitempty"`
}

func (s DescribeKnowledgeSpaceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetACLMode() *string {
	return s.ACLMode
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetDBName() *string {
	return s.DBName
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetKnowledgeBaseCount() *int32 {
	return s.KnowledgeBaseCount
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetLLMModel() *string {
	return s.LLMModel
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetRerankModel() *string {
	return s.RerankModel
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetShardSize() *int32 {
	return s.ShardSize
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetShardingStrategyConfig() *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig {
	return s.ShardingStrategyConfig
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetTotalDocs() *int32 {
	return s.TotalDocs
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) GetTotalSizeBytes() *int64 {
	return s.TotalSizeBytes
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetACLMode(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.ACLMode = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetCreationTime(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetDBClusterId(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetDBName(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.DBName = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetDBType(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetDescription(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetEmbeddingDimension(v int32) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.EmbeddingDimension = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetEmbeddingModel(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.EmbeddingModel = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetKnowledgeBaseCount(v int32) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.KnowledgeBaseCount = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetKnowledgeSpaceId(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetLLMModel(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.LLMModel = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetName(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetOSSBucket(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.OSSBucket = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetRequestId(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetRerankModel(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.RerankModel = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetShardSize(v int32) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.ShardSize = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetShardingStrategyConfig(v *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.ShardingStrategyConfig = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetStatus(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetStrategy(v string) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.Strategy = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetTotalDocs(v int32) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.TotalDocs = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) SetTotalSizeBytes(v int64) *DescribeKnowledgeSpaceAttributeResponseBody {
	s.TotalSizeBytes = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBody) Validate() error {
	if s.ShardingStrategyConfig != nil {
		if err := s.ShardingStrategyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig struct {
	// The default chunking strategy. This strategy is used when no rule is matched.
	DefaultStrategy *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy `json:"DefaultStrategy,omitempty" xml:"DefaultStrategy,omitempty" type:"Struct"`
	// The list of override rules that are matched in order.
	Rules []*DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) GetDefaultStrategy() *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy {
	return s.DefaultStrategy
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) GetRules() []*DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules {
	return s.Rules
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) SetDefaultStrategy(v *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig {
	s.DefaultStrategy = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) SetRules(v []*DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig {
	s.Rules = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfig) Validate() error {
	if s.DefaultStrategy != nil {
		if err := s.DefaultStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy struct {
	// The parameter details.
	Parameters *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The type of the default chunking strategy. Valid values: hybrid or hierarchical.
	//
	// example:
	//
	// hybrid
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy) GetParameters() *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters {
	return s.Parameters
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy) GetType() *string {
	return s.Type
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy) SetParameters(v *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy {
	s.Parameters = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy) SetType(v string) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy {
	s.Type = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategy) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters struct {
	// The maximum number of tokens in a single chunk.
	//
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// Specifies whether to merge adjacent small chunks under the same heading.
	//
	// example:
	//
	// true
	MergePeers *bool `json:"MergePeers,omitempty" xml:"MergePeers,omitempty"`
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters) GetMergePeers() *bool {
	return s.MergePeers
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters) SetMaxTokens(v int32) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters {
	s.MaxTokens = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters) SetMergePeers(v bool) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters {
	s.MergePeers = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigDefaultStrategyParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules struct {
	// The content type. Currently, table is supported.
	Match *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch `json:"Match,omitempty" xml:"Match,omitempty" type:"Struct"`
	// The chunking strategy.
	Strategy *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules) GetMatch() *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch {
	return s.Match
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules) GetStrategy() *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy {
	return s.Strategy
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules) SetMatch(v *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules {
	s.Match = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules) SetStrategy(v *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules {
	s.Strategy = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRules) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch struct {
	// The content type. Currently, table is supported.
	//
	// example:
	//
	// table
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch) GetContentType() *string {
	return s.ContentType
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch) SetContentType(v string) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch {
	s.ContentType = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesMatch) Validate() error {
	return dara.Validate(s)
}

type DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy struct {
	// The parameter details.
	Parameters *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The chunking strategy type used when a rule is matched.
	//
	// example:
	//
	// hierarchical
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy) GetParameters() *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters {
	return s.Parameters
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy) GetType() *string {
	return s.Type
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy) SetParameters(v *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy {
	s.Parameters = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy) SetType(v string) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy {
	s.Type = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategy) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters struct {
	// The Markdown table processing mode. Valid values: auto, on, or off.
	//
	// example:
	//
	// auto
	MarkdownTables *string `json:"MarkdownTables,omitempty" xml:"MarkdownTables,omitempty"`
	// The maximum number of tokens in a single chunk for matched content.
	//
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters) GetMarkdownTables() *string {
	return s.MarkdownTables
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters) SetMarkdownTables(v string) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters {
	s.MarkdownTables = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters) SetMaxTokens(v int32) *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters {
	s.MaxTokens = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponseBodyShardingStrategyConfigRulesStrategyParameters) Validate() error {
	return dara.Validate(s)
}
