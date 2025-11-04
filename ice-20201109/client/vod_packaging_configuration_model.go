// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVodPackagingConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationName(v string) *VodPackagingConfiguration
	GetConfigurationName() *string
	SetCreateTime(v string) *VodPackagingConfiguration
	GetCreateTime() *string
	SetDescription(v string) *VodPackagingConfiguration
	GetDescription() *string
	SetGroupName(v string) *VodPackagingConfiguration
	GetGroupName() *string
	SetPackageConfig(v *VodPackagingConfig) *VodPackagingConfiguration
	GetPackageConfig() *VodPackagingConfig
	SetProtocol(v string) *VodPackagingConfiguration
	GetProtocol() *string
}

type VodPackagingConfiguration struct {
	ConfigurationName *string             `json:"ConfigurationName,omitempty" xml:"ConfigurationName,omitempty"`
	CreateTime        *string             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string             `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName         *string             `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	PackageConfig     *VodPackagingConfig `json:"PackageConfig,omitempty" xml:"PackageConfig,omitempty"`
	Protocol          *string             `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s VodPackagingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s VodPackagingConfiguration) GoString() string {
	return s.String()
}

func (s *VodPackagingConfiguration) GetConfigurationName() *string {
	return s.ConfigurationName
}

func (s *VodPackagingConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *VodPackagingConfiguration) GetDescription() *string {
	return s.Description
}

func (s *VodPackagingConfiguration) GetGroupName() *string {
	return s.GroupName
}

func (s *VodPackagingConfiguration) GetPackageConfig() *VodPackagingConfig {
	return s.PackageConfig
}

func (s *VodPackagingConfiguration) GetProtocol() *string {
	return s.Protocol
}

func (s *VodPackagingConfiguration) SetConfigurationName(v string) *VodPackagingConfiguration {
	s.ConfigurationName = &v
	return s
}

func (s *VodPackagingConfiguration) SetCreateTime(v string) *VodPackagingConfiguration {
	s.CreateTime = &v
	return s
}

func (s *VodPackagingConfiguration) SetDescription(v string) *VodPackagingConfiguration {
	s.Description = &v
	return s
}

func (s *VodPackagingConfiguration) SetGroupName(v string) *VodPackagingConfiguration {
	s.GroupName = &v
	return s
}

func (s *VodPackagingConfiguration) SetPackageConfig(v *VodPackagingConfig) *VodPackagingConfiguration {
	s.PackageConfig = v
	return s
}

func (s *VodPackagingConfiguration) SetProtocol(v string) *VodPackagingConfiguration {
	s.Protocol = &v
	return s
}

func (s *VodPackagingConfiguration) Validate() error {
	if s.PackageConfig != nil {
		if err := s.PackageConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
