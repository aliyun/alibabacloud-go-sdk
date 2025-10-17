// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigurationOverrides interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurations(v []*ConfigurationOverridesConfigurations) *ConfigurationOverrides
	GetConfigurations() []*ConfigurationOverridesConfigurations
}

type ConfigurationOverrides struct {
	Configurations []*ConfigurationOverridesConfigurations `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
}

func (s ConfigurationOverrides) String() string {
	return dara.Prettify(s)
}

func (s ConfigurationOverrides) GoString() string {
	return s.String()
}

func (s *ConfigurationOverrides) GetConfigurations() []*ConfigurationOverridesConfigurations {
	return s.Configurations
}

func (s *ConfigurationOverrides) SetConfigurations(v []*ConfigurationOverridesConfigurations) *ConfigurationOverrides {
	s.Configurations = v
	return s
}

func (s *ConfigurationOverrides) Validate() error {
	if s.Configurations != nil {
		for _, item := range s.Configurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ConfigurationOverridesConfigurations struct {
	ConfigFileName  *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	ConfigItemKey   *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s ConfigurationOverridesConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ConfigurationOverridesConfigurations) GoString() string {
	return s.String()
}

func (s *ConfigurationOverridesConfigurations) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *ConfigurationOverridesConfigurations) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *ConfigurationOverridesConfigurations) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *ConfigurationOverridesConfigurations) SetConfigFileName(v string) *ConfigurationOverridesConfigurations {
	s.ConfigFileName = &v
	return s
}

func (s *ConfigurationOverridesConfigurations) SetConfigItemKey(v string) *ConfigurationOverridesConfigurations {
	s.ConfigItemKey = &v
	return s
}

func (s *ConfigurationOverridesConfigurations) SetConfigItemValue(v string) *ConfigurationOverridesConfigurations {
	s.ConfigItemValue = &v
	return s
}

func (s *ConfigurationOverridesConfigurations) Validate() error {
	return dara.Validate(s)
}
