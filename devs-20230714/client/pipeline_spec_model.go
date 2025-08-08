// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPipelineSpec interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *Context) *PipelineSpec
	GetContext() *Context
	SetTemplateName(v string) *PipelineSpec
	GetTemplateName() *string
	SetTemplateSpec(v *PipelineTemplateSpec) *PipelineSpec
	GetTemplateSpec() *PipelineTemplateSpec
}

type PipelineSpec struct {
	Context *Context `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// my-pipeline-template
	TemplateName *string               `json:"templateName,omitempty" xml:"templateName,omitempty"`
	TemplateSpec *PipelineTemplateSpec `json:"templateSpec,omitempty" xml:"templateSpec,omitempty"`
}

func (s PipelineSpec) String() string {
	return dara.Prettify(s)
}

func (s PipelineSpec) GoString() string {
	return s.String()
}

func (s *PipelineSpec) GetContext() *Context {
	return s.Context
}

func (s *PipelineSpec) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PipelineSpec) GetTemplateSpec() *PipelineTemplateSpec {
	return s.TemplateSpec
}

func (s *PipelineSpec) SetContext(v *Context) *PipelineSpec {
	s.Context = v
	return s
}

func (s *PipelineSpec) SetTemplateName(v string) *PipelineSpec {
	s.TemplateName = &v
	return s
}

func (s *PipelineSpec) SetTemplateSpec(v *PipelineTemplateSpec) *PipelineSpec {
	s.TemplateSpec = v
	return s
}

func (s *PipelineSpec) Validate() error {
	return dara.Validate(s)
}
