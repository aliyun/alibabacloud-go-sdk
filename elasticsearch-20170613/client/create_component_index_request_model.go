// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMeta(v map[string]interface{}) *CreateComponentIndexRequest
	GetMeta() map[string]interface{}
	SetTemplate(v *CreateComponentIndexRequestTemplate) *CreateComponentIndexRequest
	GetTemplate() *CreateComponentIndexRequestTemplate
}

type CreateComponentIndexRequest struct {
	// example:
	//
	// {       "description": "set number of shards to one"   }
	Meta     map[string]interface{}               `json:"_meta,omitempty" xml:"_meta,omitempty"`
	Template *CreateComponentIndexRequestTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s CreateComponentIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateComponentIndexRequest) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *CreateComponentIndexRequest) GetTemplate() *CreateComponentIndexRequestTemplate {
	return s.Template
}

func (s *CreateComponentIndexRequest) SetMeta(v map[string]interface{}) *CreateComponentIndexRequest {
	s.Meta = v
	return s
}

func (s *CreateComponentIndexRequest) SetTemplate(v *CreateComponentIndexRequestTemplate) *CreateComponentIndexRequest {
	s.Template = v
	return s
}

func (s *CreateComponentIndexRequest) Validate() error {
	return dara.Validate(s)
}

type CreateComponentIndexRequestTemplate struct {
	// example:
	//
	// {}
	Aliases map[string]interface{} `json:"aliases,omitempty" xml:"aliases,omitempty"`
	// example:
	//
	// { 			"properties": { 				"@timestamp": { 					"type": "date" 				} 			} 		}
	Mappings map[string]interface{} `json:"mappings,omitempty" xml:"mappings,omitempty"`
	// example:
	//
	// { 			"index.number_of_replicas": 0 		}
	Settings map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s CreateComponentIndexRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentIndexRequestTemplate) GoString() string {
	return s.String()
}

func (s *CreateComponentIndexRequestTemplate) GetAliases() map[string]interface{} {
	return s.Aliases
}

func (s *CreateComponentIndexRequestTemplate) GetMappings() map[string]interface{} {
	return s.Mappings
}

func (s *CreateComponentIndexRequestTemplate) GetSettings() map[string]interface{} {
	return s.Settings
}

func (s *CreateComponentIndexRequestTemplate) SetAliases(v map[string]interface{}) *CreateComponentIndexRequestTemplate {
	s.Aliases = v
	return s
}

func (s *CreateComponentIndexRequestTemplate) SetMappings(v map[string]interface{}) *CreateComponentIndexRequestTemplate {
	s.Mappings = v
	return s
}

func (s *CreateComponentIndexRequestTemplate) SetSettings(v map[string]interface{}) *CreateComponentIndexRequestTemplate {
	s.Settings = v
	return s
}

func (s *CreateComponentIndexRequestTemplate) Validate() error {
	return dara.Validate(s)
}
