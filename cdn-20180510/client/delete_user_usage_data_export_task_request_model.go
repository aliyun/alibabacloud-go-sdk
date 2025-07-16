// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserUsageDataExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DeleteUserUsageDataExportTaskRequest
	GetTaskId() *string
}

type DeleteUserUsageDataExportTaskRequest struct {
	// The ID of the task to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteUserUsageDataExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserUsageDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserUsageDataExportTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteUserUsageDataExportTaskRequest) SetTaskId(v string) *DeleteUserUsageDataExportTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteUserUsageDataExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
