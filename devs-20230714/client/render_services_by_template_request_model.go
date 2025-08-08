// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenderServicesByTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v map[string]interface{}) *RenderServicesByTemplateRequest
	GetParameters() map[string]interface{}
	SetProjectName(v string) *RenderServicesByTemplateRequest
	GetProjectName() *string
	SetServiceNameChanges(v map[string]*string) *RenderServicesByTemplateRequest
	GetServiceNameChanges() map[string]*string
	SetTemplateName(v string) *RenderServicesByTemplateRequest
	GetTemplateName() *string
	SetVariableValues(v *TemplateVariableValueMap) *RenderServicesByTemplateRequest
	GetVariableValues() *TemplateVariableValueMap
}

type RenderServicesByTemplateRequest struct {
	// example:
	//
	// {"region":"cn-hangzhou"}
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// my-project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
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

func (s RenderServicesByTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s RenderServicesByTemplateRequest) GoString() string {
	return s.String()
}

func (s *RenderServicesByTemplateRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *RenderServicesByTemplateRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *RenderServicesByTemplateRequest) GetServiceNameChanges() map[string]*string {
	return s.ServiceNameChanges
}

func (s *RenderServicesByTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *RenderServicesByTemplateRequest) GetVariableValues() *TemplateVariableValueMap {
	return s.VariableValues
}

func (s *RenderServicesByTemplateRequest) SetParameters(v map[string]interface{}) *RenderServicesByTemplateRequest {
	s.Parameters = v
	return s
}

func (s *RenderServicesByTemplateRequest) SetProjectName(v string) *RenderServicesByTemplateRequest {
	s.ProjectName = &v
	return s
}

func (s *RenderServicesByTemplateRequest) SetServiceNameChanges(v map[string]*string) *RenderServicesByTemplateRequest {
	s.ServiceNameChanges = v
	return s
}

func (s *RenderServicesByTemplateRequest) SetTemplateName(v string) *RenderServicesByTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *RenderServicesByTemplateRequest) SetVariableValues(v *TemplateVariableValueMap) *RenderServicesByTemplateRequest {
	s.VariableValues = v
	return s
}

func (s *RenderServicesByTemplateRequest) Validate() error {
	return dara.Validate(s)
}
