// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateConfig interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v map[string]*string) *TemplateConfig
	GetParameters() map[string]*string
	SetServiceNameChanges(v map[string]*string) *TemplateConfig
	GetServiceNameChanges() map[string]*string
	SetTemplateName(v string) *TemplateConfig
	GetTemplateName() *string
	SetVariableValues(v *TemplateVariableValueMap) *TemplateConfig
	GetVariableValues() *TemplateVariableValueMap
}

type TemplateConfig struct {
	// example:
	//
	// {"region":"cn-hangzhou"}
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// {"svc1":"svc2"}
	ServiceNameChanges map[string]*string `json:"serviceNameChanges,omitempty" xml:"serviceNameChanges,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// start-springboot
	TemplateName   *string                   `json:"templateName,omitempty" xml:"templateName,omitempty"`
	VariableValues *TemplateVariableValueMap `json:"variableValues,omitempty" xml:"variableValues,omitempty"`
}

func (s TemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s TemplateConfig) GoString() string {
	return s.String()
}

func (s *TemplateConfig) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *TemplateConfig) GetServiceNameChanges() map[string]*string {
	return s.ServiceNameChanges
}

func (s *TemplateConfig) GetTemplateName() *string {
	return s.TemplateName
}

func (s *TemplateConfig) GetVariableValues() *TemplateVariableValueMap {
	return s.VariableValues
}

func (s *TemplateConfig) SetParameters(v map[string]*string) *TemplateConfig {
	s.Parameters = v
	return s
}

func (s *TemplateConfig) SetServiceNameChanges(v map[string]*string) *TemplateConfig {
	s.ServiceNameChanges = v
	return s
}

func (s *TemplateConfig) SetTemplateName(v string) *TemplateConfig {
	s.TemplateName = &v
	return s
}

func (s *TemplateConfig) SetVariableValues(v *TemplateVariableValueMap) *TemplateConfig {
	s.VariableValues = v
	return s
}

func (s *TemplateConfig) Validate() error {
	return dara.Validate(s)
}
