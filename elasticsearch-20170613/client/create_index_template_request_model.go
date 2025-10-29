// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateIndexTemplateRequest
	GetClientToken() *string
	SetDataStream(v bool) *CreateIndexTemplateRequest
	GetDataStream() *bool
	SetIlmPolicy(v string) *CreateIndexTemplateRequest
	GetIlmPolicy() *string
	SetIndexPatterns(v []*string) *CreateIndexTemplateRequest
	GetIndexPatterns() []*string
	SetIndexTemplate(v string) *CreateIndexTemplateRequest
	GetIndexTemplate() *string
	SetPriority(v int32) *CreateIndexTemplateRequest
	GetPriority() *int32
	SetTemplate(v *CreateIndexTemplateRequestTemplate) *CreateIndexTemplateRequest
	GetTemplate() *CreateIndexTemplateRequestTemplate
}

type CreateIndexTemplateRequest struct {
	// example:
	//
	// E1136AE9-4E49-4585-9358-6FDD2A6D****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DataStream *bool `json:"dataStream,omitempty" xml:"dataStream,omitempty"`
	// example:
	//
	// policy-1
	IlmPolicy *string `json:"ilmPolicy,omitempty" xml:"ilmPolicy,omitempty"`
	// This parameter is required.
	IndexPatterns []*string `json:"indexPatterns,omitempty" xml:"indexPatterns,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// index-template
	IndexTemplate *string `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	// example:
	//
	// 100
	Priority *int32                              `json:"priority,omitempty" xml:"priority,omitempty"`
	Template *CreateIndexTemplateRequestTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s CreateIndexTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIndexTemplateRequest) GetDataStream() *bool {
	return s.DataStream
}

func (s *CreateIndexTemplateRequest) GetIlmPolicy() *string {
	return s.IlmPolicy
}

func (s *CreateIndexTemplateRequest) GetIndexPatterns() []*string {
	return s.IndexPatterns
}

func (s *CreateIndexTemplateRequest) GetIndexTemplate() *string {
	return s.IndexTemplate
}

func (s *CreateIndexTemplateRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateIndexTemplateRequest) GetTemplate() *CreateIndexTemplateRequestTemplate {
	return s.Template
}

func (s *CreateIndexTemplateRequest) SetClientToken(v string) *CreateIndexTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIndexTemplateRequest) SetDataStream(v bool) *CreateIndexTemplateRequest {
	s.DataStream = &v
	return s
}

func (s *CreateIndexTemplateRequest) SetIlmPolicy(v string) *CreateIndexTemplateRequest {
	s.IlmPolicy = &v
	return s
}

func (s *CreateIndexTemplateRequest) SetIndexPatterns(v []*string) *CreateIndexTemplateRequest {
	s.IndexPatterns = v
	return s
}

func (s *CreateIndexTemplateRequest) SetIndexTemplate(v string) *CreateIndexTemplateRequest {
	s.IndexTemplate = &v
	return s
}

func (s *CreateIndexTemplateRequest) SetPriority(v int32) *CreateIndexTemplateRequest {
	s.Priority = &v
	return s
}

func (s *CreateIndexTemplateRequest) SetTemplate(v *CreateIndexTemplateRequestTemplate) *CreateIndexTemplateRequest {
	s.Template = v
	return s
}

func (s *CreateIndexTemplateRequest) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIndexTemplateRequestTemplate struct {
	// example:
	//
	// {"mydata": {}}
	Aliases *string `json:"aliases,omitempty" xml:"aliases,omitempty"`
	// example:
	//
	// {"properties": {"created_at": {"type": "date","format": "EEE MMM dd HH:mm:ss Z yyyy"},"host_name": {"type": "keyword"}}}
	Mappings *string `json:"mappings,omitempty" xml:"mappings,omitempty"`
	// example:
	//
	// {\"index.refresh_interval\":\"1s\"}
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s CreateIndexTemplateRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexTemplateRequestTemplate) GoString() string {
	return s.String()
}

func (s *CreateIndexTemplateRequestTemplate) GetAliases() *string {
	return s.Aliases
}

func (s *CreateIndexTemplateRequestTemplate) GetMappings() *string {
	return s.Mappings
}

func (s *CreateIndexTemplateRequestTemplate) GetSettings() *string {
	return s.Settings
}

func (s *CreateIndexTemplateRequestTemplate) SetAliases(v string) *CreateIndexTemplateRequestTemplate {
	s.Aliases = &v
	return s
}

func (s *CreateIndexTemplateRequestTemplate) SetMappings(v string) *CreateIndexTemplateRequestTemplate {
	s.Mappings = &v
	return s
}

func (s *CreateIndexTemplateRequestTemplate) SetSettings(v string) *CreateIndexTemplateRequestTemplate {
	s.Settings = &v
	return s
}

func (s *CreateIndexTemplateRequestTemplate) Validate() error {
	return dara.Validate(s)
}
