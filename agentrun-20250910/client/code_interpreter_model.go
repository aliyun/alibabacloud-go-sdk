// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeInterpreter interface {
	dara.Model
	String() string
	GoString() string
	SetCodeInterpreterId(v string) *CodeInterpreter
	GetCodeInterpreterId() *string
	SetCodeInterpreterName(v string) *CodeInterpreter
	GetCodeInterpreterName() *string
	SetCpu(v float32) *CodeInterpreter
	GetCpu() *float32
	SetCreatedAt(v string) *CodeInterpreter
	GetCreatedAt() *string
	SetDescription(v string) *CodeInterpreter
	GetDescription() *string
	SetExecutionRoleArn(v string) *CodeInterpreter
	GetExecutionRoleArn() *string
	SetLastUpdatedAt(v string) *CodeInterpreter
	GetLastUpdatedAt() *string
	SetMemory(v int32) *CodeInterpreter
	GetMemory() *int32
	SetNetworkConfiguration(v *NetworkConfiguration) *CodeInterpreter
	GetNetworkConfiguration() *NetworkConfiguration
	SetStatus(v string) *CodeInterpreter
	GetStatus() *string
	SetStatusReason(v string) *CodeInterpreter
	GetStatusReason() *string
	SetTenantId(v string) *CodeInterpreter
	GetTenantId() *string
}

type CodeInterpreter struct {
	// 代码解释器的唯一标识符
	//
	// example:
	//
	// ci-1234567890abcdef
	CodeInterpreterId *string `json:"codeInterpreterId,omitempty" xml:"codeInterpreterId,omitempty"`
	// 代码解释器的名称，用于标识和区分不同的代码解释器实例
	//
	// example:
	//
	// my-code-interpreter
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	// example:
	//
	// 2.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 代码解释器的创建时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 代码解释器的描述信息，说明该解释器的用途和功能
	//
	// example:
	//
	// Python code interpreter for data analysis
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 此代码解释器的执行角色
	//
	// example:
	//
	// acs:ram::1760720386195983:role/CodeInterpreterExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 代码解释器的最后更新时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// 内存资源配置（单位：MB）
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// 代码解释器的网络配置信息
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 代码解释器的当前状态，如READY（就绪）、TERMINATED（已终止）等
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 当前状态的原因说明（如适用）
	//
	// example:
	//
	// Code interpreter is ready for use
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// example:
	//
	// tenant-1234567890abcdef
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CodeInterpreter) String() string {
	return dara.Prettify(s)
}

func (s CodeInterpreter) GoString() string {
	return s.String()
}

func (s *CodeInterpreter) GetCodeInterpreterId() *string {
	return s.CodeInterpreterId
}

func (s *CodeInterpreter) GetCodeInterpreterName() *string {
	return s.CodeInterpreterName
}

func (s *CodeInterpreter) GetCpu() *float32 {
	return s.Cpu
}

func (s *CodeInterpreter) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CodeInterpreter) GetDescription() *string {
	return s.Description
}

func (s *CodeInterpreter) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CodeInterpreter) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *CodeInterpreter) GetMemory() *int32 {
	return s.Memory
}

func (s *CodeInterpreter) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CodeInterpreter) GetStatus() *string {
	return s.Status
}

func (s *CodeInterpreter) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CodeInterpreter) GetTenantId() *string {
	return s.TenantId
}

func (s *CodeInterpreter) SetCodeInterpreterId(v string) *CodeInterpreter {
	s.CodeInterpreterId = &v
	return s
}

func (s *CodeInterpreter) SetCodeInterpreterName(v string) *CodeInterpreter {
	s.CodeInterpreterName = &v
	return s
}

func (s *CodeInterpreter) SetCpu(v float32) *CodeInterpreter {
	s.Cpu = &v
	return s
}

func (s *CodeInterpreter) SetCreatedAt(v string) *CodeInterpreter {
	s.CreatedAt = &v
	return s
}

func (s *CodeInterpreter) SetDescription(v string) *CodeInterpreter {
	s.Description = &v
	return s
}

func (s *CodeInterpreter) SetExecutionRoleArn(v string) *CodeInterpreter {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CodeInterpreter) SetLastUpdatedAt(v string) *CodeInterpreter {
	s.LastUpdatedAt = &v
	return s
}

func (s *CodeInterpreter) SetMemory(v int32) *CodeInterpreter {
	s.Memory = &v
	return s
}

func (s *CodeInterpreter) SetNetworkConfiguration(v *NetworkConfiguration) *CodeInterpreter {
	s.NetworkConfiguration = v
	return s
}

func (s *CodeInterpreter) SetStatus(v string) *CodeInterpreter {
	s.Status = &v
	return s
}

func (s *CodeInterpreter) SetStatusReason(v string) *CodeInterpreter {
	s.StatusReason = &v
	return s
}

func (s *CodeInterpreter) SetTenantId(v string) *CodeInterpreter {
	s.TenantId = &v
	return s
}

func (s *CodeInterpreter) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
