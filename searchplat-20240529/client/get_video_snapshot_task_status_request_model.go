// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSnapshotTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetVideoSnapshotTaskStatusRequest
	GetTaskId() *string
}

type GetVideoSnapshotTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetVideoSnapshotTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSnapshotTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetVideoSnapshotTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoSnapshotTaskStatusRequest) SetTaskId(v string) *GetVideoSnapshotTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
