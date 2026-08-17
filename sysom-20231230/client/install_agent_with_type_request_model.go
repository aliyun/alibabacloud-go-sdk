// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentWithTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*InstallAgentWithTypeRequestTag) *InstallAgentWithTypeRequest
	GetTag() []*InstallAgentWithTypeRequestTag
	SetAgentId(v string) *InstallAgentWithTypeRequest
	GetAgentId() *string
	SetAgentVersion(v string) *InstallAgentWithTypeRequest
	GetAgentVersion() *string
	SetConfigId(v string) *InstallAgentWithTypeRequest
	GetConfigId() *string
	SetInstanceType(v string) *InstallAgentWithTypeRequest
	GetInstanceType() *string
	SetInstances(v []*InstallAgentWithTypeRequestInstances) *InstallAgentWithTypeRequest
	GetInstances() []*InstallAgentWithTypeRequestInstances
}

type InstallAgentWithTypeRequest struct {
	Tag []*InstallAgentWithTypeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the component to install.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The version of the component to install.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.4.0-1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// The configuration ID of the component to install.
	//
	// example:
	//
	// f0078fbb-4213-11f0-a19b-00163e4ae208
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
	// The instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The list of instances on which to install the component.
	//
	// This parameter is required.
	Instances []*InstallAgentWithTypeRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s InstallAgentWithTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentWithTypeRequest) GoString() string {
	return s.String()
}

func (s *InstallAgentWithTypeRequest) GetTag() []*InstallAgentWithTypeRequestTag {
	return s.Tag
}

func (s *InstallAgentWithTypeRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *InstallAgentWithTypeRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *InstallAgentWithTypeRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *InstallAgentWithTypeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *InstallAgentWithTypeRequest) GetInstances() []*InstallAgentWithTypeRequestInstances {
	return s.Instances
}

func (s *InstallAgentWithTypeRequest) SetTag(v []*InstallAgentWithTypeRequestTag) *InstallAgentWithTypeRequest {
	s.Tag = v
	return s
}

func (s *InstallAgentWithTypeRequest) SetAgentId(v string) *InstallAgentWithTypeRequest {
	s.AgentId = &v
	return s
}

func (s *InstallAgentWithTypeRequest) SetAgentVersion(v string) *InstallAgentWithTypeRequest {
	s.AgentVersion = &v
	return s
}

func (s *InstallAgentWithTypeRequest) SetConfigId(v string) *InstallAgentWithTypeRequest {
	s.ConfigId = &v
	return s
}

func (s *InstallAgentWithTypeRequest) SetInstanceType(v string) *InstallAgentWithTypeRequest {
	s.InstanceType = &v
	return s
}

func (s *InstallAgentWithTypeRequest) SetInstances(v []*InstallAgentWithTypeRequestInstances) *InstallAgentWithTypeRequest {
	s.Instances = v
	return s
}

func (s *InstallAgentWithTypeRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type InstallAgentWithTypeRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InstallAgentWithTypeRequestTag) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentWithTypeRequestTag) GoString() string {
	return s.String()
}

func (s *InstallAgentWithTypeRequestTag) GetKey() *string {
	return s.Key
}

func (s *InstallAgentWithTypeRequestTag) GetValue() *string {
	return s.Value
}

func (s *InstallAgentWithTypeRequestTag) SetKey(v string) *InstallAgentWithTypeRequestTag {
	s.Key = &v
	return s
}

func (s *InstallAgentWithTypeRequestTag) SetValue(v string) *InstallAgentWithTypeRequestTag {
	s.Value = &v
	return s
}

func (s *InstallAgentWithTypeRequestTag) Validate() error {
	return dara.Validate(s)
}

type InstallAgentWithTypeRequestInstances struct {
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

func (s InstallAgentWithTypeRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentWithTypeRequestInstances) GoString() string {
	return s.String()
}

func (s *InstallAgentWithTypeRequestInstances) GetInstance() *string {
	return s.Instance
}

func (s *InstallAgentWithTypeRequestInstances) GetRegion() *string {
	return s.Region
}

func (s *InstallAgentWithTypeRequestInstances) SetInstance(v string) *InstallAgentWithTypeRequestInstances {
	s.Instance = &v
	return s
}

func (s *InstallAgentWithTypeRequestInstances) SetRegion(v string) *InstallAgentWithTypeRequestInstances {
	s.Region = &v
	return s
}

func (s *InstallAgentWithTypeRequestInstances) Validate() error {
	return dara.Validate(s)
}
