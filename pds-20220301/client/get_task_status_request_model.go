// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *GetTaskStatusRequest
	GetDriveId() *string
	SetTaskId(v string) *GetTaskStatusRequest
	GetTaskId() *string
}

type GetTaskStatusRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i:SimilarImageClustering-b67d53e7-2fe8-460f-9b95-1e93636923eb
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskStatusRequest) SetDriveId(v string) *GetTaskStatusRequest {
	s.DriveId = &v
	return s
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
