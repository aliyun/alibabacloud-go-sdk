// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowser interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserId(v string) *Browser
	GetBrowserId() *string
	SetBrowserName(v string) *Browser
	GetBrowserName() *string
	SetCpu(v float32) *Browser
	GetCpu() *float32
	SetCreatedAt(v string) *Browser
	GetCreatedAt() *string
	SetCredentialId(v string) *Browser
	GetCredentialId() *string
	SetDescription(v string) *Browser
	GetDescription() *string
	SetExecutionRoleArn(v string) *Browser
	GetExecutionRoleArn() *string
	SetLastUpdatedAt(v string) *Browser
	GetLastUpdatedAt() *string
	SetMemory(v int32) *Browser
	GetMemory() *int32
	SetNetworkConfiguration(v *NetworkConfiguration) *Browser
	GetNetworkConfiguration() *NetworkConfiguration
	SetRecording(v *BrowserRecordingConfiguration) *Browser
	GetRecording() *BrowserRecordingConfiguration
	SetStatus(v string) *Browser
	GetStatus() *string
	SetStatusReason(v string) *Browser
	GetStatusReason() *string
	SetTenantId(v string) *Browser
	GetTenantId() *string
}

type Browser struct {
	BrowserId        *string  `json:"browserId,omitempty" xml:"browserId,omitempty"`
	BrowserName      *string  `json:"browserName,omitempty" xml:"browserName,omitempty"`
	Cpu              *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CreatedAt        *string  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialId     *string  `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	Description      *string  `json:"description,omitempty" xml:"description,omitempty"`
	ExecutionRoleArn *string  `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LastUpdatedAt    *string  `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// 内存资源配置（单位：MB）
	Memory               *int32                         `json:"memory,omitempty" xml:"memory,omitempty"`
	NetworkConfiguration *NetworkConfiguration          `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	Recording            *BrowserRecordingConfiguration `json:"recording,omitempty" xml:"recording,omitempty"`
	Status               *string                        `json:"status,omitempty" xml:"status,omitempty"`
	// 当前状态的原因说明（如适用）
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	TenantId     *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s Browser) String() string {
	return dara.Prettify(s)
}

func (s Browser) GoString() string {
	return s.String()
}

func (s *Browser) GetBrowserId() *string {
	return s.BrowserId
}

func (s *Browser) GetBrowserName() *string {
	return s.BrowserName
}

func (s *Browser) GetCpu() *float32 {
	return s.Cpu
}

func (s *Browser) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Browser) GetCredentialId() *string {
	return s.CredentialId
}

func (s *Browser) GetDescription() *string {
	return s.Description
}

func (s *Browser) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *Browser) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *Browser) GetMemory() *int32 {
	return s.Memory
}

func (s *Browser) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *Browser) GetRecording() *BrowserRecordingConfiguration {
	return s.Recording
}

func (s *Browser) GetStatus() *string {
	return s.Status
}

func (s *Browser) GetStatusReason() *string {
	return s.StatusReason
}

func (s *Browser) GetTenantId() *string {
	return s.TenantId
}

func (s *Browser) SetBrowserId(v string) *Browser {
	s.BrowserId = &v
	return s
}

func (s *Browser) SetBrowserName(v string) *Browser {
	s.BrowserName = &v
	return s
}

func (s *Browser) SetCpu(v float32) *Browser {
	s.Cpu = &v
	return s
}

func (s *Browser) SetCreatedAt(v string) *Browser {
	s.CreatedAt = &v
	return s
}

func (s *Browser) SetCredentialId(v string) *Browser {
	s.CredentialId = &v
	return s
}

func (s *Browser) SetDescription(v string) *Browser {
	s.Description = &v
	return s
}

func (s *Browser) SetExecutionRoleArn(v string) *Browser {
	s.ExecutionRoleArn = &v
	return s
}

func (s *Browser) SetLastUpdatedAt(v string) *Browser {
	s.LastUpdatedAt = &v
	return s
}

func (s *Browser) SetMemory(v int32) *Browser {
	s.Memory = &v
	return s
}

func (s *Browser) SetNetworkConfiguration(v *NetworkConfiguration) *Browser {
	s.NetworkConfiguration = v
	return s
}

func (s *Browser) SetRecording(v *BrowserRecordingConfiguration) *Browser {
	s.Recording = v
	return s
}

func (s *Browser) SetStatus(v string) *Browser {
	s.Status = &v
	return s
}

func (s *Browser) SetStatusReason(v string) *Browser {
	s.StatusReason = &v
	return s
}

func (s *Browser) SetTenantId(v string) *Browser {
	s.TenantId = &v
	return s
}

func (s *Browser) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Recording != nil {
		if err := s.Recording.Validate(); err != nil {
			return err
		}
	}
	return nil
}
