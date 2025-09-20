// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTriggerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionsShrink(v string) *UpdateTriggerShrinkRequest
	GetActionsShrink() *string
	SetId(v string) *UpdateTriggerShrinkRequest
	GetId() *string
	SetInputShrink(v string) *UpdateTriggerShrinkRequest
	GetInputShrink() *string
	SetProjectName(v string) *UpdateTriggerShrinkRequest
	GetProjectName() *string
	SetTagsShrink(v string) *UpdateTriggerShrinkRequest
	GetTagsShrink() *string
}

type UpdateTriggerShrinkRequest struct {
	// The processing templates.
	ActionsShrink *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The ID of the trigger. You can obtain the ID of the trigger from the response of the [CreateTrigger](https://help.aliyun.com/document_detail/479912.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// trigger-9f72636a-0f0c-4baf-ae78-38b27b******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data source configurations.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateTriggerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTriggerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerShrinkRequest) GetActionsShrink() *string {
	return s.ActionsShrink
}

func (s *UpdateTriggerShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateTriggerShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *UpdateTriggerShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateTriggerShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateTriggerShrinkRequest) SetActionsShrink(v string) *UpdateTriggerShrinkRequest {
	s.ActionsShrink = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetId(v string) *UpdateTriggerShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetInputShrink(v string) *UpdateTriggerShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetProjectName(v string) *UpdateTriggerShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetTagsShrink(v string) *UpdateTriggerShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
