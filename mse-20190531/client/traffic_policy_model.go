// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrafficPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerSettings(v *TrafficPolicyLoadBalancerSettings) *TrafficPolicy
	GetLoadBalancerSettings() *TrafficPolicyLoadBalancerSettings
	SetTlsSetting(v *TrafficPolicyTlsSetting) *TrafficPolicy
	GetTlsSetting() *TrafficPolicyTlsSetting
}

type TrafficPolicy struct {
	LoadBalancerSettings *TrafficPolicyLoadBalancerSettings `json:"LoadBalancerSettings,omitempty" xml:"LoadBalancerSettings,omitempty" type:"Struct"`
	TlsSetting           *TrafficPolicyTlsSetting           `json:"TlsSetting,omitempty" xml:"TlsSetting,omitempty" type:"Struct"`
}

func (s TrafficPolicy) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicy) GoString() string {
	return s.String()
}

func (s *TrafficPolicy) GetLoadBalancerSettings() *TrafficPolicyLoadBalancerSettings {
	return s.LoadBalancerSettings
}

func (s *TrafficPolicy) GetTlsSetting() *TrafficPolicyTlsSetting {
	return s.TlsSetting
}

func (s *TrafficPolicy) SetLoadBalancerSettings(v *TrafficPolicyLoadBalancerSettings) *TrafficPolicy {
	s.LoadBalancerSettings = v
	return s
}

func (s *TrafficPolicy) SetTlsSetting(v *TrafficPolicyTlsSetting) *TrafficPolicy {
	s.TlsSetting = v
	return s
}

func (s *TrafficPolicy) Validate() error {
	if s.LoadBalancerSettings != nil {
		if err := s.LoadBalancerSettings.Validate(); err != nil {
			return err
		}
	}
	if s.TlsSetting != nil {
		if err := s.TlsSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrafficPolicyLoadBalancerSettings struct {
	ConsistentHashLBConfig *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig `json:"ConsistentHashLBConfig,omitempty" xml:"ConsistentHashLBConfig,omitempty" type:"Struct"`
	// example:
	//
	// RANDOM
	LoadbalancerType *string `json:"LoadbalancerType,omitempty" xml:"LoadbalancerType,omitempty"`
	WarmupDuration   *int64  `json:"WarmupDuration,omitempty" xml:"WarmupDuration,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettings) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettings) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettings) GetConsistentHashLBConfig() *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	return s.ConsistentHashLBConfig
}

func (s *TrafficPolicyLoadBalancerSettings) GetLoadbalancerType() *string {
	return s.LoadbalancerType
}

func (s *TrafficPolicyLoadBalancerSettings) GetWarmupDuration() *int64 {
	return s.WarmupDuration
}

func (s *TrafficPolicyLoadBalancerSettings) SetConsistentHashLBConfig(v *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) *TrafficPolicyLoadBalancerSettings {
	s.ConsistentHashLBConfig = v
	return s
}

func (s *TrafficPolicyLoadBalancerSettings) SetLoadbalancerType(v string) *TrafficPolicyLoadBalancerSettings {
	s.LoadbalancerType = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettings) SetWarmupDuration(v int64) *TrafficPolicyLoadBalancerSettings {
	s.WarmupDuration = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettings) Validate() error {
	if s.ConsistentHashLBConfig != nil {
		if err := s.ConsistentHashLBConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig struct {
	ConsistentHashLBType *string                                                            `json:"ConsistentHashLBType,omitempty" xml:"ConsistentHashLBType,omitempty"`
	HttpCookie           *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie `json:"HttpCookie,omitempty" xml:"HttpCookie,omitempty" type:"Struct"`
	ParameterName        *string                                                            `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetConsistentHashLBType() *string {
	return s.ConsistentHashLBType
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetHttpCookie() *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	return s.HttpCookie
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetParameterName() *string {
	return s.ParameterName
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetConsistentHashLBType(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.ConsistentHashLBType = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetHttpCookie(v *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.HttpCookie = v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetParameterName(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.ParameterName = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) Validate() error {
	if s.HttpCookie != nil {
		if err := s.HttpCookie.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 0s
	TTL *string `json:"TTL,omitempty" xml:"TTL,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetName() *string {
	return s.Name
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetPath() *string {
	return s.Path
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetTTL() *string {
	return s.TTL
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetName(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Name = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetPath(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Path = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetTTL(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.TTL = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) Validate() error {
	return dara.Validate(s)
}

type TrafficPolicyTlsSetting struct {
	CaCertContent *string `json:"CaCertContent,omitempty" xml:"CaCertContent,omitempty"`
	CertId        *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	Sni           *string `json:"Sni,omitempty" xml:"Sni,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SIMPLE
	TlsMode *string `json:"TlsMode,omitempty" xml:"TlsMode,omitempty"`
}

func (s TrafficPolicyTlsSetting) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicyTlsSetting) GoString() string {
	return s.String()
}

func (s *TrafficPolicyTlsSetting) GetCaCertContent() *string {
	return s.CaCertContent
}

func (s *TrafficPolicyTlsSetting) GetCertId() *string {
	return s.CertId
}

func (s *TrafficPolicyTlsSetting) GetSni() *string {
	return s.Sni
}

func (s *TrafficPolicyTlsSetting) GetTlsMode() *string {
	return s.TlsMode
}

func (s *TrafficPolicyTlsSetting) SetCaCertContent(v string) *TrafficPolicyTlsSetting {
	s.CaCertContent = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetCertId(v string) *TrafficPolicyTlsSetting {
	s.CertId = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetSni(v string) *TrafficPolicyTlsSetting {
	s.Sni = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetTlsMode(v string) *TrafficPolicyTlsSetting {
	s.TlsMode = &v
	return s
}

func (s *TrafficPolicyTlsSetting) Validate() error {
	return dara.Validate(s)
}
