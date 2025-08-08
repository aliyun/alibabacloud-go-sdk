// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenderServicesByTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangedServiceNames(v map[string]*string) *RenderServicesByTemplateResponseBody
	GetChangedServiceNames() map[string]*string
	SetServices(v map[string]*ServiceConfig) *RenderServicesByTemplateResponseBody
	GetServices() map[string]*ServiceConfig
	SetVariables(v map[string]*Variable) *RenderServicesByTemplateResponseBody
	GetVariables() map[string]*Variable
}

type RenderServicesByTemplateResponseBody struct {
	ChangedServiceNames map[string]*string        `json:"changedServiceNames,omitempty" xml:"changedServiceNames,omitempty"`
	Services            map[string]*ServiceConfig `json:"services,omitempty" xml:"services,omitempty"`
	Variables           map[string]*Variable      `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s RenderServicesByTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenderServicesByTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *RenderServicesByTemplateResponseBody) GetChangedServiceNames() map[string]*string {
	return s.ChangedServiceNames
}

func (s *RenderServicesByTemplateResponseBody) GetServices() map[string]*ServiceConfig {
	return s.Services
}

func (s *RenderServicesByTemplateResponseBody) GetVariables() map[string]*Variable {
	return s.Variables
}

func (s *RenderServicesByTemplateResponseBody) SetChangedServiceNames(v map[string]*string) *RenderServicesByTemplateResponseBody {
	s.ChangedServiceNames = v
	return s
}

func (s *RenderServicesByTemplateResponseBody) SetServices(v map[string]*ServiceConfig) *RenderServicesByTemplateResponseBody {
	s.Services = v
	return s
}

func (s *RenderServicesByTemplateResponseBody) SetVariables(v map[string]*Variable) *RenderServicesByTemplateResponseBody {
	s.Variables = v
	return s
}

func (s *RenderServicesByTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
