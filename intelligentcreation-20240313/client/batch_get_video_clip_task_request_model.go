// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetVideoClipTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIdList(v []*string) *BatchGetVideoClipTaskRequest
	GetTaskIdList() []*string
}

type BatchGetVideoClipTaskRequest struct {
	TaskIdList []*string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty" type:"Repeated"`
}

func (s BatchGetVideoClipTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoClipTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskRequest) GetTaskIdList() []*string {
	return s.TaskIdList
}

func (s *BatchGetVideoClipTaskRequest) SetTaskIdList(v []*string) *BatchGetVideoClipTaskRequest {
	s.TaskIdList = v
	return s
}

func (s *BatchGetVideoClipTaskRequest) Validate() error {
	return dara.Validate(s)
}
