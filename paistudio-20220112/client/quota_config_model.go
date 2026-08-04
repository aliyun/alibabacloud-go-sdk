// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaConfig interface {
	dara.Model
	String() string
	GoString() string
	SetACS(v *ACS) *QuotaConfig
	GetACS() *ACS
	SetClusterId(v string) *QuotaConfig
	GetClusterId() *string
	SetControlPlaneClusterId(v string) *QuotaConfig
	GetControlPlaneClusterId() *string
	SetDefaultGPUDriver(v string) *QuotaConfig
	GetDefaultGPUDriver() *string
	SetEnableGPUShare(v bool) *QuotaConfig
	GetEnableGPUShare() *bool
	SetEnablePreemptSubquotaWorkloads(v bool) *QuotaConfig
	GetEnablePreemptSubquotaWorkloads() *bool
	SetEnableSelfQuotaPreemption(v bool) *QuotaConfig
	GetEnableSelfQuotaPreemption() *bool
	SetEnableSubQuotaPreemption(v bool) *QuotaConfig
	GetEnableSubQuotaPreemption() *bool
	SetEniCacheConfig(v *EniCacheConfig) *QuotaConfig
	GetEniCacheConfig() *EniCacheConfig
	SetIsEncryptedResource(v bool) *QuotaConfig
	GetIsEncryptedResource() *bool
	SetOversoldUsageConfig(v *OversoldUsageConfig) *QuotaConfig
	GetOversoldUsageConfig() *OversoldUsageConfig
	SetResourceSpecs(v []*WorkspaceSpecs) *QuotaConfig
	GetResourceSpecs() []*WorkspaceSpecs
	SetSandboxCacheConfig(v *SandboxCacheConfig) *QuotaConfig
	GetSandboxCacheConfig() *SandboxCacheConfig
	SetSelfQuotaPreemptionConfig(v *SelfQuotaPreemptionConfig) *QuotaConfig
	GetSelfQuotaPreemptionConfig() *SelfQuotaPreemptionConfig
	SetSubQuotaPreemptionConfig(v *SubQuotaPreemptionConfig) *QuotaConfig
	GetSubQuotaPreemptionConfig() *SubQuotaPreemptionConfig
	SetSupportGPUDrivers(v []*string) *QuotaConfig
	GetSupportGPUDrivers() []*string
	SetSupportRDMA(v bool) *QuotaConfig
	GetSupportRDMA() *bool
	SetUseCase(v string) *QuotaConfig
	GetUseCase() *string
	SetUserVpc(v *UserVpc) *QuotaConfig
	GetUserVpc() *UserVpc
	SetWorkloadTypes(v []*string) *QuotaConfig
	GetWorkloadTypes() []*string
}

