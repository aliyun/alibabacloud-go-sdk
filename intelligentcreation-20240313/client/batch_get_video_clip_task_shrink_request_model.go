// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetVideoClipTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIdListShrink(v string) *BatchGetVideoClipTaskShrinkRequest
	GetTaskIdListShrink() *string
}

type BatchGetVideoClipTaskShrinkRequest struct {
	TaskIdListShrink *string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty"`
}

func (s BatchGetVideoClipTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoClipTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskShrinkRequest) GetTaskIdListShrink() *string {
	return s.TaskIdListShrink
}

func (s *BatchGetVideoClipTaskShrinkRequest) SetTaskIdListShrink(v string) *BatchGetVideoClipTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

func (s *BatchGetVideoClipTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
