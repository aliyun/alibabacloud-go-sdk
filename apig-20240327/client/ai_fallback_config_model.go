// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiFallbackConfig interface {
	dara.Model
	String() string
	GoString() string
	SetOnlyRedirectUpstreamCode(v bool) *AiFallbackConfig
	GetOnlyRedirectUpstreamCode() *bool
	SetServiceConfigs(v []*AiFallbackConfigServiceConfigs) *AiFallbackConfig
	GetServiceConfigs() []*AiFallbackConfigServiceConfigs
}

type AiFallbackConfig struct {
	OnlyRedirectUpstreamCode *bool                             `json:"onlyRedirectUpstreamCode,omitempty" xml:"onlyRedirectUpstreamCode,omitempty"`
	ServiceConfigs           []*AiFallbackConfigServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s AiFallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s AiFallbackConfig) GoString() string {
	return s.String()
}

func (s *AiFallbackConfig) GetOnlyRedirectUpstreamCode() *bool {
	return s.OnlyRedirectUpstreamCode
}

func (s *AiFallbackConfig) GetServiceConfigs() []*AiFallbackConfigServiceConfigs {
	return s.ServiceConfigs
}

func (s *AiFallbackConfig) SetOnlyRedirectUpstreamCode(v bool) *AiFallbackConfig {
	s.OnlyRedirectUpstreamCode = &v
	return s
}

func (s *AiFallbackConfig) SetServiceConfigs(v []*AiFallbackConfigServiceConfigs) *AiFallbackConfig {
	s.ServiceConfigs = v
	return s
}

func (s *AiFallbackConfig) Validate() error {
	if s.ServiceConfigs != nil {
		for _, item := range s.ServiceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AiFallbackConfigServiceConfigs struct {
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
	PassThroughModelName *bool   `json:"passThroughModelName,omitempty" xml:"passThroughModelName,omitempty"`
	ServiceId            *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TargetModelName      *string `json:"targetModelName,omitempty" xml:"targetModelName,omitempty"`
}

func (s AiFallbackConfigServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s AiFallbackConfigServiceConfigs) GoString() string {
	return s.String()
}

func (s *AiFallbackConfigServiceConfigs) GetName() *string {
	return s.Name
}

func (s *AiFallbackConfigServiceConfigs) GetPassThroughModelName() *bool {
	return s.PassThroughModelName
}

func (s *AiFallbackConfigServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *AiFallbackConfigServiceConfigs) GetTargetModelName() *string {
	return s.TargetModelName
}

func (s *AiFallbackConfigServiceConfigs) SetName(v string) *AiFallbackConfigServiceConfigs {
	s.Name = &v
	return s
}

func (s *AiFallbackConfigServiceConfigs) SetPassThroughModelName(v bool) *AiFallbackConfigServiceConfigs {
	s.PassThroughModelName = &v
	return s
}

func (s *AiFallbackConfigServiceConfigs) SetServiceId(v string) *AiFallbackConfigServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *AiFallbackConfigServiceConfigs) SetTargetModelName(v string) *AiFallbackConfigServiceConfigs {
	s.TargetModelName = &v
	return s
}

func (s *AiFallbackConfigServiceConfigs) Validate() error {
	return dara.Validate(s)
}
