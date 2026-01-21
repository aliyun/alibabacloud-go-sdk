// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiPlatformConfig(v *CreateAgentPlatformRequestAiPlatformConfig) *CreateAgentPlatformRequest
	GetAiPlatformConfig() *CreateAgentPlatformRequestAiPlatformConfig
	SetDBClusterId(v string) *CreateAgentPlatformRequest
	GetDBClusterId() *string
	SetName(v string) *CreateAgentPlatformRequest
	GetName() *string
	SetRegionId(v string) *CreateAgentPlatformRequest
	GetRegionId() *string
}

type CreateAgentPlatformRequest struct {
	// This parameter is required.
	AiPlatformConfig *CreateAgentPlatformRequestAiPlatformConfig `json:"AiPlatformConfig,omitempty" xml:"AiPlatformConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testplatform
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAgentPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentPlatformRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentPlatformRequest) GetAiPlatformConfig() *CreateAgentPlatformRequestAiPlatformConfig {
	return s.AiPlatformConfig
}

func (s *CreateAgentPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgentPlatformRequest) GetName() *string {
	return s.Name
}

func (s *CreateAgentPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgentPlatformRequest) SetAiPlatformConfig(v *CreateAgentPlatformRequestAiPlatformConfig) *CreateAgentPlatformRequest {
	s.AiPlatformConfig = v
	return s
}

func (s *CreateAgentPlatformRequest) SetDBClusterId(v string) *CreateAgentPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgentPlatformRequest) SetName(v string) *CreateAgentPlatformRequest {
	s.Name = &v
	return s
}

func (s *CreateAgentPlatformRequest) SetRegionId(v string) *CreateAgentPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAgentPlatformRequest) Validate() error {
	if s.AiPlatformConfig != nil {
		if err := s.AiPlatformConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentPlatformRequestAiPlatformConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 3760d3**************************
	ServeApiKey *string `json:"ServeApiKey,omitempty" xml:"ServeApiKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://111.xx.xx.xx:8100/inferenceservice/emb
	ServeEmbeddingEndpoint *string `json:"ServeEmbeddingEndpoint,omitempty" xml:"ServeEmbeddingEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Qwen3-Embedding-8B
	ServeEmbeddingModelName *string `json:"ServeEmbeddingModelName,omitempty" xml:"ServeEmbeddingModelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://111.xx.xx.xx:8100/inferenceservice/base
	ServeEndpoint *string `json:"ServeEndpoint,omitempty" xml:"ServeEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Qwen3-235B-A22B-Instruct-2507
	ServeModelName *string `json:"ServeModelName,omitempty" xml:"ServeModelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// large
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s CreateAgentPlatformRequestAiPlatformConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentPlatformRequestAiPlatformConfig) GoString() string {
	return s.String()
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) GetServeApiKey() *string {
	return s.ServeApiKey
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) GetServeEmbeddingEndpoint() *string {
	return s.ServeEmbeddingEndpoint
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) GetServeEmbeddingModelName() *string {
	return s.ServeEmbeddingModelName
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) GetServeEndpoint() *string {
	return s.ServeEndpoint
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) GetServeModelName() *string {
	return s.ServeModelName
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) GetSpecName() *string {
	return s.SpecName
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) SetServeApiKey(v string) *CreateAgentPlatformRequestAiPlatformConfig {
	s.ServeApiKey = &v
	return s
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) SetServeEmbeddingEndpoint(v string) *CreateAgentPlatformRequestAiPlatformConfig {
	s.ServeEmbeddingEndpoint = &v
	return s
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) SetServeEmbeddingModelName(v string) *CreateAgentPlatformRequestAiPlatformConfig {
	s.ServeEmbeddingModelName = &v
	return s
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) SetServeEndpoint(v string) *CreateAgentPlatformRequestAiPlatformConfig {
	s.ServeEndpoint = &v
	return s
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) SetServeModelName(v string) *CreateAgentPlatformRequestAiPlatformConfig {
	s.ServeModelName = &v
	return s
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) SetSpecName(v string) *CreateAgentPlatformRequestAiPlatformConfig {
	s.SpecName = &v
	return s
}

func (s *CreateAgentPlatformRequestAiPlatformConfig) Validate() error {
	return dara.Validate(s)
}
