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
	// example:
	//
	// {"incident":{"title":"CPU使用率过高","severity":"P1"}}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 待渲染的模板列表，每项包含模板标识和模板内容
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"key":"title","template":"云监控报警: {{.incident.title}}"},{"key":"content","template":">- 级别: {{.incident.severity}}"}]
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
	// 模板标识，前端自定义，原样返回。建议使用 title / content 等语义名称
	//
	// This parameter is required.
	//
	// example:
	//
	// title
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Go template 语法的模板内容字符串
	//
	// This parameter is required.
	//
	// example:
	//
	// 云监控报警: {{.incident.title}}
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
