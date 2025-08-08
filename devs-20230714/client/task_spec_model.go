// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskSpec interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *Context) *TaskSpec
	GetContext() *Context
	SetTemplateName(v string) *TaskSpec
	GetTemplateName() *string
}

type TaskSpec struct {
	Context *Context `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// my-task-template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s TaskSpec) String() string {
	return dara.Prettify(s)
}

func (s TaskSpec) GoString() string {
	return s.String()
}

func (s *TaskSpec) GetContext() *Context {
	return s.Context
}

func (s *TaskSpec) GetTemplateName() *string {
	return s.TemplateName
}

func (s *TaskSpec) SetContext(v *Context) *TaskSpec {
	s.Context = v
	return s
}

func (s *TaskSpec) SetTemplateName(v string) *TaskSpec {
	s.TemplateName = &v
	return s
}

func (s *TaskSpec) Validate() error {
	return dara.Validate(s)
}
