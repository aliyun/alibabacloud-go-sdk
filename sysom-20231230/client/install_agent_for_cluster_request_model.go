// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentForClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *InstallAgentForClusterRequest
	GetAgentId() *string
	SetAgentVersion(v string) *InstallAgentForClusterRequest
	GetAgentVersion() *string
	SetClusterId(v string) *InstallAgentForClusterRequest
	GetClusterId() *string
	SetConfigId(v string) *InstallAgentForClusterRequest
	GetConfigId() *string
	SetGrayscaleConfig(v string) *InstallAgentForClusterRequest
	GetGrayscaleConfig() *string
}

type InstallAgentForClusterRequest struct {
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// example:
	//
	// 3.4.0-1
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// example:
	//
	// c9d7f3fc3d42942afbcb65c1100ffb19d
	ClusterId       *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	ConfigId        *string `json:"config_id,omitempty" xml:"config_id,omitempty"`
	GrayscaleConfig *string `json:"grayscale_config,omitempty" xml:"grayscale_config,omitempty"`
}

func (s InstallAgentForClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentForClusterRequest) GoString() string {
	return s.String()
}

func (s *InstallAgentForClusterRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *InstallAgentForClusterRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *InstallAgentForClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallAgentForClusterRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *InstallAgentForClusterRequest) GetGrayscaleConfig() *string {
	return s.GrayscaleConfig
}

func (s *InstallAgentForClusterRequest) SetAgentId(v string) *InstallAgentForClusterRequest {
	s.AgentId = &v
	return s
}

func (s *InstallAgentForClusterRequest) SetAgentVersion(v string) *InstallAgentForClusterRequest {
	s.AgentVersion = &v
	return s
}

func (s *InstallAgentForClusterRequest) SetClusterId(v string) *InstallAgentForClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallAgentForClusterRequest) SetConfigId(v string) *InstallAgentForClusterRequest {
	s.ConfigId = &v
	return s
}

func (s *InstallAgentForClusterRequest) SetGrayscaleConfig(v string) *InstallAgentForClusterRequest {
	s.GrayscaleConfig = &v
	return s
}

func (s *InstallAgentForClusterRequest) Validate() error {
	return dara.Validate(s)
}
