// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskExec interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *Context) *TaskExec
	GetContext() *Context
	SetName(v string) *TaskExec
	GetName() *string
	SetRunAfters(v []*RunAfter) *TaskExec
	GetRunAfters() []*RunAfter
	SetTaskTemplate(v string) *TaskExec
	GetTaskTemplate() *string
}

type TaskExec struct {
	Context *Context `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// task-1
	Name      *string     `json:"name,omitempty" xml:"name,omitempty"`
	RunAfters []*RunAfter `json:"runAfters,omitempty" xml:"runAfters,omitempty" type:"Repeated"`
	// example:
	//
	// serverless-runner
	TaskTemplate *string `json:"taskTemplate,omitempty" xml:"taskTemplate,omitempty"`
}

func (s TaskExec) String() string {
	return dara.Prettify(s)
}

func (s TaskExec) GoString() string {
	return s.String()
}

func (s *TaskExec) GetContext() *Context {
	return s.Context
}

func (s *TaskExec) GetName() *string {
	return s.Name
}

func (s *TaskExec) GetRunAfters() []*RunAfter {
	return s.RunAfters
}

func (s *TaskExec) GetTaskTemplate() *string {
	return s.TaskTemplate
}

func (s *TaskExec) SetContext(v *Context) *TaskExec {
	s.Context = v
	return s
}

func (s *TaskExec) SetName(v string) *TaskExec {
	s.Name = &v
	return s
}

func (s *TaskExec) SetRunAfters(v []*RunAfter) *TaskExec {
	s.RunAfters = v
	return s
}

func (s *TaskExec) SetTaskTemplate(v string) *TaskExec {
	s.TaskTemplate = &v
	return s
}

func (s *TaskExec) Validate() error {
	return dara.Validate(s)
}
