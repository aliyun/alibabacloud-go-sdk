// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseFileShardingStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest
	GetFileId() *string
	SetInheritSpaceStrategy(v bool) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest
	GetInheritSpaceStrategy() *bool
	SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest
	GetKnowledgeBaseId() *string
	SetRegionId(v string) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest
	GetRegionId() *string
	SetShardingStrategyConfigShrink(v string) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest
	GetShardingStrategyConfigShrink() *string
}

type UpdateKnowledgeBaseFileShardingStrategyShrinkRequest struct {
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
	ShardingStrategyConfigShrink *string `json:"ShardingStrategyConfig,omitempty" xml:"ShardingStrategyConfig,omitempty"`
}

func (s UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) GetFileId() *string {
	return s.FileId
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) GetInheritSpaceStrategy() *bool {
	return s.InheritSpaceStrategy
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) GetShardingStrategyConfigShrink() *string {
	return s.ShardingStrategyConfigShrink
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) SetFileId(v string) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest {
	s.FileId = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) SetInheritSpaceStrategy(v bool) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest {
	s.InheritSpaceStrategy = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) SetRegionId(v string) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) SetShardingStrategyConfigShrink(v string) *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest {
	s.ShardingStrategyConfigShrink = &v
	return s
}

func (s *UpdateKnowledgeBaseFileShardingStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
