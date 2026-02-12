// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryCollectionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateMemoryCollectionInput
	GetDescription() *string
	SetEmbedderConfig(v *EmbedderConfig) *UpdateMemoryCollectionInput
	GetEmbedderConfig() *EmbedderConfig
	SetEnableConversationHistory(v bool) *UpdateMemoryCollectionInput
	GetEnableConversationHistory() *bool
	SetEnableConversationState(v bool) *UpdateMemoryCollectionInput
	GetEnableConversationState() *bool
	SetExecutionRoleArn(v string) *UpdateMemoryCollectionInput
	GetExecutionRoleArn() *string
	SetLlmConfig(v *LLMConfig) *UpdateMemoryCollectionInput
	GetLlmConfig() *LLMConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateMemoryCollectionInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetVectorStoreConfig(v *VectorStoreConfig) *UpdateMemoryCollectionInput
	GetVectorStoreConfig() *VectorStoreConfig
}

type UpdateMemoryCollectionInput struct {
	Description               *string               `json:"description,omitempty" xml:"description,omitempty"`
	EmbedderConfig            *EmbedderConfig       `json:"embedderConfig,omitempty" xml:"embedderConfig,omitempty"`
	EnableConversationHistory *bool                 `json:"enableConversationHistory,omitempty" xml:"enableConversationHistory,omitempty"`
	EnableConversationState   *bool                 `json:"enableConversationState,omitempty" xml:"enableConversationState,omitempty"`
	ExecutionRoleArn          *string               `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LlmConfig                 *LLMConfig            `json:"llmConfig,omitempty" xml:"llmConfig,omitempty"`
	NetworkConfiguration      *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	VectorStoreConfig         *VectorStoreConfig    `json:"vectorStoreConfig,omitempty" xml:"vectorStoreConfig,omitempty"`
}

func (s UpdateMemoryCollectionInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryCollectionInput) GoString() string {
	return s.String()
}

func (s *UpdateMemoryCollectionInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateMemoryCollectionInput) GetEmbedderConfig() *EmbedderConfig {
	return s.EmbedderConfig
}

func (s *UpdateMemoryCollectionInput) GetEnableConversationHistory() *bool {
	return s.EnableConversationHistory
}

func (s *UpdateMemoryCollectionInput) GetEnableConversationState() *bool {
	return s.EnableConversationState
}

func (s *UpdateMemoryCollectionInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateMemoryCollectionInput) GetLlmConfig() *LLMConfig {
	return s.LlmConfig
}

func (s *UpdateMemoryCollectionInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateMemoryCollectionInput) GetVectorStoreConfig() *VectorStoreConfig {
	return s.VectorStoreConfig
}

func (s *UpdateMemoryCollectionInput) SetDescription(v string) *UpdateMemoryCollectionInput {
	s.Description = &v
	return s
}

func (s *UpdateMemoryCollectionInput) SetEmbedderConfig(v *EmbedderConfig) *UpdateMemoryCollectionInput {
	s.EmbedderConfig = v
	return s
}

func (s *UpdateMemoryCollectionInput) SetEnableConversationHistory(v bool) *UpdateMemoryCollectionInput {
	s.EnableConversationHistory = &v
	return s
}

func (s *UpdateMemoryCollectionInput) SetEnableConversationState(v bool) *UpdateMemoryCollectionInput {
	s.EnableConversationState = &v
	return s
}

func (s *UpdateMemoryCollectionInput) SetExecutionRoleArn(v string) *UpdateMemoryCollectionInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateMemoryCollectionInput) SetLlmConfig(v *LLMConfig) *UpdateMemoryCollectionInput {
	s.LlmConfig = v
	return s
}

func (s *UpdateMemoryCollectionInput) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateMemoryCollectionInput {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateMemoryCollectionInput) SetVectorStoreConfig(v *VectorStoreConfig) *UpdateMemoryCollectionInput {
	s.VectorStoreConfig = v
	return s
}

func (s *UpdateMemoryCollectionInput) Validate() error {
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
