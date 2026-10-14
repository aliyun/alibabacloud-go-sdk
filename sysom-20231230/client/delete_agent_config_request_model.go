// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DeleteAgentConfigRequest
	GetAgentId() *string
	SetConfigId(v string) *DeleteAgentConfigRequest
	GetConfigId() *string
}

type DeleteAgentConfigRequest struct {
	// The ID of the component associated with the configuration to delete.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The ID of the component configuration to delete.
	//
	// example:
	//
	// 4d5ecf02-3d01-11f0-8c90-00163e1029af
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
}

func (s DeleteAgentConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentConfigRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteAgentConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteAgentConfigRequest) SetAgentId(v string) *DeleteAgentConfigRequest {
	s.AgentId = &v
	return s
}

func (s *DeleteAgentConfigRequest) SetConfigId(v string) *DeleteAgentConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteAgentConfigRequest) Validate() error {
	return dara.Validate(s)
}
