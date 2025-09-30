// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIdsShrink(v string) *UpdateVideoAnalysisTasksShrinkRequest
	GetTaskIdsShrink() *string
	SetTaskStatus(v string) *UpdateVideoAnalysisTasksShrinkRequest
	GetTaskStatus() *string
}

type UpdateVideoAnalysisTasksShrinkRequest struct {
	// This parameter is required.
	TaskIdsShrink *string `json:"taskIds,omitempty" xml:"taskIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CANCELED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s UpdateVideoAnalysisTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTasksShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *UpdateVideoAnalysisTasksShrinkRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *UpdateVideoAnalysisTasksShrinkRequest) SetTaskIdsShrink(v string) *UpdateVideoAnalysisTasksShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *UpdateVideoAnalysisTasksShrinkRequest) SetTaskStatus(v string) *UpdateVideoAnalysisTasksShrinkRequest {
	s.TaskStatus = &v
	return s
}

func (s *UpdateVideoAnalysisTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
