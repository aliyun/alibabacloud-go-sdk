// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateKnowledgeSpaceRequest
	GetDescription() *string
	SetKnowledgeSpaceId(v string) *UpdateKnowledgeSpaceRequest
	GetKnowledgeSpaceId() *string
	SetLLMModel(v string) *UpdateKnowledgeSpaceRequest
	GetLLMModel() *string
	SetName(v string) *UpdateKnowledgeSpaceRequest
	GetName() *string
	SetRegionId(v string) *UpdateKnowledgeSpaceRequest
	GetRegionId() *string
	SetRerankModel(v string) *UpdateKnowledgeSpaceRequest
	GetRerankModel() *string
	SetShardingStrategyConfig(v *UpdateKnowledgeSpaceRequestShardingStrategyConfig) *UpdateKnowledgeSpaceRequest
	GetShardingStrategyConfig() *UpdateKnowledgeSpaceRequestShardingStrategyConfig
}

type UpdateKnowledgeSpaceRequest struct {
	// example:
	//
	// testDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// example:
	//
	// qwen3.6-plus
	LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// qwen3-rerank
	RerankModel            *string                                            `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	ShardingStrategyConfig *UpdateKnowledgeSpaceRequestShardingStrategyConfig `json:"ShardingStrategyConfig,omitempty" xml:"ShardingStrategyConfig,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeSpaceRequest) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *UpdateKnowledgeSpaceRequest) GetLLMModel() *string {
	return s.LLMModel
}

func (s *UpdateKnowledgeSpaceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeSpaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKnowledgeSpaceRequest) GetRerankModel() *string {
	return s.RerankModel
}

func (s *UpdateKnowledgeSpaceRequest) GetShardingStrategyConfig() *UpdateKnowledgeSpaceRequestShardingStrategyConfig {
	return s.ShardingStrategyConfig
}

func (s *UpdateKnowledgeSpaceRequest) SetDescription(v string) *UpdateKnowledgeSpaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequest) SetKnowledgeSpaceId(v string) *UpdateKnowledgeSpaceRequest {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequest) SetLLMModel(v string) *UpdateKnowledgeSpaceRequest {
	s.LLMModel = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequest) SetName(v string) *UpdateKnowledgeSpaceRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequest) SetRegionId(v string) *UpdateKnowledgeSpaceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequest) SetRerankModel(v string) *UpdateKnowledgeSpaceRequest {
	s.RerankModel = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequest) SetShardingStrategyConfig(v *UpdateKnowledgeSpaceRequestShardingStrategyConfig) *UpdateKnowledgeSpaceRequest {
	s.ShardingStrategyConfig = v
	return s
}

func (s *UpdateKnowledgeSpaceRequest) Validate() error {
	if s.ShardingStrategyConfig != nil {
		if err := s.ShardingStrategyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeSpaceRequestShardingStrategyConfig struct {
	DefaultStrategy *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy `json:"DefaultStrategy,omitempty" xml:"DefaultStrategy,omitempty" type:"Struct"`
	Rules           []*UpdateKnowledgeSpaceRequestShardingStrategyConfigRules         `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfig) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfig) GetDefaultStrategy() *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy {
	return s.DefaultStrategy
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfig) GetRules() []*UpdateKnowledgeSpaceRequestShardingStrategyConfigRules {
	return s.Rules
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfig) SetDefaultStrategy(v *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy) *UpdateKnowledgeSpaceRequestShardingStrategyConfig {
	s.DefaultStrategy = v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfig) SetRules(v []*UpdateKnowledgeSpaceRequestShardingStrategyConfigRules) *UpdateKnowledgeSpaceRequestShardingStrategyConfig {
	s.Rules = v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfig) Validate() error {
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

type UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy struct {
	Parameters *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// example:
	//
	// hybrid
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy) GetParameters() *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters {
	return s.Parameters
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy) GetType() *string {
	return s.Type
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy) SetParameters(v *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters) *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy {
	s.Parameters = v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy) SetType(v string) *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy {
	s.Type = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategy) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters struct {
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// example:
	//
	// true
	MergePeers *bool `json:"MergePeers,omitempty" xml:"MergePeers,omitempty"`
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters) GetMergePeers() *bool {
	return s.MergePeers
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters) SetMaxTokens(v int32) *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters {
	s.MaxTokens = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters) SetMergePeers(v bool) *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters {
	s.MergePeers = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigDefaultStrategyParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeSpaceRequestShardingStrategyConfigRules struct {
	Match    *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch    `json:"Match,omitempty" xml:"Match,omitempty" type:"Struct"`
	Strategy *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRules) GetMatch() *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch {
	return s.Match
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRules) GetStrategy() *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy {
	return s.Strategy
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRules) SetMatch(v *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch) *UpdateKnowledgeSpaceRequestShardingStrategyConfigRules {
	s.Match = v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRules) SetStrategy(v *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy) *UpdateKnowledgeSpaceRequestShardingStrategyConfigRules {
	s.Strategy = v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRules) Validate() error {
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

type UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch struct {
	// example:
	//
	// table
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch) SetContentType(v string) *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch {
	s.ContentType = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesMatch) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy struct {
	Parameters *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// example:
	//
	// hierarchical
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy) GetParameters() *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters {
	return s.Parameters
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy) GetType() *string {
	return s.Type
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy) SetParameters(v *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters) *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy {
	s.Parameters = v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy) SetType(v string) *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy {
	s.Type = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategy) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters struct {
	// example:
	//
	// auto
	MarkdownTables *string `json:"MarkdownTables,omitempty" xml:"MarkdownTables,omitempty"`
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters) GetMarkdownTables() *string {
	return s.MarkdownTables
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters) SetMarkdownTables(v string) *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters {
	s.MarkdownTables = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters) SetMaxTokens(v int32) *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters {
	s.MaxTokens = &v
	return s
}

func (s *UpdateKnowledgeSpaceRequestShardingStrategyConfigRulesStrategyParameters) Validate() error {
	return dara.Validate(s)
}
