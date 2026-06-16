// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *CreateContextStoreRequestConfig) *CreateContextStoreRequest
	GetConfig() *CreateContextStoreRequestConfig
	SetContextStoreName(v string) *CreateContextStoreRequest
	GetContextStoreName() *string
	SetContextType(v string) *CreateContextStoreRequest
	GetContextType() *string
	SetDescription(v string) *CreateContextStoreRequest
	GetDescription() *string
	SetClientToken(v string) *CreateContextStoreRequest
	GetClientToken() *string
}

type CreateContextStoreRequest struct {
	Config *CreateContextStoreRequestConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// my-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// experience
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// example:
	//
	// 我的上下文库
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateContextStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateContextStoreRequest) GetConfig() *CreateContextStoreRequestConfig {
	return s.Config
}

func (s *CreateContextStoreRequest) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *CreateContextStoreRequest) GetContextType() *string {
	return s.ContextType
}

func (s *CreateContextStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateContextStoreRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateContextStoreRequest) SetConfig(v *CreateContextStoreRequestConfig) *CreateContextStoreRequest {
	s.Config = v
	return s
}

func (s *CreateContextStoreRequest) SetContextStoreName(v string) *CreateContextStoreRequest {
	s.ContextStoreName = &v
	return s
}

func (s *CreateContextStoreRequest) SetContextType(v string) *CreateContextStoreRequest {
	s.ContextType = &v
	return s
}

func (s *CreateContextStoreRequest) SetDescription(v string) *CreateContextStoreRequest {
	s.Description = &v
	return s
}

func (s *CreateContextStoreRequest) SetClientToken(v string) *CreateContextStoreRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateContextStoreRequest) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateContextStoreRequestConfig struct {
	// example:
	//
	// {"userId":"user_id","sessionId":"session_id"}
	MetadataField map[string]*string `json:"metadataField,omitempty" xml:"metadataField,omitempty"`
	// example:
	//
	// 1d
	MiningInterval *string `json:"miningInterval,omitempty" xml:"miningInterval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["order-service","payment-service"]
	ServiceNames []*string                              `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
	Source       *CreateContextStoreRequestConfigSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s CreateContextStoreRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreRequestConfig) GoString() string {
	return s.String()
}

func (s *CreateContextStoreRequestConfig) GetMetadataField() map[string]*string {
	return s.MetadataField
}

func (s *CreateContextStoreRequestConfig) GetMiningInterval() *string {
	return s.MiningInterval
}

func (s *CreateContextStoreRequestConfig) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *CreateContextStoreRequestConfig) GetSource() *CreateContextStoreRequestConfigSource {
	return s.Source
}

func (s *CreateContextStoreRequestConfig) SetMetadataField(v map[string]*string) *CreateContextStoreRequestConfig {
	s.MetadataField = v
	return s
}

func (s *CreateContextStoreRequestConfig) SetMiningInterval(v string) *CreateContextStoreRequestConfig {
	s.MiningInterval = &v
	return s
}

func (s *CreateContextStoreRequestConfig) SetServiceNames(v []*string) *CreateContextStoreRequestConfig {
	s.ServiceNames = v
	return s
}

func (s *CreateContextStoreRequestConfig) SetSource(v *CreateContextStoreRequestConfigSource) *CreateContextStoreRequestConfig {
	s.Source = v
	return s
}

func (s *CreateContextStoreRequestConfig) Validate() error {
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateContextStoreRequestConfigSource struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CreateContextStoreRequestConfigSource) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreRequestConfigSource) GoString() string {
	return s.String()
}

func (s *CreateContextStoreRequestConfigSource) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *CreateContextStoreRequestConfigSource) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateContextStoreRequestConfigSource) SetAgentSpace(v string) *CreateContextStoreRequestConfigSource {
	s.AgentSpace = &v
	return s
}

func (s *CreateContextStoreRequestConfigSource) SetStartTime(v string) *CreateContextStoreRequestConfigSource {
	s.StartTime = &v
	return s
}

func (s *CreateContextStoreRequestConfigSource) Validate() error {
	return dara.Validate(s)
}
