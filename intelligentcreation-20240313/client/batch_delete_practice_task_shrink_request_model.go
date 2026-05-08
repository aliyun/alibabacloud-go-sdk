// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePracticeTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdempotentId(v string) *BatchDeletePracticeTaskShrinkRequest
	GetIdempotentId() *string
	SetTaskIdsShrink(v string) *BatchDeletePracticeTaskShrinkRequest
	GetTaskIdsShrink() *string
}

type BatchDeletePracticeTaskShrinkRequest struct {
	// example:
	//
	// 1234567890
	IdempotentId  *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	TaskIdsShrink *string `json:"taskIds,omitempty" xml:"taskIds,omitempty"`
}

func (s BatchDeletePracticeTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePracticeTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeletePracticeTaskShrinkRequest) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *BatchDeletePracticeTaskShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *BatchDeletePracticeTaskShrinkRequest) SetIdempotentId(v string) *BatchDeletePracticeTaskShrinkRequest {
	s.IdempotentId = &v
	return s
}

func (s *BatchDeletePracticeTaskShrinkRequest) SetTaskIdsShrink(v string) *BatchDeletePracticeTaskShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *BatchDeletePracticeTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
