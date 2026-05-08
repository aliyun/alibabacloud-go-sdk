// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetProjectTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIdList(v []*string) *BatchGetProjectTaskRequest
	GetTaskIdList() []*string
}

type BatchGetProjectTaskRequest struct {
	TaskIdList []*string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty" type:"Repeated"`
}

func (s BatchGetProjectTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskRequest) GetTaskIdList() []*string {
	return s.TaskIdList
}

func (s *BatchGetProjectTaskRequest) SetTaskIdList(v []*string) *BatchGetProjectTaskRequest {
	s.TaskIdList = v
	return s
}

func (s *BatchGetProjectTaskRequest) Validate() error {
	return dara.Validate(s)
}
