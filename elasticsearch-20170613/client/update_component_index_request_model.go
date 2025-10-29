// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMeta(v map[string]interface{}) *UpdateComponentIndexRequest
	GetMeta() map[string]interface{}
	SetTemplate(v *UpdateComponentIndexRequestTemplate) *UpdateComponentIndexRequest
	GetTemplate() *UpdateComponentIndexRequestTemplate
}

type UpdateComponentIndexRequest struct {
	// example:
	//
	// { "description": "set number of shards to one" }
	Meta     map[string]interface{}               `json:"_meta,omitempty" xml:"_meta,omitempty"`
	Template *UpdateComponentIndexRequestTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s UpdateComponentIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentIndexRequest) GoString() string {
	return s.String()
}

func (s *UpdateComponentIndexRequest) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *UpdateComponentIndexRequest) GetTemplate() *UpdateComponentIndexRequestTemplate {
	return s.Template
}

func (s *UpdateComponentIndexRequest) SetMeta(v map[string]interface{}) *UpdateComponentIndexRequest {
	s.Meta = v
	return s
}

func (s *UpdateComponentIndexRequest) SetTemplate(v *UpdateComponentIndexRequestTemplate) *UpdateComponentIndexRequest {
	s.Template = v
	return s
}

func (s *UpdateComponentIndexRequest) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComponentIndexRequestTemplate struct {
	// example:
	//
	// {}
	Aliases map[string]interface{} `json:"aliases,omitempty" xml:"aliases,omitempty"`
	// example:
	//
	// { "properties": { "@timestamp": { "type": "date" } } }
	Mappings map[string]interface{} `json:"mappings,omitempty" xml:"mappings,omitempty"`
	// example:
	//
	// { "index.number_of_replicas": 0 }
	Settings map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s UpdateComponentIndexRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentIndexRequestTemplate) GoString() string {
	return s.String()
}

func (s *UpdateComponentIndexRequestTemplate) GetAliases() map[string]interface{} {
	return s.Aliases
}

func (s *UpdateComponentIndexRequestTemplate) GetMappings() map[string]interface{} {
	return s.Mappings
}

func (s *UpdateComponentIndexRequestTemplate) GetSettings() map[string]interface{} {
	return s.Settings
}

func (s *UpdateComponentIndexRequestTemplate) SetAliases(v map[string]interface{}) *UpdateComponentIndexRequestTemplate {
	s.Aliases = v
	return s
}

func (s *UpdateComponentIndexRequestTemplate) SetMappings(v map[string]interface{}) *UpdateComponentIndexRequestTemplate {
	s.Mappings = v
	return s
}

func (s *UpdateComponentIndexRequestTemplate) SetSettings(v map[string]interface{}) *UpdateComponentIndexRequestTemplate {
	s.Settings = v
	return s
}

func (s *UpdateComponentIndexRequestTemplate) Validate() error {
	return dara.Validate(s)
}
