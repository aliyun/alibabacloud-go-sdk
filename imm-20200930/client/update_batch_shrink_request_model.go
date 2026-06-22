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
	// A list of data processing templates.
	ActionsShrink *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The ID of the batch processing task. For more information, see [Create a batch processing task](https://help.aliyun.com/document_detail/606694.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// batch-4eb9223f-3e88-42d3-a578-3f2852******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data source configuration.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Custom tags used to search for and filter asynchronous tasks.
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
