// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryTaskStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskCode(v string) *BatchQueryTaskStatusShrinkRequest
	GetTaskCode() *string
	SetTaskIdsShrink(v string) *BatchQueryTaskStatusShrinkRequest
	GetTaskIdsShrink() *string
}

type BatchQueryTaskStatusShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// EssayCorrectionTask
	TaskCode *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
	// example:
	//
	// ["xxxx1","xxxx2"]
	TaskIdsShrink *string `json:"taskIds,omitempty" xml:"taskIds,omitempty"`
}

func (s BatchQueryTaskStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryTaskStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryTaskStatusShrinkRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *BatchQueryTaskStatusShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *BatchQueryTaskStatusShrinkRequest) SetTaskCode(v string) *BatchQueryTaskStatusShrinkRequest {
	s.TaskCode = &v
	return s
}

func (s *BatchQueryTaskStatusShrinkRequest) SetTaskIdsShrink(v string) *BatchQueryTaskStatusShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *BatchQueryTaskStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
