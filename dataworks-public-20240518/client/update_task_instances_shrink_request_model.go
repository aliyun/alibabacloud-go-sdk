// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateTaskInstancesShrinkRequest
	GetComment() *string
	SetTaskInstancesShrink(v string) *UpdateTaskInstancesShrinkRequest
	GetTaskInstancesShrink() *string
}

type UpdateTaskInstancesShrinkRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The instances.
	TaskInstancesShrink *string `json:"TaskInstances,omitempty" xml:"TaskInstances,omitempty"`
}

func (s UpdateTaskInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateTaskInstancesShrinkRequest) GetTaskInstancesShrink() *string {
	return s.TaskInstancesShrink
}

func (s *UpdateTaskInstancesShrinkRequest) SetComment(v string) *UpdateTaskInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *UpdateTaskInstancesShrinkRequest) SetTaskInstancesShrink(v string) *UpdateTaskInstancesShrinkRequest {
	s.TaskInstancesShrink = &v
	return s
}

func (s *UpdateTaskInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
