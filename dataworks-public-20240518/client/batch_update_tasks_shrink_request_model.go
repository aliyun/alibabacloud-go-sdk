// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *BatchUpdateTasksShrinkRequest
	GetComment() *string
	SetTasksShrink(v string) *BatchUpdateTasksShrinkRequest
	GetTasksShrink() *string
}

type BatchUpdateTasksShrinkRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The tasks.
	TasksShrink *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
}

func (s BatchUpdateTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *BatchUpdateTasksShrinkRequest) GetTasksShrink() *string {
	return s.TasksShrink
}

func (s *BatchUpdateTasksShrinkRequest) SetComment(v string) *BatchUpdateTasksShrinkRequest {
	s.Comment = &v
	return s
}

func (s *BatchUpdateTasksShrinkRequest) SetTasksShrink(v string) *BatchUpdateTasksShrinkRequest {
	s.TasksShrink = &v
	return s
}

func (s *BatchUpdateTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
