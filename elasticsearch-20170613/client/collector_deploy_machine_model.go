// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectorDeployMachine interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *CollectorDeployMachine
	GetConfigType() *string
	SetGroupId(v string) *CollectorDeployMachine
	GetGroupId() *string
	SetInstanceId(v string) *CollectorDeployMachine
	GetInstanceId() *string
	SetMachines(v []*CollectorDeployMachineMachines) *CollectorDeployMachine
	GetMachines() []*CollectorDeployMachineMachines
	SetSuccessPodsCount(v string) *CollectorDeployMachine
	GetSuccessPodsCount() *string
	SetTotalPodsCount(v string) *CollectorDeployMachine
	GetTotalPodsCount() *string
	SetType(v string) *CollectorDeployMachine
	GetType() *string
}

type CollectorDeployMachine struct {
	// This parameter is required.
	//
	// example:
	//
	// collectorDeployMachine
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// example:
	//
	// default_ct-cn-f3t0dq5p97199ru3z
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// ce93d7566df2141f490f0f60f646252c3
	InstanceId *string                           `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Machines   []*CollectorDeployMachineMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SuccessPodsCount *string `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	// example:
	//
	// 2
	TotalPodsCount *string `json:"totalPodsCount,omitempty" xml:"totalPodsCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ACKCluster
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CollectorDeployMachine) String() string {
	return dara.Prettify(s)
}

func (s CollectorDeployMachine) GoString() string {
	return s.String()
}

func (s *CollectorDeployMachine) GetConfigType() *string {
	return s.ConfigType
}

func (s *CollectorDeployMachine) GetGroupId() *string {
	return s.GroupId
}

func (s *CollectorDeployMachine) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CollectorDeployMachine) GetMachines() []*CollectorDeployMachineMachines {
	return s.Machines
}

func (s *CollectorDeployMachine) GetSuccessPodsCount() *string {
	return s.SuccessPodsCount
}

func (s *CollectorDeployMachine) GetTotalPodsCount() *string {
	return s.TotalPodsCount
}

func (s *CollectorDeployMachine) GetType() *string {
	return s.Type
}

func (s *CollectorDeployMachine) SetConfigType(v string) *CollectorDeployMachine {
	s.ConfigType = &v
	return s
}

func (s *CollectorDeployMachine) SetGroupId(v string) *CollectorDeployMachine {
	s.GroupId = &v
	return s
}

func (s *CollectorDeployMachine) SetInstanceId(v string) *CollectorDeployMachine {
	s.InstanceId = &v
	return s
}

func (s *CollectorDeployMachine) SetMachines(v []*CollectorDeployMachineMachines) *CollectorDeployMachine {
	s.Machines = v
	return s
}

func (s *CollectorDeployMachine) SetSuccessPodsCount(v string) *CollectorDeployMachine {
	s.SuccessPodsCount = &v
	return s
}

func (s *CollectorDeployMachine) SetTotalPodsCount(v string) *CollectorDeployMachine {
	s.TotalPodsCount = &v
	return s
}

func (s *CollectorDeployMachine) SetType(v string) *CollectorDeployMachine {
	s.Type = &v
	return s
}

func (s *CollectorDeployMachine) Validate() error {
	return dara.Validate(s)
}

type CollectorDeployMachineMachines struct {
	// example:
	//
	// UNINSTALLED
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	// example:
	//
	// i-xs34****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CollectorDeployMachineMachines) String() string {
	return dara.Prettify(s)
}

func (s CollectorDeployMachineMachines) GoString() string {
	return s.String()
}

func (s *CollectorDeployMachineMachines) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *CollectorDeployMachineMachines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CollectorDeployMachineMachines) SetAgentStatus(v string) *CollectorDeployMachineMachines {
	s.AgentStatus = &v
	return s
}

func (s *CollectorDeployMachineMachines) SetInstanceId(v string) *CollectorDeployMachineMachines {
	s.InstanceId = &v
	return s
}

func (s *CollectorDeployMachineMachines) Validate() error {
	return dara.Validate(s)
}
