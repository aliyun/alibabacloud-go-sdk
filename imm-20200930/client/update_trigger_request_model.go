// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*UpdateTriggerRequestActions) *UpdateTriggerRequest
	GetActions() []*UpdateTriggerRequestActions
	SetId(v string) *UpdateTriggerRequest
	GetId() *string
	SetInput(v *Input) *UpdateTriggerRequest
	GetInput() *Input
	SetProjectName(v string) *UpdateTriggerRequest
	GetProjectName() *string
	SetTags(v map[string]interface{}) *UpdateTriggerRequest
	GetTags() map[string]interface{}
}

type UpdateTriggerRequest struct {
	// The processing templates.
	Actions []*UpdateTriggerRequestActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The ID of the trigger. You can obtain the ID of the trigger from the response of the [CreateTrigger](https://help.aliyun.com/document_detail/479912.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// trigger-9f72636a-0f0c-4baf-ae78-38b27b******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data source configurations.
	Input *Input `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {"test": "val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequest) GetActions() []*UpdateTriggerRequestActions {
	return s.Actions
}

func (s *UpdateTriggerRequest) GetId() *string {
	return s.Id
}

func (s *UpdateTriggerRequest) GetInput() *Input {
	return s.Input
}

func (s *UpdateTriggerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateTriggerRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateTriggerRequest) SetActions(v []*UpdateTriggerRequestActions) *UpdateTriggerRequest {
	s.Actions = v
	return s
}

func (s *UpdateTriggerRequest) SetId(v string) *UpdateTriggerRequest {
	s.Id = &v
	return s
}

func (s *UpdateTriggerRequest) SetInput(v *Input) *UpdateTriggerRequest {
	s.Input = v
	return s
}

func (s *UpdateTriggerRequest) SetProjectName(v string) *UpdateTriggerRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateTriggerRequest) SetTags(v map[string]interface{}) *UpdateTriggerRequest {
	s.Tags = v
	return s
}

func (s *UpdateTriggerRequest) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTriggerRequestActions struct {
	// The template name.
	//
	// example:
	//
	// doc/convert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template parameters.
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s UpdateTriggerRequestActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateTriggerRequestActions) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequestActions) GetName() *string {
	return s.Name
}

func (s *UpdateTriggerRequestActions) GetParameters() []*string {
	return s.Parameters
}

func (s *UpdateTriggerRequestActions) SetName(v string) *UpdateTriggerRequestActions {
	s.Name = &v
	return s
}

func (s *UpdateTriggerRequestActions) SetParameters(v []*string) *UpdateTriggerRequestActions {
	s.Parameters = v
	return s
}

func (s *UpdateTriggerRequestActions) Validate() error {
	return dara.Validate(s)
}
