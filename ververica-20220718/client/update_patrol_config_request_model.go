// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePatrolConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCron(v string) *UpdatePatrolConfigRequest
	GetCron() *string
	SetEnabled(v bool) *UpdatePatrolConfigRequest
	GetEnabled() *bool
	SetScopeConfig(v *UpdatePatrolConfigRequestScopeConfig) *UpdatePatrolConfigRequest
	GetScopeConfig() *UpdatePatrolConfigRequestScopeConfig
	SetScopeType(v string) *UpdatePatrolConfigRequest
	GetScopeType() *string
	SetTimezone(v string) *UpdatePatrolConfigRequest
	GetTimezone() *string
}

type UpdatePatrolConfigRequest struct {
	// The cron expression that defines the inspection scheduling time.
	//
	// example:
	//
	// 0 2 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// Specifies whether to enable the inspection.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The inspection scope configuration.
	ScopeConfig *UpdatePatrolConfigRequestScopeConfig `json:"scopeConfig,omitempty" xml:"scopeConfig,omitempty" type:"Struct"`
	// The inspection scope type.
	//
	// example:
	//
	// ALL
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s UpdatePatrolConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatrolConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdatePatrolConfigRequest) GetCron() *string {
	return s.Cron
}

func (s *UpdatePatrolConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePatrolConfigRequest) GetScopeConfig() *UpdatePatrolConfigRequestScopeConfig {
	return s.ScopeConfig
}

func (s *UpdatePatrolConfigRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *UpdatePatrolConfigRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdatePatrolConfigRequest) SetCron(v string) *UpdatePatrolConfigRequest {
	s.Cron = &v
	return s
}

func (s *UpdatePatrolConfigRequest) SetEnabled(v bool) *UpdatePatrolConfigRequest {
	s.Enabled = &v
	return s
}

func (s *UpdatePatrolConfigRequest) SetScopeConfig(v *UpdatePatrolConfigRequestScopeConfig) *UpdatePatrolConfigRequest {
	s.ScopeConfig = v
	return s
}

func (s *UpdatePatrolConfigRequest) SetScopeType(v string) *UpdatePatrolConfigRequest {
	s.ScopeType = &v
	return s
}

func (s *UpdatePatrolConfigRequest) SetTimezone(v string) *UpdatePatrolConfigRequest {
	s.Timezone = &v
	return s
}

func (s *UpdatePatrolConfigRequest) Validate() error {
	if s.ScopeConfig != nil {
		if err := s.ScopeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePatrolConfigRequestScopeConfig struct {
	// The list of deployment IDs. This parameter is valid only when scopeType is set to DEPLOYMENTS.
	DeploymentIds []*string `json:"deploymentIds,omitempty" xml:"deploymentIds,omitempty" type:"Repeated"`
	// The tag mapping. This parameter is valid only when scopeType is set to TAGS. The key is the tag name, and the value is a list of tag values.
	Tags map[string][]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s UpdatePatrolConfigRequestScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatrolConfigRequestScopeConfig) GoString() string {
	return s.String()
}

func (s *UpdatePatrolConfigRequestScopeConfig) GetDeploymentIds() []*string {
	return s.DeploymentIds
}

func (s *UpdatePatrolConfigRequestScopeConfig) GetTags() map[string][]*string {
	return s.Tags
}

func (s *UpdatePatrolConfigRequestScopeConfig) SetDeploymentIds(v []*string) *UpdatePatrolConfigRequestScopeConfig {
	s.DeploymentIds = v
	return s
}

func (s *UpdatePatrolConfigRequestScopeConfig) SetTags(v map[string][]*string) *UpdatePatrolConfigRequestScopeConfig {
	s.Tags = v
	return s
}

func (s *UpdatePatrolConfigRequestScopeConfig) Validate() error {
	return dara.Validate(s)
}
