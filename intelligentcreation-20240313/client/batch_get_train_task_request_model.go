// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetTrainTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunMainId(v string) *BatchGetTrainTaskRequest
	GetAliyunMainId() *string
	SetTaskIdList(v []*string) *BatchGetTrainTaskRequest
	GetTaskIdList() []*string
}

type BatchGetTrainTaskRequest struct {
	// example:
	//
	// 1524004782431111
	AliyunMainId *string   `json:"aliyunMainId,omitempty" xml:"aliyunMainId,omitempty"`
	TaskIdList   []*string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty" type:"Repeated"`
}

func (s BatchGetTrainTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskRequest) GetAliyunMainId() *string {
	return s.AliyunMainId
}

func (s *BatchGetTrainTaskRequest) GetTaskIdList() []*string {
	return s.TaskIdList
}

func (s *BatchGetTrainTaskRequest) SetAliyunMainId(v string) *BatchGetTrainTaskRequest {
	s.AliyunMainId = &v
	return s
}

func (s *BatchGetTrainTaskRequest) SetTaskIdList(v []*string) *BatchGetTrainTaskRequest {
	s.TaskIdList = v
	return s
}

func (s *BatchGetTrainTaskRequest) Validate() error {
	return dara.Validate(s)
}
