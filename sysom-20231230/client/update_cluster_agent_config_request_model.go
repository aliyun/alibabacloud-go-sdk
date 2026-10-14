// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAgentConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdateClusterAgentConfigRequest
	GetAgentId() *string
	SetClusterId(v string) *UpdateClusterAgentConfigRequest
	GetClusterId() *string
	SetConfigId(v string) *UpdateClusterAgentConfigRequest
	GetConfigId() *string
}

type UpdateClusterAgentConfigRequest struct {
	// The ID of the component whose configuration is to be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The ID of the cluster. You can call the ListClusters operation to obtain the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c8b90bc1b2ef0436f81213c032b75c5f3
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The ID of the component configuration to install on the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// f0078fbb-4213-11f0-a19b-00163e4ae208
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
}

func (s UpdateClusterAgentConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAgentConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterAgentConfigRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateClusterAgentConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterAgentConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateClusterAgentConfigRequest) SetAgentId(v string) *UpdateClusterAgentConfigRequest {
	s.AgentId = &v
	return s
}

func (s *UpdateClusterAgentConfigRequest) SetClusterId(v string) *UpdateClusterAgentConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterAgentConfigRequest) SetConfigId(v string) *UpdateClusterAgentConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateClusterAgentConfigRequest) Validate() error {
	return dara.Validate(s)
}
