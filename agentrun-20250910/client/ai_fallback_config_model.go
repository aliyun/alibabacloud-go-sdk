// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiFallbackConfig interface {
	dara.Model
	String() string
	GoString() string
	SetServiceConfigs(v []*AiFallbackServiceConfig) *AiFallbackConfig
	GetServiceConfigs() []*AiFallbackServiceConfig
}

type AiFallbackConfig struct {
	ServiceConfigs []*AiFallbackServiceConfig `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s AiFallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s AiFallbackConfig) GoString() string {
	return s.String()
}

func (s *AiFallbackConfig) GetServiceConfigs() []*AiFallbackServiceConfig {
	return s.ServiceConfigs
}

func (s *AiFallbackConfig) SetServiceConfigs(v []*AiFallbackServiceConfig) *AiFallbackConfig {
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
