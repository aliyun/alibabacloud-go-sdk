// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPolarClawSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UninstallPolarClawSkillsRequest
	GetAgentId() *string
	SetApplicationId(v string) *UninstallPolarClawSkillsRequest
	GetApplicationId() *string
	SetRegionId(v string) *UninstallPolarClawSkillsRequest
	GetRegionId() *string
	SetScope(v string) *UninstallPolarClawSkillsRequest
	GetScope() *string
	SetUninstallSkillNames(v string) *UninstallPolarClawSkillsRequest
	GetUninstallSkillNames() *string
}

type UninstallPolarClawSkillsRequest struct {
	// The task-level target agent (workspace). If not specified, the default agent is used.
	//
	// example:
	//
	// test1
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query region details.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task-level scope. Valid values: workspace (default) and global. This parameter is mutually exclusive with AgentId.
	//
	// example:
	//
	// global
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The names of the skills to uninstall.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["polarclawhub","github"]
	UninstallSkillNames *string `json:"UninstallSkillNames,omitempty" xml:"UninstallSkillNames,omitempty"`
}

func (s UninstallPolarClawSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallPolarClawSkillsRequest) GoString() string {
	return s.String()
}

func (s *UninstallPolarClawSkillsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UninstallPolarClawSkillsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UninstallPolarClawSkillsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UninstallPolarClawSkillsRequest) GetScope() *string {
	return s.Scope
}

func (s *UninstallPolarClawSkillsRequest) GetUninstallSkillNames() *string {
	return s.UninstallSkillNames
}

func (s *UninstallPolarClawSkillsRequest) SetAgentId(v string) *UninstallPolarClawSkillsRequest {
	s.AgentId = &v
	return s
}

func (s *UninstallPolarClawSkillsRequest) SetApplicationId(v string) *UninstallPolarClawSkillsRequest {
	s.ApplicationId = &v
	return s
}

func (s *UninstallPolarClawSkillsRequest) SetRegionId(v string) *UninstallPolarClawSkillsRequest {
	s.RegionId = &v
	return s
}

func (s *UninstallPolarClawSkillsRequest) SetScope(v string) *UninstallPolarClawSkillsRequest {
	s.Scope = &v
	return s
}

func (s *UninstallPolarClawSkillsRequest) SetUninstallSkillNames(v string) *UninstallPolarClawSkillsRequest {
	s.UninstallSkillNames = &v
	return s
}

func (s *UninstallPolarClawSkillsRequest) Validate() error {
	return dara.Validate(s)
}
