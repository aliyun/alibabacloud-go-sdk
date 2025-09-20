// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkflowParameter interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *WorkflowParameter
	GetName() *string
	SetValue(v string) *WorkflowParameter
	GetValue() *string
}

type WorkflowParameter struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WorkflowParameter) String() string {
	return dara.Prettify(s)
}

func (s WorkflowParameter) GoString() string {
	return s.String()
}

func (s *WorkflowParameter) GetName() *string {
	return s.Name
}

func (s *WorkflowParameter) GetValue() *string {
	return s.Value
}

func (s *WorkflowParameter) SetName(v string) *WorkflowParameter {
	s.Name = &v
	return s
}

func (s *WorkflowParameter) SetValue(v string) *WorkflowParameter {
	s.Value = &v
	return s
}

func (s *WorkflowParameter) Validate() error {
	return dara.Validate(s)
}
