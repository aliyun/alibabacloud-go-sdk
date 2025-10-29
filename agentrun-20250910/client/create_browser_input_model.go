// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrowserInput interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserName(v string) *CreateBrowserInput
	GetBrowserName() *string
	SetCpu(v float32) *CreateBrowserInput
	GetCpu() *float32
	SetCredentialId(v string) *CreateBrowserInput
	GetCredentialId() *string
	SetDescription(v string) *CreateBrowserInput
	GetDescription() *string
	SetExecutionRoleArn(v string) *CreateBrowserInput
	GetExecutionRoleArn() *string
	SetMemory(v int32) *CreateBrowserInput
	GetMemory() *int32
	SetNetworkConfiguration(v *NetworkConfiguration) *CreateBrowserInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetSessionIdleTimeoutSeconds(v int32) *CreateBrowserInput
	GetSessionIdleTimeoutSeconds() *int32
}

type CreateBrowserInput struct {
	// This parameter is required.
	//
	// example:
	//
	// my-browser
	BrowserName *string `json:"browserName,omitempty" xml:"browserName,omitempty"`
	// CPU资源配置（单位：核）
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// cred-1234567890abcdef
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// example:
	//
	// Web automation browser for testing
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// acs:ram::1760720386195983:role/BrowserExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 内存资源配置（单位：MB）
	//
	// This parameter is required.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// This parameter is required.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 会话的空闲超时时间，单位为秒。实例没有会话请求后处于空闲状态，空闲态为闲置计费模式，超过此超时时间后会话自动过期，不可继续使用
	//
	// example:
	//
	// 3600
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
}

func (s CreateBrowserInput) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInput) GoString() string {
	return s.String()
}

func (s *CreateBrowserInput) GetBrowserName() *string {
	return s.BrowserName
}

func (s *CreateBrowserInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateBrowserInput) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateBrowserInput) GetDescription() *string {
	return s.Description
}

func (s *CreateBrowserInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateBrowserInput) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateBrowserInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateBrowserInput) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *CreateBrowserInput) SetBrowserName(v string) *CreateBrowserInput {
	s.BrowserName = &v
	return s
}

func (s *CreateBrowserInput) SetCpu(v float32) *CreateBrowserInput {
	s.Cpu = &v
	return s
}

func (s *CreateBrowserInput) SetCredentialId(v string) *CreateBrowserInput {
	s.CredentialId = &v
	return s
}

func (s *CreateBrowserInput) SetDescription(v string) *CreateBrowserInput {
	s.Description = &v
	return s
}

func (s *CreateBrowserInput) SetExecutionRoleArn(v string) *CreateBrowserInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateBrowserInput) SetMemory(v int32) *CreateBrowserInput {
	s.Memory = &v
	return s
}

func (s *CreateBrowserInput) SetNetworkConfiguration(v *NetworkConfiguration) *CreateBrowserInput {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateBrowserInput) SetSessionIdleTimeoutSeconds(v int32) *CreateBrowserInput {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *CreateBrowserInput) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
