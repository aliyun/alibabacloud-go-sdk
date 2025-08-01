// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentForClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpgradeAgentForClusterRequest
	GetAgentId() *string
	SetAgentVersion(v string) *UpgradeAgentForClusterRequest
	GetAgentVersion() *string
	SetClusterId(v string) *UpgradeAgentForClusterRequest
	GetClusterId() *string
}

type UpgradeAgentForClusterRequest struct {
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
	// c1c187fd513cb41a19876bac0e6b05212
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
}

func (s UpgradeAgentForClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentForClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpgradeAgentForClusterRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UpgradeAgentForClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeAgentForClusterRequest) SetAgentId(v string) *UpgradeAgentForClusterRequest {
	s.AgentId = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) SetAgentVersion(v string) *UpgradeAgentForClusterRequest {
	s.AgentVersion = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) SetClusterId(v string) *UpgradeAgentForClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) Validate() error {
	return dara.Validate(s)
}
