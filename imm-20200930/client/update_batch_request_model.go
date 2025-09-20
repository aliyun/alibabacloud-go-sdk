// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*UpdateBatchRequestActions) *UpdateBatchRequest
	GetActions() []*UpdateBatchRequestActions
	SetId(v string) *UpdateBatchRequest
	GetId() *string
	SetInput(v *Input) *UpdateBatchRequest
	GetInput() *Input
	SetProjectName(v string) *UpdateBatchRequest
	GetProjectName() *string
	SetTags(v map[string]interface{}) *UpdateBatchRequest
	GetTags() map[string]interface{}
}

type UpdateBatchRequest struct {
	// The processing templates.
	Actions []*UpdateBatchRequestActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The ID of the batch processing task. You can obtain the ID of the batch processing task from the response of the [CreateBatch](https://help.aliyun.com/document_detail/606694.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// batch-4eb9223f-3e88-42d3-a578-3f2852******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input data source.
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
	// {"key":"val"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchRequest) GetActions() []*UpdateBatchRequestActions {
	return s.Actions
}

func (s *UpdateBatchRequest) GetId() *string {
	return s.Id
}

func (s *UpdateBatchRequest) GetInput() *Input {
	return s.Input
}

func (s *UpdateBatchRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateBatchRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateBatchRequest) SetActions(v []*UpdateBatchRequestActions) *UpdateBatchRequest {
	s.Actions = v
	return s
}

func (s *UpdateBatchRequest) SetId(v string) *UpdateBatchRequest {
	s.Id = &v
	return s
}

func (s *UpdateBatchRequest) SetInput(v *Input) *UpdateBatchRequest {
	s.Input = v
	return s
}

func (s *UpdateBatchRequest) SetProjectName(v string) *UpdateBatchRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateBatchRequest) SetTags(v map[string]interface{}) *UpdateBatchRequest {
	s.Tags = v
	return s
}

func (s *UpdateBatchRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchRequestActions struct {
	// The name of the template.
	//
	// example:
	//
	// doc/convert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template parameters.
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s UpdateBatchRequestActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchRequestActions) GoString() string {
	return s.String()
}

func (s *UpdateBatchRequestActions) GetName() *string {
	return s.Name
}

func (s *UpdateBatchRequestActions) GetParameters() []*string {
	return s.Parameters
}

func (s *UpdateBatchRequestActions) SetName(v string) *UpdateBatchRequestActions {
	s.Name = &v
	return s
}

func (s *UpdateBatchRequestActions) SetParameters(v []*string) *UpdateBatchRequestActions {
	s.Parameters = v
	return s
}

func (s *UpdateBatchRequestActions) Validate() error {
	return dara.Validate(s)
}
