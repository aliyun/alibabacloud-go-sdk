// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentTriggerCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeploymentTriggerCmd
	GetAccountId() *string
	SetApplicationType(v string) *DeploymentTriggerCmd
	GetApplicationType() *string
	SetCpu(v int32) *DeploymentTriggerCmd
	GetCpu() *int32
	SetDescription(v string) *DeploymentTriggerCmd
	GetDescription() *string
	SetImageTag(v string) *DeploymentTriggerCmd
	GetImageTag() *string
	SetInstanceCount(v int32) *DeploymentTriggerCmd
	GetInstanceCount() *int32
	SetIsMonitorEnable(v int32) *DeploymentTriggerCmd
	GetIsMonitorEnable() *int32
	SetIsServiceGroupEnable(v int32) *DeploymentTriggerCmd
	GetIsServiceGroupEnable() *int32
	SetMemory(v int32) *DeploymentTriggerCmd
	GetMemory() *int32
	SetPreStop(v string) *DeploymentTriggerCmd
	GetPreStop() *string
	SetServiceGroupId(v int64) *DeploymentTriggerCmd
	GetServiceGroupId() *int64
	SetTimeout(v int32) *DeploymentTriggerCmd
	GetTimeout() *int32
	SetTimes(v int32) *DeploymentTriggerCmd
	GetTimes() *int32
	SetType(v string) *DeploymentTriggerCmd
	GetType() *string
}

type DeploymentTriggerCmd struct {
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// DAPR
	ApplicationType *string `json:"applicationType,omitempty" xml:"applicationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 员工管理
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// v0.1
	ImageTag *string `json:"imageTag,omitempty" xml:"imageTag,omitempty"`
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
	Memory  *int32  `json:"memory,omitempty" xml:"memory,omitempty"`
	PreStop *string `json:"preStop,omitempty" xml:"preStop,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// example:
	//
	// 123
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

func (s DeploymentTriggerCmd) String() string {
	return dara.Prettify(s)
}

func (s DeploymentTriggerCmd) GoString() string {
	return s.String()
}

func (s *DeploymentTriggerCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *DeploymentTriggerCmd) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DeploymentTriggerCmd) GetCpu() *int32 {
	return s.Cpu
}

func (s *DeploymentTriggerCmd) GetDescription() *string {
	return s.Description
}

func (s *DeploymentTriggerCmd) GetImageTag() *string {
	return s.ImageTag
}

func (s *DeploymentTriggerCmd) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DeploymentTriggerCmd) GetIsMonitorEnable() *int32 {
	return s.IsMonitorEnable
}

func (s *DeploymentTriggerCmd) GetIsServiceGroupEnable() *int32 {
	return s.IsServiceGroupEnable
}

func (s *DeploymentTriggerCmd) GetMemory() *int32 {
	return s.Memory
}

func (s *DeploymentTriggerCmd) GetPreStop() *string {
	return s.PreStop
}

func (s *DeploymentTriggerCmd) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *DeploymentTriggerCmd) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeploymentTriggerCmd) GetTimes() *int32 {
	return s.Times
}

func (s *DeploymentTriggerCmd) GetType() *string {
	return s.Type
}

func (s *DeploymentTriggerCmd) SetAccountId(v string) *DeploymentTriggerCmd {
	s.AccountId = &v
	return s
}

func (s *DeploymentTriggerCmd) SetApplicationType(v string) *DeploymentTriggerCmd {
	s.ApplicationType = &v
	return s
}

func (s *DeploymentTriggerCmd) SetCpu(v int32) *DeploymentTriggerCmd {
	s.Cpu = &v
	return s
}

func (s *DeploymentTriggerCmd) SetDescription(v string) *DeploymentTriggerCmd {
	s.Description = &v
	return s
}

func (s *DeploymentTriggerCmd) SetImageTag(v string) *DeploymentTriggerCmd {
	s.ImageTag = &v
	return s
}

func (s *DeploymentTriggerCmd) SetInstanceCount(v int32) *DeploymentTriggerCmd {
	s.InstanceCount = &v
	return s
}

func (s *DeploymentTriggerCmd) SetIsMonitorEnable(v int32) *DeploymentTriggerCmd {
	s.IsMonitorEnable = &v
	return s
}

func (s *DeploymentTriggerCmd) SetIsServiceGroupEnable(v int32) *DeploymentTriggerCmd {
	s.IsServiceGroupEnable = &v
	return s
}

func (s *DeploymentTriggerCmd) SetMemory(v int32) *DeploymentTriggerCmd {
	s.Memory = &v
	return s
}

func (s *DeploymentTriggerCmd) SetPreStop(v string) *DeploymentTriggerCmd {
	s.PreStop = &v
	return s
}

func (s *DeploymentTriggerCmd) SetServiceGroupId(v int64) *DeploymentTriggerCmd {
	s.ServiceGroupId = &v
	return s
}

func (s *DeploymentTriggerCmd) SetTimeout(v int32) *DeploymentTriggerCmd {
	s.Timeout = &v
	return s
}

func (s *DeploymentTriggerCmd) SetTimes(v int32) *DeploymentTriggerCmd {
	s.Times = &v
	return s
}

func (s *DeploymentTriggerCmd) SetType(v string) *DeploymentTriggerCmd {
	s.Type = &v
	return s
}

func (s *DeploymentTriggerCmd) Validate() error {
	return dara.Validate(s)
}
