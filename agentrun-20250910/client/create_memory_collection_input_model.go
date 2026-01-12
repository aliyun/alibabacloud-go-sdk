// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryCollectionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateMemoryCollectionInput
	GetDescription() *string
	SetEmbedderConfig(v *EmbedderConfig) *CreateMemoryCollectionInput
	GetEmbedderConfig() *EmbedderConfig
	SetExecutionRoleArn(v string) *CreateMemoryCollectionInput
	GetExecutionRoleArn() *string
	SetLlmConfig(v *LLMConfig) *CreateMemoryCollectionInput
	GetLlmConfig() *LLMConfig
	SetMemoryCollectionName(v string) *CreateMemoryCollectionInput
	GetMemoryCollectionName() *string
	SetNetworkConfiguration(v *NetworkConfiguration) *CreateMemoryCollectionInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetType(v string) *CreateMemoryCollectionInput
	GetType() *string
	SetVectorStoreConfig(v *VectorStoreConfig) *CreateMemoryCollectionInput
	GetVectorStoreConfig() *VectorStoreConfig
}

type CreateMemoryCollectionInput struct {
	Description          *string               `json:"description,omitempty" xml:"description,omitempty"`
	EmbedderConfig       *EmbedderConfig       `json:"embedderConfig,omitempty" xml:"embedderConfig,omitempty"`
	ExecutionRoleArn     *string               `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LlmConfig            *LLMConfig            `json:"llmConfig,omitempty" xml:"llmConfig,omitempty"`
	MemoryCollectionName *string               `json:"memoryCollectionName,omitempty" xml:"memoryCollectionName,omitempty"`
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	Type                 *string               `json:"type,omitempty" xml:"type,omitempty"`
	VectorStoreConfig    *VectorStoreConfig    `json:"vectorStoreConfig,omitempty" xml:"vectorStoreConfig,omitempty"`
}

func (s CreateMemoryCollectionInput) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryCollectionInput) GoString() string {
	return s.String()
}

func (s *CreateMemoryCollectionInput) GetDescription() *string {
	return s.Description
}

func (s *CreateMemoryCollectionInput) GetEmbedderConfig() *EmbedderConfig {
	return s.EmbedderConfig
}

func (s *CreateMemoryCollectionInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateMemoryCollectionInput) GetLlmConfig() *LLMConfig {
	return s.LlmConfig
}

func (s *CreateMemoryCollectionInput) GetMemoryCollectionName() *string {
	return s.MemoryCollectionName
}

func (s *CreateMemoryCollectionInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateMemoryCollectionInput) GetType() *string {
	return s.Type
}

func (s *CreateMemoryCollectionInput) GetVectorStoreConfig() *VectorStoreConfig {
	return s.VectorStoreConfig
}

func (s *CreateMemoryCollectionInput) SetDescription(v string) *CreateMemoryCollectionInput {
	s.Description = &v
	return s
}

func (s *CreateMemoryCollectionInput) SetEmbedderConfig(v *EmbedderConfig) *CreateMemoryCollectionInput {
	s.EmbedderConfig = v
	return s
}

func (s *CreateMemoryCollectionInput) SetExecutionRoleArn(v string) *CreateMemoryCollectionInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateMemoryCollectionInput) SetLlmConfig(v *LLMConfig) *CreateMemoryCollectionInput {
	s.LlmConfig = v
	return s
}

func (s *CreateMemoryCollectionInput) SetMemoryCollectionName(v string) *CreateMemoryCollectionInput {
	s.MemoryCollectionName = &v
	return s
}

func (s *CreateMemoryCollectionInput) SetNetworkConfiguration(v *NetworkConfiguration) *CreateMemoryCollectionInput {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateMemoryCollectionInput) SetType(v string) *CreateMemoryCollectionInput {
	s.Type = &v
	return s
}

func (s *CreateMemoryCollectionInput) SetVectorStoreConfig(v *VectorStoreConfig) *CreateMemoryCollectionInput {
	s.VectorStoreConfig = v
	return s
}

func (s *CreateMemoryCollectionInput) Validate() error {
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
