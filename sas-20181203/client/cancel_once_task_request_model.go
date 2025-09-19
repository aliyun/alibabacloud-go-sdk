// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOnceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *CancelOnceTaskRequest
	GetTaskId() *string
}

type CancelOnceTaskRequest struct {
	// The ID of the task.
	//
	// >  You can call the [GenerateOnceTask](~~GenerateOnceTask~~) operation to query the IDs of tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// ec9c0d88f36cc27765a98c554ee2****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelOnceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelOnceTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelOnceTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelOnceTaskRequest) SetTaskId(v string) *CancelOnceTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelOnceTaskRequest) Validate() error {
	return dara.Validate(s)
}
