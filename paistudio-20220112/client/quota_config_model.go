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
	SetOversoldUsageInfo(v *OversoldUsageConfig) *QuotaConfig
	GetOversoldUsageInfo() *OversoldUsageConfig
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
}

type QuotaConfig struct {
	ACS *ACS `json:"ACS,omitempty" xml:"ACS,omitempty"`
	// example:
	//
	// ceeb37xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 470.199.02
	DefaultGPUDriver               *string                    `json:"DefaultGPUDriver,omitempty" xml:"DefaultGPUDriver,omitempty"`
	EnableGPUShare                 *bool                      `json:"EnableGPUShare,omitempty" xml:"EnableGPUShare,omitempty"`
	EnablePreemptSubquotaWorkloads *bool                      `json:"EnablePreemptSubquotaWorkloads,omitempty" xml:"EnablePreemptSubquotaWorkloads,omitempty"`
	EnableSelfQuotaPreemption      *bool                      `json:"EnableSelfQuotaPreemption,omitempty" xml:"EnableSelfQuotaPreemption,omitempty"`
	EnableSubQuotaPreemption       *bool                      `json:"EnableSubQuotaPreemption,omitempty" xml:"EnableSubQuotaPreemption,omitempty"`
	EniCacheConfig                 *EniCacheConfig            `json:"EniCacheConfig,omitempty" xml:"EniCacheConfig,omitempty"`
	OversoldUsageInfo              *OversoldUsageConfig       `json:"OversoldUsageInfo,omitempty" xml:"OversoldUsageInfo,omitempty"`
	ResourceSpecs                  []*WorkspaceSpecs          `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	SandboxCacheConfig             *SandboxCacheConfig        `json:"SandboxCacheConfig,omitempty" xml:"SandboxCacheConfig,omitempty"`
	SelfQuotaPreemptionConfig      *SelfQuotaPreemptionConfig `json:"SelfQuotaPreemptionConfig,omitempty" xml:"SelfQuotaPreemptionConfig,omitempty"`
	SubQuotaPreemptionConfig       *SubQuotaPreemptionConfig  `json:"SubQuotaPreemptionConfig,omitempty" xml:"SubQuotaPreemptionConfig,omitempty"`
	SupportGPUDrivers              []*string                  `json:"SupportGPUDrivers,omitempty" xml:"SupportGPUDrivers,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SupportRDMA *bool    `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
	UseCase     *string  `json:"UseCase,omitempty" xml:"UseCase,omitempty"`
	UserVpc     *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
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

func (s *QuotaConfig) GetOversoldUsageInfo() *OversoldUsageConfig {
	return s.OversoldUsageInfo
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

func (s *QuotaConfig) SetACS(v *ACS) *QuotaConfig {
	s.ACS = v
	return s
}

func (s *QuotaConfig) SetClusterId(v string) *QuotaConfig {
	s.ClusterId = &v
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

func (s *QuotaConfig) SetOversoldUsageInfo(v *OversoldUsageConfig) *QuotaConfig {
	s.OversoldUsageInfo = v
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
	if s.OversoldUsageInfo != nil {
		if err := s.OversoldUsageInfo.Validate(); err != nil {
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
