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
	// The unique identifier for the CodeInterpreter.
	//
	// example:
	//
	// ci-1234567890abcdef
	CodeInterpreterId *string `json:"codeInterpreterId,omitempty" xml:"codeInterpreterId,omitempty"`
	// A user-defined name for the CodeInterpreter instance.
	//
	// example:
	//
	// my-code-interpreter
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	// The number of CPU cores allocated to the instance.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The creation timestamp for the CodeInterpreter, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The description of the CodeInterpreter.
	//
	// example:
	//
	// Python code interpreter for data analysis
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ARN of the execution role for this CodeInterpreter.
	//
	// example:
	//
	// acs:ram::1760720386195983:role/CodeInterpreterExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// The last update timestamp for the CodeInterpreter, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// The memory allocated to the instance, in MB.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// The network configuration for the CodeInterpreter.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// The current status of the CodeInterpreter, such as READY or TERMINATED.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the current status, if applicable.
	//
	// example:
	//
	// Code interpreter is ready for use
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The unique identifier for the tenant.
	//
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
