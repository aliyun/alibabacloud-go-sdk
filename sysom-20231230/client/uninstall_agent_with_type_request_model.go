// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentWithTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UninstallAgentWithTypeRequest
	GetAgentId() *string
	SetAgentVersion(v string) *UninstallAgentWithTypeRequest
	GetAgentVersion() *string
	SetInstanceType(v string) *UninstallAgentWithTypeRequest
	GetInstanceType() *string
	SetInstances(v []*UninstallAgentWithTypeRequestInstances) *UninstallAgentWithTypeRequest
	GetInstances() []*UninstallAgentWithTypeRequestInstances
}

type UninstallAgentWithTypeRequest struct {
	// The ID of the component to uninstall.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The version of the component to uninstall.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.4.0-1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The list of instances from which to uninstall the component.
	//
	// This parameter is required.
	Instances []*UninstallAgentWithTypeRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s UninstallAgentWithTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentWithTypeRequest) GoString() string {
	return s.String()
}

func (s *UninstallAgentWithTypeRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UninstallAgentWithTypeRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UninstallAgentWithTypeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UninstallAgentWithTypeRequest) GetInstances() []*UninstallAgentWithTypeRequestInstances {
	return s.Instances
}

func (s *UninstallAgentWithTypeRequest) SetAgentId(v string) *UninstallAgentWithTypeRequest {
	s.AgentId = &v
	return s
}

func (s *UninstallAgentWithTypeRequest) SetAgentVersion(v string) *UninstallAgentWithTypeRequest {
	s.AgentVersion = &v
	return s
}

func (s *UninstallAgentWithTypeRequest) SetInstanceType(v string) *UninstallAgentWithTypeRequest {
	s.InstanceType = &v
	return s
}

func (s *UninstallAgentWithTypeRequest) SetInstances(v []*UninstallAgentWithTypeRequestInstances) *UninstallAgentWithTypeRequest {
	s.Instances = v
	return s
}

func (s *UninstallAgentWithTypeRequest) Validate() error {
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

type UninstallAgentWithTypeRequestInstances struct {
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

func (s UninstallAgentWithTypeRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentWithTypeRequestInstances) GoString() string {
	return s.String()
}

func (s *UninstallAgentWithTypeRequestInstances) GetInstance() *string {
	return s.Instance
}

func (s *UninstallAgentWithTypeRequestInstances) GetRegion() *string {
	return s.Region
}

func (s *UninstallAgentWithTypeRequestInstances) SetInstance(v string) *UninstallAgentWithTypeRequestInstances {
	s.Instance = &v
	return s
}

func (s *UninstallAgentWithTypeRequestInstances) SetRegion(v string) *UninstallAgentWithTypeRequestInstances {
	s.Region = &v
	return s
}

func (s *UninstallAgentWithTypeRequestInstances) Validate() error {
	return dara.Validate(s)
}
