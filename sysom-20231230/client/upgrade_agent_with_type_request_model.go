// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentWithTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *UpgradeAgentWithTypeRequest
	GetXDebugId() *string
	SetAgentId(v string) *UpgradeAgentWithTypeRequest
	GetAgentId() *string
	SetAgentVersion(v string) *UpgradeAgentWithTypeRequest
	GetAgentVersion() *string
	SetInstanceType(v string) *UpgradeAgentWithTypeRequest
	GetInstanceType() *string
	SetInstances(v []*UpgradeAgentWithTypeRequestInstances) *UpgradeAgentWithTypeRequest
	GetInstances() []*UpgradeAgentWithTypeRequestInstances
	SetXSysomInvokeSource(v string) *UpgradeAgentWithTypeRequest
	GetXSysomInvokeSource() *string
}

type UpgradeAgentWithTypeRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The ID of the component to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The version of the component to update to.
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
	// The list of instances for which to update the component.
	//
	// This parameter is required.
	Instances          []*UpgradeAgentWithTypeRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	XSysomInvokeSource *string                                 `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s UpgradeAgentWithTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentWithTypeRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAgentWithTypeRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *UpgradeAgentWithTypeRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpgradeAgentWithTypeRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UpgradeAgentWithTypeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpgradeAgentWithTypeRequest) GetInstances() []*UpgradeAgentWithTypeRequestInstances {
	return s.Instances
}

func (s *UpgradeAgentWithTypeRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *UpgradeAgentWithTypeRequest) SetXDebugId(v string) *UpgradeAgentWithTypeRequest {
	s.XDebugId = &v
	return s
}

func (s *UpgradeAgentWithTypeRequest) SetAgentId(v string) *UpgradeAgentWithTypeRequest {
	s.AgentId = &v
	return s
}

func (s *UpgradeAgentWithTypeRequest) SetAgentVersion(v string) *UpgradeAgentWithTypeRequest {
	s.AgentVersion = &v
	return s
}

func (s *UpgradeAgentWithTypeRequest) SetInstanceType(v string) *UpgradeAgentWithTypeRequest {
	s.InstanceType = &v
	return s
}

func (s *UpgradeAgentWithTypeRequest) SetInstances(v []*UpgradeAgentWithTypeRequestInstances) *UpgradeAgentWithTypeRequest {
	s.Instances = v
	return s
}

func (s *UpgradeAgentWithTypeRequest) SetXSysomInvokeSource(v string) *UpgradeAgentWithTypeRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *UpgradeAgentWithTypeRequest) Validate() error {
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

type UpgradeAgentWithTypeRequestInstances struct {
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

func (s UpgradeAgentWithTypeRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentWithTypeRequestInstances) GoString() string {
	return s.String()
}

func (s *UpgradeAgentWithTypeRequestInstances) GetInstance() *string {
	return s.Instance
}

func (s *UpgradeAgentWithTypeRequestInstances) GetRegion() *string {
	return s.Region
}

func (s *UpgradeAgentWithTypeRequestInstances) SetInstance(v string) *UpgradeAgentWithTypeRequestInstances {
	s.Instance = &v
	return s
}

func (s *UpgradeAgentWithTypeRequestInstances) SetRegion(v string) *UpgradeAgentWithTypeRequestInstances {
	s.Region = &v
	return s
}

func (s *UpgradeAgentWithTypeRequestInstances) Validate() error {
	return dara.Validate(s)
}
