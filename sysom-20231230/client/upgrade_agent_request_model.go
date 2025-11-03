// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpgradeAgentRequest
	GetAgentId() *string
	SetAgentVersion(v string) *UpgradeAgentRequest
	GetAgentVersion() *string
	SetInstances(v []*UpgradeAgentRequestInstances) *UpgradeAgentRequest
	GetInstances() []*UpgradeAgentRequestInstances
}

type UpgradeAgentRequest struct {
	// This parameter is required.
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// This parameter is required.
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// This parameter is required.
	Instances []*UpgradeAgentRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s UpgradeAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpgradeAgentRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UpgradeAgentRequest) GetInstances() []*UpgradeAgentRequestInstances {
	return s.Instances
}

func (s *UpgradeAgentRequest) SetAgentId(v string) *UpgradeAgentRequest {
	s.AgentId = &v
	return s
}

func (s *UpgradeAgentRequest) SetAgentVersion(v string) *UpgradeAgentRequest {
	s.AgentVersion = &v
	return s
}

func (s *UpgradeAgentRequest) SetInstances(v []*UpgradeAgentRequestInstances) *UpgradeAgentRequest {
	s.Instances = v
	return s
}

func (s *UpgradeAgentRequest) Validate() error {
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

type UpgradeAgentRequestInstances struct {
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

func (s UpgradeAgentRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentRequestInstances) GoString() string {
	return s.String()
}

func (s *UpgradeAgentRequestInstances) GetInstance() *string {
	return s.Instance
}

func (s *UpgradeAgentRequestInstances) GetRegion() *string {
	return s.Region
}

func (s *UpgradeAgentRequestInstances) SetInstance(v string) *UpgradeAgentRequestInstances {
	s.Instance = &v
	return s
}

func (s *UpgradeAgentRequestInstances) SetRegion(v string) *UpgradeAgentRequestInstances {
	s.Region = &v
	return s
}

func (s *UpgradeAgentRequestInstances) Validate() error {
	return dara.Validate(s)
}
