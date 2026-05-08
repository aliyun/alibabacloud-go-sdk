// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetTrainTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunMainId(v string) *BatchGetTrainTaskShrinkRequest
	GetAliyunMainId() *string
	SetTaskIdListShrink(v string) *BatchGetTrainTaskShrinkRequest
	GetTaskIdListShrink() *string
}

type BatchGetTrainTaskShrinkRequest struct {
	// example:
	//
	// 1524004782431111
	AliyunMainId     *string `json:"aliyunMainId,omitempty" xml:"aliyunMainId,omitempty"`
	TaskIdListShrink *string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty"`
}

func (s BatchGetTrainTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetTrainTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskShrinkRequest) GetAliyunMainId() *string {
	return s.AliyunMainId
}

func (s *BatchGetTrainTaskShrinkRequest) GetTaskIdListShrink() *string {
	return s.TaskIdListShrink
}

func (s *BatchGetTrainTaskShrinkRequest) SetAliyunMainId(v string) *BatchGetTrainTaskShrinkRequest {
	s.AliyunMainId = &v
	return s
}

func (s *BatchGetTrainTaskShrinkRequest) SetTaskIdListShrink(v string) *BatchGetTrainTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

func (s *BatchGetTrainTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
