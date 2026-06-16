// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeInterpreterInput interface {
	dara.Model
	String() string
	GoString() string
	SetCodeInterpreterName(v string) *CreateCodeInterpreterInput
	GetCodeInterpreterName() *string
	SetCpu(v float32) *CreateCodeInterpreterInput
	GetCpu() *float32
	SetCredentialId(v string) *CreateCodeInterpreterInput
	GetCredentialId() *string
	SetDescription(v string) *CreateCodeInterpreterInput
	GetDescription() *string
	SetExecutionRoleArn(v string) *CreateCodeInterpreterInput
	GetExecutionRoleArn() *string
	SetMemory(v int32) *CreateCodeInterpreterInput
	GetMemory() *int32
	SetNetworkConfiguration(v *NetworkConfiguration) *CreateCodeInterpreterInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetSessionIdleTimeoutSeconds(v int32) *CreateCodeInterpreterInput
	GetSessionIdleTimeoutSeconds() *int32
}

type CreateCodeInterpreterInput struct {
	// The name of the code interpreter. Use this to identify and distinguish code interpreter instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-code-interpreter
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	// The amount of CPU to allocate, in cores.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The credential ID used for authentication.
	//
	// example:
	//
	// cred-1234567890abcdef
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// A description of the code interpreter.
	//
	// example:
	//
	// Python code interpreter for data analysis
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the execution role for the code interpreter.
	//
	// example:
	//
	// acs:ram::1760720386195983:role/CodeInterpreterExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// The amount of memory to allocate, in megabytes (MB).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// Specifies the network configuration for the code interpreter, including VPC and security group settings.
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// The idle timeout for a session, in seconds. If an instance has no new requests for this duration, its session expires and cannot be reused.
	//
	// example:
	//
	// 3600
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
}

func (s CreateCodeInterpreterInput) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeInterpreterInput) GoString() string {
	return s.String()
}

func (s *CreateCodeInterpreterInput) GetCodeInterpreterName() *string {
	return s.CodeInterpreterName
}

func (s *CreateCodeInterpreterInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateCodeInterpreterInput) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateCodeInterpreterInput) GetDescription() *string {
	return s.Description
}

func (s *CreateCodeInterpreterInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateCodeInterpreterInput) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateCodeInterpreterInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateCodeInterpreterInput) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *CreateCodeInterpreterInput) SetCodeInterpreterName(v string) *CreateCodeInterpreterInput {
	s.CodeInterpreterName = &v
	return s
}

func (s *CreateCodeInterpreterInput) SetCpu(v float32) *CreateCodeInterpreterInput {
	s.Cpu = &v
	return s
}

func (s *CreateCodeInterpreterInput) SetCredentialId(v string) *CreateCodeInterpreterInput {
	s.CredentialId = &v
	return s
}

func (s *CreateCodeInterpreterInput) SetDescription(v string) *CreateCodeInterpreterInput {
	s.Description = &v
	return s
}

func (s *CreateCodeInterpreterInput) SetExecutionRoleArn(v string) *CreateCodeInterpreterInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateCodeInterpreterInput) SetMemory(v int32) *CreateCodeInterpreterInput {
	s.Memory = &v
	return s
}

func (s *CreateCodeInterpreterInput) SetNetworkConfiguration(v *NetworkConfiguration) *CreateCodeInterpreterInput {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateCodeInterpreterInput) SetSessionIdleTimeoutSeconds(v int32) *CreateCodeInterpreterInput {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *CreateCodeInterpreterInput) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
