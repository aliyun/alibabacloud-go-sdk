// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployConfigInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v int32) *DeployConfigInfo
	GetCpu() *int32
	SetDeploymentId(v int64) *DeployConfigInfo
	GetDeploymentId() *int64
	SetDeploymentType(v string) *DeployConfigInfo
	GetDeploymentType() *string
	SetEnv(v string) *DeployConfigInfo
	GetEnv() *string
	SetFinish(v bool) *DeployConfigInfo
	GetFinish() *bool
	SetInstanceCount(v int32) *DeployConfigInfo
	GetInstanceCount() *int32
	SetIsMonitorEnable(v int32) *DeployConfigInfo
	GetIsMonitorEnable() *int32
	SetIsServiceGroupEnable(v int32) *DeployConfigInfo
	GetIsServiceGroupEnable() *int32
	SetMemory(v int32) *DeployConfigInfo
	GetMemory() *int32
	SetPipelineId(v string) *DeployConfigInfo
	GetPipelineId() *string
	SetPreStop(v string) *DeployConfigInfo
	GetPreStop() *string
	SetServiceGroupId(v int64) *DeployConfigInfo
	GetServiceGroupId() *int64
	SetServiceId(v int64) *DeployConfigInfo
	GetServiceId() *int64
	SetTimeout(v int32) *DeployConfigInfo
	GetTimeout() *int32
	SetTimes(v int32) *DeployConfigInfo
	GetTimes() *int32
}

type DeployConfigInfo struct {
	// example:
	//
	// 2
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 1
	DeploymentId *int64 `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// AUTO
	DeploymentType *string `json:"deploymentType,omitempty" xml:"deploymentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// true
	Finish *bool `json:"finish,omitempty" xml:"finish,omitempty"`
	// example:
	//
	// 3
	InstanceCount   *int32 `json:"instanceCount,omitempty" xml:"instanceCount,omitempty"`
	IsMonitorEnable *int32 `json:"isMonitorEnable,omitempty" xml:"isMonitorEnable,omitempty"`
	// example:
	//
	// 1
	IsServiceGroupEnable *int32 `json:"isServiceGroupEnable,omitempty" xml:"isServiceGroupEnable,omitempty"`
	// example:
	//
	// 4
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// 1234
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PreStop    *string `json:"preStop,omitempty" xml:"preStop,omitempty"`
	// example:
	//
	// 1
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// 123
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// example:
	//
	// 1
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s DeployConfigInfo) String() string {
	return dara.Prettify(s)
}

func (s DeployConfigInfo) GoString() string {
	return s.String()
}

func (s *DeployConfigInfo) GetCpu() *int32 {
	return s.Cpu
}

func (s *DeployConfigInfo) GetDeploymentId() *int64 {
	return s.DeploymentId
}

func (s *DeployConfigInfo) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DeployConfigInfo) GetEnv() *string {
	return s.Env
}

func (s *DeployConfigInfo) GetFinish() *bool {
	return s.Finish
}

func (s *DeployConfigInfo) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DeployConfigInfo) GetIsMonitorEnable() *int32 {
	return s.IsMonitorEnable
}

func (s *DeployConfigInfo) GetIsServiceGroupEnable() *int32 {
	return s.IsServiceGroupEnable
}

func (s *DeployConfigInfo) GetMemory() *int32 {
	return s.Memory
}

func (s *DeployConfigInfo) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DeployConfigInfo) GetPreStop() *string {
	return s.PreStop
}

func (s *DeployConfigInfo) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *DeployConfigInfo) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *DeployConfigInfo) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployConfigInfo) GetTimes() *int32 {
	return s.Times
}

func (s *DeployConfigInfo) SetCpu(v int32) *DeployConfigInfo {
	s.Cpu = &v
	return s
}

func (s *DeployConfigInfo) SetDeploymentId(v int64) *DeployConfigInfo {
	s.DeploymentId = &v
	return s
}

func (s *DeployConfigInfo) SetDeploymentType(v string) *DeployConfigInfo {
	s.DeploymentType = &v
	return s
}

func (s *DeployConfigInfo) SetEnv(v string) *DeployConfigInfo {
	s.Env = &v
	return s
}

func (s *DeployConfigInfo) SetFinish(v bool) *DeployConfigInfo {
	s.Finish = &v
	return s
}

func (s *DeployConfigInfo) SetInstanceCount(v int32) *DeployConfigInfo {
	s.InstanceCount = &v
	return s
}

func (s *DeployConfigInfo) SetIsMonitorEnable(v int32) *DeployConfigInfo {
	s.IsMonitorEnable = &v
	return s
}

func (s *DeployConfigInfo) SetIsServiceGroupEnable(v int32) *DeployConfigInfo {
	s.IsServiceGroupEnable = &v
	return s
}

func (s *DeployConfigInfo) SetMemory(v int32) *DeployConfigInfo {
	s.Memory = &v
	return s
}

func (s *DeployConfigInfo) SetPipelineId(v string) *DeployConfigInfo {
	s.PipelineId = &v
	return s
}

func (s *DeployConfigInfo) SetPreStop(v string) *DeployConfigInfo {
	s.PreStop = &v
	return s
}

func (s *DeployConfigInfo) SetServiceGroupId(v int64) *DeployConfigInfo {
	s.ServiceGroupId = &v
	return s
}

func (s *DeployConfigInfo) SetServiceId(v int64) *DeployConfigInfo {
	s.ServiceId = &v
	return s
}

func (s *DeployConfigInfo) SetTimeout(v int32) *DeployConfigInfo {
	s.Timeout = &v
	return s
}

func (s *DeployConfigInfo) SetTimes(v int32) *DeployConfigInfo {
	s.Times = &v
	return s
}

func (s *DeployConfigInfo) Validate() error {
	return dara.Validate(s)
}
