// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAiServiceConfig(v *AiServiceConfig) *ServiceConfig
	GetAiServiceConfig() *AiServiceConfig
	SetName(v string) *ServiceConfig
	GetName() *string
}

type ServiceConfig struct {
	AiServiceConfig *AiServiceConfig `json:"aiServiceConfig,omitempty" xml:"aiServiceConfig,omitempty"`
	Name            *string          `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s ServiceConfig) GoString() string {
	return s.String()
}

func (s *ServiceConfig) GetAiServiceConfig() *AiServiceConfig {
	return s.AiServiceConfig
}

func (s *ServiceConfig) GetName() *string {
	return s.Name
}

func (s *ServiceConfig) SetAiServiceConfig(v *AiServiceConfig) *ServiceConfig {
	s.AiServiceConfig = v
	return s
}

func (s *ServiceConfig) SetName(v string) *ServiceConfig {
	s.Name = &v
	return s
}

func (s *ServiceConfig) Validate() error {
	if s.AiServiceConfig != nil {
		if err := s.AiServiceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
