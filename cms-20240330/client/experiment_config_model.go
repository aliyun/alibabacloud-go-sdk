// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExperimentConfig interface {
  dara.Model
  String() string
  GoString() string
  SetLabel(v string) *ExperimentConfig
  GetLabel() *string 
  SetModelName(v string) *ExperimentConfig
  GetModelName() *string 
  SetModelParameters(v *ModelParameters) *ExperimentConfig
  GetModelParameters() *ModelParameters 
  SetModelProvider(v string) *ExperimentConfig
  GetModelProvider() *string 
  SetName(v string) *ExperimentConfig
  GetName() *string 
  SetPromptTemplate(v []*PromptTemplateItem) *ExperimentConfig
  GetPromptTemplate() []*PromptTemplateItem 
}

type ExperimentConfig struct {
  Label *string `json:"label,omitempty" xml:"label,omitempty"`
  ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
  ModelParameters *ModelParameters `json:"modelParameters,omitempty" xml:"modelParameters,omitempty"`
  ModelProvider *string `json:"modelProvider,omitempty" xml:"modelProvider,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  PromptTemplate []*PromptTemplateItem `json:"promptTemplate,omitempty" xml:"promptTemplate,omitempty" type:"Repeated"`
}

func (s ExperimentConfig) String() string {
  return dara.Prettify(s)
}

func (s ExperimentConfig) GoString() string {
  return s.String()
}

func (s *ExperimentConfig) GetLabel() *string  {
  return s.Label
}

func (s *ExperimentConfig) GetModelName() *string  {
  return s.ModelName
}

func (s *ExperimentConfig) GetModelParameters() *ModelParameters  {
  return s.ModelParameters
}

func (s *ExperimentConfig) GetModelProvider() *string  {
  return s.ModelProvider
}

func (s *ExperimentConfig) GetName() *string  {
  return s.Name
}

func (s *ExperimentConfig) GetPromptTemplate() []*PromptTemplateItem  {
  return s.PromptTemplate
}

func (s *ExperimentConfig) SetLabel(v string) *ExperimentConfig {
  s.Label = &v
  return s
}

func (s *ExperimentConfig) SetModelName(v string) *ExperimentConfig {
  s.ModelName = &v
  return s
}

func (s *ExperimentConfig) SetModelParameters(v *ModelParameters) *ExperimentConfig {
  s.ModelParameters = v
  return s
}

func (s *ExperimentConfig) SetModelProvider(v string) *ExperimentConfig {
  s.ModelProvider = &v
  return s
}

func (s *ExperimentConfig) SetName(v string) *ExperimentConfig {
  s.Name = &v
  return s
}

func (s *ExperimentConfig) SetPromptTemplate(v []*PromptTemplateItem) *ExperimentConfig {
  s.PromptTemplate = v
  return s
}

func (s *ExperimentConfig) Validate() error {
  if s.ModelParameters != nil {
    if err := s.ModelParameters.Validate(); err != nil {
      return err
    }
  }
  if s.PromptTemplate != nil {
    for _, item := range s.PromptTemplate {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

