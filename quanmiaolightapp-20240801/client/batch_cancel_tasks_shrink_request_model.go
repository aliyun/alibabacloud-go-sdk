// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskCode(v string) *BatchCancelTasksShrinkRequest
	GetTaskCode() *string
	SetTaskIdsShrink(v string) *BatchCancelTasksShrinkRequest
	GetTaskIdsShrink() *string
}

type BatchCancelTasksShrinkRequest struct {
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

func (s BatchCancelTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCancelTasksShrinkRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *BatchCancelTasksShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *BatchCancelTasksShrinkRequest) SetTaskCode(v string) *BatchCancelTasksShrinkRequest {
	s.TaskCode = &v
	return s
}

func (s *BatchCancelTasksShrinkRequest) SetTaskIdsShrink(v string) *BatchCancelTasksShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *BatchCancelTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
