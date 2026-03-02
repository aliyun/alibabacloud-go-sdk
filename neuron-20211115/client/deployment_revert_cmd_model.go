// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentRevertCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeploymentRevertCmd
	GetAccountId() *string
	SetCpu(v int32) *DeploymentRevertCmd
	GetCpu() *int32
	SetDeploymentId(v int64) *DeploymentRevertCmd
	GetDeploymentId() *int64
	SetDescription(v string) *DeploymentRevertCmd
	GetDescription() *string
	SetInstanceCount(v int32) *DeploymentRevertCmd
	GetInstanceCount() *int32
	SetIsMonitorEnable(v int32) *DeploymentRevertCmd
	GetIsMonitorEnable() *int32
	SetIsServiceGroupEnable(v int32) *DeploymentRevertCmd
	GetIsServiceGroupEnable() *int32
	SetMemory(v int32) *DeploymentRevertCmd
	GetMemory() *int32
	SetTimeout(v int32) *DeploymentRevertCmd
	GetTimeout() *int32
	SetTimes(v int32) *DeploymentRevertCmd
	GetTimes() *int32
	SetType(v string) *DeploymentRevertCmd
	GetType() *string
}

type DeploymentRevertCmd struct {
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeploymentId *int64 `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// 员工管理
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	InstanceCount        *int32 `json:"instanceCount,omitempty" xml:"instanceCount,omitempty"`
	IsMonitorEnable      *int32 `json:"isMonitorEnable,omitempty" xml:"isMonitorEnable,omitempty"`
	IsServiceGroupEnable *int32 `json:"isServiceGroupEnable,omitempty" xml:"isServiceGroupEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AUTO
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeploymentRevertCmd) String() string {
	return dara.Prettify(s)
}

func (s DeploymentRevertCmd) GoString() string {
	return s.String()
}

func (s *DeploymentRevertCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *DeploymentRevertCmd) GetCpu() *int32 {
	return s.Cpu
}

func (s *DeploymentRevertCmd) GetDeploymentId() *int64 {
	return s.DeploymentId
}

func (s *DeploymentRevertCmd) GetDescription() *string {
	return s.Description
}

func (s *DeploymentRevertCmd) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DeploymentRevertCmd) GetIsMonitorEnable() *int32 {
	return s.IsMonitorEnable
}

func (s *DeploymentRevertCmd) GetIsServiceGroupEnable() *int32 {
	return s.IsServiceGroupEnable
}

func (s *DeploymentRevertCmd) GetMemory() *int32 {
	return s.Memory
}

func (s *DeploymentRevertCmd) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeploymentRevertCmd) GetTimes() *int32 {
	return s.Times
}

func (s *DeploymentRevertCmd) GetType() *string {
	return s.Type
}

func (s *DeploymentRevertCmd) SetAccountId(v string) *DeploymentRevertCmd {
	s.AccountId = &v
	return s
}

func (s *DeploymentRevertCmd) SetCpu(v int32) *DeploymentRevertCmd {
	s.Cpu = &v
	return s
}

func (s *DeploymentRevertCmd) SetDeploymentId(v int64) *DeploymentRevertCmd {
	s.DeploymentId = &v
	return s
}

func (s *DeploymentRevertCmd) SetDescription(v string) *DeploymentRevertCmd {
	s.Description = &v
	return s
}

func (s *DeploymentRevertCmd) SetInstanceCount(v int32) *DeploymentRevertCmd {
	s.InstanceCount = &v
	return s
}

func (s *DeploymentRevertCmd) SetIsMonitorEnable(v int32) *DeploymentRevertCmd {
	s.IsMonitorEnable = &v
	return s
}

func (s *DeploymentRevertCmd) SetIsServiceGroupEnable(v int32) *DeploymentRevertCmd {
	s.IsServiceGroupEnable = &v
	return s
}

func (s *DeploymentRevertCmd) SetMemory(v int32) *DeploymentRevertCmd {
	s.Memory = &v
	return s
}

func (s *DeploymentRevertCmd) SetTimeout(v int32) *DeploymentRevertCmd {
	s.Timeout = &v
	return s
}

func (s *DeploymentRevertCmd) SetTimes(v int32) *DeploymentRevertCmd {
	s.Times = &v
	return s
}

func (s *DeploymentRevertCmd) SetType(v string) *DeploymentRevertCmd {
	s.Type = &v
	return s
}

func (s *DeploymentRevertCmd) Validate() error {
	return dara.Validate(s)
}
