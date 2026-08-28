// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerPatrolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScopeConfig(v *TriggerPatrolRequestScopeConfig) *TriggerPatrolRequest
	GetScopeConfig() *TriggerPatrolRequestScopeConfig
	SetScopeType(v string) *TriggerPatrolRequest
	GetScopeType() *string
}

type TriggerPatrolRequest struct {
	// The inspection scope configuration.
	ScopeConfig *TriggerPatrolRequestScopeConfig `json:"scopeConfig,omitempty" xml:"scopeConfig,omitempty" type:"Struct"`
	// The inspection scope type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s TriggerPatrolRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerPatrolRequest) GoString() string {
	return s.String()
}

func (s *TriggerPatrolRequest) GetScopeConfig() *TriggerPatrolRequestScopeConfig {
	return s.ScopeConfig
}

func (s *TriggerPatrolRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *TriggerPatrolRequest) SetScopeConfig(v *TriggerPatrolRequestScopeConfig) *TriggerPatrolRequest {
	s.ScopeConfig = v
	return s
}

func (s *TriggerPatrolRequest) SetScopeType(v string) *TriggerPatrolRequest {
	s.ScopeType = &v
	return s
}

func (s *TriggerPatrolRequest) Validate() error {
	if s.ScopeConfig != nil {
		if err := s.ScopeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TriggerPatrolRequestScopeConfig struct {
	// The list of deployment IDs. This parameter is valid only when scopeType is set to DEPLOYMENTS.
	DeploymentIds []*string `json:"deploymentIds,omitempty" xml:"deploymentIds,omitempty" type:"Repeated"`
	// The tag mapping. This parameter is valid only when scopeType is set to TAGS. The key is the tag name, and the value is a list of tag values.
	Tags map[string][]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s TriggerPatrolRequestScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s TriggerPatrolRequestScopeConfig) GoString() string {
	return s.String()
}

func (s *TriggerPatrolRequestScopeConfig) GetDeploymentIds() []*string {
	return s.DeploymentIds
}

func (s *TriggerPatrolRequestScopeConfig) GetTags() map[string][]*string {
	return s.Tags
}

func (s *TriggerPatrolRequestScopeConfig) SetDeploymentIds(v []*string) *TriggerPatrolRequestScopeConfig {
	s.DeploymentIds = v
	return s
}

func (s *TriggerPatrolRequestScopeConfig) SetTags(v map[string][]*string) *TriggerPatrolRequestScopeConfig {
	s.Tags = v
	return s
}

func (s *TriggerPatrolRequestScopeConfig) Validate() error {
	return dara.Validate(s)
}
