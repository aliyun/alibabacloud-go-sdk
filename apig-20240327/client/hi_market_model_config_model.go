// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketModelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetModelAPIConfig(v *HiMarketModelConfigModelAPIConfig) *HiMarketModelConfig
	GetModelAPIConfig() *HiMarketModelConfigModelAPIConfig
}

type HiMarketModelConfig struct {
	ModelAPIConfig *HiMarketModelConfigModelAPIConfig `json:"modelAPIConfig,omitempty" xml:"modelAPIConfig,omitempty" type:"Struct"`
}

func (s HiMarketModelConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketModelConfig) GoString() string {
	return s.String()
}

func (s *HiMarketModelConfig) GetModelAPIConfig() *HiMarketModelConfigModelAPIConfig {
	return s.ModelAPIConfig
}

func (s *HiMarketModelConfig) SetModelAPIConfig(v *HiMarketModelConfigModelAPIConfig) *HiMarketModelConfig {
	s.ModelAPIConfig = v
	return s
}

func (s *HiMarketModelConfig) Validate() error {
	if s.ModelAPIConfig != nil {
		if err := s.ModelAPIConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HiMarketModelConfigModelAPIConfig struct {
	AiProtocols   []*string            `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	ModelCategory *string              `json:"modelCategory,omitempty" xml:"modelCategory,omitempty"`
	Routes        []*HiMarketHttpRoute `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s HiMarketModelConfigModelAPIConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketModelConfigModelAPIConfig) GoString() string {
	return s.String()
}

func (s *HiMarketModelConfigModelAPIConfig) GetAiProtocols() []*string {
	return s.AiProtocols
}

func (s *HiMarketModelConfigModelAPIConfig) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *HiMarketModelConfigModelAPIConfig) GetRoutes() []*HiMarketHttpRoute {
	return s.Routes
}

func (s *HiMarketModelConfigModelAPIConfig) SetAiProtocols(v []*string) *HiMarketModelConfigModelAPIConfig {
	s.AiProtocols = v
	return s
}

func (s *HiMarketModelConfigModelAPIConfig) SetModelCategory(v string) *HiMarketModelConfigModelAPIConfig {
	s.ModelCategory = &v
	return s
}

func (s *HiMarketModelConfigModelAPIConfig) SetRoutes(v []*HiMarketHttpRoute) *HiMarketModelConfigModelAPIConfig {
	s.Routes = v
	return s
}

func (s *HiMarketModelConfigModelAPIConfig) Validate() error {
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
