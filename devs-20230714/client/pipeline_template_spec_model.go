// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPipelineTemplateSpec interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *Context) *PipelineTemplateSpec
	GetContext() *Context
	SetTasks(v []*TaskExec) *PipelineTemplateSpec
	GetTasks() []*TaskExec
}

type PipelineTemplateSpec struct {
	Context *Context    `json:"context,omitempty" xml:"context,omitempty"`
	Tasks   []*TaskExec `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
}

func (s PipelineTemplateSpec) String() string {
	return dara.Prettify(s)
}

func (s PipelineTemplateSpec) GoString() string {
	return s.String()
}

func (s *PipelineTemplateSpec) GetContext() *Context {
	return s.Context
}

func (s *PipelineTemplateSpec) GetTasks() []*TaskExec {
	return s.Tasks
}

func (s *PipelineTemplateSpec) SetContext(v *Context) *PipelineTemplateSpec {
	s.Context = v
	return s
}

func (s *PipelineTemplateSpec) SetTasks(v []*TaskExec) *PipelineTemplateSpec {
	s.Tasks = v
	return s
}

func (s *PipelineTemplateSpec) Validate() error {
	return dara.Validate(s)
}
