// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentInstanceMemoryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentInstanceIds(v []*string) *ModifyAgentInstanceMemoryConfigRequest
	GetAgentInstanceIds() []*string
	SetEnabled(v bool) *ModifyAgentInstanceMemoryConfigRequest
	GetEnabled() *bool
}

type ModifyAgentInstanceMemoryConfigRequest struct {
	// The list of agent instance IDs. You can specify 1 to 100 IDs, and the IDs must be unique. Each instance must belong to the current account, support memory management, and have the memory feature configured.
	AgentInstanceIds []*string `json:"AgentInstanceIds,omitempty" xml:"AgentInstanceIds,omitempty" type:"Repeated"`
	// Required. Specifies whether to enable the memory feature. Valid values:
	//
	// - true: Enables the memory feature.
	//
	// - false: Disables the memory feature.
	//
	// Before you enable the feature, ensure that the tenant-level memory store is enabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ModifyAgentInstanceMemoryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentInstanceMemoryConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgentInstanceMemoryConfigRequest) GetAgentInstanceIds() []*string {
	return s.AgentInstanceIds
}

func (s *ModifyAgentInstanceMemoryConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyAgentInstanceMemoryConfigRequest) SetAgentInstanceIds(v []*string) *ModifyAgentInstanceMemoryConfigRequest {
	s.AgentInstanceIds = v
	return s
}

func (s *ModifyAgentInstanceMemoryConfigRequest) SetEnabled(v bool) *ModifyAgentInstanceMemoryConfigRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyAgentInstanceMemoryConfigRequest) Validate() error {
	return dara.Validate(s)
}