type QuotaConfig struct {
	// The ACS-related configurations.
	ACS *ACS `json:"ACS,omitempty" xml:"ACS,omitempty"`
	// The ID of the cluster where the quota resides.
	//
	// example:
	//
	// ceeb3724255364***
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ControlPlaneClusterId *string `json:"ControlPlaneClusterId,omitempty" xml:"ControlPlaneClusterId,omitempty"`
	// The default GPU driver version for the resource quota.
	//
	// example:
	//
	// 470.199.02
	DefaultGPUDriver *string `json:"DefaultGPUDriver,omitempty" xml:"DefaultGPUDriver,omitempty"`
	EnableGPUShare   *bool   `json:"EnableGPUShare,omitempty" xml:"EnableGPUShare,omitempty"`
	// Specifies whether workloads in sub-quotas can be preempted.
	//
	// example:
	//
	// false
	EnablePreemptSubquotaWorkloads *bool `json:"EnablePreemptSubquotaWorkloads,omitempty" xml:"EnablePreemptSubquotaWorkloads,omitempty"`
	// Specifies whether guaranteed resources within this quota can be preempted.
	EnableSelfQuotaPreemption *bool `json:"EnableSelfQuotaPreemption,omitempty" xml:"EnableSelfQuotaPreemption,omitempty"`
	// Specifies whether resources in sub-quotas can be preempted.
	EnableSubQuotaPreemption *bool                `json:"EnableSubQuotaPreemption,omitempty" xml:"EnableSubQuotaPreemption,omitempty"`
	EniCacheConfig           *EniCacheConfig      `json:"EniCacheConfig,omitempty" xml:"EniCacheConfig,omitempty"`
	IsEncryptedResource      *bool                `json:"IsEncryptedResource,omitempty" xml:"IsEncryptedResource,omitempty"`
	OversoldUsageConfig      *OversoldUsageConfig `json:"OversoldUsageConfig,omitempty" xml:"OversoldUsageConfig,omitempty"`
	// The resource specification templates.
	ResourceSpecs             []*WorkspaceSpecs          `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	SandboxCacheConfig        *SandboxCacheConfig        `json:"SandboxCacheConfig,omitempty" xml:"SandboxCacheConfig,omitempty"`
	SelfQuotaPreemptionConfig *SelfQuotaPreemptionConfig `json:"SelfQuotaPreemptionConfig,omitempty" xml:"SelfQuotaPreemptionConfig,omitempty"`
	// The configuration for the sub-quota preemption task.
	SubQuotaPreemptionConfig *SubQuotaPreemptionConfig `json:"SubQuotaPreemptionConfig,omitempty" xml:"SubQuotaPreemptionConfig,omitempty"`
	// The GPU driver versions supported by the resource quota.
	SupportGPUDrivers []*string `json:"SupportGPUDrivers,omitempty" xml:"SupportGPUDrivers,omitempty" type:"Repeated"`
	// Specifies whether RDMA is supported.
	//
	// example:
	//
	// false
	SupportRDMA *bool   `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
	UseCase     *string `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
	// The user VPC information.
	UserVpc       *UserVpc  `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	WorkloadTypes []*string `json:"WorkloadTypes,omitempty" xml:"WorkloadTypes,omitempty" type:"Repeated"`
}

func (s QuotaConfig) String() string {
	return dara.Prettify(s)
}

func (s QuotaConfig) GoString() string {
	return s.String()
}

func (s *QuotaConfig) GetACS() *ACS {
	return s.ACS
}

func (s *QuotaConfig) GetClusterId() *string {
	return s.ClusterId
}

func (s *QuotaConfig) GetControlPlaneClusterId() *string {
	return s.ControlPlaneClusterId
}

func (s *QuotaConfig) GetDefaultGPUDriver() *string {
	return s.DefaultGPUDriver
}

func (s *QuotaConfig) GetEnableGPUShare() *bool {
	return s.EnableGPUShare
}

func (s *QuotaConfig) GetEnablePreemptSubquotaWorkloads() *bool {
	return s.EnablePreemptSubquotaWorkloads
}

func (s *QuotaConfig) GetEnableSelfQuotaPreemption() *bool {
	return s.EnableSelfQuotaPreemption
}

func (s *QuotaConfig) GetEnableSubQuotaPreemption() *bool {
	return s.EnableSubQuotaPreemption
}

func (s *QuotaConfig) GetEniCacheConfig() *EniCacheConfig {
	return s.EniCacheConfig
}

func (s *QuotaConfig) GetIsEncryptedResource() *bool {
	return s.IsEncryptedResource
}

func (s *QuotaConfig) GetOversoldUsageConfig() *OversoldUsageConfig {
	return s.OversoldUsageConfig
}

func (s *QuotaConfig) GetResourceSpecs() []*WorkspaceSpecs {
	return s.ResourceSpecs
}

func (s *QuotaConfig) GetSandboxCacheConfig() *SandboxCacheConfig {
	return s.SandboxCacheConfig
}

func (s *QuotaConfig) GetSelfQuotaPreemptionConfig() *SelfQuotaPreemptionConfig {
	return s.SelfQuotaPreemptionConfig
}

func (s *QuotaConfig) GetSubQuotaPreemptionConfig() *SubQuotaPreemptionConfig {
	return s.SubQuotaPreemptionConfig
}

func (s *QuotaConfig) GetSupportGPUDrivers() []*string {
	return s.SupportGPUDrivers
}

func (s *QuotaConfig) GetSupportRDMA() *bool {
	return s.SupportRDMA
}

