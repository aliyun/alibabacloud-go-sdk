// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInfoConfig interface {
	dara.Model
	String() string
	GoString() string
	SetModelFeatures(v *ModelFeatures) *ModelInfoConfig
	GetModelFeatures() *ModelFeatures
	SetModelName(v string) *ModelInfoConfig
	GetModelName() *string
	SetModelParameterRules(v []*ModelParameterRule) *ModelInfoConfig
	GetModelParameterRules() []*ModelParameterRule
	SetModelProperties(v *ModelProperties) *ModelInfoConfig
	GetModelProperties() *ModelProperties
}

type ModelInfoConfig struct {
	ModelFeatures       *ModelFeatures        `json:"modelFeatures,omitempty" xml:"modelFeatures,omitempty"`
	ModelName           *string               `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ModelParameterRules []*ModelParameterRule `json:"modelParameterRules,omitempty" xml:"modelParameterRules,omitempty" type:"Repeated"`
	ModelProperties     *ModelProperties      `json:"modelProperties,omitempty" xml:"modelProperties,omitempty"`
}

func (s ModelInfoConfig) String() string {
	return dara.Prettify(s)
}

func (s ModelInfoConfig) GoString() string {
	return s.String()
}

func (s *ModelInfoConfig) GetModelFeatures() *ModelFeatures {
	return s.ModelFeatures
}

func (s *ModelInfoConfig) GetModelName() *string {
	return s.ModelName
}

func (s *ModelInfoConfig) GetModelParameterRules() []*ModelParameterRule {
	return s.ModelParameterRules
}

func (s *ModelInfoConfig) GetModelProperties() *ModelProperties {
	return s.ModelProperties
}

func (s *ModelInfoConfig) SetModelFeatures(v *ModelFeatures) *ModelInfoConfig {
	s.ModelFeatures = v
	return s
}

func (s *ModelInfoConfig) SetModelName(v string) *ModelInfoConfig {
	s.ModelName = &v
	return s
}

func (s *ModelInfoConfig) SetModelParameterRules(v []*ModelParameterRule) *ModelInfoConfig {
	s.ModelParameterRules = v
	return s
}

func (s *ModelInfoConfig) SetModelProperties(v *ModelProperties) *ModelInfoConfig {
	s.ModelProperties = v
	return s
}

func (s *ModelInfoConfig) Validate() error {
	if s.ModelFeatures != nil {
		if err := s.ModelFeatures.Validate(); err != nil {
			return err
		}
	}
	if s.ModelParameterRules != nil {
		for _, item := range s.ModelParameterRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModelProperties != nil {
		if err := s.ModelProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}
