// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseFileShardingStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *UpdateKnowledgeBaseFileShardingStrategyRequest
	GetFileId() *string
	SetInheritSpaceStrategy(v bool) *UpdateKnowledgeBaseFileShardingStrategyRequest
	GetInheritSpaceStrategy() *bool
	SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseFileShardingStrategyRequest
	GetKnowledgeBaseId() *string
	SetRegionId(v string) *UpdateKnowledgeBaseFileShardingStrategyRequest
	GetRegionId() *string
	SetShardingStrategyConfig(v *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) *UpdateKnowledgeBaseFileShardingStrategyRequest
	GetShardingStrategyConfig() *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig
}

type UpdateKnowledgeBaseFileShardingStrategyRequest struct {
	// The unique ID of the knowledge base file.
	//
	// This parameter is required.
	//
	// example:
	//
	// e347ddb8-49bb-5c66-94bc-fa05cedaeac8
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Specifies whether to restore inheritance of the chunking strategy from the knowledge space. When this parameter is set to true, ShardingStrategyConfig cannot be specified at the same time.
	//
	// example:
	//
	// false
	InheritSpaceStrategy *bool `json:"InheritSpaceStrategy,omitempty" xml:"InheritSpaceStrategy,omitempty"`
	// The unique ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-2zesv6l6a63xsrym
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the region where the knowledge base resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The file-level chunking strategy configuration. This parameter is required when InheritSpaceStrategy is not set to true.
	ShardingStrategyConfig *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig `json:"ShardingStrategyConfig,omitempty" xml:"ShardingStrategyConfig,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) GetFileId() *string {
	return s.FileId
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) GetInheritSpaceStrategy() *bool {
	return s.InheritSpaceStrategy
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) GetShardingStrategyConfig() *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig {
	return s.ShardingStrategyConfig
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) SetFileId(v string) *UpdateKnowledgeBaseFileShardingStrategyRequest {
	s.FileId = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) SetInheritSpaceStrategy(v bool) *UpdateKnowledgeBaseFileShardingStrategyRequest {
	s.InheritSpaceStrategy = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseFileShardingStrategyRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) SetRegionId(v string) *UpdateKnowledgeBaseFileShardingStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) SetShardingStrategyConfig(v *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) *UpdateKnowledgeBaseFileShardingStrategyRequest {
	s.ShardingStrategyConfig = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequest) Validate() error {
	if s.ShardingStrategyConfig != nil {
		if err := s.ShardingStrategyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig struct {
	// The default chunking strategy. This strategy is used when no rule is matched.
	//
	// This parameter is required.
	DefaultStrategy *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy `json:"DefaultStrategy,omitempty" xml:"DefaultStrategy,omitempty" type:"Struct"`
	// The list of override rules that are matched in order. Currently, a maximum of one exact-match rule with ContentType set to table is supported.
	Rules []*UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) GetDefaultStrategy() *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy {
	return s.DefaultStrategy
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) GetRules() []*UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules {
	return s.Rules
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) SetDefaultStrategy(v *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig {
	s.DefaultStrategy = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) SetRules(v []*UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig {
	s.Rules = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfig) Validate() error {
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

type UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy struct {
	// The parameters of the default chunking strategy. MaxTokens and MergePeers are supported only when Type is set to hybrid.
	Parameters *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The type of the default chunking strategy. Valid values:
	//
	// - hybrid: Splits by document structure and limits the token count.
	//
	// - hierarchical: Splits only by document structure.
	//
	// This parameter is required.
	//
	// example:
	//
	// hybrid
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy) GetParameters() *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters {
	return s.Parameters
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy) GetType() *string {
	return s.Type
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy) SetParameters(v *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy {
	s.Parameters = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy) SetType(v string) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy {
	s.Type = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters struct {
	// The maximum number of tokens per chunk. The value must be a positive integer. This parameter takes effect only when Type is set to hybrid.
	//
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// Specifies whether to merge adjacent small chunks under the same heading. This parameter takes effect only when Type is set to hybrid.
	//
	// example:
	//
	// true
	MergePeers *bool `json:"MergePeers,omitempty" xml:"MergePeers,omitempty"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters) GetMergePeers() *bool {
	return s.MergePeers
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters) SetMaxTokens(v int32) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters {
	s.MaxTokens = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters) SetMergePeers(v bool) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters {
	s.MergePeers = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules struct {
	// The rule match condition. Currently, only exact matching by content type for table content is supported.
	//
	// This parameter is required.
	Match *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch `json:"Match,omitempty" xml:"Match,omitempty" type:"Struct"`
	// The chunking strategy to use when the rule is matched.
	//
	// This parameter is required.
	Strategy *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules) GetMatch() *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch {
	return s.Match
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules) GetStrategy() *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy {
	return s.Strategy
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules) SetMatch(v *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules {
	s.Match = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules) SetStrategy(v *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules {
	s.Strategy = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules) Validate() error {
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

type UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch struct {
	// The content type. Currently, only table is supported, which matches content that is parsed as tables.
	//
	// example:
	//
	// table
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch) SetContentType(v string) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch {
	s.ContentType = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy struct {
	// The chunking strategy parameters of the override rule. MaxTokens takes effect only when Type is set to hybrid. MarkdownTables supports auto, on, or off.
	Parameters *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The chunking strategy type of the override rule. Valid values:
	//
	// - hybrid
	//
	// - hierarchical
	//
	// This parameter is required.
	//
	// example:
	//
	// hierarchical
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy) GetParameters() *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters {
	return s.Parameters
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy) GetType() *string {
	return s.Type
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy) SetParameters(v *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy {
	s.Parameters = v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy) SetType(v string) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy {
	s.Type = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategy) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters struct {
	// The Markdown table processing mode. Valid values:
	//
	// - auto: Automatically determines the processing mode.
	//
	// - on: Forcefully enables Markdown table processing.
	//
	// - off: Disables Markdown table processing.
	//
	// example:
	//
	// auto
	MarkdownTables *string `json:"MarkdownTables,omitempty" xml:"MarkdownTables,omitempty"`
	// The maximum number of tokens per chunk for matched content. The value must be a positive integer. This parameter takes effect only when Type is set to hybrid.
	//
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters) GetMarkdownTables() *string {
	return s.MarkdownTables
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters) SetMarkdownTables(v string) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters {
	s.MarkdownTables = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters) SetMaxTokens(v int32) *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters {
	s.MaxTokens = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters) Validate() error {
	return dara.Validate(s)
}
