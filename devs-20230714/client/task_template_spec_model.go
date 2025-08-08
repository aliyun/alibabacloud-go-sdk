// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskTemplateSpec interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *Context) *TaskTemplateSpec
	GetContext() *Context
	SetDescription(v string) *TaskTemplateSpec
	GetDescription() *string
	SetExecuteCondition(v *Condition) *TaskTemplateSpec
	GetExecuteCondition() *Condition
	SetWorker(v *TaskWorker) *TaskTemplateSpec
	GetWorker() *TaskWorker
}

type TaskTemplateSpec struct {
	Context *Context `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// build&deploy.
	Description      *string     `json:"description,omitempty" xml:"description,omitempty"`
	ExecuteCondition *Condition  `json:"executeCondition,omitempty" xml:"executeCondition,omitempty"`
	Worker           *TaskWorker `json:"worker,omitempty" xml:"worker,omitempty"`
}

func (s TaskTemplateSpec) String() string {
	return dara.Prettify(s)
}

func (s TaskTemplateSpec) GoString() string {
	return s.String()
}

func (s *TaskTemplateSpec) GetContext() *Context {
	return s.Context
}

func (s *TaskTemplateSpec) GetDescription() *string {
	return s.Description
}

func (s *TaskTemplateSpec) GetExecuteCondition() *Condition {
	return s.ExecuteCondition
}

func (s *TaskTemplateSpec) GetWorker() *TaskWorker {
	return s.Worker
}

func (s *TaskTemplateSpec) SetContext(v *Context) *TaskTemplateSpec {
	s.Context = v
	return s
}

func (s *TaskTemplateSpec) SetDescription(v string) *TaskTemplateSpec {
	s.Description = &v
	return s
}

func (s *TaskTemplateSpec) SetExecuteCondition(v *Condition) *TaskTemplateSpec {
	s.ExecuteCondition = v
	return s
}

func (s *TaskTemplateSpec) SetWorker(v *TaskWorker) *TaskTemplateSpec {
	s.Worker = v
	return s
}

func (s *TaskTemplateSpec) Validate() error {
	return dara.Validate(s)
}
