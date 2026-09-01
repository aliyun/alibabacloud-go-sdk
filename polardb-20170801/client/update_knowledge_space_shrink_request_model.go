// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeSpaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateKnowledgeSpaceShrinkRequest
	GetDescription() *string
	SetKnowledgeSpaceId(v string) *UpdateKnowledgeSpaceShrinkRequest
	GetKnowledgeSpaceId() *string
	SetLLMModel(v string) *UpdateKnowledgeSpaceShrinkRequest
	GetLLMModel() *string
	SetName(v string) *UpdateKnowledgeSpaceShrinkRequest
	GetName() *string
	SetRegionId(v string) *UpdateKnowledgeSpaceShrinkRequest
	GetRegionId() *string
	SetRerankModel(v string) *UpdateKnowledgeSpaceShrinkRequest
	GetRerankModel() *string
	SetShardingStrategyConfigShrink(v string) *UpdateKnowledgeSpaceShrinkRequest
	GetShardingStrategyConfigShrink() *string
}

type UpdateKnowledgeSpaceShrinkRequest struct {
	// The description of the knowledge space. The description can be up to 512 characters in length.
	//
	// example:
	//
	// testDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the knowledge space.
	//
	// This parameter is required.
	//
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// The name of the large language model.
	//
	// example:
	//
	// qwen3.6-plus
	LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
	// The name of the knowledge space. The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the reranking model.
	//
	// example:
	//
	// qwen3-rerank
	RerankModel *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// The default chunking strategy configuration for the knowledge space. Both simple strategies and composite strategies that match by content type are supported.
	ShardingStrategyConfigShrink *string `json:"ShardingStrategyConfig,omitempty" xml:"ShardingStrategyConfig,omitempty"`
}

func (s UpdateKnowledgeSpaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeSpaceShrinkRequest) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *UpdateKnowledgeSpaceShrinkRequest) GetLLMModel() *string {
	return s.LLMModel
}

func (s *UpdateKnowledgeSpaceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeSpaceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKnowledgeSpaceShrinkRequest) GetRerankModel() *string {
	return s.RerankModel
}

func (s *UpdateKnowledgeSpaceShrinkRequest) GetShardingStrategyConfigShrink() *string {
	return s.ShardingStrategyConfigShrink
}

func (s *UpdateKnowledgeSpaceShrinkRequest) SetDescription(v string) *UpdateKnowledgeSpaceShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeSpaceShrinkRequest) SetKnowledgeSpaceId(v string) *UpdateKnowledgeSpaceShrinkRequest {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *UpdateKnowledgeSpaceShrinkRequest) SetLLMModel(v string) *UpdateKnowledgeSpaceShrinkRequest {
	s.LLMModel = &v
	return s
}

func (s *UpdateKnowledgeSpaceShrinkRequest) SetName(v string) *UpdateKnowledgeSpaceShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeSpaceShrinkRequest) SetRegionId(v string) *UpdateKnowledgeSpaceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKnowledgeSpaceShrinkRequest) SetRerankModel(v string) *UpdateKnowledgeSpaceShrinkRequest {
	s.RerankModel = &v
	return s
}

func (s *UpdateKnowledgeSpaceShrinkRequest) SetShardingStrategyConfigShrink(v string) *UpdateKnowledgeSpaceShrinkRequest {
	s.ShardingStrategyConfigShrink = &v
	return s
}

func (s *UpdateKnowledgeSpaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
