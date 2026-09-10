// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *GetDataCheckConfigRequest
	GetTaskId() *int64
}

type GetDataCheckConfigRequest struct {
	// The task ID that uniquely identifies a task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetDataCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDataCheckConfigRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetDataCheckConfigRequest) SetTaskId(v int64) *GetDataCheckConfigRequest {
	s.TaskId = &v
	return s
}

func (s *GetDataCheckConfigRequest) Validate() error {
	return dara.Validate(s)
}
