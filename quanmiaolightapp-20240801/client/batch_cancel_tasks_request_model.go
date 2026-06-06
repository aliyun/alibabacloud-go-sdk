// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskCode(v string) *BatchCancelTasksRequest
	GetTaskCode() *string
	SetTaskIds(v []*string) *BatchCancelTasksRequest
	GetTaskIds() []*string
}

type BatchCancelTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// EssayCorrectionTask
	TaskCode *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
	// example:
	//
	// ["xxxx1","xxxx2"]
	TaskIds []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s BatchCancelTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTasksRequest) GoString() string {
	return s.String()
}

func (s *BatchCancelTasksRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *BatchCancelTasksRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *BatchCancelTasksRequest) SetTaskCode(v string) *BatchCancelTasksRequest {
	s.TaskCode = &v
	return s
}

func (s *BatchCancelTasksRequest) SetTaskIds(v []*string) *BatchCancelTasksRequest {
	s.TaskIds = v
	return s
}

func (s *BatchCancelTasksRequest) Validate() error {
	return dara.Validate(s)
}
