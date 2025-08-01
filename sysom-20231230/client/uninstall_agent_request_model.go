// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UninstallAgentRequest
	GetAgentId() *string
	SetAgentVersion(v string) *UninstallAgentRequest
	GetAgentVersion() *string
	SetInstances(v []*UninstallAgentRequestInstances) *UninstallAgentRequest
	GetInstances() []*UninstallAgentRequestInstances
}

type UninstallAgentRequest struct {
	// This parameter is required.
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// This parameter is required.
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// This parameter is required.
	Instances []*UninstallAgentRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s UninstallAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentRequest) GoString() string {
	return s.String()
}

func (s *UninstallAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UninstallAgentRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UninstallAgentRequest) GetInstances() []*UninstallAgentRequestInstances {
	return s.Instances
}

func (s *UninstallAgentRequest) SetAgentId(v string) *UninstallAgentRequest {
	s.AgentId = &v
	return s
}

func (s *UninstallAgentRequest) SetAgentVersion(v string) *UninstallAgentRequest {
	s.AgentVersion = &v
	return s
}

func (s *UninstallAgentRequest) SetInstances(v []*UninstallAgentRequestInstances) *UninstallAgentRequest {
	s.Instances = v
	return s
}

func (s *UninstallAgentRequest) Validate() error {
	return dara.Validate(s)
}

type UninstallAgentRequestInstances struct {
	// This parameter is required.
	//
	// example:
	//
	// i-wz9b9vucz1iubsz8sjqo
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s UninstallAgentRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentRequestInstances) GoString() string {
	return s.String()
}

func (s *UninstallAgentRequestInstances) GetInstance() *string {
	return s.Instance
}

func (s *UninstallAgentRequestInstances) GetRegion() *string {
	return s.Region
}

func (s *UninstallAgentRequestInstances) SetInstance(v string) *UninstallAgentRequestInstances {
	s.Instance = &v
	return s
}

func (s *UninstallAgentRequestInstances) SetRegion(v string) *UninstallAgentRequestInstances {
	s.Region = &v
	return s
}

func (s *UninstallAgentRequestInstances) Validate() error {
	return dara.Validate(s)
}
