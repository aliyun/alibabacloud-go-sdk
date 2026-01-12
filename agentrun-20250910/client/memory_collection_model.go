// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMemoryCollection interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *MemoryCollection
	GetCreatedAt() *string
	SetDescription(v string) *MemoryCollection
	GetDescription() *string
	SetEmbedderConfig(v *EmbedderConfig) *MemoryCollection
	GetEmbedderConfig() *EmbedderConfig
	SetExecutionRoleArn(v string) *MemoryCollection
	GetExecutionRoleArn() *string
	SetLastUpdatedAt(v string) *MemoryCollection
	GetLastUpdatedAt() *string
	SetLlmConfig(v *LLMConfig) *MemoryCollection
	GetLlmConfig() *LLMConfig
	SetMemoryCollectionId(v string) *MemoryCollection
	GetMemoryCollectionId() *string
	SetMemoryCollectionName(v string) *MemoryCollection
	GetMemoryCollectionName() *string
	SetNetworkConfiguration(v *NetworkConfiguration) *MemoryCollection
	GetNetworkConfiguration() *NetworkConfiguration
	SetVectorStoreConfig(v *VectorStoreConfig) *MemoryCollection
	GetVectorStoreConfig() *VectorStoreConfig
}

type MemoryCollection struct {
	CreatedAt            *string               `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description          *string               `json:"description,omitempty" xml:"description,omitempty"`
	EmbedderConfig       *EmbedderConfig       `json:"embedderConfig,omitempty" xml:"embedderConfig,omitempty"`
	ExecutionRoleArn     *string               `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LastUpdatedAt        *string               `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	LlmConfig            *LLMConfig            `json:"llmConfig,omitempty" xml:"llmConfig,omitempty"`
	MemoryCollectionId   *string               `json:"memoryCollectionId,omitempty" xml:"memoryCollectionId,omitempty"`
	MemoryCollectionName *string               `json:"memoryCollectionName,omitempty" xml:"memoryCollectionName,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	VectorStoreConfig    *VectorStoreConfig    `json:"vectorStoreConfig,omitempty" xml:"vectorStoreConfig,omitempty"`
}

func (s MemoryCollection) String() string {
	return dara.Prettify(s)
}

func (s MemoryCollection) GoString() string {
	return s.String()
}

func (s *MemoryCollection) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *MemoryCollection) GetDescription() *string {
	return s.Description
}

func (s *MemoryCollection) GetEmbedderConfig() *EmbedderConfig {
	return s.EmbedderConfig
}

func (s *MemoryCollection) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *MemoryCollection) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *MemoryCollection) GetLlmConfig() *LLMConfig {
	return s.LlmConfig
}

func (s *MemoryCollection) GetMemoryCollectionId() *string {
	return s.MemoryCollectionId
}

func (s *MemoryCollection) GetMemoryCollectionName() *string {
	return s.MemoryCollectionName
}

func (s *MemoryCollection) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *MemoryCollection) GetVectorStoreConfig() *VectorStoreConfig {
	return s.VectorStoreConfig
}

func (s *MemoryCollection) SetCreatedAt(v string) *MemoryCollection {
	s.CreatedAt = &v
	return s
}

func (s *MemoryCollection) SetDescription(v string) *MemoryCollection {
	s.Description = &v
	return s
}

func (s *MemoryCollection) SetEmbedderConfig(v *EmbedderConfig) *MemoryCollection {
	s.EmbedderConfig = v
	return s
}

func (s *MemoryCollection) SetExecutionRoleArn(v string) *MemoryCollection {
	s.ExecutionRoleArn = &v
	return s
}

func (s *MemoryCollection) SetLastUpdatedAt(v string) *MemoryCollection {
	s.LastUpdatedAt = &v
	return s
}

func (s *MemoryCollection) SetLlmConfig(v *LLMConfig) *MemoryCollection {
	s.LlmConfig = v
	return s
}

func (s *MemoryCollection) SetMemoryCollectionId(v string) *MemoryCollection {
	s.MemoryCollectionId = &v
	return s
}

func (s *MemoryCollection) SetMemoryCollectionName(v string) *MemoryCollection {
	s.MemoryCollectionName = &v
	return s
}

func (s *MemoryCollection) SetNetworkConfiguration(v *NetworkConfiguration) *MemoryCollection {
	s.NetworkConfiguration = v
	return s
}

func (s *MemoryCollection) SetVectorStoreConfig(v *VectorStoreConfig) *MemoryCollection {
	s.VectorStoreConfig = v
	return s
}

func (s *MemoryCollection) Validate() error {
	if s.EmbedderConfig != nil {
		if err := s.EmbedderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LlmConfig != nil {
		if err := s.LlmConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.VectorStoreConfig != nil {
		if err := s.VectorStoreConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
