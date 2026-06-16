// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *UpdateContextStoreRequestConfig) *UpdateContextStoreRequest
	GetConfig() *UpdateContextStoreRequestConfig
	SetContextType(v string) *UpdateContextStoreRequest
	GetContextType() *string
	SetDescription(v string) *UpdateContextStoreRequest
	GetDescription() *string
	SetClientToken(v string) *UpdateContextStoreRequest
	GetClientToken() *string
}

type UpdateContextStoreRequest struct {
	Config *UpdateContextStoreRequestConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s UpdateContextStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreRequest) GetConfig() *UpdateContextStoreRequestConfig {
	return s.Config
}

func (s *UpdateContextStoreRequest) GetContextType() *string {
	return s.ContextType
}

func (s *UpdateContextStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateContextStoreRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextStoreRequest) SetConfig(v *UpdateContextStoreRequestConfig) *UpdateContextStoreRequest {
	s.Config = v
	return s
}

func (s *UpdateContextStoreRequest) SetContextType(v string) *UpdateContextStoreRequest {
	s.ContextType = &v
	return s
}

func (s *UpdateContextStoreRequest) SetDescription(v string) *UpdateContextStoreRequest {
	s.Description = &v
	return s
}

func (s *UpdateContextStoreRequest) SetClientToken(v string) *UpdateContextStoreRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextStoreRequest) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextStoreRequestConfig struct {
	// example:
	//
	// {"userId":"user_id","sessionId":"session_id"}
	MetadataField map[string]*string                     `json:"metadataField,omitempty" xml:"metadataField,omitempty"`
	Source        *UpdateContextStoreRequestConfigSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s UpdateContextStoreRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreRequestConfig) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreRequestConfig) GetMetadataField() map[string]*string {
	return s.MetadataField
}

func (s *UpdateContextStoreRequestConfig) GetSource() *UpdateContextStoreRequestConfigSource {
	return s.Source
}

func (s *UpdateContextStoreRequestConfig) SetMetadataField(v map[string]*string) *UpdateContextStoreRequestConfig {
	s.MetadataField = v
	return s
}

func (s *UpdateContextStoreRequestConfig) SetSource(v *UpdateContextStoreRequestConfigSource) *UpdateContextStoreRequestConfig {
	s.Source = v
	return s
}

func (s *UpdateContextStoreRequestConfig) Validate() error {
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextStoreRequestConfigSource struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// 2026-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateContextStoreRequestConfigSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreRequestConfigSource) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreRequestConfigSource) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *UpdateContextStoreRequestConfigSource) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateContextStoreRequestConfigSource) SetAgentSpace(v string) *UpdateContextStoreRequestConfigSource {
	s.AgentSpace = &v
	return s
}

func (s *UpdateContextStoreRequestConfigSource) SetStartTime(v string) *UpdateContextStoreRequestConfigSource {
	s.StartTime = &v
	return s
}

func (s *UpdateContextStoreRequestConfigSource) Validate() error {
	return dara.Validate(s)
}
