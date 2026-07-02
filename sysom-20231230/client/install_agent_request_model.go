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
	// The ID of the component to install.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// The version of the component to install.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.4.0-1
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// The installation type. Valid values:
	//
	// - InstallAndUpgrade: installs the component if it does not exist, or updates it if it exists.
	//
	// - OnlyInstallNotHasAgent: installs the component if it does not exist, or takes no action if it exists.
	//
	// - OnlyUpgradeHasAgent: takes no action if the component does not exist, or updates it if it exists.
	//
	// - OnlyInstallWithoutStart: installs the component only without starting the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstallAndUpgrade
	InstallType *string `json:"install_type,omitempty" xml:"install_type,omitempty"`
	// The list of instances on which to install the component.
	//
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
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallAgentRequestInstances struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-wz9b9vucz1iubsz8sjqo
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The region ID.
	//
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
