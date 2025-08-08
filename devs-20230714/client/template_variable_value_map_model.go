// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateVariableValueMap interface {
	dara.Model
	String() string
	GoString() string
	SetServices(v map[string]map[string]interface{}) *TemplateVariableValueMap
	GetServices() map[string]map[string]interface{}
	SetShared(v map[string]interface{}) *TemplateVariableValueMap
	GetShared() map[string]interface{}
}

type TemplateVariableValueMap struct {
	Services map[string]map[string]interface{} `json:"services,omitempty" xml:"services,omitempty"`
	// example:
	//
	// {"TEST_KEY":"new_value"}
	Shared map[string]interface{} `json:"shared,omitempty" xml:"shared,omitempty"`
}

func (s TemplateVariableValueMap) String() string {
	return dara.Prettify(s)
}

func (s TemplateVariableValueMap) GoString() string {
	return s.String()
}

func (s *TemplateVariableValueMap) GetServices() map[string]map[string]interface{} {
	return s.Services
}

func (s *TemplateVariableValueMap) GetShared() map[string]interface{} {
	return s.Shared
}

func (s *TemplateVariableValueMap) SetServices(v map[string]map[string]interface{}) *TemplateVariableValueMap {
	s.Services = v
	return s
}

func (s *TemplateVariableValueMap) SetShared(v map[string]interface{}) *TemplateVariableValueMap {
	s.Shared = v
	return s
}

func (s *TemplateVariableValueMap) Validate() error {
	return dara.Validate(s)
}
