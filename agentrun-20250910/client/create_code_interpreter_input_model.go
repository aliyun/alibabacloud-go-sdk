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
	// 代码解释器的名称，用于标识和区分不同的代码解释器实例
	//
	// This parameter is required.
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	// CPU资源配置（单位：核数）
	//
	// This parameter is required.
	Cpu          *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CredentialId *string  `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// 代码解释器的描述信息，说明该解释器的用途和功能
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 此代码解释器的执行角色
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 内存资源配置（单位：MB）
	//
	// This parameter is required.
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// 代码解释器的网络配置，包括VPC、安全组等网络访问设置
	//
	// This parameter is required.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 会话的空闲超时时间，单位为秒。实例没有会话请求后处于空闲状态，空闲态为闲置计费模式，超过此超时时间后会话自动过期，不可继续使用
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
	return dara.Validate(s)
}
