// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiPlatformConfig(v *ModifyAgentPlatformRequestAiPlatformConfig) *ModifyAgentPlatformRequest
	GetAiPlatformConfig() *ModifyAgentPlatformRequestAiPlatformConfig
	SetDBClusterId(v string) *ModifyAgentPlatformRequest
	GetDBClusterId() *string
	SetName(v string) *ModifyAgentPlatformRequest
	GetName() *string
	SetRegionId(v string) *ModifyAgentPlatformRequest
	GetRegionId() *string
}

type ModifyAgentPlatformRequest struct {
	// The parameters required for upgrading or downgrading the metric platform.
	AiPlatformConfig *ModifyAgentPlatformRequestAiPlatformConfig `json:"AiPlatformConfig,omitempty" xml:"AiPlatformConfig,omitempty" type:"Struct"`
	// The instance cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the metric platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_platform
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// > You can call the DescribeRegions operation to query the region ID of a specified Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAgentPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentPlatformRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgentPlatformRequest) GetAiPlatformConfig() *ModifyAgentPlatformRequestAiPlatformConfig {
	return s.AiPlatformConfig
}

func (s *ModifyAgentPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAgentPlatformRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAgentPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAgentPlatformRequest) SetAiPlatformConfig(v *ModifyAgentPlatformRequestAiPlatformConfig) *ModifyAgentPlatformRequest {
	s.AiPlatformConfig = v
	return s
}

func (s *ModifyAgentPlatformRequest) SetDBClusterId(v string) *ModifyAgentPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAgentPlatformRequest) SetName(v string) *ModifyAgentPlatformRequest {
	s.Name = &v
	return s
}

func (s *ModifyAgentPlatformRequest) SetRegionId(v string) *ModifyAgentPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAgentPlatformRequest) Validate() error {
	if s.AiPlatformConfig != nil {
		if err := s.AiPlatformConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAgentPlatformRequestAiPlatformConfig struct {
	// The API key of the model for the inference service that the metric analysis platform depends on.
	//
	// example:
	//
	// 3760d3**************************
	ServeApiKey *string `json:"ServeApiKey,omitempty" xml:"ServeApiKey,omitempty"`
	// The endpoint of the embedding model for the inference service that the metric analysis platform depends on.
	//
	// example:
	//
	// http://111.xx.xx.xx:8100/inferenceservice/emb
	ServeEmbeddingEndpoint *string `json:"ServeEmbeddingEndpoint,omitempty" xml:"ServeEmbeddingEndpoint,omitempty"`
	// The name of the embedding model for the inference service that the metric analysis platform depends on.
	//
	// example:
	//
	// Qwen3-Embedding-8B
	ServeEmbeddingModelName *string `json:"ServeEmbeddingModelName,omitempty" xml:"ServeEmbeddingModelName,omitempty"`
	// The endpoint of the base model for the inference service that the metric analysis platform depends on.
	//
	// example:
	//
	// http://111.xx.xx.xx:8100/inferenceservice/base
	ServeEndpoint *string `json:"ServeEndpoint,omitempty" xml:"ServeEndpoint,omitempty"`
	// The name of the base model for the inference service that the metric analysis platform depends on.
	//
	// example:
	//
	// Qwen3-235B-A22B-Instruct-2507
	ServeModelName *string `json:"ServeModelName,omitempty" xml:"ServeModelName,omitempty"`
	// The specification of the metric analysis platform.
	//
	// example:
	//
	// large
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s ModifyAgentPlatformRequestAiPlatformConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentPlatformRequestAiPlatformConfig) GoString() string {
	return s.String()
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) GetServeApiKey() *string {
	return s.ServeApiKey
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) GetServeEmbeddingEndpoint() *string {
	return s.ServeEmbeddingEndpoint
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) GetServeEmbeddingModelName() *string {
	return s.ServeEmbeddingModelName
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) GetServeEndpoint() *string {
	return s.ServeEndpoint
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) GetServeModelName() *string {
	return s.ServeModelName
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) GetSpecName() *string {
	return s.SpecName
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) SetServeApiKey(v string) *ModifyAgentPlatformRequestAiPlatformConfig {
	s.ServeApiKey = &v
	return s
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) SetServeEmbeddingEndpoint(v string) *ModifyAgentPlatformRequestAiPlatformConfig {
	s.ServeEmbeddingEndpoint = &v
	return s
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) SetServeEmbeddingModelName(v string) *ModifyAgentPlatformRequestAiPlatformConfig {
	s.ServeEmbeddingModelName = &v
	return s
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) SetServeEndpoint(v string) *ModifyAgentPlatformRequestAiPlatformConfig {
	s.ServeEndpoint = &v
	return s
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) SetServeModelName(v string) *ModifyAgentPlatformRequestAiPlatformConfig {
	s.ServeModelName = &v
	return s
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) SetSpecName(v string) *ModifyAgentPlatformRequestAiPlatformConfig {
	s.SpecName = &v
	return s
}

func (s *ModifyAgentPlatformRequestAiPlatformConfig) Validate() error {
	return dara.Validate(s)
}
