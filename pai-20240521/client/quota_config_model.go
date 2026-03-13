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
	SetEnablePreemptSubquotaWorkloads(v bool) *QuotaConfig
	GetEnablePreemptSubquotaWorkloads() *bool
	SetResourceSpecs(v []*WorkspaceSpecs) *QuotaConfig
	GetResourceSpecs() []*WorkspaceSpecs
	SetSupportGPUDrivers(v []*string) *QuotaConfig
	GetSupportGPUDrivers() []*string
	SetSupportRDMA(v bool) *QuotaConfig
	GetSupportRDMA() *bool
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
	DefaultGPUDriver               *string           `json:"DefaultGPUDriver,omitempty" xml:"DefaultGPUDriver,omitempty"`
	EnablePreemptSubquotaWorkloads *bool             `json:"EnablePreemptSubquotaWorkloads,omitempty" xml:"EnablePreemptSubquotaWorkloads,omitempty"`
	ResourceSpecs                  []*WorkspaceSpecs `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	SupportGPUDrivers              []*string         `json:"SupportGPUDrivers,omitempty" xml:"SupportGPUDrivers,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SupportRDMA *bool    `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
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

func (s *QuotaConfig) GetEnablePreemptSubquotaWorkloads() *bool {
	return s.EnablePreemptSubquotaWorkloads
}

func (s *QuotaConfig) GetResourceSpecs() []*WorkspaceSpecs {
	return s.ResourceSpecs
}

func (s *QuotaConfig) GetSupportGPUDrivers() []*string {
	return s.SupportGPUDrivers
}

func (s *QuotaConfig) GetSupportRDMA() *bool {
	return s.SupportRDMA
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

func (s *QuotaConfig) SetEnablePreemptSubquotaWorkloads(v bool) *QuotaConfig {
	s.EnablePreemptSubquotaWorkloads = &v
	return s
}

func (s *QuotaConfig) SetResourceSpecs(v []*WorkspaceSpecs) *QuotaConfig {
	s.ResourceSpecs = v
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
	if s.ResourceSpecs != nil {
		for _, item := range s.ResourceSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}
