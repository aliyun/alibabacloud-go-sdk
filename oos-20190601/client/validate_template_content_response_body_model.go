// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTemplateContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutputs(v string) *ValidateTemplateContentResponseBody
	GetOutputs() *string
	SetParameters(v string) *ValidateTemplateContentResponseBody
	GetParameters() *string
	SetRamRole(v string) *ValidateTemplateContentResponseBody
	GetRamRole() *string
	SetRequestId(v string) *ValidateTemplateContentResponseBody
	GetRequestId() *string
	SetTasks(v []*ValidateTemplateContentResponseBodyTasks) *ValidateTemplateContentResponseBody
	GetTasks() []*ValidateTemplateContentResponseBodyTasks
}

type ValidateTemplateContentResponseBody struct {
	// The outputs of the template.
	//
	// example:
	//
	// {}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The parameters of the template.
	//
	// example:
	//
	// { "Status": { "Description": "(Required) The status of the Ecs instance.", "Type": "String" } }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The RAM role.
	//
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D5EE9591-1F2D-573E-8751-7F08BBB388D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task defined in the template.
	Tasks []*ValidateTemplateContentResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ValidateTemplateContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateContentResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponseBody) GetOutputs() *string {
	return s.Outputs
}

func (s *ValidateTemplateContentResponseBody) GetParameters() *string {
	return s.Parameters
}

func (s *ValidateTemplateContentResponseBody) GetRamRole() *string {
	return s.RamRole
}

func (s *ValidateTemplateContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateTemplateContentResponseBody) GetTasks() []*ValidateTemplateContentResponseBodyTasks {
	return s.Tasks
}

func (s *ValidateTemplateContentResponseBody) SetOutputs(v string) *ValidateTemplateContentResponseBody {
	s.Outputs = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetParameters(v string) *ValidateTemplateContentResponseBody {
	s.Parameters = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetRamRole(v string) *ValidateTemplateContentResponseBody {
	s.RamRole = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetRequestId(v string) *ValidateTemplateContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetTasks(v []*ValidateTemplateContentResponseBodyTasks) *ValidateTemplateContentResponseBody {
	s.Tasks = v
	return s
}

func (s *ValidateTemplateContentResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ValidateTemplateContentResponseBodyTasks struct {
	// The description of the task.
	//
	// example:
	//
	// (Required) The status of the Ecs instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// foo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The outputs of the task.
	//
	// example:
	//
	// .instanceId
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The properties of the task.
	//
	// example:
	//
	// {"API": "DescribeInstances","Parameters": {"Status": "{{ Status }}"},"Service": "Ecs"}
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// ACS::ExecuteAPI
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ValidateTemplateContentResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateContentResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *ValidateTemplateContentResponseBodyTasks) GetName() *string {
	return s.Name
}

func (s *ValidateTemplateContentResponseBodyTasks) GetOutputs() *string {
	return s.Outputs
}

func (s *ValidateTemplateContentResponseBodyTasks) GetProperties() *string {
	return s.Properties
}

func (s *ValidateTemplateContentResponseBodyTasks) GetType() *string {
	return s.Type
}

func (s *ValidateTemplateContentResponseBodyTasks) SetDescription(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetName(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetOutputs(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Outputs = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetProperties(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Properties = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetType(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Type = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
