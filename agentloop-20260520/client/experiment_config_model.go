// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExperimentConfig interface {
  dara.Model
  String() string
  GoString() string
  SetEndpointConnectorId(v string) *ExperimentConfig
  GetEndpointConnectorId() *string 
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
  SetRequestBodyTemplate(v string) *ExperimentConfig
  GetRequestBodyTemplate() *string 
  SetRequestHeaderTemplate(v string) *ExperimentConfig
  GetRequestHeaderTemplate() *string 
  SetRequestMethod(v string) *ExperimentConfig
  GetRequestMethod() *string 
}

type ExperimentConfig struct {
  // The endpoint connector ID. This parameter is required in user/agent mode.
  EndpointConnectorId *string `json:"endpointConnectorId,omitempty" xml:"endpointConnectorId,omitempty"`
  // The experiment label (A/B/C/D/E).
  Label *string `json:"label,omitempty" xml:"label,omitempty"`
  // The model name. You can set this parameter to agent in agent scenarios.
  ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
  // The model parameters.
  ModelParameters *ModelParameters `json:"modelParameters,omitempty" xml:"modelParameters,omitempty"`
  // The model provider. Valid values: dashscope (default), user, and agent.
  ModelProvider *string `json:"modelProvider,omitempty" xml:"modelProvider,omitempty"`
  // The experiment name.
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // The prompt message template. Supports {{variable name}} placeholders.
  PromptTemplate []*PromptTemplateItem `json:"promptTemplate,omitempty" xml:"promptTemplate,omitempty" type:"Repeated"`
  // The agent request body template. This parameter is required in agent mode. Supports {{variable name}} placeholders.
  RequestBodyTemplate *string `json:"requestBodyTemplate,omitempty" xml:"requestBodyTemplate,omitempty"`
  // The agent request header template. This parameter is optional in agent mode. Supports {{variable name}} placeholders.
  // 
  // example:
  // 
  // {"Content-Type: application/json" }
  RequestHeaderTemplate *string `json:"requestHeaderTemplate,omitempty" xml:"requestHeaderTemplate,omitempty"`
  // The agent request method. Valid values: POST (default) and GET.
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty"`
}

func (s ExperimentConfig) String() string {
  return dara.Prettify(s)
}

func (s ExperimentConfig) GoString() string {
  return s.String()
}

func (s *ExperimentConfig) GetEndpointConnectorId() *string  {
  return s.EndpointConnectorId
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

func (s *ExperimentConfig) GetRequestBodyTemplate() *string  {
  return s.RequestBodyTemplate
}

func (s *ExperimentConfig) GetRequestHeaderTemplate() *string  {
  return s.RequestHeaderTemplate
}

func (s *ExperimentConfig) GetRequestMethod() *string  {
  return s.RequestMethod
}

func (s *ExperimentConfig) SetEndpointConnectorId(v string) *ExperimentConfig {
  s.EndpointConnectorId = &v
  return s
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

func (s *ExperimentConfig) SetRequestBodyTemplate(v string) *ExperimentConfig {
  s.RequestBodyTemplate = &v
  return s
}

func (s *ExperimentConfig) SetRequestHeaderTemplate(v string) *ExperimentConfig {
  s.RequestHeaderTemplate = &v
  return s
}

func (s *ExperimentConfig) SetRequestMethod(v string) *ExperimentConfig {
  s.RequestMethod = &v
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

