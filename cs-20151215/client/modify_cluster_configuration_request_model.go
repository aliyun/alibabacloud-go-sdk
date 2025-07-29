// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomizeConfig(v []*ModifyClusterConfigurationRequestCustomizeConfig) *ModifyClusterConfigurationRequest
	GetCustomizeConfig() []*ModifyClusterConfigurationRequestCustomizeConfig
}

type ModifyClusterConfigurationRequest struct {
	// The custom configurations.
	CustomizeConfig []*ModifyClusterConfigurationRequestCustomizeConfig `json:"customize_config,omitempty" xml:"customize_config,omitempty" type:"Repeated"`
}

func (s ModifyClusterConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationRequest) GetCustomizeConfig() []*ModifyClusterConfigurationRequestCustomizeConfig {
	return s.CustomizeConfig
}

func (s *ModifyClusterConfigurationRequest) SetCustomizeConfig(v []*ModifyClusterConfigurationRequestCustomizeConfig) *ModifyClusterConfigurationRequest {
	s.CustomizeConfig = v
	return s
}

func (s *ModifyClusterConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterConfigurationRequestCustomizeConfig struct {
	// The custom configurations.
	Configs []*ModifyClusterConfigurationRequestCustomizeConfigConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// The name of the component.
	//
	// example:
	//
	// kube-apiserver
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModifyClusterConfigurationRequestCustomizeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterConfigurationRequestCustomizeConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationRequestCustomizeConfig) GetConfigs() []*ModifyClusterConfigurationRequestCustomizeConfigConfigs {
	return s.Configs
}

func (s *ModifyClusterConfigurationRequestCustomizeConfig) GetName() *string {
	return s.Name
}

func (s *ModifyClusterConfigurationRequestCustomizeConfig) SetConfigs(v []*ModifyClusterConfigurationRequestCustomizeConfigConfigs) *ModifyClusterConfigurationRequestCustomizeConfig {
	s.Configs = v
	return s
}

func (s *ModifyClusterConfigurationRequestCustomizeConfig) SetName(v string) *ModifyClusterConfigurationRequestCustomizeConfig {
	s.Name = &v
	return s
}

func (s *ModifyClusterConfigurationRequestCustomizeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterConfigurationRequestCustomizeConfigConfigs struct {
	// The name of the configuration item.
	//
	// example:
	//
	// MaxRequestsInflight
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ModifyClusterConfigurationRequestCustomizeConfigConfigs) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterConfigurationRequestCustomizeConfigConfigs) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationRequestCustomizeConfigConfigs) GetKey() *string {
	return s.Key
}

func (s *ModifyClusterConfigurationRequestCustomizeConfigConfigs) GetValue() *string {
	return s.Value
}

func (s *ModifyClusterConfigurationRequestCustomizeConfigConfigs) SetKey(v string) *ModifyClusterConfigurationRequestCustomizeConfigConfigs {
	s.Key = &v
	return s
}

func (s *ModifyClusterConfigurationRequestCustomizeConfigConfigs) SetValue(v string) *ModifyClusterConfigurationRequestCustomizeConfigConfigs {
	s.Value = &v
	return s
}

func (s *ModifyClusterConfigurationRequestCustomizeConfigConfigs) Validate() error {
	return dara.Validate(s)
}
