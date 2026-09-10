// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *GetDataCheckTaskConfigRequest
	GetTaskId() *int64
}

type GetDataCheckTaskConfigRequest struct {
	// The ID of the data validation task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetDataCheckTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskConfigRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetDataCheckTaskConfigRequest) SetTaskId(v int64) *GetDataCheckTaskConfigRequest {
	s.TaskId = &v
	return s
}

func (s *GetDataCheckTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
