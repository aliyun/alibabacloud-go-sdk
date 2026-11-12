// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAgentJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ImportAgentJobsRequest
	GetAgentName() *string
	SetClusterId(v string) *ImportAgentJobsRequest
	GetClusterId() *string
	SetMigrateStrategy(v int32) *ImportAgentJobsRequest
	GetMigrateStrategy() *int32
}

type ImportAgentJobsRequest struct {
	// agentname
	//
	// This parameter is required.
	//
	// example:
	//
	// Name
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2027179f-20b1-4e0b-841b-d86f2bc7ebf7
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The strategy for synchronizing scheduled tasks. Valid values:
	//
	// - 1: Migrates both the content and the scheduling time of local scheduled tasks from the agent to the cloud.
	//
	// - 2: Migrates only the scheduling time to the cloud. When the scheduled time arrives, the platform notifies the agent, and the agent executes the task based on the local task content. The task content is always retained locally.
	//
	// example:
	//
	// 1
	MigrateStrategy *int32 `json:"MigrateStrategy,omitempty" xml:"MigrateStrategy,omitempty"`
}

func (s ImportAgentJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportAgentJobsRequest) GoString() string {
	return s.String()
}

func (s *ImportAgentJobsRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ImportAgentJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ImportAgentJobsRequest) GetMigrateStrategy() *int32 {
	return s.MigrateStrategy
}

func (s *ImportAgentJobsRequest) SetAgentName(v string) *ImportAgentJobsRequest {
	s.AgentName = &v
	return s
}

func (s *ImportAgentJobsRequest) SetClusterId(v string) *ImportAgentJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportAgentJobsRequest) SetMigrateStrategy(v int32) *ImportAgentJobsRequest {
	s.MigrateStrategy = &v
	return s
}

func (s *ImportAgentJobsRequest) Validate() error {
	return dara.Validate(s)
}
