// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetProjectTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIdListShrink(v string) *BatchGetProjectTaskShrinkRequest
	GetTaskIdListShrink() *string
}

type BatchGetProjectTaskShrinkRequest struct {
	TaskIdListShrink *string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty"`
}

func (s BatchGetProjectTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetProjectTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskShrinkRequest) GetTaskIdListShrink() *string {
	return s.TaskIdListShrink
}

func (s *BatchGetProjectTaskShrinkRequest) SetTaskIdListShrink(v string) *BatchGetProjectTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

func (s *BatchGetProjectTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
