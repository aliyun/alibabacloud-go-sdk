// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSegmentationTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetVideoSegmentationTaskStatusRequest
	GetTaskId() *string
}

type GetVideoSegmentationTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetVideoSegmentationTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSegmentationTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetVideoSegmentationTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoSegmentationTaskStatusRequest) SetTaskId(v string) *GetVideoSegmentationTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
