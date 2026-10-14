// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAgentConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdateInstanceAgentConfigRequest
	GetAgentId() *string
	SetConfigId(v string) *UpdateInstanceAgentConfigRequest
	GetConfigId() *string
	SetInstances(v []*UpdateInstanceAgentConfigRequestInstances) *UpdateInstanceAgentConfigRequest
	GetInstances() []*UpdateInstanceAgentConfigRequestInstances
}

type UpdateInstanceAgentConfigRequest struct {
	// The ID of the component whose configuration is to be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The ID of the component configuration to be applied to the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// f0078fbb-4213-11f0-a19b-00163e4ae208
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
	// The instance information.
	Instances []*UpdateInstanceAgentConfigRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s UpdateInstanceAgentConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAgentConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAgentConfigRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateInstanceAgentConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateInstanceAgentConfigRequest) GetInstances() []*UpdateInstanceAgentConfigRequestInstances {
	return s.Instances
}

func (s *UpdateInstanceAgentConfigRequest) SetAgentId(v string) *UpdateInstanceAgentConfigRequest {
	s.AgentId = &v
	return s
}

func (s *UpdateInstanceAgentConfigRequest) SetConfigId(v string) *UpdateInstanceAgentConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateInstanceAgentConfigRequest) SetInstances(v []*UpdateInstanceAgentConfigRequestInstances) *UpdateInstanceAgentConfigRequest {
	s.Instances = v
	return s
}

func (s *UpdateInstanceAgentConfigRequest) Validate() error {
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

type UpdateInstanceAgentConfigRequestInstances struct {
	// The instance ID.
	//
	// example:
	//
	// i-wz9b9vucz1iubsz8sjqo
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s UpdateInstanceAgentConfigRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAgentConfigRequestInstances) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAgentConfigRequestInstances) GetInstance() *string {
	return s.Instance
}

func (s *UpdateInstanceAgentConfigRequestInstances) GetRegion() *string {
	return s.Region
}

func (s *UpdateInstanceAgentConfigRequestInstances) SetInstance(v string) *UpdateInstanceAgentConfigRequestInstances {
	s.Instance = &v
	return s
}

func (s *UpdateInstanceAgentConfigRequestInstances) SetRegion(v string) *UpdateInstanceAgentConfigRequestInstances {
	s.Region = &v
	return s
}

func (s *UpdateInstanceAgentConfigRequestInstances) Validate() error {
	return dara.Validate(s)
}
