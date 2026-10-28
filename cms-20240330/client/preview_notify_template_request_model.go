// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewNotifyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *PreviewNotifyTemplateRequest
	GetData() interface{}
	SetTemplates(v []*PreviewNotifyTemplateRequestTemplates) *PreviewNotifyTemplateRequest
	GetTemplates() []*PreviewNotifyTemplateRequestTemplates
}

type PreviewNotifyTemplateRequest struct {
	// The rendering context data in key-value pair format. This data source is used to populate Go template placeholders in the templates, such as {{.incident.title}}.
	//
	// example:
	//
	// {"incident":{"title":"High CPU utilization","severity":"P1"}}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// The list of templates to be rendered. Each item includes a template identifier and the template content.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"key":"title","template":"CloudMonitor alert: {{.incident.title}}"},{"key":"content","template":">- Severity: {{.incident.severity}}"}]
	Templates []*PreviewNotifyTemplateRequestTemplates `json:"templates,omitempty" xml:"templates,omitempty" type:"Repeated"`
}

func (s PreviewNotifyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewNotifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *PreviewNotifyTemplateRequest) GetData() interface{} {
	return s.Data
}

func (s *PreviewNotifyTemplateRequest) GetTemplates() []*PreviewNotifyTemplateRequestTemplates {
	return s.Templates
}

func (s *PreviewNotifyTemplateRequest) SetData(v interface{}) *PreviewNotifyTemplateRequest {
	s.Data = v
	return s
}

func (s *PreviewNotifyTemplateRequest) SetTemplates(v []*PreviewNotifyTemplateRequestTemplates) *PreviewNotifyTemplateRequest {
	s.Templates = v
	return s
}

func (s *PreviewNotifyTemplateRequest) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewNotifyTemplateRequestTemplates struct {
	// The template identifier. This identifier is defined by the frontend and returned as-is. Use semantic names such as title or content.
	//
	// This parameter is required.
	//
	// example:
	//
	// title
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The template content string in Go template syntax.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudMonitor alert: {{.incident.title}}
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
}

func (s PreviewNotifyTemplateRequestTemplates) String() string {
	return dara.Prettify(s)
}

func (s PreviewNotifyTemplateRequestTemplates) GoString() string {
	return s.String()
}

func (s *PreviewNotifyTemplateRequestTemplates) GetKey() *string {
	return s.Key
}

func (s *PreviewNotifyTemplateRequestTemplates) GetTemplate() *string {
	return s.Template
}

func (s *PreviewNotifyTemplateRequestTemplates) SetKey(v string) *PreviewNotifyTemplateRequestTemplates {
	s.Key = &v
	return s
}

func (s *PreviewNotifyTemplateRequestTemplates) SetTemplate(v string) *PreviewNotifyTemplateRequestTemplates {
	s.Template = &v
	return s
}

func (s *PreviewNotifyTemplateRequestTemplates) Validate() error {
	return dara.Validate(s)
}
