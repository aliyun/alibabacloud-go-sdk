// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *InstallAgentRequest
	GetAgentId() *string
	SetAgentVersion(v string) *InstallAgentRequest
	GetAgentVersion() *string
	SetInstallType(v string) *InstallAgentRequest
	GetInstallType() *string
	SetInstances(v []*InstallAgentRequestInstances) *InstallAgentRequest
	GetInstances() []*InstallAgentRequestInstances
}

type InstallAgentRequest struct {
	// This parameter is required.
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// This parameter is required.
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// InstallAndUpgrade
	InstallType *string `json:"install_type,omitempty" xml:"install_type,omitempty"`
	// This parameter is required.
	Instances []*InstallAgentRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s InstallAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *InstallAgentRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *InstallAgentRequest) GetInstallType() *string {
	return s.InstallType
}

func (s *InstallAgentRequest) GetInstances() []*InstallAgentRequestInstances {
	return s.Instances
}

func (s *InstallAgentRequest) SetAgentId(v string) *InstallAgentRequest {
	s.AgentId = &v
	return s
}

func (s *InstallAgentRequest) SetAgentVersion(v string) *InstallAgentRequest {
	s.AgentVersion = &v
	return s
}

func (s *InstallAgentRequest) SetInstallType(v string) *InstallAgentRequest {
	s.InstallType = &v
	return s
}

func (s *InstallAgentRequest) SetInstances(v []*InstallAgentRequestInstances) *InstallAgentRequest {
	s.Instances = v
	return s
}

func (s *InstallAgentRequest) Validate() error {
	return dara.Validate(s)
}

type InstallAgentRequestInstances struct {
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

func (s InstallAgentRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentRequestInstances) GoString() string {
	return s.String()
}

func (s *InstallAgentRequestInstances) GetInstance() *string {
	return s.Instance
}

func (s *InstallAgentRequestInstances) GetRegion() *string {
	return s.Region
}

func (s *InstallAgentRequestInstances) SetInstance(v string) *InstallAgentRequestInstances {
	s.Instance = &v
	return s
}

func (s *InstallAgentRequestInstances) SetRegion(v string) *InstallAgentRequestInstances {
	s.Region = &v
	return s
}

func (s *InstallAgentRequestInstances) Validate() error {
	return dara.Validate(s)
}
