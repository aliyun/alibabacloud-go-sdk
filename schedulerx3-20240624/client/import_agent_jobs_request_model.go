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
	// This parameter is required.
	//
	// example:
	//
	// 名称
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2027179f-20b1-4e0b-841b-d86f2bc7ebf7
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