func (s *QuotaConfig) GetUseCase() *string {
	return s.UseCase
}

func (s *QuotaConfig) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *QuotaConfig) GetWorkloadTypes() []*string {
	return s.WorkloadTypes
}

func (s *QuotaConfig) SetACS(v *ACS) *QuotaConfig {
	s.ACS = v
	return s
}

func (s *QuotaConfig) SetClusterId(v string) *QuotaConfig {
	s.ClusterId = &v
	return s
}

func (s *QuotaConfig) SetControlPlaneClusterId(v string) *QuotaConfig {
	s.ControlPlaneClusterId = &v
	return s
}

func (s *QuotaConfig) SetDefaultGPUDriver(v string) *QuotaConfig {
	s.DefaultGPUDriver = &v
	return s
}

func (s *QuotaConfig) SetEnableGPUShare(v bool) *QuotaConfig {
	s.EnableGPUShare = &v
	return s
}

func (s *QuotaConfig) SetEnablePreemptSubquotaWorkloads(v bool) *QuotaConfig {
	s.EnablePreemptSubquotaWorkloads = &v
	return s
}

func (s *QuotaConfig) SetEnableSelfQuotaPreemption(v bool) *QuotaConfig {
	s.EnableSelfQuotaPreemption = &v
	return s
}

func (s *QuotaConfig) SetEnableSubQuotaPreemption(v bool) *QuotaConfig {
	s.EnableSubQuotaPreemption = &v
	return s
}

func (s *QuotaConfig) SetEniCacheConfig(v *EniCacheConfig) *QuotaConfig {
	s.EniCacheConfig = v
	return s
}

func (s *QuotaConfig) SetIsEncryptedResource(v bool) *QuotaConfig {
	s.IsEncryptedResource = &v
	return s
}

func (s *QuotaConfig) SetOversoldUsageConfig(v *OversoldUsageConfig) *QuotaConfig {
	s.OversoldUsageConfig = v
	return s
}

func (s *QuotaConfig) SetResourceSpecs(v []*WorkspaceSpecs) *QuotaConfig {
	s.ResourceSpecs = v
	return s
}

func (s *QuotaConfig) SetSandboxCacheConfig(v *SandboxCacheConfig) *QuotaConfig {
	s.SandboxCacheConfig = v
	return s
}

func (s *QuotaConfig) SetSelfQuotaPreemptionConfig(v *SelfQuotaPreemptionConfig) *QuotaConfig {
	s.SelfQuotaPreemptionConfig = v
	return s
}

func (s *QuotaConfig) SetSubQuotaPreemptionConfig(v *SubQuotaPreemptionConfig) *QuotaConfig {
	s.SubQuotaPreemptionConfig = v
	return s
}

func (s *QuotaConfig) SetSupportGPUDrivers(v []*string) *QuotaConfig {
	s.SupportGPUDrivers = v
	return s
}

func (s *QuotaConfig) SetSupportRDMA(v bool) *QuotaConfig {
	s.SupportRDMA = &v
	return s
}

func (s *QuotaConfig) SetUseCase(v string) *QuotaConfig {
	s.UseCase = &v
	return s
}

func (s *QuotaConfig) SetUserVpc(v *UserVpc) *QuotaConfig {
	s.UserVpc = v
	return s
}

func (s *QuotaConfig) SetWorkloadTypes(v []*string) *QuotaConfig {
	s.WorkloadTypes = v
	return s
}

func (s *QuotaConfig) Validate() error {
	if s.ACS != nil {
		if err := s.ACS.Validate(); err != nil {
			return err
		}
	}
	if s.EniCacheConfig != nil {
		if err := s.EniCacheConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OversoldUsageConfig != nil {
		if err := s.OversoldUsageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceSpecs != nil {
		for _, item := range s.ResourceSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SandboxCacheConfig != nil {
		if err := s.SandboxCacheConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SelfQuotaPreemptionConfig != nil {
		if err := s.SelfQuotaPreemptionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SubQuotaPreemptionConfig != nil {
		if err := s.SubQuotaPreemptionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}
