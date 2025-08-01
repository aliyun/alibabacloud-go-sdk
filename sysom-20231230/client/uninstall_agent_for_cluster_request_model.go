// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentForClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UninstallAgentForClusterRequest
	GetAgentId() *string
	SetAgentVersion(v string) *UninstallAgentForClusterRequest
	GetAgentVersion() *string
	SetClusterId(v string) *UninstallAgentForClusterRequest
	GetClusterId() *string
}

type UninstallAgentForClusterRequest struct {
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
	// c822f83bb45994ddbac9326b4c2f04f35
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
}

func (s UninstallAgentForClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentForClusterRequest) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UninstallAgentForClusterRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UninstallAgentForClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UninstallAgentForClusterRequest) SetAgentId(v string) *UninstallAgentForClusterRequest {
	s.AgentId = &v
	return s
}

func (s *UninstallAgentForClusterRequest) SetAgentVersion(v string) *UninstallAgentForClusterRequest {
	s.AgentVersion = &v
	return s
}

func (s *UninstallAgentForClusterRequest) SetClusterId(v string) *UninstallAgentForClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallAgentForClusterRequest) Validate() error {
	return dara.Validate(s)
}
