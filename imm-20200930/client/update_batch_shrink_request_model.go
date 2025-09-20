// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionsShrink(v string) *UpdateBatchShrinkRequest
	GetActionsShrink() *string
	SetId(v string) *UpdateBatchShrinkRequest
	GetId() *string
	SetInputShrink(v string) *UpdateBatchShrinkRequest
	GetInputShrink() *string
	SetProjectName(v string) *UpdateBatchShrinkRequest
	GetProjectName() *string
	SetTagsShrink(v string) *UpdateBatchShrinkRequest
	GetTagsShrink() *string
}

type UpdateBatchShrinkRequest struct {
	// The processing templates.
	ActionsShrink *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The ID of the batch processing task. You can obtain the ID of the batch processing task from the response of the [CreateBatch](https://help.aliyun.com/document_detail/606694.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// batch-4eb9223f-3e88-42d3-a578-3f2852******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input data source.
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
	// {"key":"val"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateBatchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchShrinkRequest) GetActionsShrink() *string {
	return s.ActionsShrink
}

func (s *UpdateBatchShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateBatchShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *UpdateBatchShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateBatchShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateBatchShrinkRequest) SetActionsShrink(v string) *UpdateBatchShrinkRequest {
	s.ActionsShrink = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetId(v string) *UpdateBatchShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetInputShrink(v string) *UpdateBatchShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetProjectName(v string) *UpdateBatchShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetTagsShrink(v string) *UpdateBatchShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateBatchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
