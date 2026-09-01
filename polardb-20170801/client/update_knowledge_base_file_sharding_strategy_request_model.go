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
	// This parameter is required.
	//
	// example:
	//
	// e347ddb8-49bb-5c66-94bc-fa05cedaeac8
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// false
	InheritSpaceStrategy *bool `json:"InheritSpaceStrategy,omitempty" xml:"InheritSpaceStrategy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pkb-2zesv6l6a63xsrym
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId               *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// This parameter is required.
	DefaultStrategy *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategy `json:"DefaultStrategy,omitempty" xml:"DefaultStrategy,omitempty" type:"Struct"`
	Rules           []*UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRules         `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
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
	Parameters *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigDefaultStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
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
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
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
	// This parameter is required.
	Match *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesMatch `json:"Match,omitempty" xml:"Match,omitempty" type:"Struct"`
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
	Parameters *UpdateKnowledgeBaseFileShardingStrategyRequestShardingStrategyConfigRulesStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
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
	// example:
	//
	// auto
	MarkdownTables *string `json:"MarkdownTables,omitempty" xml:"MarkdownTables,omitempty"`
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
